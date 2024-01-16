package cgnode

import (
	"github.com/onkeypress-llc/codegen/cg/cgtmp"
)

type NodeSelfOutputInterface interface {
	ToString() (string, error)
}

type NodeSelfOutput[D NodeSelfOutputInterface] struct {
	name     string
	instance D
	imports  *ImportSet
}

func (o *NodeSelfOutput[D]) Name() string {
	return o.name
}

func (o *NodeSelfOutput[D]) ToString() (string, error) {
	return o.instance.ToString()
}

func (o *NodeSelfOutput[D]) Data() D {
	return o.instance
}

func (o *NodeSelfOutput[D]) Templates() *cgtmp.Templates {
	return nil
}

func (o *NodeSelfOutput[D]) UsedImports() (*ImportSet, error) {
	return o.imports, nil
}

func (o *NodeSelfOutput[D]) SetUsedImports(imports *ImportSet) *NodeSelfOutput[D] {
	o.imports = imports
	return o
}

func (o *NodeSelfOutput[D]) SetName(name string) *NodeSelfOutput[D] {
	o.name = name
	return o
}

func (o *NodeSelfOutput[D]) ToInterface() NodeOutputInterface {
	return o
}

func SelfOutput[D NodeSelfOutputInterface](instance D) *NodeSelfOutput[D] {
	return &NodeSelfOutput[D]{name: GetTypeString(instance), instance: instance}
}
