package cgfile_test

import (
	"embed"
	"fmt"
	"strings"
	"testing"

	"github.com/onkeypress-llc/codegen/cg"
	"github.com/onkeypress-llc/codegen/cg/cgfile"
	"github.com/onkeypress-llc/codegen/cg/cgnode"
)

var testFileName = "foo.go"
var testFilePath = "sample/path"

var testDestination = cgfile.NewDestination(testFileName, testFilePath)
var testPackageName = "foopackage"

//go:embed test_snapshots/file/*.snap
var fileSnapshotFS embed.FS

func snapshot_b(t *testing.T, fileName string) string {
	return snapshot(t, fileSnapshotFS, fmt.Sprintf("file/%s", fileName))
}

func TestFileIsNode(t *testing.T) {
	if !isNode[cgfile.Data](cgfile.NewFile(testDestination)) {
		t.Errorf("File not valid node")
	}
	if !isNode[cgfile.Data](cgfile.NewGeneratedFile(testDestination)) {
		t.Errorf("Generated file not valid node")
	}
	if !isNode[cgfile.Data](cgfile.NewPartiallyGeneratedFile(testDestination)) {
		t.Errorf("Partially generated file not valid node")
	}
}

func TestFileImplementsNodeInterface(t *testing.T) {
	file := cgfile.NewFile(testDestination).Package(testPackageName)
	output, err := cg.NodeToString(ctx(), file)
	if err != nil {
		t.Error(err)
	}
	expected := snapshot_b(t, "basic.snap")
	if output != expected {
		t.Error(fileExpectation(expected, output))
	}
}

func TestFileHeader(t *testing.T) {
	headerText := "Here is some header text to inject into the template"
	file := cgfile.NewFile(testDestination).Package(testPackageName).SetHeader(headerText)
	output, err := cg.NodeToString(ctx(), file)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(output, headerText) {
		t.Errorf("Expected output to contain header text: %s. File content\n\n%s", headerText, output)
	}
}

// verify different types meet the interface requirement
func isNode[T any](n cgnode.NodeInterface[*cgfile.Data]) bool {
	return true
}

// func TestFileOutputCopiesPackageName(t *testing.T) {}
// func TestFileOutputComputeTemplates(t *testing.T)  {}
// func TestFileOutputComputeImports(t *testing.T)    {}
