package cg_test

import (
	"github.com/onkeypress-llc/codegen/cg"
	"github.com/onkeypress-llc/codegen/cg/cgfs"
)

// verify different types meet the interface requirement
func isNode[T any](n cg.NodeInterface) bool {
	return true
}

func ctx() cg.GeneratorContextInterface {
	return cg.NewContextWithFS(cgfs.NewMapFS())
}
