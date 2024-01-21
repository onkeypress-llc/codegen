package cgi

type TemplateInterface interface {
	FullName() string
	NameWithExtension() string
}
