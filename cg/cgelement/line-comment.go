package cgelement

import (
	"fmt"

	"github.com/onkeypress-llc/codegen/cg/cgcontext"
	"github.com/onkeypress-llc/codegen/cg/cgnode"
)

type LineComment struct {
	comment           string
	noSpaceAfterSlash bool
}

const DefaultCommentPrefix = " "

func (c *LineComment) ToString(ctx cgcontext.Interface) (string, error) {
	separator := " "
	if c.noSpaceAfterSlash {
		separator = ""
	}
	return fmt.Sprintf("//%s%s\n", separator, c.comment), nil
}

func (c *LineComment) NoSpaceAfterSlash() *LineComment {
	return c.SetSpaceAfterSlash(false)
}

func (c *LineComment) SetSpaceAfterSlash(value bool) *LineComment {
	c.noSpaceAfterSlash = !value
	return c
}

func (c *LineComment) Generate(ctx cgcontext.Interface) (cgnode.NodeOutputInterface, error) {
	return cgnode.StringOutput(c), nil
}

func NewLineComment(comment string, args ...interface{}) *LineComment {
	return &LineComment{comment: fmt.Sprintf(comment, args...)}
}
