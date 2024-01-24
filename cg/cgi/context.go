package cgi

import (
	"embed"
)

// interface for the context object of a generation run
type ContextInterface interface {
	// the filesystem output will be written to
	FS() FSInterface
	// the filesystem to be used when loading templates to generate the output
	TemplateFS() embed.FS
	// the command used to generate the output
	CommandString(string) ContextInterface
	// the name of the generation process creating the output
	AttributionString(string) ContextInterface
	// retrieve CommandString value
	GetCommandString() string
	// retrieve AttributionString value
	GetAttributionString() string
	// Given an import, compute what namespace it should use within the current context
	GetImportNamespace(ImportInterface) (string, error)
	// Creates a new context with the file argument's package name and import namespaces applied
	WithinFile(string, ImportSetInterface) ContextInterface
}
