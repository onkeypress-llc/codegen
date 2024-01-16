package cgnode

import (
	"reflect"

	"github.com/onkeypress-llc/codegen/cg/cgtmp"
)

type NodeOutputInterface interface {
	Name() string
	ToString() (string, error)
	Templates() *cgtmp.Templates
	UsedImports() (*ImportSet, error)
}

type NodeOutput struct {
}

func (o *NodeOutput) Name() string {
	return GetTypeString(o)
}

func (o *NodeOutput) ToString() (string, error) {
	return "", nil
}

func (o *NodeOutput) Templates() *cgtmp.Templates {
	return cgtmp.NewSet()
}

func (o *NodeOutput) UsedImports() (*ImportSet, error) {
	return nil, nil
}

func (o *NodeOutput) ToInterface() NodeOutputInterface {
	return o
}

func MergedImports(nodes ...NodeOutputInterface) (*ImportSet, error) {
	set := NewImportSet()
	for i := range nodes {
		usedImports, err := nodes[i].UsedImports()
		if err != nil {
			return nil, err
		}
		result, err := set.MergeWith(usedImports)
		if err != nil {
			return nil, err
		}
		set = result
	}

	return set, nil
}

func GetTypeString(instance interface{}) string {
	if t := reflect.TypeOf(instance); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
