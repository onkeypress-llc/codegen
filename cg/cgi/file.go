package cgi

type FileInterface interface {
	Save(ContextInterface) error
}
