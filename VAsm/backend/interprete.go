package backend

import (
	parser "VAsm/backend/analizador/parser"
	"VAsm/backend/visitor"
	"VAsm/frontend/symbols"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type TreeShapeListener struct {
	*parser.BasegramaticaListener
}

func Run(entrada string) {
	input := antlr.NewInputStream(entrada)
	lexer := parser.NewgramaticaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewgramaticaParser(stream)
	tree := p.Init()
	fmt.Println(tree.ToStringTree(p.RuleNames, p))
	// Initialize the symbol table as needed; here we use an empty TablaSimbolos as an example
	simbolTable := symbols.NewTablaSimbolos()
	visitor := visitor.NewEvalVisitor(simbolTable)
	visitor.Visit(tree)
}
