package main

import (
	"bytes"
	"log"
	"text/template"
)

const FileTemplate = `
package {{.Package}}

{{$len := len .Imports}}{{ if ne $len 0 }}import ({{range $k,$v := .Imports}}
    "{{$k}}" {{end}}
){{ end }}

{{.Content}}
`

/*
file template
*/
type FileTemplateContent struct {
	Package string
	Imports map[string]interface{}
	Content string
}

func (c FileTemplateContent) generateContent() ([]byte, error) {
	// file content container
	b := bytes.NewBuffer([]byte{})

	// new file template
	t, err := template.New("").Parse(FileTemplate)
	if err != nil {
		log.Fatal(err)
	}
	// generate file content write to file content container
	if err := t.Execute(b, c); err != nil {
		log.Fatal(err)
	}

	return b.Bytes(), nil
}
