package cgnode

import (
	"reflect"

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

func (o *NodeOutput[D]) SetUsedImports(imports cgi.ImportSetInterface) *NodeOutput[D] {
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
	return cgtmp.ExecuteTemplate(ctx, obj)
}

func NodeObjectDataToString[D cgi.NodeWithStringOutput](ctx cgi.ContextInterface, obj *NodeOutput[D]) (string, error) {
	return obj.data.ToString(ctx)
}

func GetTypeString(instance interface{}) string {
	if t := reflect.TypeOf(instance); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
