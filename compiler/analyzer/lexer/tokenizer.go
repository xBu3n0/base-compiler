package lexer

import "github.com/xBu3n0/base-compiler/compiler"

type Tokenizer interface {
	GenToken(lexeme compiler.Lexeme, span compiler.Span) compiler.Token
}
