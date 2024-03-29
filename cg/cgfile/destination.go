package cgfile

import (
	"path"

	"github.com/onkeypress-llc/codegen/cg/cgi"
)

type Destination struct {
	Name string
	Path string
}

func NewDestination(name, path string) *Destination {
	return &Destination{Name: name, Path: path}
}

func (d *Destination) Write(context cgi.ContextInterface, content string) error {
	filename := path.Join(d.Path, d.Name)
	return context.FS().Write(filename, content)
}
