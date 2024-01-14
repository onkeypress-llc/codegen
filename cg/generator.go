package cg

func NodeToString(context GeneratorContextInterface, node NodeInterface) (string, error) {
	output, err := node.Generate(context)
	if err != nil {
		return "", err
	}
	return output.ToString()
}
