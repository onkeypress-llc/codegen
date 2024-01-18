package cg

import (
	"github.com/onkeypress-llc/codegen/cg/cgcontext"
	"github.com/onkeypress-llc/codegen/cg/cgnode"
)

func NodeToString[D any](context cgcontext.Interface, node cgnode.NodeInterface[D]) (string, error) {
	output, err := node.Generate(context)
	if err != nil {
		return "", err
	}
	return output.ToString(context)
}
