package cgnode

import (
	"github.com/onkeypress-llc/codegen/cg/cgcontext"
)

type NodeInterface interface {
	Generate(cgcontext.Interface) (NodeOutputInterface, error)
}
