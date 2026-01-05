package lexer

import (
	"github.com/xBu3n0/base-compiler/compiler"
	"github.com/xBu3n0/base-compiler/compiler/source"
)

type Scanner interface {
	// (Lexeme, [EOF|InvalidLexeme|nil])
	NextLexeme(stream source.Stream) (compiler.Lexeme, compiler.Span, error)
}
