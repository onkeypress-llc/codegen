package cgnode

import (
	"embed"

	"github.com/onkeypress-llc/codegen/cg/cgtmp"
)

type NodeTemplateOutput[D any] struct {
	name      string
	data      *D
	templates *cgtmp.Templates
	fs        embed.FS
}

func (o *NodeTemplateOutput[D]) Name() string {
	return o.name
}

func (o *NodeTemplateOutput[D]) ToString() (string, error) {
	return cgtmp.ExecuteTemplates(o.data, o.fs, o.templates)
}

func (o *NodeTemplateOutput[D]) Templates() *cgtmp.Templates {
	return o.templates
}

func (o *NodeTemplateOutput[D]) UsedImports() (*ImportSet, error) {
	return nil, nil
}

func (o *NodeTemplateOutput[D]) SetName(name string) *NodeTemplateOutput[D] {
	o.name = name
	return o
}

func (o *NodeTemplateOutput[D]) SetTemplates(templates *cgtmp.Templates) *NodeTemplateOutput[D] {
	o.templates = templates
	return o
}

func (o *NodeTemplateOutput[D]) ToInterface() NodeOutputInterface {
	return o
}

func TemplateOutput[D any](templateFS embed.FS, data *D, templates *cgtmp.Templates) *NodeTemplateOutput[D] {
	return &NodeTemplateOutput[D]{name: GetTypeString(data), data: data, templates: templates, fs: templateFS}
}
