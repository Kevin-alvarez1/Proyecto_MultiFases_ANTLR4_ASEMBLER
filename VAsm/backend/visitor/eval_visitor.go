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

var contadorBloques int = 0

type EvalVisitor struct {
	*parser.BasegramaticaVisitor
	Entorno   map[string]interface{}
	Console   *strings.Builder
	OutputASM strings.Builder
	Tabla     *symbols.TablaSimbolos
}

func NewEvalVisitor(tabla *symbols.TablaSimbolos) *EvalVisitor {
	return &EvalVisitor{
		BasegramaticaVisitor: &parser.BasegramaticaVisitor{},
		Tabla:                tabla,
		Console:              &strings.Builder{},
		Entorno:              make(map[string]interface{}),
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
	if ctx.DeclaracionMultipleSimple() != nil {
		return v.Visit(ctx.DeclaracionMultipleSimple())
	}
	if ctx.DeclaracionMultipleSinTipo() != nil {
		return v.Visit(ctx.DeclaracionMultipleSinTipo())
	}
	if ctx.AsignacionMultiple() != nil {
		return v.Visit(ctx.AsignacionMultiple())
	}
	if ctx.Bloque() != nil {
		return v.Visit(ctx.Bloque())
	}
	return nil
}

// ========= DECLARACIONES =========
func (v *EvalVisitor) VisitDeclaracionMultiple(ctx *parser.DeclaracionMultipleContext) interface{} {
	tipo := ctx.Tipos().GetText()
	valores := obtenerValores(ctx.ListaExpr(), v)
	ids := obtenerIDs(ctx.ListaIDS())

	if len(ids) != len(valores) {
		msg := fmt.Sprintf("Error: cantidad de variables (%d) no coincide con valores (%d)",
			len(ids), len(valores))
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), msg)
		return nil
	}

	for _, id := range ids {
		if yaDeclaradoEnAmbitoActual(id, tipo, v.Tabla) {
			msg := fmt.Sprintf("Error: variable '%s' ya fue declarada con otro tipo en este ámbito", id)
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), msg)
			return nil
		}
	}

	if err := traducciones.ProcesarDeclaracionMultiple(ids, valores, tipo, v.Tabla, &v.OutputASM); err != nil {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), err.Error())
	}
	return nil
}

func (v *EvalVisitor) VisitDeclaracionMultipleSimple(ctx *parser.DeclaracionMultipleSimpleContext) interface{} {
	if ctx.Tipos() == nil || ctx.ListaIDS() == nil {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "Declaración incompleta")
		return nil
	}

	ids := obtenerIDs(ctx.ListaIDS())
	tipo := ctx.Tipos().GetText()

	var valorDefault interface{}
	switch strings.ToLower(tipo) {
	case "int":
		valorDefault = 0
	case "float":
		valorDefault = 0.0
	case "string":
		valorDefault = ""
	case "bool":
		valorDefault = false
	default:
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "Tipo desconocido: "+tipo)
		return nil
	}

	// Verificación de cada variable
	for _, id := range ids {
		if yaDeclaradoEnAmbitoActual(id, tipo, v.Tabla) {
			msg := fmt.Sprintf("Error: variable '%s' ya fue declarada con otro tipo en este ámbito", id)
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), msg)
			return nil
		}
	}

	valores := make([]interface{}, len(ids))
	for i := range valores {
		valores[i] = valorDefault
	}

	if err := traducciones.ProcesarDeclaracionMultiple(ids, valores, tipo, v.Tabla, &v.OutputASM); err != nil {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), err.Error())
	}

	return nil
}

func (v *EvalVisitor) VisitDeclaracionMultipleSinTipo(ctx *parser.DeclaracionMultipleSinTipoContext) interface{} {
	listaExpr := ctx.ListaExpr()
	valores := obtenerValores(listaExpr, v)
	ids := obtenerIDs(ctx.ListaIDS())

	if len(ids) != len(valores) {
		msg := fmt.Sprintf("Error semántico: cantidad de variables (%d) no coincide con valores (%d)",
			len(ids), len(valores))
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), msg)
		return nil
	}

	for i, id := range ids {
		valor := valores[i]
		tipo := inferirTipo(valor)
		if yaDeclaradoEnAmbitoActual(id, tipo, v.Tabla) {
			msg := fmt.Sprintf("Error: variable '%s' ya fue declarada con otro tipo en este ámbito", id)
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), msg)
			continue
		}
		// Guardar variable y tipo en entorno y tabla
		v.Entorno[id] = valor
		v.Tabla.DeclararVariable(id, tipo, valor, ctx, v.Tabla.EntornoActual.Nombre)

		err := traducciones.GenerarCodigoDeclaracionSinTipo(id, tipo, valor, &v.OutputASM)
		if err != nil {
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), err.Error())
		}

		fmt.Printf("Variable declarada: %s tipo: %s con valor: %v\n", id, tipo, valor)
	}

	return nil
}

// ========= ASIGNACIONES =========
func (v *EvalVisitor) VisitAsignacionMultiple(ctx *parser.AsignacionMultipleContext) interface{} {
	ids := obtenerIDs(ctx.ListaIDS())
	valores := obtenerValores(ctx.ListaExpr(), v)

	if len(ids) != len(valores) {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "Número de variables y valores no coincide en asignación múltiple")
		return nil
	}

	for i, id := range ids {
		valor := valores[i]
		valorInferido := inferirTipo(valor)

		simbolo := v.Tabla.EntornoActual.BuscarSimbolo(id)
		if simbolo == nil {
			msg := fmt.Sprintf("Error: variable '%s' no ha sido declarada en este ámbito", id)
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), msg)
			continue
		}

		if !tiposCompatibles(simbolo.TipoDato, valorInferido) {
			msg := fmt.Sprintf("Error: tipo incompatible en asignación. '%s' es de tipo '%s', no '%s'",
				id, simbolo.TipoDato, valorInferido)
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), msg)
			continue
		}

		// Actualización
		simbolo.Valor = valor
		v.Entorno[id] = valor
		fmt.Printf("Asignación: %s = %v (tipo: %s)\n", id, valor, valorInferido)

		// Generar ASM (opcional)
		err := traducciones.GenerarCodigoDeclaracionSinTipo(id, simbolo.TipoDato, valor, &v.OutputASM)
		if err != nil {
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), err.Error())
		}
	}

	return nil
}

// ========= FUNCIONES =========
func (v *EvalVisitor) VisitBloque(ctx *parser.BloqueContext) interface{} {
	// Crear un nuevo entorno hijo
	nuevo := symbols.NewEntorno(v.Tabla.EntornoActual, generarNombreUnico("bloque", v.Tabla.EntornoActual.Nombre))

	v.Tabla.EntornoActual.Hijos = append(v.Tabla.EntornoActual.Hijos, nuevo)

	entornoAnterior := v.Tabla.EntornoActual
	v.Tabla.EntornoActual = nuevo

	// Visitar todas las instrucciones y expresiones dentro del bloque
	for _, instr := range ctx.AllInstruccion() {
		v.Visit(instr)
	}
	for _, expr := range ctx.AllExpresion() {
		v.Visit(expr)
	}

	// Restaurar entorno anterior
	v.Tabla.EntornoActual = entornoAnterior

	return nil
}

// ========= FUNCIONES EMBEBIDAS =========
// gramatica para imprimir
func (v *EvalVisitor) VisitPrint(ctx *parser.PrintContext) interface{} {
	expresiones := ctx.ListaExpr().AllExpresion()

	for _, expr := range expresiones {
		val := v.Visit(expr)
		texto := valorAString(val)

		traducciones.GenerarCodigoPrint(texto, false)
	}
	return nil
}

func (v *EvalVisitor) VisitPrintln(ctx *parser.PrintlnContext) interface{} {
	expresiones := ctx.ListaExpr().AllExpresion()

	for _, expr := range expresiones {
		val := v.Visit(expr)
		texto := valorAString(val)

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

	asm := fmt.Sprintf(`// Atoi("%s") → %d
mov x0, #%d

`, strVal, intVal, intVal)
	v.OutputASM.WriteString(asm)

	return intVal
}

func (v *EvalVisitor) VisitFnParseToFloat(ctx *parser.FnParseToFloatContext) interface{} {
	if ctx.ListaExpr() == nil {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "ParseFloat requiere un argumento")
		return nil
	}

	expresiones := ctx.ListaExpr().AllExpresion()
	if len(expresiones) != 1 {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "ParseFloat requiere exactamente un argumento")
		return nil
	}

	val := v.Visit(expresiones[0])
	strVal, ok := val.(string)
	if !ok {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), fmt.Sprintf("ParseFloat espera un string, pero recibió: %T", val))
		return nil
	}

	if matched, _ := regexp.MatchString(`^-?\d+(\.\d+)?$`, strVal); !matched {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), fmt.Sprintf("ParseFloat: '%s' no es un número flotante válido", strVal))
		return nil
	}

	floatVal, err := strconv.ParseFloat(strVal, 64)
	if err != nil {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), fmt.Sprintf("ParseFloat: error al convertir '%s' a float", strVal))
		return nil
	}

	traducciones.GenerarCodigoParseFloat(strVal, floatVal, &v.OutputASM)

	return floatVal
}

func (v *EvalVisitor) VisitFnTypeOf(ctx *parser.FnTypeOfContext) interface{} {
	if ctx.ListaExpr() == nil {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "TypeOf requiere un argumento")
		return "undefined"
	}

	expresiones := ctx.ListaExpr().AllExpresion()
	if len(expresiones) != 1 {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "TypeOf requiere exactamente un argumento")
		return "undefined"
	}

	val := v.Visit(expresiones[0])
	if val == nil {
		return "nil"
	}

	tipo := inferirTipo(val)

	traducciones.GenerarCodigoTypeOf(tipo)

	return tipo
}

// ========= GENERADOR DE ASM =========

func (v *EvalVisitor) GenerarASMFinal() string {
	asm := traducciones.EmitirCodigoCompleto()
	v.OutputASM.WriteString(asm)

	fmt.Println("=== Código ensamblador generado ===")
	fmt.Println(v.OutputASM.String())
	v.Console.WriteString(v.OutputASM.String())

	return asm
}

// ========= EXPRESIONES =========
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

// ========= FUNCIONES AUXILIARES =========

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
func inferirTipo(valor interface{}) string {
	switch v := valor.(type) {
	case int:
		return "int"
	case float64:
		return "float"
	case string:
		return "string"
	case bool:
		return "bool"
	case []interface{}:
		if len(v) == 0 {
			return "slice(unknown)"
		}
		return "slice(" + inferirTipo(v[0]) + ")"
	case [][]interface{}:
		if len(v) == 0 || len(v[0]) == 0 {
			return "slice(slice(unknown))"
		}
		return "slice(slice(" + inferirTipo(v[0][0]) + "))"
	default:
		return fmt.Sprintf("%T", valor)
	}
}
func yaDeclaradoEnAmbitoActual(id string, tipo string, tabla *symbols.TablaSimbolos) bool {
	simbolo, existe := tabla.EntornoActual.Simbolos[id]
	if !existe {
		return false
	}
	// Si ya existe y el tipo es distinto, entonces no es válido
	return simbolo.TipoDato != tipo
}

func tiposCompatibles(tipoVar string, tipoValor string) bool {
	if tipoVar == tipoValor {
		return true
	}
	// Permitir asignar int a float
	if tipoVar == "float64" && tipoValor == "int" {
		return true
	}

	return false
}

func generarNombreUnico(base string, padre string) string {
	contadorBloques++
	return fmt.Sprintf("%s_%s_%d", base, padre, contadorBloques)
}
