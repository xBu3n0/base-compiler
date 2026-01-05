package syntax

import (
	"github.com/xBu3n0/base-compiler/compiler"
)

type Parser interface {
	Parse(tokens *[]compiler.Token) compiler.AST
}
