package cgelement

import (
	"fmt"
	"strings"

	"github.com/onkeypress-llc/codegen/cg/cgi"
	"github.com/onkeypress-llc/codegen/cg/cgnode"
)

type ImportSet struct {
	imports map[string]cgi.ImportInterface
}

func NewImportSet(imports ...cgi.ImportInterface) *ImportSet {
	set := &ImportSet{
		imports: make(map[string]cgi.ImportInterface, len(imports)),
	}
	result, err := set.AddValues(imports)
	if err != nil {
		panic(err)
	}
	return result
}

func (s *ImportSet) Add(imports ...cgi.ImportInterface) (*ImportSet, error) {
	return s.AddValues(imports)
}

func (s *ImportSet) AddValues(imports []cgi.ImportInterface) (*ImportSet, error) {
	for i := range imports {
		imp := imports[i]
		if imp == nil {
			continue
		}
		key := imp.Label()
		existing, ok := s.imports[key]
		if !ok {
			s.imports[imp.Label()] = imp
		} else if !existing.Equals(imp) {
			return nil, fmt.Errorf("Unable to merge import %+v into existing set, %+v has conflicting name", imp, existing)
		}
	}
	return s, nil
}

func (s *ImportSet) ImportMap() map[string]cgi.ImportInterface {
	return s.imports
}

func (s *ImportSet) GetNamespaceForImport(imp cgi.ImportInterface) string {
	return ""
}

func (s *ImportSet) Generate(ctx cgi.ContextInterface) (cgi.NodeOutputInterface, error) {
	return cgnode.StringOutput[*ImportSet](s), nil
}

func (s *ImportSet) ToString(ctx cgi.ContextInterface) (string, error) {
	length := len(s.imports)
	if length < 1 {
		return "", nil
	}
	statements := make([]string, length)
	i := 0
	for _, value := range s.imports {
		statements[i] = value.String()
		i++
	}
	if length == 1 {
		return fmt.Sprintf("import %s", statements[0]), nil
	}
	return fmt.Sprintf("import (\n%s\n)", strings.Join(statements, "\n")), nil
}

func (s *ImportSet) MergeWith(other cgi.ImportSetInterface) (*ImportSet, error) {
	if other == nil {
		return s, nil
	}
	values := make([]cgi.ImportInterface, len(other.ImportMap()))
	i := 0
	for _, value := range other.ImportMap() {
		values[i] = value
		i++
	}
	return s.AddValues(values)
}
