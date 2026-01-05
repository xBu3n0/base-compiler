package compiler

type NodeType int

type AST struct {
	Root *AstNode
}

type AstNode interface {
	Span() Span
}
