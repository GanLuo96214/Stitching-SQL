package main

import (
	"database/sql"
	"github.com/ganluo960214/StringCase"
	"log"
	"strings"
)

type postgresTable struct {
	name    string
	columns []postgresTableColumn
}
type postgresTableColumn struct {
	AttName     string
	AttNotNull  bool
	TypName     string
	AttTypMod   int32
	TypLen      int32
	TypArray    int32
	EleTypName  sql.NullString
	EleTypLen   sql.NullInt32
	Description sql.NullString
}

type ErrorUnsupportedPostgresDataType struct {
	PostgresDataType string
}

func (t ErrorUnsupportedPostgresDataType) Error() string {
	return "not support postgres data type: " + t.PostgresDataType
}

func (c postgresTableColumn) GoType() (string, error) {
	isArray := false
	if c.TypArray == 0 {
		isArray = true
	}

	typeName := c.TypName
	if isArray {
		typeName = c.EleTypName.String
	}

	attTypMod, typeLength := c.AttTypMod, c.TypLen
	if isArray {
		typeLength = c.EleTypLen.Int32
	}

	length := typeLength
	if length == -1 {
		length = attTypMod
	}

	b := strings.Builder{}
	if isArray {
		b.WriteString("[]")
	}

	switch typeName {
	case "char":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "aclitem":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "int8":
		if c.AttNotNull {
			b.WriteString("int64")
		} else {
			b.WriteString("sql.NullInt64")
		}
	case "bit":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "varbit":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "bool":
		if c.AttNotNull {
			b.WriteString("bool")
		} else {
			b.WriteString("sql.NullBool")
		}
	case "box":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "bytea":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "bpchar":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "cid":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "cidr":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "circle":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "date":
		if c.AttNotNull {
			b.WriteString("time.Time")
		} else {
			b.WriteString("sql.NullTime")
		}
	case "daterange":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "float8":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "gtsvector":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "inet":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "int2vector":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "int4range":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "int8range":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "int4":
		if c.AttNotNull {
			b.WriteString("int32")
		} else {
			b.WriteString("sql.NullInt32")
		}
	case "interval":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "json":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "jsonb":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "jsonpath":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "line":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "lseg":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "macaddr":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "macaddr8":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "money":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "name":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "numeric":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "numrange":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "oid":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "oidvector":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "path":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "pg_lsn":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "point":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "polygon":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "float4":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "refcursor":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "regclass":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "regconfig":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "regdictionary":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "regnamespace":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "regoper":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "regoperator":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "regproc":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "regprocedure":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "regrole":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "regtype":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "int2":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "text":
		if c.AttNotNull {
			b.WriteString("string")
		} else {
			b.WriteString("sql.NullString")
		}
	case "tid":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "timetz":
		if c.AttNotNull {
			b.WriteString("time.Time")
		} else {
			b.WriteString("sql.NullTime")
		}
	case "time":
		if c.AttNotNull {
			b.WriteString("time.Time")
		} else {
			b.WriteString("sql.NullTime")
		}
	case "timestamptz":
		if c.AttNotNull {
			b.WriteString("time.Time")
		} else {
			b.WriteString("sql.NullTime")
		}
	case "timestamp":
		if c.AttNotNull {
			b.WriteString("time.Time")
		} else {
			b.WriteString("sql.NullTime")
		}
	case "tsquery":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "tsrange":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "tstzrange":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "tsvector":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "txid_snapshot":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "uuid":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "xid":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	case "xml":
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	default:
		return "", ErrorUnsupportedPostgresDataType{PostgresDataType: typeName}
	}

	return b.String(), nil
}

func postgres(source, schema string) []tableContent {

	db, err := sql.Open("postgres", source)
	if err != nil {
		log.Fatal("connection error")
	}

	// ping database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// schema exists
	schemaIsExists := false
	if err := db.QueryRow("select exists(select 1 from pg_namespace where nspname = $1)", schema).Scan(&schemaIsExists); err != nil {
		log.Fatal(err)
	}

	if schemaIsExists == false {
		log.Fatal("schema not exists")
	}

	// tables
	tables := make([]postgresTable, 0)
	rows, err := db.Query("select tablename from pg_tables where schemaname = $1", schema)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		table := postgresTable{}
		if err := rows.Scan(&table.name); err != nil {
			log.Fatal(err)
		}
		tables = append(tables, table)
	}

	// columns
	for i := range tables {
		columns := make([]postgresTableColumn, 0)
		rows, err := db.Query(`
select
	pg_attribute.attname
	,pg_attribute.attnotnull
	,pg_type.typname
	,pg_attribute.atttypmod
	,pg_type.typlen
	,pg_type.typarray
	,ele_type.typname ele_typname
	,ele_type.typlen ele_typlen
	,pg_description.description
from pg_attribute
left join pg_class on pg_class.oid = pg_attribute.attrelid
left join pg_type on pg_type.oid = pg_attribute.atttypid
left join pg_type ele_type on ele_type.oid = pg_type.typelem
left join pg_namespace on pg_namespace.oid = pg_class.relnamespace
left join pg_description on pg_description.objoid = pg_class.oid and pg_description.objsubid = pg_attribute.attnum
where
pg_namespace.nspname = $1
and pg_class.relname = $2
and attnum > 0
and attisdropped = false;`, schema, tables[i].name)
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			column := postgresTableColumn{}
			if err := rows.Scan(
				&column.AttName,
				&column.AttNotNull,
				&column.TypName,
				&column.AttTypMod,
				&column.TypLen,
				&column.TypArray,
				&column.EleTypName,
				&column.EleTypLen,
				&column.Description); err != nil {
				log.Fatal(err)
			}
			columns = append(columns, column)
		}
		tables[i].columns = columns
	}

	cs := make([]tableContent, 0, len(tables))
	for _, table := range tables {
		t := TableTemplateContent{
			TypeName:    StringCase.ToUpperCamelCase(table.name),
			StructField: []TableTemplateContentStructField{},
		}

		for _, column := range table.columns {
			goType, err := column.GoType()
			if err != nil {
				log.Fatal(err)
			}
			t.StructField = append(t.StructField, TableTemplateContentStructField{
				FieldName: StringCase.ToUpperCamelCase(column.AttName),
				Type:      goType,
				Comment:   column.Description.String,
			})
		}

		c, err := t.generateContent()
		if err != nil {
			log.Fatal(err)
		}

		cs = append(cs, tableContent{
			typeName:  t.TypeName,
			tableName: table.name,
			content:   c,
		})
	}

	return cs
}
