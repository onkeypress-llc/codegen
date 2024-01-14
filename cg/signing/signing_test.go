package signing_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/onkeypress-llc/codegen/cg/signing"
)

func TestDefaultMatchPattern(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
	}()
	v := signing.DefaultMatchPattern()
	if v == nil {
		t.Error("Failed to create match pattern")
	}
}

func TestSignString(t *testing.T) {
	tokenName, docMessage := "token-name", "doc message"

	s := signing.New(tokenName, docMessage)
	value := fmt.Sprintf("Content %s", s.SigningToken())

	signed_value, err := s.SignString(value)
	if err != nil {
		t.Error(err)
	}
	if strings.Contains(signed_value, s.SigningToken()) {
		t.Errorf("Expected signed value [%s] to not contain token [%s]", signed_value, s.SigningToken())
	}
	if !strings.Contains(signed_value, fmt.Sprintf("@%s", tokenName)) {
		t.Errorf("Expected token name to be present [@%s]", tokenName)
	}
	if strings.Contains(signed_value, docMessage) {
		t.Errorf("Expected docMessage to not be present in [%s]", signed_value)
	}
}

func TestIsSigned(t *testing.T) {
	tokenName, docMessage := "token-name", "doc message"

	s := signing.New(tokenName, docMessage)
	if is_signed, err := s.IsSigned(""); err != nil || is_signed {
		t.Log(err)
		t.Error("Expected empty string to not register as signed.")
	}

	value := fmt.Sprintf("Content %s", s.SigningToken())
	if is_signed, err := s.IsSigned(value); err != nil || is_signed {
		t.Log(err)
		t.Errorf("Expected unsigned string to not register as signed. %s", value)
	}

	signed_value, err := s.SignString(value)
	if err != nil {
		t.Error(err)
	}
	is_signed, err := s.IsSigned(signed_value)
	if err != nil {
		t.Error(err)
	}
	if !is_signed {
		t.Errorf("Expected signed value to be detected")
	}
}

func TestVerify(t *testing.T) {
	tokenName, docMessage := "token-name", "doc message"

	s := signing.New(tokenName, docMessage)
	if valid, err := s.Verify(""); err == nil || valid {
		t.Log(err)
		t.Error("Expected empty string to not to verify")
	}
	token := s.SigningToken()
	value := fmt.Sprintf("Content %s\n\n%s\n\nmore content\n\n%s\n\nend", token, token, token)

	signed_value, err := s.SignString(value)
	if err != nil {
		t.Error(err)
	}
	if valid, err := s.Verify(signed_value); err != nil || !valid {
		t.Log(err)
		t.Errorf("Expected signed value to be valid: %s", signed_value)
	}

	second_signed := fmt.Sprintf("%s%s", signed_value, fakeSign())
	if valid, err := s.Verify(second_signed); err != nil || valid {
		t.Log(err)
		t.Errorf("Expected string with fake signature value to not be valid: %s", second_signed)
	}
}

func fakeSign() string {
	return signatureWrapper(fakeSignature())
}

func signatureWrapper(value string) string {
	return fmt.Sprintf("SignedSource<<%s>>", value)
}

// create 128 a's in a row to pattern natch an encoded sha512 checksum
func fakeSignature() string {
	return strings.Repeat("a", 128)
}
