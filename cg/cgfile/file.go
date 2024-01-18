package cgfile

import (
	"fmt"

	"github.com/onkeypress-llc/codegen/cg"
	"github.com/onkeypress-llc/codegen/cg/cgcontext"
	"github.com/onkeypress-llc/codegen/cg/cgnode"
	"github.com/onkeypress-llc/codegen/cg/cgtmp"
	"github.com/onkeypress-llc/codegen/cg/signing"
)

type TemplateFile struct {
	// header comment for file
	headerString string
	// where the file should be written on the OS
	destination *Destination
	// package at the top of file
	packageName string
	// manually added imports to file
	imports *cgnode.ImportSet
	// set of content in the file
	contents []cgnode.NodeInterface[any]

	allowWriteIfFormatFails bool

	signer signing.SignedStringInterface
}

type Data struct {
	// header comment
	Header cgnode.NodeOutputInterface[*cg.Docblock]
	// package of this file
	PackageName string
	// union of imports used in the file
	Imports *cgnode.ImportSet
	// set of content in the file
	Contents []cgnode.NodeOutputInterface[any]
}

func NewFile(destination *Destination) *TemplateFile {
	return &TemplateFile{imports: cgnode.NewImportSet(), destination: destination, contents: []cgnode.NodeInterface[any]{}}
}

func (f *TemplateFile) UsedImports() (*cgnode.ImportSet, error) {
	return f.imports, nil
}

func GoFileTemplates() *cgtmp.Templates {
	return cgtmp.NewSet().AddTemplate(cgtmp.New("go-file-contents"), cgtmp.New("go-file-content"))
}

func (f *TemplateFile) Generate(c cgcontext.Interface) (cgnode.NodeOutputInterface[*Data], error) {
	imports, err := cgnode.NewImportSet().MergeWith(f.imports)
	if err != nil {
		return nil, err
	}
	contents := make([]cgnode.NodeOutputInterface[any], len(f.contents))
	templates := cgtmp.NewSet().AddTemplates(GoFileTemplates())
	// from file content blocks, aggragate imports, template names, and content
	// add imports from file content blocks
	for i := range f.contents {
		content := f.contents[i]
		contentOutput, err := content.Generate(c)
		if err != nil {
			return nil, err
		}
		contents[i] = contentOutput
		templates = templates.AddTemplates(contentOutput.Templates())
		nodeImports, err := contentOutput.UsedImports()
		if err != nil {
			return nil, err
		}
		imports, err = imports.MergeWith(nodeImports)
		if err != nil {
			return nil, err
		}
	}
	fileHeader := cg.NewDocBlock(f.HeaderString())
	headerOutput, err := fileHeader.Generate(c)
	if err != nil {
		return nil, err
	}

	return cgnode.Output(
		cgtmp.New("go-file"),
		&Data{
			Header:      headerOutput,
			PackageName: f.packageName,
			Imports:     imports,
			Contents:    contents,
		},
	).SetTemplates(templates), nil
}

func (f *TemplateFile) Save(c cgcontext.Interface) (string, error) {
	panic("Not implemented")
}

func (f *TemplateFile) Package(packageName string) *TemplateFile {
	f.packageName = packageName
	return f
}

func (f *TemplateFile) Contents(contents []cgnode.NodeInterface[any]) *TemplateFile {
	f.contents = contents
	return f
}

func (f *TemplateFile) AllowWriteIfFormatFails() *TemplateFile {
	f.allowWriteIfFormatFails = true
	return f
}

func (f *TemplateFile) ImportsUsed(imports *cgnode.ImportSet) (*TemplateFile, error) {
	f.imports = imports
	return f, nil
}

func (f *TemplateFile) SetHeader(header string) *TemplateFile {
	f.headerString = header
	return f
}

func (f *TemplateFile) HeaderString() string {
	value, signer := f.headerString, f.signer
	if signer == nil {
		return value
	}
	return signer.DocBlock(value)
}

func FormatFile(context cgcontext.Interface, file *TemplateFile) (string, error) {
	value, err := cg.NodeToString(context, file)
	if err != nil {
		return "", err
	}
	return cg.FormatGoString(value)
}

func SignFile(context cgcontext.Interface, file *TemplateFile) (string, error) {
	formattedValue, err := FormatFile(context, file)
	if err != nil {
		return "", err
	}
	if file.signer == nil {
		return "", fmt.Errorf("File signing not configured")
	}
	return file.signer.SignString(formattedValue)
}

func MaybeSignFile(context cgcontext.Interface, file *TemplateFile) (string, error) {
	if file.signer != nil {
		return SignFile(context, file)
	}
	return FormatFile(context, file)
}

func Save(context cgcontext.Interface, file *TemplateFile) error {
	output, err := MaybeSignFile(context, file)
	if err != nil {
		return err
	}
	return file.destination.Write(context, output)
}

func (f *TemplateFile) ToInterface() cgnode.NodeInterface[*Data] {
	return f
}
