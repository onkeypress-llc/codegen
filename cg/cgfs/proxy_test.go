package cgfs_test

import (
	"testing"

	"github.com/onkeypress-llc/codegen/cg/cgfs"
)

func TestProxyFS(t *testing.T) {
	filename, filevalue := "file.foo", "file value string"
	fs := cgfs.NewProxyFS()
	content, err := fs.Read("")
	if err != nil {
		t.Error(err)
	}
	if content != "" {
		t.Error("Expected empty string")
	}
	if exists := fs.Exists(filename); exists {
		t.Errorf("File %s not expected to exist", filename)
	}
	if err := fs.Write(filename, filevalue); err != nil {
		t.Error(err)
	}
}
