package visitor

import (
	parser "VAsm/backend/analizador/parser"
	"VAsm/frontend/symbols"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type EvalVisitor struct {
	*parser.BasegramaticaVisitor
	Console *strings.Builder
}

func NewEvalVisitor(tabla *symbols.TablaSimbolos) *EvalVisitor {
	return &EvalVisitor{
		BasegramaticaVisitor: &parser.BasegramaticaVisitor{},
		Console:              &strings.Builder{},
	}
}

func (v *EvalVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		return tree.Accept(v)
	}
}

// inicio de la gramatica
func (v *EvalVisitor) VisitInit(ctx *parser.InitContext) interface{} {
	var resultados []interface{}
	for _, instr := range ctx.AllInstrucciones() {
		res := v.Visit(instr)
		if res != nil {
			resultados = append(resultados, res)
		}
	}

	return resultados
}

// visitar instrucciones
func (v *EvalVisitor) VisitInstrucciones(ctx *parser.InstruccionesContext) interface{} {
	for _, inst := range ctx.AllInstruccion() {
		v.Visit(inst)
	}
	return nil
}

// visitar instruccion
func (v *EvalVisitor) VisitInstruccion(ctx *parser.InstruccionContext) interface{} {
	if ctx.Print_() != nil {
		return v.Visit(ctx.Print_())
	}
	return nil
}

// gramatica para imprimir
func (v *EvalVisitor) VisitPrint(ctx *parser.PrintContext) interface{} {
	var output strings.Builder
	expresiones := ctx.ListaExpr().AllExpresion()

	for i, expr := range expresiones {
		val := v.Visit(expr)
		fmt.Printf("[DEBUG PRINT] Argumento %d: Valor %#v (Tipo: %T)\n", i, val, val)
		output.WriteString(valorAString(val))
		if i < len(expresiones)-1 {
			output.WriteString(" ")
		}
	}

	v.Console.WriteString(output.String())
	return output.String()
}
func valorAString(val interface{}) string {
	switch v := val.(type) {
	case nil:
		return "<nil>"
	case string:
		return v
	case int, float64, bool:
		return fmt.Sprintf("%v", v)
	case []interface{}:
		var partes []string
		for _, elem := range v {
			partes = append(partes, valorAString(elem))
		}
		return "[" + strings.Join(partes, ", ") + "]"
	case [][]interface{}:
		var filas []string
		for _, fila := range v {
			if fila == nil {
				filas = append(filas, "<nil>")
				continue
			}
			var sub []string
			for _, elem := range fila {
				sub = append(sub, valorAString(elem))
			}
			filas = append(filas, "["+strings.Join(sub, ", ")+"]")
		}
		return "[" + strings.Join(filas, ", ") + "]"
	case map[string]interface{}:
		var campos []string
		for k, valCampo := range v {
			campos = append(campos, fmt.Sprintf("%s: %s", k, valorAString(valCampo)))
		}
		return "{" + strings.Join(campos, ", ") + "}"
	default:
		return fmt.Sprintf("<valor desconocido: %v (%T)>", v, v)
	}
}

// gramatica para las expresiones
func (v *EvalVisitor) VisitExpresion(ctx *parser.ExpresionContext) interface{} {
	/* EXPRESIONES NATIVAS */
	if ctx.CADENA() != nil {
		text := ctx.CADENA().GetText()
		return text[1 : len(text)-1]
	}

	// booleano true
	if ctx.TRUE() != nil {
		return true
	}

	// booleano false
	if ctx.FALSE() != nil {
		return false
	}

	if ctx.ENTERO() != nil {
		val, err := strconv.Atoi(ctx.ENTERO().GetText())
		if err != nil {
			panic("error al convertir entero: " + err.Error())
		}
		return val
	}

	if ctx.DECIMAL() != nil {
		val, err := strconv.ParseFloat(ctx.DECIMAL().GetText(), 64)
		if err != nil {
			panic("error al convertir decimal: " + err.Error())
		}
		return val
	}

	if ctx.PAR1() != nil && ctx.PAR2() != nil {
		return v.Visit(ctx.Expresion(0))
	}

	return 0
}
