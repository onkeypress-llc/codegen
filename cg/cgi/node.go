package cgi

type NodeInterface interface {
	Generate(ContextInterface) (NodeOutputInterface, error)
}

type NodeWithName interface {
	// the canonical name of the output data
	Name() string
}

type NodeWithImports interface {
	// set of all imports necessary
	UsedImports() (ImportSetInterface, error)
}

type NodeWithTemplate interface {
	// primary template to use
	Template() TemplateInterface
	// set of all templates used
	Templates() TemplateSetInterface
}

type NodeWithUntypedData interface {
	UntypedData() any
}

type NodeWithData[D any] interface {
	// the data to be used
	Data() D
}

type NodeWithStringOutput interface {
	ToString(ContextInterface) (string, error)
}

type NodeOutputInterface interface {
	NodeWithName
	NodeWithImports
	NodeWithTemplate
	NodeWithStringOutput
	NodeWithUntypedData
}

type NodeOutputWithTypedDataInterface[D any] interface {
	NodeOutputInterface
	NodeWithData[D]
}
