package signing

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"

	"github.com/onkeypress-llc/codegen/cg/cge"
	"github.com/onkeypress-llc/codegen/cg/cgi"
)

var defaultToken = "<<SignedSource::%1*zOeVoAZle#+L!plEphiEmie@I>>"

var partiallyGeneratedTokenName = "partially-generated"
var generatedTokenName = "generated"

var generatedDocMessage = "This file is generated. Do not modify it manually!"

type SignedString struct {
	token        string
	preprocess   func(string) string
	tokenName    string
	matchPattern *regexp.Regexp
	docMessage   string
	signingType  cge.SigningType
}

type MatchResult struct {
	tokenName string
	signature string
}

func DefaultMatchPattern() *regexp.Regexp {
	return regexp.MustCompile("@(?P<token_name>\\S+) (?:SignedSource<<(?P<signature>[a-f0-9]{128})>>)")
}

func New(tokenName, docMessage string) *SignedString {
	return &SignedString{preprocess: preprocess, tokenName: tokenName, matchPattern: DefaultMatchPattern(), token: defaultToken, docMessage: docMessage, signingType: cge.Full}
}

func NewGeneratedString() cgi.SignedStringInterface {
	return New(generatedTokenName, generatedDocMessage)
}

func NewPartiallyGeneratedString(begin, end string) cgi.SignedStringInterface {
	return New(partiallyGeneratedTokenName, partiallyGeneratedDocMessage(begin, end)).setSigningType(cge.Partial)
}

func (s *SignedString) SigningToken() string {
	return fmt.Sprintf("@%s %s", s.tokenName, s.token)
}

func (s *SignedString) Pattern() *regexp.Regexp {
	return s.matchPattern
}

func (s *SignedString) TokenName() string {
	return s.tokenName
}

func sign(value string) string {
	hash := sha512.New()
	hash.Write([]byte(value))
	return hex.EncodeToString(hash.Sum(nil))
}

func (s *SignedString) SignString(value string) (string, error) {
	signature := sign(s.preprocess(value))

	output := strings.ReplaceAll(value, s.token, replacementToken(signature))
	if output == value {
		return "", fmt.Errorf("To sign a file you must embed a signing token within its contents\ne.g.\n%s\nshould appear in\n%s\n\n", s.token, value)
	}

	return output, nil
}

func (s *SignedString) IsSigned(value string) (bool, error) {
	matches, err := namedMatches(s.matchPattern, value)
	if err != nil {
		return false, err
	}
	// if any matches have the token name return true
	for i := range matches {
		if matches[i].tokenName == s.tokenName {
			return true, nil
		}
	}
	return false, nil
}

func (s *SignedString) Verify(maybe_signed string) (bool, error) {
	matches, err := namedMatches(s.matchPattern, maybe_signed)
	if err != nil {
		return false, err
	}
	var candidate *MatchResult
	for i := range matches {
		match := matches[i]
		if match.tokenName != s.tokenName {
			continue
		}
		if candidate == nil {
			candidate = match
		} else if candidate.signature != match.signature {
			return false, fmt.Errorf("Detected multiple signatures for token %s", s.tokenName)
		}
	}
	if candidate == nil {
		return false, fmt.Errorf("No signatures detected for token %s", s.tokenName)
	}
	// there is at least one signature for this token name
	value := strings.ReplaceAll(maybe_signed, replacementToken(candidate.signature), s.token)
	value = s.preprocess(value)

	return sign(value) == candidate.signature, nil
}

func (s *SignedString) HasValidSignature(value string) (bool, error) {
	if signed, err := s.IsSigned(value); signed || err != nil {
		return false, err
	}
	verified, err := s.Verify(value)
	if err != nil {
		return false, err
	}
	return verified, nil
}

func (s *SignedString) DocBlock(comment string) string {
	sections := []string{s.docMessage}
	if comment != "" {
		sections = append(sections, comment)
	}
	return strings.Join(append(sections, s.SigningToken()), "\n\n")
}

func (s *SignedString) SigningType() cge.SigningType {
	return s.signingType
}

func (s *SignedString) setSigningType(v cge.SigningType) *SignedString {
	s.signingType = v
	return s
}

func replacementToken(value string) string {
	return fmt.Sprintf("SignedSource<<%s>>", value)
}

func preprocess(s string) string {
	return s
}

func namedMatches(re *regexp.Regexp, value string) ([]*MatchResult, error) {
	matches := re.FindAllStringSubmatch(value, -1)
	results := make([]*MatchResult, len(matches))
	// find indexes
	names := re.SubexpNames()
	tokenIndex, signatureIndex := -1, -1
	for i := range names {
		if names[i] == "token_name" {
			tokenIndex = i
		} else if names[i] == "signature" {
			signatureIndex = i
		}
	}
	if tokenIndex < 0 || signatureIndex < 0 {
		return nil, fmt.Errorf("Unable to locate index for signature or token name in expression")
	}
	for i := range matches {
		match := matches[i]
		results[i] = &MatchResult{
			tokenName: match[tokenIndex],
			signature: match[signatureIndex],
		}
	}
	return results, nil
}

func partiallyGeneratedDocMessage(begin, end string) string {
	return fmt.Sprintf("This file is partially generated. Only make modifications between %s and %s designators.", begin, end)
}
