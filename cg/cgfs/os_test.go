package cgfs_test

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"testing"

	"github.com/onkeypress-llc/codegen/cg/cgfs"
)

var testFSDirectory = "fstestdir"

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	cleanup()
	os.Exit(code)
}

func setup() {
	if _, err := os.Stat(testFSDirectory); !errors.Is(err, fs.ErrNotExist) {
		fmt.Printf("Detected %s, removing...", testFSDirectory)
		cleanup()
	}
	err := os.Mkdir(testFSDirectory, 0755)
	if err != nil {
		fmt.Printf("Unable to create directory %s, %s", testFSDirectory, err)
		os.Exit(1)
	}
}

func cleanup() {
	err := os.RemoveAll(testFSDirectory)
	if err != nil {
		fmt.Printf("Failed to remove directory %s, %s", testFSDirectory, err)
	}
}

func osFileName(name string) string {
	return path.Join(testFSDirectory, name)
}

func TestOSFS(t *testing.T) {
	fs := cgfs.NewOSFS()
	filename, filevalue := osFileName("file.foo"), "some value for file"
	if err := testReadCase(fs, readCase(filename, "", true)); err != nil {
		t.Error(err)
	}
	err := fs.Write(filename, filevalue)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(filename)
	if err := testExistsCase(fs, existsCase(filename, true)); err != nil {
		t.Error(err)
	}
	if err := testReadCase(fs, readCase(filename, filevalue, false)); err != nil {
		t.Error(err)
	}
}

func TestOSFSCreateError(t *testing.T) {
	fs := cgfs.NewOSFS()
	filename, filevalue := osFileName("dir"), "some value for file"
	if err := os.Mkdir(filename, 0755); err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(filename)
	if err := fs.Write(filename, filevalue); err == nil {
		t.Errorf("Expected write to directory %s to produce error", filename)
	}
}
