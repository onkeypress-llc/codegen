package cgelement_test

import (
	"fmt"
	"testing"

	"github.com/onkeypress-llc/codegen/cg/cgelement"
	"github.com/onkeypress-llc/codegen/cg/cgi"
)

var simpleSource string = "fmt"
var simpleSourceDefaultLabel string = "fmt"
var pathedSource string = "github.com/onkeypress-llc/codegen/cg"
var pathedSourceDefaultLabel string = "cg"

type ImportFieldsTestData struct {
	input                cgi.ImportInterface
	expectedLabel        string
	expectedSource       string
	expectedDefaultLabel string
}

func TestImportFields(t *testing.T) {
	testCases := []*ImportFieldsTestData{
		{
			input:                cgelement.NewImport(simpleSource),
			expectedLabel:        simpleSourceDefaultLabel,
			expectedSource:       simpleSource,
			expectedDefaultLabel: simpleSourceDefaultLabel,
		},
		{
			input:                cgelement.NewImport(pathedSource),
			expectedLabel:        pathedSourceDefaultLabel,
			expectedSource:       pathedSource,
			expectedDefaultLabel: pathedSourceDefaultLabel,
		},
		{
			input:                cgelement.NewImport(simpleSource).SetLabel("foo"),
			expectedLabel:        "foo",
			expectedSource:       simpleSource,
			expectedDefaultLabel: simpleSourceDefaultLabel,
		},
		{
			input:                cgelement.NewImport(pathedSource).SetLabel("foo"),
			expectedLabel:        "foo",
			expectedSource:       pathedSource,
			expectedDefaultLabel: pathedSourceDefaultLabel,
		},
	}
	for i := range testCases {
		if err := checkImportFields(testCases[i]); err != nil {
			t.Error(err)
		}
	}
}

func checkImportFields(data *ImportFieldsTestData) error {
	imp := data.input
	if imp.Label() != data.expectedLabel {
		return fmt.Errorf("expected label to be %s, got %s", data.expectedLabel, imp.Label())
	}
	if imp.Source() != data.expectedSource {
		return fmt.Errorf("expected source to be %s, got %s", data.expectedSource, imp.Source())
	}
	if imp.DefaultLabel() != data.expectedDefaultLabel {
		return fmt.Errorf("expected default label to be %s, got %s", data.expectedDefaultLabel, imp.DefaultLabel())
	}
	return nil
}

type ImportEquivalenceTestData struct {
	compareValue cgi.ImportInterface
	withValue    cgi.ImportInterface
	expected     bool
}

func TestImportEquals(t *testing.T) {
	label0 := "foo"
	label1 := "bar"
	unlabeledTestingImport := cgelement.NewImport(simpleSource)
	labeledTestingImport0 := cgelement.NewImport(simpleSource).SetLabel(label0)
	labeledPathedImport0 := cgelement.NewImport(pathedSource).SetLabel(label0)
	labeledPathedImport1 := cgelement.NewImport(pathedSource).SetLabel(label1)
	testCases := []*ImportEquivalenceTestData{
		// self equivalence
		{
			compareValue: unlabeledTestingImport,
			withValue:    unlabeledTestingImport,
			expected:     true,
		},
		// same source, different label not equivalent
		{
			compareValue: unlabeledTestingImport,
			withValue:    labeledTestingImport0,
			expected:     false,
		},
		// different label and source not equivalent
		{
			compareValue: labeledTestingImport0,
			withValue:    labeledPathedImport0,
			expected:     false,
		},
		// same label, different source not equivalent
		{
			compareValue: labeledTestingImport0,
			withValue:    labeledPathedImport1,
			expected:     false,
		},
	}
	for i := range testCases {
		data := testCases[i]
		if data.compareValue.Equals(data.withValue) != data.expected {
			t.Errorf("Expected imports %v and %v to have equivalence result: %t", data.compareValue, data.withValue, data.expected)
		}
	}
}

type ImportStringTestData struct {
	value    cgi.ImportInterface
	expected string
}

func TestImportString(t *testing.T) {
	label0 := "foo"
	testCases := []*ImportStringTestData{
		{
			value:    cgelement.NewImport(simpleSource),
			expected: fmt.Sprintf("\"%s\"", simpleSource),
		},
		{
			value:    cgelement.NewImport(simpleSource).SetLabel(label0),
			expected: fmt.Sprintf("%s \"%s\"", label0, simpleSource),
		},
		{
			value:    cgelement.NewImport(pathedSource),
			expected: fmt.Sprintf("\"%s\"", pathedSource),
		},
		{
			value:    cgelement.NewImport(pathedSource).SetLabel(label0),
			expected: fmt.Sprintf("%s \"%s\"", label0, pathedSource),
		},
	}
	for i := range testCases {
		data := testCases[i]
		if data.value.String() != data.expected {
			t.Errorf("Expected import %v to output string %s, got %s", data.value, data.expected, data.value.String())
		}
	}
}
