package cgtmp

import (
	"bytes"
	"embed"
	"html/template"

	"github.com/onkeypress-llc/codegen/cg/cgi"
)

//go:embed templates/*.tmp
var templateFS embed.FS

func TemplateFS() embed.FS {
	return templateFS
}

func ExecuteTemplate(ctx cgi.ContextInterface, obj cgi.NodeOutputInterface) (string, error) {
	templates := NewSet(obj.Template()).AddTemplates(obj.Templates())
	data := obj.UntypedData()

	tmp, err := template.New("").Funcs(map[string]any{
		// convenience method for
		"context": func() cgi.ContextInterface {
			return ctx
		},
		"stringify": func(o cgi.NodeWithStringOutput) (string, error) {
			if o == nil {
				return "", nil
			}
			return o.ToString(ctx)
		},
	}).ParseFS(ctx.TemplateFS(), templates.Names()...)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	err = tmp.ExecuteTemplate(&buffer, obj.Template().NameWithExtension(), data)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
