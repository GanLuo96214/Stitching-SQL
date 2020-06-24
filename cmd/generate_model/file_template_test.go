package main

import "testing"

func TestFileTemplateContent_GenerateFileContent(t *testing.T) {
	c := FileTemplateContent{
		Package: "this_is_package",
		Content: "this_is_content",
	}

	_, err := c.generateContent()
	if err != nil {
		t.Fatal(err)
	}
}
