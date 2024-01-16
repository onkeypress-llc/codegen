package cg

import (
	"github.com/onkeypress-llc/codegen/cg/cgcontext"
	"github.com/onkeypress-llc/codegen/cg/cgnode"
)

func NodeToString(context cgcontext.Interface, node cgnode.NodeInterface) (string, error) {
	output, err := node.Generate(context)
	if err != nil {
		return "", err
	}
	return output.ToString()
}
