package cgfile_test

import (
	"embed"
	"fmt"
	"testing"

	"github.com/onkeypress-llc/codegen/cg"
	"github.com/onkeypress-llc/codegen/cg/cgfs"
)

func snapshot(t *testing.T, fs embed.FS, name string) string {
	data, err := fs.ReadFile(fmt.Sprintf("test_snapshots/%s", name))
	if err != nil {
		t.Error(err)
	}
	return string(data)
}

func fileExpectation(expected, actual string) string {
	return fmt.Sprintf("Expected %s got %s", filePrint(expected), filePrint(actual))
}

func filePrint(content string) string {
	return fmt.Sprintf("\n---\n%s\n---\n", content)
}

func ctx() cg.GeneratorContextInterface {
	return cg.NewContextWithFS(cgfs.NewMapFS())
}
