package cgnode

import (
	"github.com/onkeypress-llc/codegen/cg/cgcontext"
)

type NodeSet struct {
	nodes []NodeInterface
}

func (n *NodeSet) Generate(ctx cgcontext.Interface) (NodeOutputInterface, error) {
	return nil, nil
}

func (s *NodeSet) ToInterface() NodeInterface {
	return s
}

func NewSet(nodes []NodeInterface) *NodeSet {
	return &NodeSet{nodes: nodes}
}
