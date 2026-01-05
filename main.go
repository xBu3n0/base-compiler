package main

import (
	"fmt"
	"os"

	"github.com/xBu3n0/base-compiler/compiler/lang/posfix"
	"github.com/xBu3n0/base-compiler/compiler/source"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println(`
		=======================================
		Informe o arquivo fonte para compilar:

		ex: "go run . main.src"
		=======================================
		`)

		return
	}

	sourceFile := args[1]

	/** Inicializacao */
	// Ler arquivo
	sourceStream, err := source.NewFileStream(sourceFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	/** Compiler */
	comp := posfix.NewCompiler()

	ast, err := comp.Compile(sourceStream)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ast, err)
}
