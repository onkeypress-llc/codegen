package cgfile

import "go/format"

func FormatGoString(content string) (string, error) {
	formatted, err := format.Source([]byte(content))
	return string(formatted), err
}
