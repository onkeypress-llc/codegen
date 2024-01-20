package cgfile_test

import (
	"embed"
	"fmt"
	"testing"

	"github.com/onkeypress-llc/codegen/cg/cgfile"
)

//go:embed test_snapshots/generated_file/*.snap
var generatedFileSnapshotFS embed.FS

func snapshot_g(t *testing.T, fileName string) string {
	return snapshot(t, generatedFileSnapshotFS, fmt.Sprintf("generated_file/%s", fileName))
}

func TestGeneratedFile(t *testing.T) {
	f := cgfile.NewFile(testDestination).Package(testPackageName)
	result, err := cgfile.MaybeSignFile(ctx(), f)
	if err != nil {
		t.Error(err)
	}
	expected := snapshot_g(t, "basic.snap")
	if result != expected {
		t.Errorf("Expected \n\n%s\n\ngot\n\n%s", expected, result)
	}
}
