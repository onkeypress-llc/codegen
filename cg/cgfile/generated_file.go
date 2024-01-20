package cgfile

import "github.com/onkeypress-llc/codegen/cg/signing"

func NewFile(destination *Destination) *File {
	file := newFile(destination)
	file.signer = signing.NewGeneratedString()
	return file
}
