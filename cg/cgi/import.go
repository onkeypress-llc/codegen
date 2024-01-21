package cgi

type ImportInterface interface {
	Source() string
	Label() string
	Equals(ImportInterface) bool
	String() string
	DefaultLabel() string
}
