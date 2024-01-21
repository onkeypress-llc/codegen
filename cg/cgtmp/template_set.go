package cgtmp

import "github.com/onkeypress-llc/codegen/cg/cgi"

type TemplateSet struct {
	hash        map[string]bool
	orderedList []string
}

func (s *TemplateSet) Size() int {
	return len(s.orderedList)
}

func (s *TemplateSet) AddTemplates(container cgi.TemplateSetInterface) *TemplateSet {
	if container == nil {
		return s
	}
	return s.AddStrings(container.Names()...)
}

func (s *TemplateSet) AddTemplate(list ...cgi.TemplateInterface) *TemplateSet {
	names := make([]string, len(list))
	for i := range list {
		t := list[i]
		names[i] = t.FullName()
	}
	return s.AddStrings(names...)
}

func (s *TemplateSet) AddStrings(list ...string) *TemplateSet {
	for i := range list {
		t := list[i]
		_, ok := s.hash[t]
		if !ok {
			s.hash[t] = true
			s.orderedList = append(s.orderedList, t)
		}
	}
	return s
}

func (s *TemplateSet) Names() []string {
	return s.orderedList
}

func NewSet(templates ...cgi.TemplateInterface) *TemplateSet {
	instance := &TemplateSet{hash: make(map[string]bool), orderedList: make([]string, 0)}
	return instance.AddTemplate(templates...)
}
