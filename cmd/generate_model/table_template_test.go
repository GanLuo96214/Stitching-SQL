package main

import "testing"

func TestTableTemplateContent_GenerateFileContent(t *testing.T) {
	c := TableTemplateContent{
		TypeName:    "TestTableTemplateContent_GenerateFileContent",
		StructField: nil,
	}

	_, err := c.generateContent()
	if err != nil {
		t.Fatal(err)
	}
}
