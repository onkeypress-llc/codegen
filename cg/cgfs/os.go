package cgfs

import (
	"errors"
	"os"
)

type OSFS struct{}

var osFS = &OSFS{}

func (fs *OSFS) Read(filename string) (string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(f), nil
}

func (fs *OSFS) Write(filename, content string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	return err
}

func (fs *OSFS) Exists(filename string) bool {
	_, err := os.Stat(filename)
	return !errors.Is(err, os.ErrNotExist)
}

func NewOSFS() FSInterface {
	return osFS
}
