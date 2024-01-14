package cg

import "embed"

//go:embed templates/*.tmp
var templateFS embed.FS

type NodeInterface interface {
	Generate(GeneratorContextInterface) (NodeOutputInterface, error)
	UsedImports() *TemplateImportSet
}

type NodeOutputInterface interface {
	ToString() (string, error)
}

type NodeTemplateOutput[D any] struct {
	data          *D
	templateNames []string
	fs            embed.FS
}

func (o *NodeTemplateOutput[D]) ToString() (string, error) {
	return ExecuteTemplates(o.data, o.fs, o.templateNames...)
}

func TemplateOutput[D any](data *D, templateNames ...string) NodeOutputInterface {
	return &NodeTemplateOutput[D]{data: data, templateNames: templateNames, fs: templateFS}
}

type NodeStringOutput[D any] struct {
	data     *D
	toString func(*D) (string, error)
}

func (o *NodeStringOutput[D]) ToString() (string, error) {
	return o.toString(o.data)
}

func (o *NodeStringOutput[D]) ToFormattedString() (string, error) {
	value, err := o.ToString()
	if err != nil {
		return "", err
	}
	return FormatGoString(value)
}

func (o *NodeStringOutput[D]) Data() *D {
	return o.data
}

func StringOutput[D any](data *D, toString func(*D) (string, error)) NodeOutputInterface {
	return &NodeStringOutput[D]{data: data, toString: toString}
}
