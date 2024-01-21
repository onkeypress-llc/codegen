package cgi

import (
	"regexp"

	"github.com/onkeypress-llc/codegen/cg/cge"
)

type SignedStringInterface interface {
	SigningToken() string
	Pattern() *regexp.Regexp
	TokenName() string
	SignString(string) (string, error)
	IsSigned(string) (bool, error)
	Verify(string) (bool, error)
	HasValidSignature(string) (bool, error)
	DocBlock(string) string
	SigningType() cge.SigningType
}
