package cgnode

import (
	"fmt"
	"strings"

	"github.com/onkeypress-llc/codegen/cg/cgi"
)

type Import struct {
	source       string
	label        string
	defaultLabel string
}

func NewImport(source string) *Import {
	defaultLabel := GetImportDefaultLabel(source)
	return &Import{
		source:       source,
		defaultLabel: defaultLabel,
		label:        defaultLabel,
	}
}

func GetImportDefaultLabel(source string) string {
	parts := strings.Split(source, "/")
	return parts[len(parts)-1]
}

func (i *Import) Label() string {
	return i.label
}

func (i *Import) Source() string {
	return i.source
}

func (i *Import) DefaultLabel() string {
	return i.defaultLabel
}

func (i *Import) SetLabel(label string) *Import {
	i.label = label
	return i
}

func (i *Import) String() string {
	prefix := ""
	if i.defaultLabel != i.label {
		prefix = fmt.Sprintf("%s ", i.label)
	}
	return fmt.Sprintf("%s\"%s\"", prefix, i.source)
}

func (i *Import) Equals(compareTo cgi.ImportInterface) bool {
	return i.label == compareTo.Label() && i.source == compareTo.Source()
}
