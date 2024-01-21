package cgnode

import "github.com/onkeypress-llc/codegen/cg/cgi"

type StringNode struct {
	toString func(cgi.ContextInterface) (string, error)
}

func (n *StringNode) ToString(ctx cgi.ContextInterface) (string, error) {
	return n.toString(ctx)
}

func (n *StringNode) Generate(cgi.ContextInterface) (cgi.NodeOutputInterface, error) {
	return StringOutput(n), nil
}

func (n *StringNode) toInterface() cgi.NodeInterface {
	return n
}

func NewStringNode(toString func(cgi.ContextInterface) (string, error)) *StringNode {
	return &StringNode{toString: toString}
}
