package cgelement

import (
	"github.com/onkeypress-llc/codegen/cg/cgi"
	"github.com/onkeypress-llc/codegen/cg/cgnode"
)

type RawText struct {
	text string
}

func (r *RawText) ToString(ctx cgi.ContextInterface) (string, error) {
	return r.text, nil
}

func (r *RawText) Generate(ctx cgi.ContextInterface) (cgi.NodeOutputInterface, error) {
	return cgnode.StringOutput(r), nil
}

func NewRawText(text string) *RawText {
	return &RawText{text: text}
}
