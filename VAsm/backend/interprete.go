package backend

import (
	"fmt"
	parser "VAsm/backend/analizador/parser"
	"VAsm/backend/visitor"
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
	visitor := visitor.NewEvalVisitor()
	visitor.Visit(tree)
}