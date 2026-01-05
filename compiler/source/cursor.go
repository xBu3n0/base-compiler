package source

type Cursor struct {
	Value rune
	Row   int
	Col   int
}

func NewCursor(value rune, row int, col int) Cursor {
	return Cursor{
		Value: value,
		Row:   row,
		Col:   col,
	}
}
