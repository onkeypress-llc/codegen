package cg

type TemplateType struct {
	Name   string
	Import *TemplateImport
}

func (t *TemplateType) String() string {
	return t.Name
}
