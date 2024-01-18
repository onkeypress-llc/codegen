package cgfunc

import (
	"github.com/onkeypress-llc/codegen/cg/cgcontext"
	"github.com/onkeypress-llc/codegen/cg/cgnode"
	"github.com/onkeypress-llc/codegen/cg/cgtmp"
)

type TemplateFunction struct {
	Name         string
	ClassMember  *FunctionNameType
	Arguments    []*FunctionNameType
	ReturnValues []*FunctionNameType
	Content      cgnode.NodeInterface[any]
}

type FunctionNameType struct {
	Name string
	Type *cgnode.Type
}

func (f *TemplateFunction) Generate(ctx cgcontext.Interface) (cgnode.NodeOutputInterface[*TemplateFunctionOutput], error) {
	return cgnode.StringOutput(&TemplateFunctionOutput{}), nil
}

func (f *TemplateFunction) ToInterface() cgnode.NodeInterface[*TemplateFunctionOutput] {
	return f
}

type TemplateFunctionOutput struct{}

func (o *TemplateFunctionOutput) Name() string {
	return "TemplateFunctionOutput"
}

func (o *TemplateFunctionOutput) ToString(ctx cgcontext.Interface) (string, error) {
	return "", nil
}

func (o *TemplateFunctionOutput) UsedImports() (*cgnode.ImportSet, error) {
	return cgnode.NewImportSet(), nil
}

func (o *TemplateFunctionOutput) Templates() *cgtmp.Templates {
	return nil
}
