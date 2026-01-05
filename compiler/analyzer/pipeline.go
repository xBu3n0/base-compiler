package analyzer

import (
	"github.com/xBu3n0/base-compiler/compiler"
	"github.com/xBu3n0/base-compiler/compiler/source"
)

// scan -> tokenize -> parse -> validate
type Pipeline interface {
	Compile(stream source.Stream) (compiler.AST, error)
}
