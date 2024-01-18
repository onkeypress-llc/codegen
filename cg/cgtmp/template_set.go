package cgtmp

type Templates struct {
	hash        map[string]bool
	orderedList []string
}

func (s *Templates) Size() int {
	return len(s.orderedList)
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
		names[i] = t.FullName()
	}
	return s.AddStrings(names...)
}

func (s *Templates) AddStrings(list ...string) *Templates {
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

func (s *Templates) Names() []string {
	return s.orderedList
}

func NewSet(templates ...*Template) *Templates {
	instance := &Templates{hash: make(map[string]bool), orderedList: make([]string, 0)}
	return instance.AddTemplate(templates...)
}
