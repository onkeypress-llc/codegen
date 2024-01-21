package cgnode

import "github.com/onkeypress-llc/codegen/cg/cgcontext"

func NodeToString(context cgcontext.Interface, node NodeInterface) (string, error) {
	output, err := node.Generate(context)
	if err != nil {
		return "", err
	}
	return output.ToString(context)
}
