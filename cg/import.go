package cg

import (
	"fmt"
	"strings"
)

type TemplateImport struct {
	Source       string
	Label        string
	defaultLabel string
}

func NewTemplateImport(source string) *TemplateImport {
	return NewLabeledTemplateImport(source, "")
}

func NewLabeledTemplateImport(source, label string) *TemplateImport {
	defaultLabel := GetImportDefaultLabel(source)
	if len(label) == 0 {
		label = defaultLabel
	}
	return &TemplateImport{
		Source:       source,
		defaultLabel: defaultLabel,
		Label:        label,
	}
}

func GetImportDefaultLabel(source string) string {
	parts := strings.Split(source, "/")
	return parts[len(parts)-1]
}

func (i *TemplateImport) DefaultLabel() string {
	return i.defaultLabel
}

func (i *TemplateImport) String() string {
	prefix := ""
	if i.defaultLabel != i.Label {
		prefix = fmt.Sprintf("%s ", i.Label)
	}
	return fmt.Sprintf("%s\"%s\"", prefix, i.Source)
}

func (i *TemplateImport) Equals(compareTo *TemplateImport) bool {
	return i.Label == compareTo.Label && i.Source == compareTo.Source
}
