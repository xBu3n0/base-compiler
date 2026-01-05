package postfix

import (
	"errors"
	"strings"

	"github.com/xBu3n0/base-compiler/compiler"
	"github.com/xBu3n0/base-compiler/compiler/analyzer/lexer"
	"github.com/xBu3n0/base-compiler/compiler/source"
)

type CScanner struct {
}

func NewScanner() lexer.Scanner {
	return &CScanner{}
}

func (scanner *CScanner) NextLexeme(stream source.Stream) (compiler.Lexeme, compiler.Span, error) {
	lexeme := ""
	row := -1
	col := -1

	for cursor, err := stream.NextRune(); err == nil; cursor, err = stream.NextRune() {
		// Skipar todos as runas que nao inicializam um lexema
		if strings.ContainsAny(string(cursor.Value), " \t\r\n") {
			continue
		}

		lexeme = string(cursor.Value)
		row = cursor.Row
		col = cursor.Col
		break
	}

	if len(lexeme) == 0 {
		return compiler.Lexeme{}, compiler.Span{}, errors.New("EOF")
	}

	return compiler.Lexeme{
			Value: lexeme,
		}, compiler.Span{
			File:   "file.src",
			Row:    row,
			Column: col,
			Offset: len(lexeme),
		}, nil
}
