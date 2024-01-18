package cgtmp

import (
	"fmt"
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

func (t *Template) NameWithExtension() string {
	return fmt.Sprintf("%s.%s", t.Name, t.extension)
}

func (t *Template) FullName() string {
	return fmt.Sprintf("%s/%s", t.path, t.NameWithExtension())
}
