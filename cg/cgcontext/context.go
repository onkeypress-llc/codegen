package cgcontext

import (
	"embed"

	"github.com/onkeypress-llc/codegen/cg/cgfs"
)

// interface for the context object of a generation run
type Interface interface {
	FS() cgfs.FSInterface
	TemplateFS() embed.FS
	CommandString(string) Interface
	AttributionString(string) Interface
	GetCommandString() string
	GetAttributionString() string
}

// implementation for the context interface
type Context struct {
	// the full command used to generate the output as you would find it after a go:generate comment
	generateCommandString string
	// the name of the process responsible for generating output
	generatedAttributionString string

	// filestystem to read/write generated files
	fs cgfs.FSInterface
	// filesystem to load templates
	templateFiles embed.FS
}

func (c *Context) CommandString(value string) Interface {
	c.generateCommandString = value
	return c
}

func (c *Context) AttributionString(value string) Interface {
	c.generatedAttributionString = value
	return c
}

func (c *Context) GetCommandString() string {
	return c.generateCommandString
}

func (c *Context) GetAttributionString() string {
	return c.generatedAttributionString
}

func (c *Context) FS() cgfs.FSInterface {
	return c.fs
}

func (c *Context) TemplateFS() embed.FS {
	return c.templateFiles
}

func (c *Context) SetFS(fs cgfs.FSInterface) *Context {
	c.fs = fs
	return c
}

func (c *Context) SetTemplateFS(fs embed.FS) *Context {
	c.templateFiles = fs
	return c
}

func New(templateFiles embed.FS) *Context {
	return &Context{fs: cgfs.NewOSFS(), templateFiles: templateFiles}
}
