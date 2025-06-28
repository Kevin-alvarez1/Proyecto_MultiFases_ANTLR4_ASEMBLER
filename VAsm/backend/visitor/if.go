package visitor

import (
	parser "VAsm/backend/analizador/parser"
	"VAsm/backend/analizador/traducciones"
	"fmt"
	"sync/atomic"

	"github.com/antlr4-go/antlr/v4"
)

var etiquetaCounter int64 = 0

func nuevaEtiqueta(prefix string) string {
	id := atomic.AddInt64(&etiquetaCounter, 1)
	return fmt.Sprintf("%s_%d", prefix, id)
}

type If struct{}

func NewIf() *If {
	return &If{}
}
func (i *If) Ejecutar(
	visit func(tree antlr.Tree) interface{},
	generarCondicion func(ctx parser.IExpresionContext, etiqueta string),
	ctx *parser.SifContext,
) interface{} {
	idIf := atomic.AddInt64(&etiquetaCounter, 1)
	etiquetaFin := fmt.Sprintf("if_%d_end", idIf)

	// --- IF principal ---
	if ctx.Expresion() != nil {
		etiquetaElse := fmt.Sprintf("if_%d_else", idIf)
		traducciones.TextBuilder.WriteString(fmt.Sprintf("\n//  ====== IF ====== %d\n", idIf))
		generarCondicion(ctx.Expresion(), etiquetaElse)

		// Bloque if verdadero
		visit(ctx.Bloque(0))

		if len(ctx.AllElseifPart()) > 0 || ctx.ELSE() != nil {
			traducciones.TextBuilder.WriteString(fmt.Sprintf("    b %s\n", etiquetaFin))
		}

		traducciones.TextBuilder.WriteString(fmt.Sprintf("%s:\n", etiquetaElse))
	}

	// --- ELSE IFs ---
	cantidadElseIf := len(ctx.AllElseifPart())
	for i := 0; i < cantidadElseIf; i++ {
		elseif := ctx.ElseifPart(i)
		etiquetaElseIf := fmt.Sprintf("if_%d_else_if_%d", idIf, i+1)

		traducciones.TextBuilder.WriteString(fmt.Sprintf("\n// ELSE IF %d\n", i+1))
		generarCondicion(elseif.Expresion(), etiquetaElseIf)

		visit(elseif.Bloque())
		traducciones.TextBuilder.WriteString(fmt.Sprintf("    b %s\n", etiquetaFin))
		traducciones.TextBuilder.WriteString(fmt.Sprintf("%s:\n", etiquetaElseIf))
	}

	// --- ELSE ---
	if ctx.ELSE() != nil {
		traducciones.TextBuilder.WriteString("\n// ELSE\n")
		visit(ctx.Bloque(cantidadElseIf + 1))
	}

	// --- FIN IF ---
	traducciones.TextBuilder.WriteString(fmt.Sprintf("\n%s:\n", etiquetaFin))
	return nil
}
