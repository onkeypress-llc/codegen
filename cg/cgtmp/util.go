package cgtmp

import (
	"embed"
)

//go:embed templates/*.tmp
var templateFS embed.FS

func TemplateFS() embed.FS {
	return templateFS
}
