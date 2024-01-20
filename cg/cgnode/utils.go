package cgnode

import "github.com/onkeypress-llc/codegen/cg/cgcontext"

func NodeToString[D any](context cgcontext.Interface, node NodeInterface[D]) (string, error) {
	output, err := node.Generate(context)
	if err != nil {
		return "", err
	}
	return output.ToString(context)
}
