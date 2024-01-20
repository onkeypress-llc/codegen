package cg

import (
	"fmt"

	"github.com/onkeypress-llc/codegen/cg/cgcontext"
	"github.com/onkeypress-llc/codegen/cg/cgfile"
	"github.com/onkeypress-llc/codegen/cg/cgtmp"
)

func NewContext() cgcontext.Interface {
	return cgcontext.New(cgtmp.TemplateFS())
}

func Generate(ctx cgcontext.Interface, nodes ...cgfile.Interface) error {
	errors := []error{}
	for i := range nodes {
		file := nodes[i]
		if err := file.Save(ctx); err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("Generation failed with the following errors: %v", errors)
	}
	return nil
}
