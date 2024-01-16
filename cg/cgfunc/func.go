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
	Content      cgnode.NodeInterface
}

type FunctionNameType struct {
	Name string
	Type *cgnode.Type
}

func (f *TemplateFunction) Generate(ctx cgcontext.Interface) (cgnode.NodeOutputInterface, error) {
	return &TemplateFunctionOutput{}, nil
}

func (f *TemplateFunction) ToInterface() cgnode.NodeInterface {
	return f
}

type TemplateFunctionOutput struct{}

func (o *TemplateFunctionOutput) Name() string {
	return "TemplateFunctionOutput"
}

func (o *TemplateFunctionOutput) ToString() (string, error) {
	return "", nil
}

func (o *TemplateFunctionOutput) UsedImports() (*cgnode.ImportSet, error) {
	return cgnode.NewImportSet(), nil
}

func (o *TemplateFunctionOutput) Templates() *cgtmp.Templates {
	return nil
}
