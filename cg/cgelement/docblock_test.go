package cgelement_test

import (
	"testing"

	"github.com/onkeypress-llc/codegen/cg/cgelement"
	"github.com/onkeypress-llc/codegen/cg/cgnode"
)

func TestDocBlockEmptyOutput(t *testing.T) {
	result, err := cgnode.NodeToString(ctx(), cgelement.NewDocBlock(""))
	if err != nil {
		t.Error(err)
	}
	if result != "" {
		t.Errorf("Expected empty docblock to yield empty string, got %s", result)
	}
}

// func TestDocBlockSingleLineOutput(t *testing.T) {
// 	value := "some string"
// 	expected := fmt.Sprintf(`/**
//  * %s
//  */`, value)
// 	result, err := cgnode.NodeToString(ctx(), cg.NewDocBlock(value))
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if result != expected {
// 		t.Errorf("Expected value \n%s\n\ngot\n\n%s\n", expected, result)
// 	}
// }
// func TestDocBlockMultiLineOutput(t *testing.T) {
// 	line0, line1 := "first line", "second line"
// 	result, err := cgnode.NodeToString(ctx(), cg.NewDocBlock(line0, line1))
// 	expected := fmt.Sprintf(`/**
//  * %s
//  * %s
//  */`, line0, line1)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if result != expected {
// 		t.Errorf("Expected value \n%s\n\ngot\n\n%s\n", expected, result)
// 	}
// }
