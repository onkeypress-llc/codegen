package cgfile

import "github.com/onkeypress-llc/codegen/cg/signing"

func NewGeneratedFile(destination *Destination) *TemplateFile {
	file := NewFile(destination)
	file.signer = signing.NewGeneratedString()
	file.createHeader = func() string {
		return file.signer.DocBlock(file.HeaderString())
	}
	return file
}
