package posfix

import (
	"github.com/xBu3n0/base-compiler/compiler"
	"github.com/xBu3n0/base-compiler/compiler/analyzer/semantic"
)

type CAnalyzer struct {
}

func NewAnalyzer() semantic.Analyzer {
	return &CAnalyzer{}
}

func (a *CAnalyzer) Validate(ast *compiler.AST) bool {
	return true
}
