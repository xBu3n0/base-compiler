package source

type Stream interface {
	NextRune() (Cursor, error)
}
