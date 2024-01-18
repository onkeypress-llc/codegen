package cgnode

import (
	"fmt"
	"strings"

	"github.com/onkeypress-llc/codegen/cg/cgcontext"
)

type ImportSet struct {
	imports map[string]*Import
}

func NewImportSet(imports ...*Import) *ImportSet {
	set := &ImportSet{
		imports: make(map[string]*Import, len(imports)),
	}
	result, err := set.AddValues(imports)
	if err != nil {
		panic(err)
	}
	return result
}

func (s *ImportSet) Add(imports ...*Import) (*ImportSet, error) {
	return s.AddValues(imports)
}

func (s *ImportSet) AddValues(imports []*Import) (*ImportSet, error) {
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

func (s *ImportSet) Generate(ctx cgcontext.Interface) (NodeOutputInterface[*ImportSet], error) {
	return StringOutput[*ImportSet](s), nil
}

func (s *ImportSet) ToString(ctx cgcontext.Interface) (string, error) {
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

func (s *ImportSet) MergeWith(other *ImportSet) (*ImportSet, error) {
	if other == nil {
		return s, nil
	}
	values := make([]*Import, len(other.imports))
	i := 0
	for _, value := range other.imports {
		values[i] = value
		i++
	}
	return s.AddValues(values)
}
