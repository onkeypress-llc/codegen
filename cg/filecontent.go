package cg

type TemplateFileContent interface {
	Imports() *TemplateImportSet
	Templates() *Templates
	Generate() (TemplateFileContentOutput, error)
}

type TemplateFileContentOutput interface {
	Name() string
	Data() any
}
