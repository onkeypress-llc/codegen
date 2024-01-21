package cgfs_test

import (
	"fmt"
	"testing"

	"github.com/onkeypress-llc/codegen/cg/cgfs"
)

func TestMapFSCreation(t *testing.T) {
	fs := cgfs.NewMapFS()
	if !isFSInterface(fs) {
		t.Error("Expected cgi.FSInterface")
	}
}

func TestMapFSOperations(t *testing.T) {
	filename := "file.foo"
	filevalue := "file value"
	file0 := filename
	file1 := fmt.Sprintf("some/file/path/%s", filename)
	fs := cgfs.NewMapFS()
	if err := testExistsCases(fs, existsCase(file0, false), existsCase(file1, false)); err != nil {
		t.Error(err)
	}
	err := fs.Write(file0, filevalue)
	if err != nil {
		t.Error(err)
	}
	if err := testExistsCases(fs, existsCase(file0, true), existsCase(file1, false)); err != nil {
		t.Error(err)
	}
	if err := testReadCases(fs, readCase(file0, filevalue, false), readCase(file1, "", true)); err != nil {
		t.Error(err)
	}
	fs = fs.Remove(file0)
	if err := testExistsCases(fs, existsCase(file0, false), existsCase(file1, false)); err != nil {
		t.Error(err)
	}
	if err := testReadCases(fs, readCase(file0, "", true), readCase(file1, "", true)); err != nil {
		t.Error(err)
	}
}
func TestMapFSSet(t *testing.T) {
	filename, filevalue := "file.foo", "filevalue string"
	fs := cgfs.NewMapFS().Set(cgfs.NewMapFile(filename, filevalue))
	if err := testExistsCases(fs, existsCase(filename, true), existsCase("secondfile.foo", false)); err != nil {
		t.Error(err)
	}
	if err := testReadCases(fs, readCase(filename, filevalue, false), readCase("secondfile.foo", "", true)); err != nil {
		t.Error(err)
	}
}
