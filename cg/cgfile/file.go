package cgfile

import (
	"fmt"

	"github.com/onkeypress-llc/codegen/cg"
	"github.com/onkeypress-llc/codegen/cg/signing"
)

var goFileTemplates = []string{"go-file", "go-file-contents", "go-file-content"}

type TemplateFile struct {
	// header comment for file
	headerString string
	// where the file should be written on the OS
	destination *Destination
	// package at the top of file
	packageName string
	// manually added imports to file
	imports *cg.TemplateImportSet
	// set of content in the file
	contents []cg.TemplateFileContent

	allowWriteIfFormatFails bool

	signer       signing.SignedStringInterface
	createHeader func() string
}

type TemplateFileOutputData struct {
	// header comment
	Header cg.NodeOutputInterface
	// package of this file
	PackageName string
	// union of imports used in the file
	Imports *cg.TemplateImportSet
	// set of content in the file
	Contents []cg.TemplateFileContentOutput
}

func NewFile(destination *Destination) *TemplateFile {
	return &TemplateFile{imports: cg.NewImportSet(), destination: destination, contents: []cg.TemplateFileContent{}, createHeader: createHeader}
}

func (f *TemplateFile) UsedImports() *cg.TemplateImportSet {
	return f.imports
}

func (f *TemplateFile) Generate(c cg.GeneratorContextInterface) (cg.NodeOutputInterface, error) {
	imports, err := cg.NewImportSet().MergeValuesInFrom(f.imports)
	if err != nil {
		return nil, err
	}
	contents := make([]cg.TemplateFileContentOutput, len(f.contents))
	templates := cg.NewTemplates()
	// from file content blocks, aggragate imports, template names, and content

	// add imports from file content blocks
	for i := range f.contents {
		content := f.contents[i]
		imports, err = imports.MergeValuesInFrom(content.Imports())

		templates.AddTemplates(content.Templates())
		// TODO: capture template names used by this content block
		contentOutput, err := content.Generate()
		if err != nil {
			return nil, err
		}
		contents[i] = contentOutput
	}
	fileHeader := cg.NewDocBlock(f.createHeader())
	headerOutput, err := fileHeader.Generate(c)
	if err != nil {
		return nil, err
	}

	templateNames := append(goFileTemplates, templates.Names()...)

	return cg.TemplateOutput(
		&TemplateFileOutputData{
			Header:      headerOutput,
			PackageName: f.packageName,
			Imports:     imports,
			Contents:    contents,
		},
		templateNames...,
	), nil
}

func (f *TemplateFile) Save(c cg.GeneratorContextInterface) (string, error) {
	panic("Not implemented")
}

func (f *TemplateFile) Package(packageName string) *TemplateFile {
	f.packageName = packageName
	return f
}

func (f *TemplateFile) Contents(contents []cg.TemplateFileContent) *TemplateFile {
	f.contents = contents
	return f
}

func (f *TemplateFile) AllowWriteIfFormatFails() *TemplateFile {
	f.allowWriteIfFormatFails = true
	return f
}

func (f *TemplateFile) Imports(imports *cg.TemplateImportSet) *TemplateFile {
	f.imports = imports
	return f
}

func (f *TemplateFile) SetHeader(header string) *TemplateFile {
	f.headerString = header
	return f
}

func (f *TemplateFile) HeaderString() string {
	return f.headerString
}

func createHeader() string { return "" }

func FormatFile(context cg.GeneratorContextInterface, file *TemplateFile) (string, error) {
	value, err := cg.NodeToString(context, file)
	if err != nil {
		return "", err
	}
	return cg.FormatGoString(value)
}

func SignFile(context cg.GeneratorContextInterface, file *TemplateFile) (string, error) {
	formattedValue, err := FormatFile(context, file)
	if err != nil {
		return "", err
	}
	if file.signer == nil {
		return "", fmt.Errorf("File signing not configured")
	}
	return file.signer.SignString(formattedValue)
}

func MaybeSignFile(context cg.GeneratorContextInterface, file *TemplateFile) (string, error) {
	if file.signer != nil {
		return SignFile(context, file)
	}
	return FormatFile(context, file)
}

func Save(context cg.GeneratorContextInterface, file *TemplateFile) error {
	output, err := MaybeSignFile(context, file)
	if err != nil {
		return err
	}
	return file.destination.Write(context, output)
}
