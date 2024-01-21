package cgelement

import (
	"github.com/onkeypress-llc/codegen/cg/cgi"
	"github.com/onkeypress-llc/codegen/cg/cgnode"
)

type StructProperty struct {
	Name string
	// todo: generics
	Type cgi.TypeInterface
}

func (p *StructProperty) asInterface() cgi.NodeInterface {
	return p
}

func (p *StructProperty) Generate(ctx cgi.ContextInterface) (cgi.NodeOutputInterface, error) {
	return cgnode.StringOutput(p), nil
}

func (p *StructProperty) ToString(ctx cgi.ContextInterface) (string, error) {

	return "", nil
}
