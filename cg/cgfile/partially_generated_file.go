package cgfile

import "github.com/onkeypress-llc/codegen/cg/signing"

const PartiallyGeneratedBegin = "BEGIN MANUAL SECTION"
const PartiallyGeneratedEnd = "END MANUAL SECTION"

func NewPartiallyGeneratedFile(destination *Destination) *TemplateFile {
	file := NewFile(destination)
	file.signer = signing.NewPartiallyGeneratedString(PartiallyGeneratedBegin, PartiallyGeneratedEnd)
	return file
}
