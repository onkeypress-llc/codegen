package cgnode

import (
	"bytes"
	"reflect"
	"text/template"

	"github.com/onkeypress-llc/codegen/cg/cgcontext"
	"github.com/onkeypress-llc/codegen/cg/cgtmp"
)

type NodeWithName interface {
	// the canonical name of the output data
	Name() string
}

type NodeWithImports interface {
	// set of all imports necessary
	UsedImports() (*ImportSet, error)
}

type NodeWithTemplate interface {
	// primary template to use
	Template() *cgtmp.Template
	// set of all templates used
	Templates() *cgtmp.Templates
}

type NodeWithData[D any] interface {
	// the data to be used
	Data() D
}

type NodeWithStringOutput interface {
	ToString(cgcontext.Interface) (string, error)
}

type NodeOutputInterface[D any] interface {
	NodeWithName
	NodeWithImports
	NodeWithTemplate
	NodeWithData[D]
	NodeWithStringOutput
}

type NodeOutput[D any] struct {
	name      string
	data      D
	template  *cgtmp.Template
	templates *cgtmp.Templates
	imports   *ImportSet

	toStringFunction func(cgcontext.Interface, *NodeOutput[D]) (string, error)
}

func (o *NodeOutput[D]) Name() string {
	return o.name
}

func (o *NodeOutput[D]) SetName(name string) *NodeOutput[D] {
	o.name = name
	return o
}

func (o *NodeOutput[D]) Templates() *cgtmp.Templates {
	return o.templates
}

func (o *NodeOutput[D]) SetTemplates(templates *cgtmp.Templates) *NodeOutput[D] {
	o.templates = templates
	return o
}

func (o *NodeOutput[D]) UsedImports() (*ImportSet, error) {
	return o.imports, nil
}

func (o *NodeOutput[D]) SetUsedImports(imports *ImportSet) *NodeOutput[D] {
	o.imports = imports
	return o
}

func (o *NodeOutput[D]) Template() *cgtmp.Template {
	return o.template
}

func (o *NodeOutput[D]) ToString(ctx cgcontext.Interface) (string, error) {
	return o.toStringFunction(ctx, o)
}

func (o *NodeOutput[D]) Data() D {
	return o.data
}

func (o *NodeOutput[D]) ToInterface() NodeOutputInterface[D] {
	return o
}

func Output[D any](template *cgtmp.Template, data D) *NodeOutput[D] {
	return &NodeOutput[D]{data: data, template: template, name: GetTypeString(data), templates: cgtmp.NewSet(), toStringFunction: NodeObjectExecuteTemplate[D]}
}

func NodeObjectExecuteTemplate[D any](ctx cgcontext.Interface, obj *NodeOutput[D]) (string, error) {
	return ExecuteTemplate(ctx, obj)
}

func NodeObjectDataToString[D NodeWithStringOutput](ctx cgcontext.Interface, obj *NodeOutput[D]) (string, error) {
	return obj.data.ToString(ctx)
}

func MergedImports(nodes ...NodeWithImports) (*ImportSet, error) {
	set := NewImportSet()
	for i := range nodes {
		usedImports, err := nodes[i].UsedImports()
		if err != nil {
			return nil, err
		}
		result, err := set.MergeWith(usedImports)
		if err != nil {
			return nil, err
		}
		set = result
	}

	return set, nil
}

func GetTypeString(instance interface{}) string {
	if t := reflect.TypeOf(instance); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func ExecuteTemplate[D any](ctx cgcontext.Interface, obj NodeOutputInterface[D]) (string, error) {
	templates := cgtmp.NewSet(obj.Template()).AddTemplates(obj.Templates())
	data := obj.Data()

	tmp, err := template.New("").Funcs(map[string]any{
		// convenience method for
		"context": func() cgcontext.Interface {
			return ctx
		},
		"stringify": func(o NodeWithStringOutput) (string, error) {
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
