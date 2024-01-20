package cgelement

func NewGoGenerateLineComment(command string) *LineComment {
	return NewLineComment("go:generate %s", command).NoSpaceAfterSlash()
}
