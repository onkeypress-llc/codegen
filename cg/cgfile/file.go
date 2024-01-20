package cgfile

import (
	"fmt"

	"github.com/onkeypress-llc/codegen/cg/cgcontext"
	"github.com/onkeypress-llc/codegen/cg/cgelement"
	"github.com/onkeypress-llc/codegen/cg/cgnode"
	"github.com/onkeypress-llc/codegen/cg/cgtmp"
	"github.com/onkeypress-llc/codegen/cg/signing"
)

const DefaultPackageName = "main"

type File struct {
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

// intentionally verbose
func NewFileWithoutGeneratorHeadersOrSigning(destination *Destination) *File {
	return newFile(destination)
}

func newFile(destination *Destination) *File {
	return &File{imports: cgnode.NewImportSet(), destination: destination, packageName: DefaultPackageName, contents: []cgnode.NodeInterface[any]{}}
}

func (f *File) UsedImports() (*cgnode.ImportSet, error) {
	return f.imports, nil
}

func GoFileTemplates() *cgtmp.Templates {
	return cgtmp.NewSet().AddTemplate(cgtmp.New("go-file-contents"), cgtmp.New("go-file-content"))
}

func (f *File) Generate(c cgcontext.Interface) (cgnode.NodeOutputInterface[*FileData], error) {
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
	fileHeader := cgelement.NewDocBlock(f.HeaderString())
	headerOutput, err := fileHeader.Generate(c)
	if err != nil {
		return nil, err
	}
	a, b := f.generatedComments(c)

	return cgnode.Output(
		cgtmp.New("go-file"),
		&FileData{
			GoGenerateCommentLine:    a,
			GeneratedFileCommentLine: b,
			Header:                   headerOutput,
			PackageName:              f.packageName,
			Imports:                  imports,
			Contents:                 contents,
		},
	).SetTemplates(templates), nil
}

func (f *File) generatedComments(ctx cgcontext.Interface) (*cgelement.LineComment, *cgelement.LineComment) {
	signer := f.signer
	if signer == nil {
		return nil, nil
	}
	generateLine := cgelement.NewGoGenerateLineComment(ctx.GetCommandString())
	if signer.SigningType() == signing.Full {
		return generateLine, cgelement.NewGeneratedFileLineComment(ctx.GetAttributionString())
	} else if signer.SigningType() == signing.Partial {
		return generateLine, nil
	}
	panic("Undefined generation case")
}

func (f *File) Save(c cgcontext.Interface) error {
	return Save(c, f)
}

func (f *File) Package(packageName string) *File {
	f.packageName = packageName
	return f
}

func (f *File) Contents(contents []cgnode.NodeInterface[any]) *File {
	f.contents = contents
	return f
}

func (f *File) AllowWriteIfFormatFails() *File {
	f.allowWriteIfFormatFails = true
	return f
}

func (f *File) ImportsUsed(imports *cgnode.ImportSet) (*File, error) {
	f.imports = imports
	return f, nil
}

func (f *File) SetHeader(header string) *File {
	f.headerString = header
	return f
}

func (f *File) HeaderString() string {
	value, signer := f.headerString, f.signer
	if signer == nil {
		return value
	}
	return signer.DocBlock(value)
}

func FormatFile(context cgcontext.Interface, file *File) (string, error) {
	value, err := cgnode.NodeToString(context, file)
	if err != nil {
		return "", err
	}
	return FormatGoString(value)
}

func SignFile(context cgcontext.Interface, file *File) (string, error) {
	formattedValue, err := FormatFile(context, file)
	if err != nil {
		return "", err
	}
	if file.signer == nil {
		return "", fmt.Errorf("File signing not configured")
	}
	return file.signer.SignString(formattedValue)
}

func MaybeSignFile(context cgcontext.Interface, file *File) (string, error) {
	if file.signer != nil {
		return SignFile(context, file)
	}
	return FormatFile(context, file)
}

func Save(context cgcontext.Interface, file *File) error {
	output, err := MaybeSignFile(context, file)
	if err != nil {
		return err
	}
	return file.destination.Write(context, output)
}

func (f *File) ToInterface() cgnode.NodeInterface[*FileData] {
	return f
}
