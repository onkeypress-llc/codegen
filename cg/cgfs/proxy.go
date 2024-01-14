package cgfs

type ProxyFS struct {
	ReadFn   func(string) (string, error)
	WriteFn  func(string, string) error
	ExistsFn func(string) bool
}

func (fs *ProxyFS) Read(filename string) (string, error) {
	return fs.ReadFn(filename)
}

func (fs *ProxyFS) Write(filename, content string) error {
	return fs.WriteFn(filename, content)
}

func (fs *ProxyFS) Exists(filename string) bool {
	return fs.ExistsFn(filename)
}

func NewProxyFS() FSInterface {
	return &ProxyFS{ReadFn: NoopReadFn, WriteFn: NoopWriteFn, ExistsFn: NoopExistsFn}
}

func NoopReadFn(filename string) (string, error) {
	return "", nil
}
func NoopWriteFn(filename, content string) error {
	return nil
}
func NoopExistsFn(filename string) bool {
	return false
}
