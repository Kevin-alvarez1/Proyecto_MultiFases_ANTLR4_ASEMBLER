package visitor

import (
	parser "VAsm/backend/analizador/parser"
	"VAsm/backend/analizador/traducciones"
	"VAsm/frontend/symbols"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type EvalVisitor struct {
	*parser.BasegramaticaVisitor
	Console   *strings.Builder
	OutputASM strings.Builder
	Tabla     *symbols.TablaSimbolos
}

func NewEvalVisitor(tabla *symbols.TablaSimbolos) *EvalVisitor {
	return &EvalVisitor{
		BasegramaticaVisitor: &parser.BasegramaticaVisitor{},
		Tabla:                tabla,
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
	traducciones.ResetearCodigoASM()
	var resultados []interface{}
	for _, instr := range ctx.AllInstrucciones() {
		res := v.Visit(instr)
		if res != nil {
			resultados = append(resultados, res)
		}
	}
	v.GenerarASMFinal()
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
	if ctx.Println_() != nil {
		return v.Visit(ctx.Println_())
	}
	if ctx.DeclaracionMultiple() != nil {
		return v.Visit(ctx.DeclaracionMultiple())
	}
	return nil
}

// gramatica para asignaciones y declaraciones
func (v *EvalVisitor) VisitDeclaracionMultiple(ctx *parser.DeclaracionMultipleContext) interface{} {
	tipo := ctx.Tipos().GetText()
	valores := obtenerValores(ctx.ListaExpr(), v)
	ids := obtenerIDs(ctx.ListaIDS())
	fmt.Printf("Valores obtenidos: %v\n", valores)

	if len(ids) != len(valores) {
		msg := fmt.Sprintf("Error: cantidad de variables (%d) no coincide con valores (%d)",
			len(ids), len(valores))
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), msg)
		return nil
	}

	if err := traducciones.ProcesarDeclaracionMultiple(ids, valores, tipo, v.Tabla, &v.OutputASM); err != nil {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), err.Error())
	}

	return nil
}

// gramatica para imprimir
func (v *EvalVisitor) VisitPrint(ctx *parser.PrintContext) interface{} {
	expresiones := ctx.ListaExpr().AllExpresion()

	for _, expr := range expresiones {
		val := v.Visit(expr)
		texto := valorAString(val)

		// Generar ensamblador para ese texto
		traducciones.GenerarCodigoPrint(texto, false)
	}
	return nil
}

func (v *EvalVisitor) VisitPrintln(ctx *parser.PrintlnContext) interface{} {
	expresiones := ctx.ListaExpr().AllExpresion()

	for _, expr := range expresiones {
		val := v.Visit(expr)
		texto := valorAString(val)

		// Generar ensamblador con salto de línea explícito
		traducciones.GenerarCodigoPrint(texto, true)
	}
	return nil
}

func (v *EvalVisitor) VisitFnAtoi(ctx *parser.FnAtoiContext) interface{} {
	if ctx.ListaExpr() == nil {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "Atoi requiere un argumento")
		return nil
	}

	expresiones := ctx.ListaExpr().AllExpresion()
	if len(expresiones) != 1 {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "Atoi requiere exactamente un argumento")
		return nil
	}

	val := v.Visit(expresiones[0])

	strVal, ok := val.(string)
	if !ok {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Atoi espera un string, pero recibió: %T", val))
		return nil
	}

	if matched, _ := regexp.MatchString(`^-?\d+$`, strVal); !matched {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Atoi: la cadena '%s' no es un número entero válido", strVal))
		return nil
	}

	intVal, err := strconv.Atoi(strVal)
	if err != nil {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Atoi: error al convertir '%s' a entero", strVal))
		return nil
	}

	// ✅ Generar código ensamblador directamente
	asm := fmt.Sprintf(`// Atoi("%s") → %d
mov x0, #%d

`, strVal, intVal, intVal)
	v.OutputASM.WriteString(asm)

	return intVal
}

func (v *EvalVisitor) GenerarASMFinal() string {
	asm := traducciones.EmitirCodigoCompleto()
	v.OutputASM.WriteString(asm)

	fmt.Println("=== Código ensamblador generado ===")
	fmt.Println(v.OutputASM.String())
	v.Console.WriteString(v.OutputASM.String())

	return asm
}

// gramatica para las expresiones
func (v *EvalVisitor) VisitExpresion(ctx *parser.ExpresionContext) interface{} {
	/* EXPRESIONES NATIVAS */
	if ctx.IDENTIFICADOR() != nil {
		id := ctx.GetText()
		simbolo, existe := v.Tabla.EntornoActual.Simbolos[id]
		if !existe {
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "Variable no definida: "+id)
			return nil
		}
		return simbolo.Valor
	}

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
	// funciones embebidas

	if ctx.FnAtoi() != nil {
		return v.Visit(ctx.FnAtoi())
	}
	if ctx.FnParseToFloat() != nil {
		return v.Visit(ctx.FnParseToFloat())
	}
	if ctx.FnTypeOf() != nil {
		return v.Visit(ctx.FnTypeOf())
	}

	return 0
}

func obtenerValores(ctx parser.IListaExprContext, visitor *EvalVisitor) []interface{} {
	valores := []interface{}{}

	if ctx == nil {
		return valores
	}

	listaCtx, ok := ctx.(*parser.ListaExprContext)
	if !ok {
		return valores
	}

	// Evaluar todas las expresiones directamente
	for _, expr := range listaCtx.AllExpresion() {
		valores = append(valores, visitor.Visit(expr))
	}

	return valores
}

func valorAString(val interface{}) string {
	switch v := val.(type) {
	case int:
		return strconv.Itoa(v)
	case float64:
		return fmt.Sprintf("%f", v)
	case bool:
		return fmt.Sprintf("%t", v)
	case string:
		return v
	default:
		return fmt.Sprintf("%v", val)
	}
}

func obtenerIDs(ctx parser.IListaIDSContext) []string {
	ids := []string{}

	if ctx == nil {
		return ids
	}

	// Type assertion al contexto concreto
	listaCtx, ok := ctx.(*parser.ListaIDSContext)
	if !ok {
		return ids
	}

	// Obtener TODOS los identificadores directamente
	for _, idToken := range listaCtx.AllIDENTIFICADOR() {
		ids = append(ids, idToken.GetText())
	}

	return ids
}

type VisitContext interface {
	OutputASMBuilder() *strings.Builder
	TablaSimbolos() *symbols.TablaSimbolos
}

func (v *EvalVisitor) OutputASMBuilder() *strings.Builder {
	return &v.OutputASM
}

func (v *EvalVisitor) TablaSimbolos() *symbols.TablaSimbolos {
	return v.Tabla
}
