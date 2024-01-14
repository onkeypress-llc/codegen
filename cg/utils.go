package cg

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"text/template"
)

func FormatGoString(content string) (string, error) {
	formatted, err := format.Source([]byte(content))
	return string(formatted), err
}

func getTemplateFile(name string) string {
	return fmt.Sprintf("templates/%s.tmp", name)
}

func getTemplateFiles(names []string) []string {
	fileNames := make([]string, len(names))
	for i := range names {
		fileNames[i] = getTemplateFile(names[i])
	}
	return fileNames
}

func ExecuteTemplates(object any, templateFS embed.FS, templateNames ...string) (string, error) {
	tmp, err := template.ParseFS(templateFS, getTemplateFiles(templateNames)...)
	if err != nil {
		return "", err
	}
	var buffer bytes.Buffer
	err = tmp.Execute(&buffer, object)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
