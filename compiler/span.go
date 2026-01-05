package compiler

type Span struct {
	File   string
	Row    int
	Column int
	Offset int
}
