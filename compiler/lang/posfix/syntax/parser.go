package posfix

import (
	"github.com/xBu3n0/base-compiler/compiler"
	"github.com/xBu3n0/base-compiler/compiler/analyzer/syntax"
)

type CParser struct{}

func NewParser() syntax.Parser {
	return &CParser{}
}

func (parser *CParser) Parse(tokens *[]compiler.Token) compiler.AST {
	return compiler.AST{}
}
