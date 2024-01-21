package cgelement_test

import (
	"testing"

	"github.com/onkeypress-llc/codegen/cg/cgelement"
	"github.com/onkeypress-llc/codegen/cg/cgnode"
)

func isNode(i cgnode.NodeInterface) bool { return true }

func TestRawTextInterface(t *testing.T) {
	element := cgelement.NewRawText("")
	if !isNode(element) {

	}
}
