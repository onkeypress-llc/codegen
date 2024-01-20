package cgfile

import "github.com/onkeypress-llc/codegen/cg/cgcontext"

type Interface interface {
	Save(cgcontext.Interface) error
}
