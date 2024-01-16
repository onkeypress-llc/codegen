package cgnode

import (
	"fmt"
	"strings"
)

type Import struct {
	Source       string
	Label        string
	defaultLabel string
}

func NewImport(source string) *Import {
	defaultLabel := GetImportDefaultLabel(source)
	return &Import{
		Source:       source,
		defaultLabel: defaultLabel,
		Label:        defaultLabel,
	}
}

func GetImportDefaultLabel(source string) string {
	parts := strings.Split(source, "/")
	return parts[len(parts)-1]
}

func (i *Import) DefaultLabel() string {
	return i.defaultLabel
}

func (i *Import) SetLabel(label string) *Import {
	i.Label = label
	return i
}

func (i *Import) String() string {
	prefix := ""
	if i.defaultLabel != i.Label {
		prefix = fmt.Sprintf("%s ", i.Label)
	}
	return fmt.Sprintf("%s\"%s\"", prefix, i.Source)
}

func (i *Import) Equals(compareTo *Import) bool {
	return i.Label == compareTo.Label && i.Source == compareTo.Source
}
