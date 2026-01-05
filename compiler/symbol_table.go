package compiler

import "fmt"

type SymbolTable interface {
	LookAhead()
}

type ST struct{}

func (st *ST) LookAhead() {
	fmt.Println("Teste")
}
