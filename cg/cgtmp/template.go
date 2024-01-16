package cgtmp

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"
)

type Template struct {
	Name      string
	path      string
	extension string
}

func New(name string) *Template {
	return &Template{Name: name, path: "templates", extension: "tmp"}
}

func (t *Template) Path(path string) *Template {
	t.path = path
	return t
}

func (t *Template) Extension(extension string) *Template {
	t.extension = extension
	return t
}

func (t *Template) FullName() string {
	return fmt.Sprintf("%s/%s.%s", t.path, t.Name, t.extension)
}

func ExecuteTemplates(object any, templateFS embed.FS, templates *Templates) (string, error) {
	tmp, err := template.ParseFS(templateFS, templates.Names()...)
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
