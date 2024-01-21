package cgelement

import (
	"fmt"
	"strings"

	"github.com/onkeypress-llc/codegen/cg/cgi"
	"github.com/onkeypress-llc/codegen/cg/cgnode"
)

const docblockLinePrefixBase = "\n *"

type Docblock struct {
	comment string
}

func (b *Docblock) ToString(ctx cgi.ContextInterface) (string, error) {
	lines := docblockCreateLines(b.comment)
	if len(lines) < 1 {
		return "", nil
	}
	docblockLinePrefix := fmt.Sprintf("%s ", docblockLinePrefixBase)
	return fmt.Sprintf("/**%s%s%s/", docblockLinePrefix, strings.Join(lines, docblockLinePrefix), docblockLinePrefixBase), nil

}

func (b *Docblock) ToInterface() cgi.NodeInterface {
	return b
}

func NewDocBlock(comments ...string) *Docblock {
	return &Docblock{comment: strings.Join(comments, "\n")}
}

func (b *Docblock) Generate(c cgi.ContextInterface) (cgi.NodeOutputInterface, error) {
	return cgnode.StringOutput(b), nil
}

// if there is a comment, split it by lines, else return empty array
func docblockCreateLines(comment string) []string {
	if comment == "" {
		return []string{}
	}
	return strings.Split(comment, "\n")
}
