package postfix

import (
	"fmt"

	"github.com/xBu3n0/base-compiler/compiler"
	"github.com/xBu3n0/base-compiler/compiler/analyzer"
	"github.com/xBu3n0/base-compiler/compiler/analyzer/lexer"
	"github.com/xBu3n0/base-compiler/compiler/analyzer/semantic"
	"github.com/xBu3n0/base-compiler/compiler/analyzer/syntax"
	cLexer "github.com/xBu3n0/base-compiler/compiler/lang/postfix/lexer"
	cSemantic "github.com/xBu3n0/base-compiler/compiler/lang/postfix/semantic"
	cSyntax "github.com/xBu3n0/base-compiler/compiler/lang/postfix/syntax"
	"github.com/xBu3n0/base-compiler/compiler/source"
)

type Pipeline struct {
	// Compile(stream source.Stream) analyzer.AST
	// TODO: Adicionar preprocessor
	// preprocessor preprocessor.Preprocessor
	scanner   lexer.Scanner
	tokenizer lexer.Tokenizer
	parser    syntax.Parser
	analyzer  semantic.Analyzer
}

func NewCompiler() analyzer.Pipeline {
	return &Pipeline{
		// preprocessor: NewProcessor(),
		scanner:   cLexer.NewScanner(),
		tokenizer: cLexer.NewTokenizer(),
		parser:    cSyntax.NewParser(),
		analyzer:  cSemantic.NewAnalyzer(),
	}
}

func (p *Pipeline) Compile(stream source.Stream) (compiler.AST, error) {
	/** Lexical analyzer */
	var tokens []compiler.Token

	for lexeme, span, err := p.scanner.NextLexeme(stream); ; lexeme, span, err = p.scanner.NextLexeme(stream) {
		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			fmt.Println("Erro: ", err)
			return compiler.AST{}, err
		}

		token := p.tokenizer.GenToken(lexeme, span)

		tokens = append(tokens, token)
	}

	/** Parser */
	ast := p.parser.Parse(&tokens)
	fmt.Println(tokens, ast)

	/** Analyzer */
	p.analyzer.Validate(&ast)

	return ast, nil
}
