package cgnode

import (
	"github.com/onkeypress-llc/codegen/cg/cgcontext"
)

type NodeSet struct {
	nodes []NodeInterface[any]
}

func (n *NodeSet) Generate(ctx cgcontext.Interface) (NodeOutputInterface[any], error) {
	return nil, nil
}

func (s *NodeSet) ToInterface() NodeInterface[any] {
	return s
}

func NewSet(nodes []NodeInterface[any]) *NodeSet {
	return &NodeSet{nodes: nodes}
}
