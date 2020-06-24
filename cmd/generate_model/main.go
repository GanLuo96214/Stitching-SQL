package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
)

/*
env
*/
var (
	env = struct {
		GoPackage string `validate:"required"`
	}{
		GoPackage: os.Getenv("GOPACKAGE"),
	}
)

/*
flags
*/
var (
	flags = struct {
		Type      flagsDatabaseType `validate:"required,eq=mysql|eq=postgres"`
		Source    string            `validate:"required"`
		Schema    string            `validate:"required"`
		Directory string
	}{Directory: "./"}
)

type flagsDatabaseType string

const (
	flagsDatabaseTypeMysql    flagsDatabaseType = "mysql"
	flagsDatabaseTypePostgres flagsDatabaseType = "postgres"
)

func init() {
	flag.StringVar((*string)(&flags.Type), "type", "", "type; must be set; database type: mysql, postgres.")
	flag.StringVar(&flags.Source, "source", "", "source; must be set; database connection string.")
	flag.StringVar(&flags.Schema, "schema", "", "schema; must be set; schema name")
	flag.StringVar(&flags.Directory, "directory", flags.Directory, "directory; generate files save to in directory.")
}

func usage() {
	log.Print(`
generate model by database table.

-type
	support mysql, postgres. must be set.
	used driver:
		mysql    : https://github.com/go-sql-driver
		postgres : https://github.com/lib/pq

-source
	database connection string. must be set.
	example:
		mysql : [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
		ref   : https://github.com/go-sql-driver/mysql#usage
		
		postgres : postgres://[username[:password]@address/dbname[?param1=value1&...&paramN=valueN]
		ref      : https://pkg.go.dev/github.com/lib/pq?tab=doc#pkg-overview

-schema
	schema name. must be set.
	
-directory (default: ` + flags.Directory + `)
	generate files save to in directory.
`)
	os.Exit(0)
}

type tableContent struct {
	typeName  string
	tableName string
	content   []byte
}

func main() {

	log.SetFlags(log.Lshortfile)
	flag.Usage = usage
	flag.Parse()

	// check env
	if err := validate.Struct(env); err != nil {
		log.Fatal(err)
	}

	// check flags
	if err := validate.Struct(flags); err != nil {
		log.Fatal(err)
	}

	var cs []tableContent
	switch flags.Type {
	case flagsDatabaseTypeMysql:
		cs = mysql(flags.Source, flags.Schema)
	case flagsDatabaseTypePostgres:
		cs = postgres(flags.Source, flags.Schema)
	default:
		log.Fatalf("%s database not support", flags.Type)
	}

	for _, c := range cs {

		fileName := c.tableName + ".go"

		// --- table import package
		tableImports := make(map[string]interface{}, 0)
		src := fmt.Sprintf("package temp\n%s", c.content)

		astF, err := parser.ParseFile(token.NewFileSet(), "", src, 0)
		if err != nil {
			log.Fatal(err)
		}
		tableFieldTypes := make([]string, 0)
		ast.Inspect(astF, func(node ast.Node) bool {
			t, ok := node.(*ast.TypeSpec)
			if ok == false {
				return true
			}

			st, ok := t.Type.(*ast.StructType)
			if ok == false {
				return true
			}

			if t.Name.Name != c.typeName {
				return true
			}

			for _, f := range st.Fields.List {
				typeName := ""

				if i, ok := f.Type.(*ast.Ident); ok {
					typeName = i.Name
				} else {
					if s, ok := f.Type.(*ast.SelectorExpr); ok {
						if i, ok := s.X.(*ast.Ident); ok {
							typeName = i.Name + "." + s.Sel.Name
						} else {
							log.Fatal("something wrong")
						}
					} else {
						log.Fatal("something wrong")
					}
				}
				tableFieldTypes = append(tableFieldTypes, typeName)
			}
			return false
		})

		var (
			importsRef = map[string]string{
				"sql.NullTime":    "database/sql",
				"sql.NullInt64":   "database/sql",
				"sql.NullInt32":   "database/sql",
				"sql.NullBool":    "database/sql",
				"sql.NullFloat64": "database/sql",
				"sql.NullString":  "database/sql",
				"time.Time":       "time",
			}
		)

		// generate file imports
		for _, t := range tableFieldTypes {
			switch t {
			case "int32", "int64", "string", "bool":
				continue
			}

			v, ok := importsRef[t]
			if ok == false {
				log.Fatal("data type can't find match import package")
			}
			tableImports[v] = nil

		}

		// check file exists
		if _, err := os.Stat(fileName); os.IsNotExist(err) { // if not exists then create new file

			t := FileTemplateContent{
				Imports: tableImports,
				Package: env.GoPackage,
				Content: string(c.content),
			}
			fc, err := t.generateContent()
			if err != nil {
				log.Fatal(err)
			}

			ffc, err := format.Source(fc)
			if err != nil {
				log.Fatal(err)
			}

			if err := ioutil.WriteFile(fileName, ffc, 0600); err != nil {
				log.Fatal(err)
			}
		} else { // if exists only replace table struct

			fset := token.NewFileSet()
			astF, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
			if err != nil {
				log.Fatal(err)
			}

			// find file imports
			fileImports, lastImportPosition := make(map[string]interface{}, 0), token.Position{}
			for _, is := range astF.Imports {
				ast.Inspect(is, func(node ast.Node) bool {
					l, ok := node.(*ast.BasicLit)
					if ok == false {
						return true
					}

					fileImports[l.Value[1:len(l.Value)-1]] = nil
					lastImportPosition = fset.Position(node.End())
					return false
				})
			}

			importDiff := make([]string, 0)
			for k, _ := range tableImports {
				if _, ok := fileImports[k]; ok == false {
					importDiff = append(importDiff, k)
				}
			}

			// find struct begin and end position
			isFind := false
			structBeginPosition, structEndPosition := token.Position{}, token.Position{}
			ast.Inspect(astF, func(n1 ast.Node) bool {
				if n1 == nil {
					return true
				}

				d, ok := n1.(*ast.GenDecl)
				if ok == false {
					return true
				}

				ast.Inspect(d.Specs[0], func(n2 ast.Node) bool {

					t, ok := n2.(*ast.TypeSpec)
					if ok == false {
						return true
					}

					_, ok = t.Type.(*ast.StructType)
					if ok == false {
						return true
					}

					if t.Name.Name == c.typeName {
						isFind = true
					}

					return false
				})
				if isFind == false {
					return true
				}

				structBeginPosition = fset.Position(d.Pos())
				structEndPosition = fset.Position(d.End())

				return false
			})

			if !isFind {
				log.Fatalf("not find %s struct in %s", c.typeName, fileName)
			}

			if structBeginPosition.Filename != structEndPosition.Filename {
				log.Fatal("don't know how to do it")
			}

			// modify file
			f, err := os.Open(fileName)
			if err != nil {
				log.Fatal(err)
			}
			fi, err := f.Stat()
			if err != nil {
				log.Fatal(err)
			}
			s := bufio.NewScanner(f)

			b := bytes.Buffer{}
			b.Grow(int(fi.Size()))

			for i := 1; s.Scan(); i++ {
				if i == lastImportPosition.Line {
					l := s.Bytes()
					for i := 0; i < len(l); i++ {
						b.WriteByte(l[i])
					}
					for _, i := range importDiff {
						b.WriteByte('\n')
						b.WriteString(fmt.Sprintf("\"%s\"", i))
					}
					b.WriteByte('\n')
				} else if i >= structBeginPosition.Line && i <= structEndPosition.Line {
					if i == structBeginPosition.Line {
						l := s.Bytes()
						for i := 0; i < structBeginPosition.Column-1; i++ {
							b.WriteByte(l[i])
						}
						b.Write(c.content)
					} else if i == structEndPosition.Line {
						l := s.Bytes()
						for i := structEndPosition.Column - 1; i < len(l); i++ {
							b.WriteByte(l[i])
						}
						b.WriteByte('\n')
					}
				} else {
					b.Write(s.Bytes())
					b.WriteByte('\n')
				}

			}

			fs, err := format.Source(b.Bytes())
			if err != nil {
				log.Fatal(err)
			}
			if err := ioutil.WriteFile(fileName, fs, 0600); err != nil {
				log.Fatal(err)
			}
			if err := f.Close(); err != nil {
				log.Fatal(err)
			}

		}

	}

}
