package cgnode

import (
	"github.com/onkeypress-llc/codegen/cg/cgcontext"
)

type NodeInterface interface {
	Generate(cgcontext.Interface) (NodeOutputInterface, error)
}

type Node struct {
}

func (n *Node) Generate(ctx cgcontext.Interface) (NodeOutputInterface, error) {
	return nil, nil
}

func (n *Node) ToInterface() NodeInterface {
	return n
}

func New() *Node {
	return &Node{}
}
