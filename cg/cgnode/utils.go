package cgnode

import "github.com/onkeypress-llc/codegen/cg/cgi"

func NodeToString(context cgi.ContextInterface, node cgi.NodeInterface) (string, error) {
	output, err := node.Generate(context)
	if err != nil {
		return "", err
	}
	return output.ToString(context)
}
