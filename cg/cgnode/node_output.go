package cgnode

import (
	"bytes"
	"reflect"
	"text/template"

	"github.com/onkeypress-llc/codegen/cg/cgi"
	"github.com/onkeypress-llc/codegen/cg/cgtmp"
)

type NodeOutput[D any] struct {
	name      string
	data      D
	template  cgi.TemplateInterface
	templates cgi.TemplateSetInterface
	imports   cgi.ImportSetInterface

	toStringFunction func(cgi.ContextInterface, *NodeOutput[D]) (string, error)
}

func (o *NodeOutput[D]) Name() string {
	return o.name
}

func (o *NodeOutput[D]) SetName(name string) *NodeOutput[D] {
	o.name = name
	return o
}

func (o *NodeOutput[D]) Templates() cgi.TemplateSetInterface {
	return o.templates
}

func (o *NodeOutput[D]) SetTemplates(templates cgi.TemplateSetInterface) *NodeOutput[D] {
	o.templates = templates
	return o
}

func (o *NodeOutput[D]) UsedImports() (cgi.ImportSetInterface, error) {
	return o.imports, nil
}

func (o *NodeOutput[D]) SetUsedImports(imports *ImportSet) *NodeOutput[D] {
	o.imports = imports
	return o
}

func (o *NodeOutput[D]) Template() cgi.TemplateInterface {
	return o.template
}

func (o *NodeOutput[D]) ToString(ctx cgi.ContextInterface) (string, error) {
	return o.toStringFunction(ctx, o)
}

func (o *NodeOutput[D]) UntypedData() any {
	return o.data
}

func (o *NodeOutput[D]) Data() D {
	return o.data
}

func (o *NodeOutput[D]) ToInterface() cgi.NodeOutputInterface {
	return o
}

func (o *NodeOutput[D]) ToTypedInterface() cgi.NodeOutputWithTypedDataInterface[D] {
	return o
}

func Output[D any](template *cgtmp.Template, data D) *NodeOutput[D] {
	return &NodeOutput[D]{data: data, template: template, name: GetTypeString(data), templates: cgtmp.NewSet(), toStringFunction: NodeObjectExecuteTemplate[D]}
}

func NodeObjectExecuteTemplate[D any](ctx cgi.ContextInterface, obj *NodeOutput[D]) (string, error) {
	return ExecuteTemplate(ctx, obj)
}

func NodeObjectDataToString[D cgi.NodeWithStringOutput](ctx cgi.ContextInterface, obj *NodeOutput[D]) (string, error) {
	return obj.data.ToString(ctx)
}

func MergedImports(nodes ...cgi.NodeWithImports) (*ImportSet, error) {
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

func ExecuteTemplate(ctx cgi.ContextInterface, obj cgi.NodeOutputInterface) (string, error) {
	templates := cgtmp.NewSet(obj.Template()).AddTemplates(obj.Templates())
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
