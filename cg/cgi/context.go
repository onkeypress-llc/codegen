package cgi

import (
	"embed"
)

// interface for the context object of a generation run
type ContextInterface interface {
	FS() FSInterface
	TemplateFS() embed.FS
	CommandString(string) ContextInterface
	AttributionString(string) ContextInterface
	GetCommandString() string
	GetAttributionString() string
	GetImportNamespace(ImportInterface) (string, error)
}
