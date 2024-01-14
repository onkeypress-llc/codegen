package cg

import (
	"github.com/onkeypress-llc/codegen/cg/cgfs"
)

// interface for the context object of a generation run
type GeneratorContextInterface interface {
	FS() cgfs.FSInterface
}

// implementation for the context interface
type GeneratorContext struct {
	FileSystem cgfs.FSInterface
}

func (c *GeneratorContext) FS() cgfs.FSInterface {
	return c.FileSystem
}

func NewContext() GeneratorContextInterface {
	return NewContextWithFS(cgfs.NewOSFS())
}

func NewContextWithFS(fs cgfs.FSInterface) GeneratorContextInterface {
	return &GeneratorContext{FileSystem: fs}
}
