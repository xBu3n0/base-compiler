package semantic

import "github.com/xBu3n0/base-compiler/compiler"

type Analyzer interface {
	Validate(ast *compiler.AST) bool
}
