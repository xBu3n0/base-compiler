package ast

import "github.com/xBu3n0/base-compiler/compiler"

const (
	INTEGER         = iota // [0-9]
	BINARY_OPERATOR        // + | - | * | /
	PRINT                  // .
)

// Op (+ | - | * | /)
type BinaryOpNode struct {
	label string
	span  compiler.Span
	left  compiler.AstNode
	right compiler.AstNode
}

func (node *BinaryOpNode) Span() compiler.Span {
	return node.span
}

// Print (.)
type UnaryOpNode struct {
	label string
	span  compiler.Span
	child compiler.AstNode
}

func (node *UnaryOpNode) Span() compiler.Span {
	return node.span
}
