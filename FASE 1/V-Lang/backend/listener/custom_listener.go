package listener

import (
	compiler "vlang/backend/analizador/parser"
)

type CustomListener struct {
	*compiler.BasegramaticaListener
}

func NewCustomListener() *CustomListener {
	return &CustomListener{
		BasegramaticaListener: &compiler.BasegramaticaListener{},
	}
}

func (l *CustomListener) EnterPrint(ctx *compiler.PrintContext) {

}
