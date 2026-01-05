package posfix

import (
	"github.com/xBu3n0/base-compiler/compiler"
	"github.com/xBu3n0/base-compiler/compiler/analyzer/lexer"
)

const (
	NUMBER        compiler.TokenType = iota // 0-9
	PLUS                                    // '+'
	MINUS                                   // '-'
	DOT                                     // print the result
	INVALID_TOKEN                           // otherwise
)

type CTokenizer struct{}

func NewTokenizer() lexer.Tokenizer {
	return &CTokenizer{}
}

func getToken(lexeme string) compiler.TokenType {
	if lexeme >= "0" && lexeme <= "9" {
		return NUMBER
	}

	switch lexeme {
	case "+":
		return PLUS
	case "-":
		return MINUS
	case ".":
		return DOT
	default:
		return INVALID_TOKEN
	}
}

func (t *CTokenizer) GenToken(lexeme compiler.Lexeme, span compiler.Span) compiler.Token {
	return compiler.Token{
		TokenType: getToken(lexeme.Value),
		Lexeme:    lexeme,
		Span:      span,
	}
}
