package cgfile

import "github.com/onkeypress-llc/codegen/cg/signing"

const PartiallyGeneratedBegin = "BEGIN MANUAL SECTION"
const PartiallyGeneratedEnd = "END MANUAL SECTION"

func NewPartiallyGeneratedFile(destination *Destination) *File {
	file := newFile(destination)
	file.signer = signing.NewPartiallyGeneratedString(PartiallyGeneratedBegin, PartiallyGeneratedEnd)
	return file
}
