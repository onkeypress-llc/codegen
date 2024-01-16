package cgfile_test

import (
	"path"
	"testing"

	"github.com/onkeypress-llc/codegen/cg"
	"github.com/onkeypress-llc/codegen/cg/cgcontext"
	"github.com/onkeypress-llc/codegen/cg/cgfile"
	"github.com/onkeypress-llc/codegen/cg/cgfs"
)

func TestDestinationWrite(t *testing.T) {
	fs := cgfs.NewMapFS()
	ctx := cgcontext.New(cg.TemplateFS()).SetFS(fs)
	filename, filepath, content := "file.foo", "some/path", "file content to be written"
	destination := cgfile.NewDestination(filename, filepath)

	if err := destination.Write(ctx, content); err != nil {
		t.Error(err)
	}
	fullFilePath := path.Join(filepath, filename)
	if exists := fs.Exists(fullFilePath); !exists {
		t.Errorf("Expected file %s to exist", fullFilePath)
	}
	value, err := fs.Read(fullFilePath)
	if err != nil {
		t.Error(err)
	}
	if value != content {
		t.Error(fileExpectation(content, value))
	}
}
