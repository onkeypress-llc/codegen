package main

import (
	"fmt"

	"github.com/onkeypress-llc/codegen/cg"
	"github.com/onkeypress-llc/codegen/cg/cgfile"
)

func main() {
	demoDirectory := "./demo"
	err := cg.Generate(
		cg.NewContext(),
		cgfile.NewFile(cgfile.NewDestination("fully-generated.gen.go", demoDirectory)),
		cgfile.NewPartiallyGeneratedFile(cgfile.NewDestination("partially-generated.gen.go", demoDirectory)),
		cgfile.NewFileWithoutGeneratorHeadersOrSigning(cgfile.NewDestination("vanilla.gen.go", demoDirectory)),
	)
	if err != nil {
		fmt.Print(err)
	}
}
