package source

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type File struct {
	file    *bufio.Reader
	pos     int
	row     int
	col     int
	fileLen int
}

func (fs *File) NextRune() (Cursor, error) {
	rune, _, err := fs.file.ReadRune()

	if err != nil {
		return Cursor{}, err
	}

	r, c := fs.row, fs.col

	switch rune {
	case '\n':
		fs.row += 1
		fs.col = 1
	default:
		fs.col += 1
	}

	return NewCursor(rune, r, c), nil
}

func NewFileStream(file string) (Stream, error) {
	f, error := os.Open(file)
	if error != nil {
		fmt.Println("Nao foi possivel abrir o arquivo: ", file)
		return nil, error
	}

	reader := bufio.NewReader(f)
	f.Seek(0, io.SeekStart)

	fileLen, err := countRunes(f)
	f.Seek(0, io.SeekStart)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &File{
		file:    reader,
		pos:     0,
		row:     1,
		col:     1,
		fileLen: fileLen,
	}, nil
}

func countRunes(f *os.File) (int, error) {
	reader := bufio.NewReader(f)

	count := 0

	for {
		_, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return count, nil
			}
			return 0, err
		}
		count++
	}
}
