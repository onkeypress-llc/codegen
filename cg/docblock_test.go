package cg_test

import (
	"fmt"
	"testing"

	"github.com/onkeypress-llc/codegen/cg"
)

func TestDocBlockIsNode(t *testing.T) {
	if !isNode[cg.TemplateDocBlockOutputData](cg.NewDocBlock("")) {
		t.Errorf("Docblock not valid node")
	}
}

func TestDocBlockImportsEmpty(t *testing.T) {
	b := cg.NewDocBlock("")
	imports := b.UsedImports()
	if imports != nil {
		t.Errorf("Expected docblock to use no imports")
	}
}

func TestDocBlockEmptyOutput(t *testing.T) {
	result, err := cg.NodeToString(ctx(), cg.NewDocBlock(""))
	if err != nil {
		t.Error(err)
	}
	if result != "" {
		t.Errorf("Expected empty docblock to yield empty string, got %s", result)
	}
}
func TestDocBlockSingleLineOutput(t *testing.T) {
	value := "some string"
	expected := fmt.Sprintf(`/**
 * %s
 */`, value)
	result, err := cg.NodeToString(ctx(), cg.NewDocBlock(value))
	if err != nil {
		t.Error(err)
	}
	if result != expected {
		t.Errorf("Expected value \n%s\n\ngot\n\n%s\n", expected, result)
	}
}
func TestDocBlockMultiLineOutput(t *testing.T) {
	line0, line1 := "first line", "second line"
	result, err := cg.NodeToString(ctx(), cg.NewDocBlock(line0, line1))
	expected := fmt.Sprintf(`/**
 * %s
 * %s
 */`, line0, line1)
	if err != nil {
		t.Error(err)
	}
	if result != expected {
		t.Errorf("Expected value \n%s\n\ngot\n\n%s\n", expected, result)
	}
}
