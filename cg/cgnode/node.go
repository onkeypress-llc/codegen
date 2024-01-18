package cgnode

import (
	"github.com/onkeypress-llc/codegen/cg/cgcontext"
)

type NodeInterface[D any] interface {
	Generate(cgcontext.Interface) (NodeOutputInterface[D], error)
}
