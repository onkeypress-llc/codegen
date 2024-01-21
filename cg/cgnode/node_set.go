package cgnode

import "github.com/onkeypress-llc/codegen/cg/cgi"

type NodeSet struct {
	nodes []cgi.NodeInterface
}

func (n *NodeSet) Generate(ctx cgi.ContextInterface) (cgi.NodeOutputInterface, error) {
	return nil, nil
}

func (s *NodeSet) ToInterface() cgi.NodeInterface {
	return s
}

func NewSet(nodes []cgi.NodeInterface) *NodeSet {
	return &NodeSet{nodes: nodes}
}
