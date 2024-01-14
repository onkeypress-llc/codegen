package cgfs

import "fmt"

type MapFS struct {
	files map[string]string
}

func (fs *MapFS) Read(filename string) (string, error) {
	content, ok := fs.files[filename]
	if !ok {
		return "", fmt.Errorf("No file located at %s", filename)
	}
	return content, nil
}

func (fs *MapFS) Write(filename, content string) error {
	fs.files[filename] = content
	return nil
}

func (fs *MapFS) Exists(filename string) bool {
	_, ok := fs.files[filename]
	return ok
}

func (fs *MapFS) Set(files ...[2]string) *MapFS {
	for i := range files {
		fs.Write(files[i][0], files[i][1])
	}
	return fs
}

func (fs *MapFS) Remove(files ...string) *MapFS {
	for i := range files {
		delete(fs.files, files[i])
	}
	return fs
}

func NewMapFS() *MapFS {
	return &MapFS{files: make(map[string]string)}
}

func NewMapFile(name, content string) [2]string {
	return [2]string{name, content}
}
