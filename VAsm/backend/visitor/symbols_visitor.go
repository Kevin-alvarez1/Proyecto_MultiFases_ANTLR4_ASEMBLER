package visitor

import (
	compiler "VAsm/backend/analizador/parser"
	"VAsm/frontend/symbols"
)

// tabla de simbolos visitor
type SymbolVisitor struct {
	*compiler.BasegramaticaVisitor
	Table   *symbols.TablaSimbolos
	Entorno map[string]*symbols.Simbolo
}

func NewSymbolVisitor(table *symbols.TablaSimbolos) *SymbolVisitor {
	return &SymbolVisitor{
		BasegramaticaVisitor: &compiler.BasegramaticaVisitor{},
		Table:                table,
		Entorno:              make(map[string]*symbols.Simbolo),
	}
}
func (v *SymbolVisitor) VisitVariableDecl(ctx *compiler.DeclaracionMultipleContext) interface{} {
	id := ctx.GetText()
	tipo := ctx.GetText()
	valorNode := ctx.ListaExpr()

	var valor interface{}
	if valorNode != nil {
		valor = valorNode.Accept(v) // Evalúa la expresión si implementaste eso
	}

	v.Table.DeclararVariable(id, tipo, valor, ctx, v.Table.EntornoActual.Nombre)
	return nil
}

func (v *SymbolVisitor) VisitFuncDecl(ctx *compiler.DeclaracionMultipleSimpleContext) interface{} {
	id := ctx.GetText()
	tipoRetorno := ctx.GetText()

	v.Table.DeclararFuncion(id, tipoRetorno, ctx)

	v.EntrarEntorno(id)
	defer v.SalirEntorno()

	return v.VisitChildren(ctx)
}

// EntrarEntorno crea un nuevo entorno en la tabla de símbolos
func (v *SymbolVisitor) actualizarEntorno() {
	v.Entorno = make(map[string]*symbols.Simbolo)
	for _, s := range v.Table.EntornoActual.Simbolos {
		v.Entorno[s.ID] = s
	}
}

func (v *SymbolVisitor) EntrarEntorno(nombre string) {
	v.Table.NuevoEntorno(nombre)
	v.actualizarEntorno()
}

func (v *SymbolVisitor) SalirEntorno() {
	v.Table.SalirEntorno()
	v.actualizarEntorno()
}
