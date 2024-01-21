package cgi

type FSInterface interface {
	Exists(string) bool
	Read(string) (string, error)
	Write(string, string) error
}
