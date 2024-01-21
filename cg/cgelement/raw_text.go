package cgelement

import (
	"github.com/onkeypress-llc/codegen/cg/cgcontext"
	"github.com/onkeypress-llc/codegen/cg/cgnode"
)

type RawText struct {
	text string
}

func (r *RawText) ToString(ctx cgcontext.Interface) (string, error) {
	return r.text, nil
}

func (r *RawText) Generate(ctx cgcontext.Interface) (cgnode.NodeOutputInterface, error) {
	return cgnode.StringOutput(r), nil
}

func NewRawText(text string) *RawText {
	return &RawText{text: text}
}
