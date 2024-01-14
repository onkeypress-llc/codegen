package cg

import (
	"fmt"
	"strings"
)

type TemplateImportSet struct {
	imports map[string]*TemplateImport
}

func NewImportSet() *TemplateImportSet {
	return &TemplateImportSet{
		imports: make(map[string]*TemplateImport, 0),
	}
}

func NewImportSetWithValues(imports ...*TemplateImport) (*TemplateImportSet, error) {
	return NewImportSetWithArray(imports)
}

func NewImportSetWithArray(imports []*TemplateImport) (*TemplateImportSet, error) {
	return NewImportSet().AddValues(imports)
}

func (s *TemplateImportSet) Add(imports ...*TemplateImport) (*TemplateImportSet, error) {
	return s.AddValues(imports)
}

func (s *TemplateImportSet) AddValues(imports []*TemplateImport) (*TemplateImportSet, error) {
	for i := range imports {
		imp := imports[i]
		if imp == nil {
			continue
		}
		key := imp.Label
		existing, ok := s.imports[key]
		if !ok {
			s.imports[imp.Label] = imp
		} else if !existing.Equals(imp) {
			return nil, fmt.Errorf("Unable to merge import %+v into existing set, %+v has conflicting name", imp, existing)
		}
	}
	return s, nil
}

func (s *TemplateImportSet) String() string {
	length := len(s.imports)
	if length < 1 {
		return ""
	}
	statements := make([]string, length)
	i := 0
	for _, value := range s.imports {
		statements[i] = value.String()
		i++
	}
	if length == 1 {
		return fmt.Sprintf("import %s", statements[0])
	}
	return fmt.Sprintf("import (\n%s\n)", strings.Join(statements, "\n"))
}

func (s *TemplateImportSet) MergeValuesInFrom(other *TemplateImportSet) (*TemplateImportSet, error) {
	if other == nil {
		return s, nil
	}
	values := make([]*TemplateImport, len(other.imports))
	i := 0
	for _, value := range other.imports {
		values[i] = value
		i++
	}
	return s.AddValues(values)
}
