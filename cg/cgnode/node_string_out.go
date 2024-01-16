package cgnode

import (
	"github.com/onkeypress-llc/codegen/cg/cgtmp"
)

type NodeStringOutput[D any] struct {
	name        string
	data        *D
	toString    func(*D) (string, error)
	usedImports *ImportSet
}

func (o *NodeStringOutput[D]) Name() string {
	return o.name
}

func (o *NodeStringOutput[D]) ToString() (string, error) {
	return o.toString(o.data)
}

func (o *NodeStringOutput[D]) Data() *D {
	return o.data
}

func (o *NodeStringOutput[D]) Templates() *cgtmp.Templates {
	return nil
}

func (o *NodeStringOutput[D]) UsedImports() (*ImportSet, error) {
	return nil, nil
}

func (o *NodeStringOutput[D]) SetName(name string) *NodeStringOutput[D] {
	o.name = name
	return o
}

func (o *NodeStringOutput[D]) ToInterface() NodeOutputInterface {
	return o
}

func StringOutput[D any](data *D, toString func(*D) (string, error)) *NodeStringOutput[D] {
	return &NodeStringOutput[D]{name: GetTypeString(data), data: data, toString: toString}
}
