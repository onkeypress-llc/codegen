package cgelement

import (
	"fmt"

	"github.com/onkeypress-llc/codegen/cg/cgi"
)

type Type struct {
	name string
	imp  cgi.ImportInterface
}

func (t *Type) ToString(ctx cgi.ContextInterface) (string, error) {
	// the context will know whether the import has been aliased
	ns, err := ctx.GetImportNamespace(t.imp)
	if err != nil {
		return "", err
	}
	if ns != "" {
		return fmt.Sprintf("%s.%s", ns, t.name), nil
	}
	// same namespace as current package, refer by name only
	return t.name, nil
}

func (t *Type) Import(imp cgi.ImportInterface) *Type {
	t.imp = imp
	return t
}

func (t *Type) Name(name string) *Type {
	t.name = name
	return t
}

func NewType(name string) *Type {
	return &Type{name: name, imp: nil}
}
