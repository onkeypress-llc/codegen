package cg

import (
	"fmt"

	"github.com/onkeypress-llc/codegen/cg/cgcontext"
	"github.com/onkeypress-llc/codegen/cg/cgi"
	"github.com/onkeypress-llc/codegen/cg/cgtmp"
)

func NewContext() cgi.ContextInterface {
	return cgcontext.New(cgtmp.TemplateFS())
}

func Generate(ctx cgi.ContextInterface, nodes ...cgi.FileInterface) error {
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
