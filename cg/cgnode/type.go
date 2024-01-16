package cgnode

type Type struct {
	Name   string
	Import *Import
}

func (t *Type) String() string {
	return t.Name
}
