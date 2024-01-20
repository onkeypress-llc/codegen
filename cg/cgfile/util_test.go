package cgfile_test

import (
	"embed"
	"fmt"
	"testing"

	"github.com/onkeypress-llc/codegen/cg/cgcontext"
	"github.com/onkeypress-llc/codegen/cg/cgfile"
	"github.com/onkeypress-llc/codegen/cg/cgfs"
	"github.com/onkeypress-llc/codegen/cg/cgtmp"
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

func ctx() cgcontext.Interface {
	return cgcontext.New(cgtmp.TemplateFS()).SetFS(cgfs.NewMapFS())
}

func TestFormatString(t *testing.T) {
	misformattedCode := `
for i := range values {
fmt.Printf("%d", i)
}
`
	result, err := cgfile.FormatGoString(misformattedCode)
	if err != nil {
		t.Error(err)
	}
	expected := `
for i := range values {
	fmt.Printf("%d", i)
}
`
	if result != expected {
		t.Errorf("Got \n%s\nexpected\n%s\n", result, expected)
	}
}
