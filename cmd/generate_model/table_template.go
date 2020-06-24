package main

import (
	"bytes"
	"github.com/ganluo960214/StringCase"
	"log"
	"text/template"
)

const (
	TableTemplate = `type {{.TypeName}} struct { {{range $v := .StructField}}
    {{.FieldName}} {{.Type}} // {{.Comment}} {{end}}
}`
)

type TableTemplateContent struct {
	TypeName    string
	StructField []TableTemplateContentStructField
}
type TableTemplateContentStructField struct {
	FieldName string
	Type      string
	Comment   string
}

func (c TableTemplateContent) generateContent() ([]byte, error) {
	// file content container
	b := bytes.NewBuffer([]byte{})

	t, err := template.New("").Funcs(template.FuncMap{
		"stringCaseToSnakeCase": StringCase.ToSnakeCase,
	}).Parse(TableTemplate)
	if err != nil {
		return nil, err
	}

	// generate file content write to file content container
	if err := t.Execute(b, c); err != nil {
		log.Fatal(err)
	}
	return b.Bytes(), nil
}
