package visitor

import (
	parser "VAsm/backend/analizador/parser"
	"VAsm/backend/analizador/traducciones"
	"fmt"
	"sync/atomic"

	"github.com/antlr4-go/antlr/v4"
)

var etiquetaCounter int64 = 0

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
	etiquetaElse := fmt.Sprintf("if_%d_else", idIf)

	// === IF principal ===
	traducciones.TextBuilder.WriteString(fmt.Sprintf("\n//  ====== IF ====== %d\n", idIf))
	generarCondicion(ctx.Expresion(), etiquetaElse)

	valCond := visit(ctx.Expresion())
	if b, ok := valCond.(bool); ok && b {
		if res := visit(ctx.Bloque(0)); res != nil {
			if _, ok := res.(BreakSignal); ok {
				return BreakSignal{}
			}
			if _, ok := res.(ContinueSignal); ok {
				return ContinueSignal{}
			}
			if list, ok := res.([]interface{}); ok && len(list) > 0 {
				switch list[len(list)-1].(type) {
				case BreakSignal:
					return BreakSignal{}
				case ContinueSignal:
					return ContinueSignal{}
				}
			}
		}
		if len(ctx.AllElseifPart()) > 0 || ctx.ELSE() != nil {
			traducciones.TextBuilder.WriteString(fmt.Sprintf("    b %s\n", etiquetaFin))
		}
		traducciones.TextBuilder.WriteString(fmt.Sprintf("%s:\n", etiquetaElse))
		traducciones.TextBuilder.WriteString(fmt.Sprintf("\n%s:\n", etiquetaFin))
		return nil
	} else {
		traducciones.TextBuilder.WriteString(fmt.Sprintf("%s:\n", etiquetaElse))
	}

	// === ELSE IFs ===
	cantidadElseIf := len(ctx.AllElseifPart())
	for i := 0; i < cantidadElseIf; i++ {
		elseif := ctx.ElseifPart(i)
		etiquetaElseIf := fmt.Sprintf("if_%d_else_if_%d", idIf, i+1)
		traducciones.TextBuilder.WriteString(fmt.Sprintf("\n// ELSE IF %d\n", i+1))
		generarCondicion(elseif.Expresion(), etiquetaElseIf)

		valElseIf := visit(elseif.Expresion())
		if b, ok := valElseIf.(bool); ok && b {
			if res := visit(elseif.Bloque()); res != nil {
				if _, ok := res.(BreakSignal); ok {
					return BreakSignal{}
				}
				if _, ok := res.(ContinueSignal); ok {
					return ContinueSignal{}
				}
				if list, ok := res.([]interface{}); ok && len(list) > 0 {
					switch list[len(list)-1].(type) {
					case BreakSignal:
						return BreakSignal{}
					case ContinueSignal:
						return ContinueSignal{}
					}
				}
			}
			traducciones.TextBuilder.WriteString(fmt.Sprintf("    b %s\n", etiquetaFin))
			traducciones.TextBuilder.WriteString(fmt.Sprintf("%s:\n", etiquetaElseIf))
			traducciones.TextBuilder.WriteString(fmt.Sprintf("\n%s:\n", etiquetaFin))
			return nil
		} else {
			traducciones.TextBuilder.WriteString(fmt.Sprintf("%s:\n", etiquetaElseIf))
		}
	}

	// === ELSE ===
	if ctx.ELSE() != nil {
		traducciones.TextBuilder.WriteString("\n// ELSE\n")
		visit(ctx.Bloque(cantidadElseIf + 1))
	}

	// === FIN ===
	traducciones.TextBuilder.WriteString(fmt.Sprintf("\n%s:\n", etiquetaFin))
	return nil
}
