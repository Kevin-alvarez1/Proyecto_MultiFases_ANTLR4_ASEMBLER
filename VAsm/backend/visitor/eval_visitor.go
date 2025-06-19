package visitor

import (
	parser "VAsm/backend/analizador/parser"
	"fmt"
	"log"
	"strconv"
	"github.com/antlr4-go/antlr/v4"
)

type EvalVisitor struct {
	*parser.BasegramaticaVisitor
}

func NewEvalVisitor() *EvalVisitor {
	return &EvalVisitor{
		BasegramaticaVisitor: &parser.BasegramaticaVisitor{},
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
	expresiones := ctx.ListaExpr().AllExpresion()
	fmt.Println(expresiones)
	return nil
}


// gramatica para las expresiones
func (v *EvalVisitor) VisitExpresion(ctx *parser.ExpresionContext) interface{} {
		/* EXPRESIONES NATIVAS */
	if ctx.CADENA() != nil {
		text := ctx.CADENA().GetText()
		return text[1 : len(text)-1]
	}

	// booleano true
	if ctx.TRUE() != nil {return true }
	
	// booleano false
	if ctx.FALSE() != nil { return false }

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