package cg

type Template struct {
	Name string
}

func NewTemplate(name string) *Template {
	return &Template{Name: name}
}

type Templates struct {
	hash map[string]bool
}

func (s *Templates) Size() int {
	return len(s.hash)
}

func (s *Templates) AddTemplates(container *Templates) *Templates {
	if container == nil {
		return s
	}
	return s.AddStrings(container.Names()...)
}

func (s *Templates) AddTemplate(list ...*Template) *Templates {
	names := make([]string, len(list))
	for i := range list {
		t := list[i]
		names[i] = t.Name
	}
	return s.AddStrings(names...)
}

func (s *Templates) AddStrings(list ...string) *Templates {
	for i := range list {
		t := list[i]
		s.hash[t] = true
	}
	return s
}

func (s *Templates) Names() []string {
	names := make([]string, len(s.hash))
	index := 0
	for name, _ := range s.hash {
		names[index] = name
		index++
	}
	return names
}

func NewTemplates() *Templates {
	return &Templates{hash: make(map[string]bool)}
}
