package cgi

type ImportSetInterface interface {
	ImportMap() map[string]ImportInterface
	GetNamespaceForImport(ImportInterface) string
}
