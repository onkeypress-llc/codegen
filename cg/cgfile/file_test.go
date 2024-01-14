package cgfile_test

import (
	"embed"
	"fmt"
	"testing"

	"github.com/onkeypress-llc/codegen/cg"
	"github.com/onkeypress-llc/codegen/cg/cgfile"
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
	if !isNode[cgfile.TemplateFileOutputData](cgfile.NewFile(testDestination)) {
		t.Errorf("File not valid node")
	}
	if !isNode[cgfile.TemplateFileOutputData](cgfile.NewGeneratedFile(testDestination)) {
		t.Errorf("Generated file not valid node")
	}
	if !isNode[cgfile.TemplateFileOutputData](cgfile.NewPartiallyGeneratedFile(testDestination)) {
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

// verify different types meet the interface requirement
func isNode[T any](n cg.NodeInterface) bool {
	return true
}

// func TestFileOutputCopiesPackageName(t *testing.T) {}
// func TestFileOutputComputeTemplates(t *testing.T)  {}
// func TestFileOutputComputeImports(t *testing.T)    {}
