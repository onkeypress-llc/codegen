package cg

import (
	"fmt"
	"strings"
)

type TemplateDocblock struct {
	comment string
}

type TemplateDocBlockOutputData struct {
	lines []string
}

func NewDocBlock(comments ...string) *TemplateDocblock {
	return &TemplateDocblock{comment: strings.Join(comments, "\n")}
}

func (b *TemplateDocblock) UsedImports() *TemplateImportSet {
	return nil
}

func (b *TemplateDocblock) Generate(c GeneratorContextInterface) (NodeOutputInterface, error) {
	return StringOutput(&TemplateDocBlockOutputData{lines: docblockCreateLines(b.comment)}, docblockDataToString), nil
}

// if there is a comment, split it by lines, else return empty array
func docblockCreateLines(comment string) []string {
	if comment == "" {
		return []string{}
	}
	return strings.Split(comment, "\n")
}

const docblockLinePrefixBase = "\n *"

// if there are no lines, return empty string, else format docblock around lines
func docblockDataToString(b *TemplateDocBlockOutputData) (string, error) {
	if len(b.lines) < 1 {
		return "", nil
	}
	docblockLinePrefix := fmt.Sprintf("%s ", docblockLinePrefixBase)
	return fmt.Sprintf("/**%s%s%s/", docblockLinePrefix, strings.Join(b.lines, docblockLinePrefix), docblockLinePrefixBase), nil
}
