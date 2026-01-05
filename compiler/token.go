package compiler

type TokenType int

type Token struct {
	TokenType TokenType
	Lexeme    Lexeme
	Span      Span
}
