package cgfile

import (
	"fmt"

	"github.com/onkeypress-llc/codegen/cg/cge"
	"github.com/onkeypress-llc/codegen/cg/cgelement"
	"github.com/onkeypress-llc/codegen/cg/cgi"
	"github.com/onkeypress-llc/codegen/cg/cgnode"
	"github.com/onkeypress-llc/codegen/cg/cgtmp"
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
	imports cgi.ImportSetInterface
	// set of content in the file
	contents []cgi.NodeInterface

	allowWriteIfFormatFails bool

	signer cgi.SignedStringInterface
}

// intentionally verbose
func NewFileWithoutGeneratorHeadersOrSigning(destination *Destination) *File {
	return newFile(destination)
}

func newFile(destination *Destination) *File {
	return &File{imports: cgelement.NewImportSet(), destination: destination, packageName: DefaultPackageName, contents: []cgi.NodeInterface{}}
}

func (f *File) UsedImports() (cgi.ImportSetInterface, error) {
	return f.imports, nil
}

func GoFileTemplates() cgi.TemplateSetInterface {
	return cgtmp.NewSet().AddTemplate(cgtmp.New("go-file-contents"), cgtmp.New("go-file-content"))
}

func (f *File) Generate(c cgi.ContextInterface) (cgi.NodeOutputInterface, error) {
	imports, err := cgelement.NewImportSet().MergeWith(f.imports)
	if err != nil {
		return nil, err
	}
	contents := make([]cgi.NodeOutputInterface, len(f.contents))
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
	).SetTemplates(templates).ContextForToString(configureContext), nil
}

func configureContext(ctx cgi.ContextInterface, output cgi.NodeOutputWithTypedDataInterface[*FileData]) cgi.ContextInterface {
	return ctx.WithinFile(output.Data().PackageName, output.Data().Imports)
}

func (f *File) generatedComments(ctx cgi.ContextInterface) (*cgelement.LineComment, *cgelement.LineComment) {
	signer := f.signer
	if signer == nil {
		return nil, nil
	}
	generateLine := cgelement.NewGoGenerateLineComment(ctx.GetCommandString())
	if signer.SigningType() == cge.Full {
		return generateLine, cgelement.NewGeneratedFileLineComment(ctx.GetAttributionString())
	} else if signer.SigningType() == cge.Partial {
		return generateLine, nil
	}
	panic("Undefined generation case")
}

func (f *File) Save(c cgi.ContextInterface) error {
	return Save(c, f)
}

func (f *File) Package(packageName string) *File {
	f.packageName = packageName
	return f
}

func (f *File) Contents(contents []cgi.NodeInterface) *File {
	f.contents = contents
	return f
}

func (f *File) Add(nodes ...cgi.NodeInterface) *File {
	if len(nodes) > 0 {
		f.contents = append(f.contents, nodes...)

	}
	return f
}

func (f *File) AllowWriteIfFormatFails() *File {
	f.allowWriteIfFormatFails = true
	return f
}

func (f *File) ImportsUsed(imports *cgelement.ImportSet) (*File, error) {
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

func FormatFile(context cgi.ContextInterface, file *File) (string, error) {
	value, err := cgnode.NodeToString(context, file)
	if err != nil {
		return "", err
	}
	return FormatGoString(value)
}

func SignFile(context cgi.ContextInterface, file *File) (string, error) {
	formattedValue, err := FormatFile(context, file)
	if err != nil {
		return "", err
	}
	if file.signer == nil {
		return "", fmt.Errorf("File signing not configured")
	}
	return file.signer.SignString(formattedValue)
}

func MaybeSignFile(context cgi.ContextInterface, file *File) (string, error) {
	if file.signer != nil {
		return SignFile(context, file)
	}
	return FormatFile(context, file)
}

func Save(context cgi.ContextInterface, file *File) error {
	output, err := MaybeSignFile(context, file)
	if err != nil {
		return err
	}
	return file.destination.Write(context, output)
}

func (f *File) ToInterface() cgi.NodeInterface {
	return f
}
