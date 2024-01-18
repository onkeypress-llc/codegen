package cg_test

import (
	"embed"
	"fmt"
	"testing"

	"github.com/onkeypress-llc/codegen/cg"
	"github.com/onkeypress-llc/codegen/cg/cgcontext"
	"github.com/onkeypress-llc/codegen/cg/cgfs"
	"github.com/onkeypress-llc/codegen/cg/cgnode"
	"github.com/onkeypress-llc/codegen/cg/cgtmp"
)

//go:embed templates/*.tmp
var testFS embed.FS

// verify different types meet the interface requirement
func isNode[T any](n cgnode.NodeInterface[any]) bool {
	return true
}

func ctx() cgcontext.Interface {
	return cgcontext.New(testFS).SetFS(cgfs.NewMapFS())
}

type errorNode struct{}

func (e *errorNode) Generate(ctx cgcontext.Interface) (cgnode.NodeOutputInterface[any], error) {
	return nil, fmt.Errorf("Error!")
}

func (e *errorNode) UsedImports() (*cgnode.ImportSet, error) {
	return nil, nil
}

func TestNodeToStringHandlesError(t *testing.T) {
	if value, err := cg.NodeToString(ctx(), &errorNode{}); err == nil || value != "" {
		t.Errorf("Expected failure to return error and empty string")
	}
}

func TestFormatString(t *testing.T) {
	misformattedCode := `
for i := range values {
fmt.Printf("%d", i)
}
`
	result, err := cg.FormatGoString(misformattedCode)
	if err != nil {
		t.Error(err)
	}
	expected := `
for i := range values {
	fmt.Printf("%d", i)
}
`
	if result != expected {
		t.Errorf("Got \n%s\nexpected\n%s\n", result, expected)
	}
}

type templateValue struct {
	Value string
}

func newTemplateValue(value string) cgnode.NodeOutputInterface[*templateValue] {
	return cgnode.Output[*templateValue](cgtmp.New("static"), &templateValue{Value: value})
}

func TestTemplateExecution(t *testing.T) {
	context := ctx()
	value := "some value"
	output, err := cgnode.ExecuteTemplate[*templateValue](context, newTemplateValue(value))
	if err != nil {
		t.Error(err)
	}
	expected := fmt.Sprintf("Before %s After", value)
	if output != expected {
		t.Errorf("Expected %s got %s", expected, output)
	}
}
