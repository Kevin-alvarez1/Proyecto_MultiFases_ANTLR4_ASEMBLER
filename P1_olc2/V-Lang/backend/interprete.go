package backend

import (
	"fmt"
	compiler "vlang/backend/analizador/parser"
	"vlang/backend/listener" // Importa tu paquete listener

	"github.com/antlr4-go/antlr/v4"
)

type TreeShapeListener struct {
	*compiler.BasegramaticaListener
}

func Run(entrada string) {
	input := antlr.NewInputStream(entrada)
	lexer := compiler.NewgramaticaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := compiler.NewgramaticaParser(stream)
	tree := p.Init()
	fmt.Println(tree.ToStringTree(p.RuleNames, p))

	// Crear listener desde el paquete listener (que debe embeder BasegramaticaListener)
	listener := listener.NewCustomListener()

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
}
