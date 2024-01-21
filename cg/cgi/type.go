package cgi

type TypeInterface interface {
	Name() string
	Import() ImportInterface
}
