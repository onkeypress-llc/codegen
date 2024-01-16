package cg

import (
	"embed"
	"go/format"
)

//go:embed templates/*.tmp
var templateFS embed.FS

func FormatGoString(content string) (string, error) {
	formatted, err := format.Source([]byte(content))
	return string(formatted), err
}

func TemplateFS() embed.FS {
	return templateFS
}
