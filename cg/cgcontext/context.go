package cgcontext

import (
	"embed"

	"github.com/onkeypress-llc/codegen/cg/cgfs"
	"github.com/onkeypress-llc/codegen/cg/cgi"
)

// implementation for the context interface
type Context struct {
	// the full command used to generate the output as you would find it after a go:generate comment
	generateCommandString string
	// the name of the process responsible for generating output
	generatedAttributionString string

	// filestystem to read/write generated files
	fs cgi.FSInterface
	// filesystem to load templates
	templateFiles embed.FS

	// the file the context for the output
	// the current package source of the file
	currentSource string
	// the file import set being used
	currentImportSet cgi.ImportSetInterface
}

func (c *Context) CommandString(value string) cgi.ContextInterface {
	c.generateCommandString = value
	return c
}

func (c *Context) GetImportNamespace(imp cgi.ImportInterface) (string, error) {
	currentSource, currentImportSet := c.currentSource, c.currentImportSet
	// if import is for a type in the same package, no namespace qualification needed
	if imp.Source() == currentSource {
		return "", nil
	}
	return currentImportSet.GetNamespaceForImport(imp), nil
}

func (c *Context) AttributionString(value string) cgi.ContextInterface {
	c.generatedAttributionString = value
	return c
}

func (c *Context) GetCommandString() string {
	return c.generateCommandString
}

func (c *Context) GetAttributionString() string {
	return c.generatedAttributionString
}

func (c *Context) FS() cgi.FSInterface {
	return c.fs
}

func (c *Context) TemplateFS() embed.FS {
	return c.templateFiles
}

func (c *Context) SetFS(fs cgi.FSInterface) *Context {
	c.fs = fs
	return c
}

func (c *Context) SetTemplateFS(fs embed.FS) *Context {
	c.templateFiles = fs
	return c
}

func (c *Context) WithinFile(source string, importSet cgi.ImportSetInterface) cgi.ContextInterface {
	newContext := Copy(c)
	newContext.currentSource = source
	newContext.currentImportSet = importSet
	return newContext
}

func New(templateFiles embed.FS) *Context {
	return &Context{fs: cgfs.NewOSFS(), templateFiles: templateFiles}
}

func Copy(ctx *Context) *Context {
	return &Context{
		generateCommandString:      ctx.generateCommandString,
		generatedAttributionString: ctx.generatedAttributionString,
		fs:                         ctx.fs,
		templateFiles:              ctx.templateFiles,
		currentSource:              ctx.currentSource,
		currentImportSet:           ctx.currentImportSet,
	}
}
