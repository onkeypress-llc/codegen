package cgfile_test

import (
	"embed"
	"fmt"
	"testing"

	"github.com/onkeypress-llc/codegen/cg/cgfile"
)

//go:embed test_snapshots/partially_generated_file/*.snap
var partiallyGeneratedFileSnapshotFS embed.FS

func snapshot_pg(t *testing.T, fileName string) string {
	return snapshot(t, partiallyGeneratedFileSnapshotFS, fmt.Sprintf("partially_generated_file/%s", fileName))
}

func TestPartiallyGeneratedFile(t *testing.T) {

	f := cgfile.NewPartiallyGeneratedFile(testDestination).Package(testPackageName)
	result, err := cgfile.MaybeSignFile(ctx(), f)
	if err != nil {
		t.Error(err)
	}
	expected := snapshot_pg(t, "basic.snap")
	if result != expected {
		t.Errorf("Expected \n\n%s\n\ngot\n\n%s", expected, result)
	}
}
