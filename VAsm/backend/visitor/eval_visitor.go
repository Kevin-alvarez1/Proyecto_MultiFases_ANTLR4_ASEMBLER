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

type VisitContext interface {
	OutputASMBuilder() *strings.Builder
	TablaSimbolos() *symbols.TablaSimbolos
}

type Funcion struct {
	Nombre               string
	TipoRetorno          string
	Parametros           []Parametro
	Bloque               antlr.ParseTree
	ValorRetorno         interface{}
	EntornoDeDeclaracion *symbols.Entorno
}

type Parametro struct {
	Nombre string
	Tipo   string
}

type Retorno struct {
	Valor interface{}
}

type EvalVisitor struct {
	*parser.BasegramaticaVisitor
	Entorno       map[string]interface{}
	Console       *strings.Builder
	OutputASM     strings.Builder
	DataSection   strings.Builder
	Tabla         *symbols.TablaSimbolos
	Funciones     map[string][]Funcion
	NombreEntorno string
}

var funcionesGeneradas map[string]bool

func NewEvalVisitor(tabla *symbols.TablaSimbolos) *EvalVisitor {
	return &EvalVisitor{
		BasegramaticaVisitor: &parser.BasegramaticaVisitor{},
		Tabla:                tabla,
		Console:              &strings.Builder{},
		Entorno:              make(map[string]interface{}),
		Funciones:            make(map[string][]Funcion),
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
	v.Funciones = make(map[string][]Funcion)
	contadorBloques = 0
	var resultados []interface{}
	for _, instr := range ctx.AllInstrucciones() {
		if esFuncion(instr) {
			v.Visit(instr)
		} else {
			res := v.Visit(instr)
			if res != nil {
				resultados = append(resultados, res)
			}
		}
	}
	v.GenerarASMFinal()
	return resultados
}

func esFuncion(instr antlr.ParseTree) bool {
	switch instr.(type) {
	case *parser.FnSinParametroContext:
		return true
	case *parser.FnConParametroContext:
		return true
	default:
		return false
	}
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
	if ctx.FnSinParametro() != nil {
		return v.Visit(ctx.FnSinParametro())
	}
	if ctx.FnConParametro() != nil {
		return v.Visit(ctx.FnConParametro())
	}
	if ctx.LlamadaFuncionesSinParametro() != nil {
		return v.Visit(ctx.LlamadaFuncionesSinParametro())
	}
	if ctx.LlamadaFuncionesConParametro() != nil {
		return v.Visit(ctx.LlamadaFuncionesConParametro())
	}
	if ctx.Retorno() != nil {
		res := v.Visit(ctx.Retorno())
		return res
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
func (v *EvalVisitor) VisitFnSinParametro(ctx *parser.FnSinParametroContext) interface{} {
	nombre := ctx.IDENTIFICADOR().GetText()

	if _, existe := v.Funciones[nombre]; existe {
		panic("Función ya declarada: " + nombre)
	}

	tipoRetorno := "void"
	if ctx.TipoRetorno() != nil {
		tipoRetorno = ctx.TipoRetorno().GetText()
	}

	// Crear entorno único para la función: bloque_funcion_p
	entornoFuncion := symbols.NewEntorno(v.Tabla.EntornoActual, "bloque_funcion_"+nombre)
	v.Tabla.EntornoActual.Hijos = append(v.Tabla.EntornoActual.Hijos, entornoFuncion)

	fn := Funcion{
		Nombre:               nombre,
		Bloque:               ctx.BloqueFuncion(),
		TipoRetorno:          tipoRetorno,
		ValorRetorno:         nil,
		EntornoDeDeclaracion: entornoFuncion, // se guarda el entorno aquí
	}

	v.Funciones[nombre] = []Funcion{fn}

	// === Generar ensamblador ===
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("\n.global fn_%s\nfn_%s:\n", nombre, nombre))
	builder.WriteString("    stp x29, x30, [sp, #-16]!\n")
	builder.WriteString("    mov x29, sp\n")
	builder.WriteString("    mov sp, x29\n")

	// Guardar TextBuilder original
	textoGlobal := traducciones.TextBuilder
	traducciones.TextBuilder = strings.Builder{}

	entAnt := v.Tabla.EntornoActual
	v.Tabla.EntornoActual = entornoFuncion

	nombreAnt := v.NombreEntorno
	v.NombreEntorno = "bloque_funcion_" + nombre

	var retorno interface{}
	if fn.Bloque != nil {
		retorno = v.Visit(fn.Bloque)
	}

	// Si hay retorno explícito
	if ret, ok := retorno.(Retorno); ok && ret.Valor != nil {
	}

	builder.WriteString(traducciones.TextBuilder.String())
	builder.WriteString("    ldp x29, x30, [sp], #16\n")
	builder.WriteString("    ret\n")

	// Restaurar contexto
	traducciones.TextBuilder = textoGlobal
	v.Tabla.EntornoActual = entAnt
	v.NombreEntorno = nombreAnt

	// Agregar código al builder final
	traducciones.FuncionesBuilder.WriteString(builder.String())

	return nil
}

func (v *EvalVisitor) VisitLlamadaFuncionesSinParametro(ctx *parser.LlamadaFuncionesSinParametroContext) interface{} {
	nombre := ctx.IDENTIFICADOR().GetText()

	if _, existe := v.Funciones[nombre]; !existe {
		panic("Función no definida: " + nombre)
	}

	// Generar ensamblador para llamar a la función
	traducciones.TextBuilder.WriteString(fmt.Sprintf("    bl fn_%s\n", nombre))
	return nil
}

func (v *EvalVisitor) VisitFnConParametro(ctx *parser.FnConParametroContext) interface{} {
	nombre := ctx.IDENTIFICADOR().GetText()
	fmt.Printf("[VisitFnConParametro] Definiendo función '%s'\n", nombre)

	// Extraer lista de parámetros (de tipo []Parametro)
	parametros := ExtraerParametros(ctx.ListaPar())

	if _, existe := v.Funciones[nombre]; existe {
		panic(fmt.Sprintf("[VisitFnConParametro] ERROR: Función '%s' ya declarada", nombre))
	}

	tipoRetorno := "void"
	if ctx.TipoRetorno() != nil {
		tipoRetorno = ctx.TipoRetorno().GetText()
	}

	// Crear entorno para la declaración (para almacenar variables globales o contexto estático)
	entornoFuncion := symbols.NewEntorno(v.Tabla.EntornoActual, "bloque_funcion_"+nombre)
	v.Tabla.EntornoActual.Hijos = append(v.Tabla.EntornoActual.Hijos, entornoFuncion)

	fn := Funcion{
		Nombre:               nombre,
		Bloque:               ctx.BloqueFuncion(),
		TipoRetorno:          tipoRetorno,
		Parametros:           parametros,
		EntornoDeDeclaracion: entornoFuncion,
	}

	v.Funciones[nombre] = []Funcion{fn}

	fmt.Printf("[VisitFnConParametro] Función '%s' registrada con tipo de retorno '%s' y %d parámetro(s)\n",
		nombre, tipoRetorno, len(parametros))

	// === Generar ensamblador sólo de la firma de la función ===
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("\n.global fn_%s\nfn_%s:\n", nombre, nombre))
	builder.WriteString("    stp x29, x30, [sp, #-16]!\n")
	builder.WriteString("    mov x29, sp\n")
	builder.WriteString("    mov sp, x29\n")

	builder.WriteString("    ldp x29, x30, [sp], #16\n")
	builder.WriteString("    ret\n")

	traducciones.FuncionesBuilder.WriteString(builder.String())

	return nil
}

func (v *EvalVisitor) VisitLlamadaFuncionesConParametro(ctx *parser.LlamadaFuncionesConParametroContext) interface{} {
	nombre := ctx.IDENTIFICADOR().GetText()

	funciones, existe := v.Funciones[nombre]
	if !existe {
		panic("Función no definida: " + nombre)
	}
	funcion := funciones[0]

	entornoAnterior := v.Tabla.EntornoActual
	nombreAnterior := v.NombreEntorno

	entornoEjecucion := symbols.NewEntorno(entornoAnterior, "bloque_funcion_ejecucion_"+nombre)
	entornoAnterior.Hijos = append(entornoAnterior.Hijos, entornoEjecucion)

	v.Tabla.EntornoActual = entornoEjecucion
	v.NombreEntorno = "bloque_funcion_ejecucion_" + nombre

	defer func() {
		v.Tabla.EntornoActual = entornoAnterior
		v.NombreEntorno = nombreAnterior
	}()

	// Validar y declarar parámetros
	parametros := funcion.Parametros
	args := ctx.ListaExpr().AllExpresion()

	if len(parametros) != len(args) {
		panic(fmt.Sprintf("Error: función '%s' esperaba %d parámetros, recibió %d", nombre, len(parametros), len(args)))
	}

	fmt.Println("=== Parámetros declarados ===")
	for i, par := range parametros {
		id := par.Nombre
		tipo := par.Tipo
		valor := v.Visit(args[i])

		fmt.Printf("-> %s: %v (%s)\n", id, valor, tipo)
		v.Tabla.DeclararParametro(id, tipo, valor, ctx, v.NombreEntorno)

		// verificar si se guardó
		sim, ok := v.Tabla.EntornoActual.Simbolos[id]
		if ok {
			fmt.Printf("[Debug] Parámetro '%s' registrado como tipo '%s' con valor '%v'\n", sim.ID, sim.TipoSimbolo, sim.Valor)
		} else {
			fmt.Printf("[Debug] Falló el registro del parámetro '%s'\n", id)
		}
	}
	fmt.Println("=============================")

	// Ejecutar cuerpo de la función
	resultado := v.Visit(funcion.Bloque)

	if ret, ok := resultado.(Retorno); ok {
		return ret.Valor
	}

	return nil
}

func (v *EvalVisitor) VisitBloqueFuncion(ctx *parser.BloqueFuncionContext) interface{} {
	entornoAnterior := v.Tabla.EntornoActual
	nombreAnterior := v.NombreEntorno

	// Solo crear nuevo entorno si no estás en ejecución de función
	if !strings.HasPrefix(v.Tabla.EntornoActual.Nombre, "bloque_funcion_ejecucion_") {
		nombreBase := nombreAnterior
		nombreBase = strings.TrimPrefix(nombreBase, "bloque_funcion_")
		nombreBase = strings.TrimPrefix(nombreBase, "temp_ejecucion_")
		nombreBase = strings.TrimPrefix(nombreBase, "fn_")

		nombreBloque := "bloque_funcion_" + nombreBase

		nuevo := symbols.NewEntorno(entornoAnterior, nombreBloque)
		entornoAnterior.Hijos = append(entornoAnterior.Hijos, nuevo)

		v.Tabla.EntornoActual = nuevo
		v.NombreEntorno = nombreBloque
	} else {
		fmt.Printf("[VisitBloqueFuncion] Ya estás en entorno válido: %s\n", v.Tabla.EntornoActual.Nombre)
	}

	defer func() {
		v.Tabla.EntornoActual = entornoAnterior
		v.NombreEntorno = nombreAnterior
	}()

	fmt.Println("=== [VisitBloqueFuncion] Parámetros en el entorno actual ===")
	for id, sim := range v.Tabla.EntornoActual.Simbolos {
		if sim.TipoSimbolo == symbols.Parametro {
			fmt.Printf("-> %s = %v (tipo %s)\n", id, sim.Valor, sim.TipoDato)
		}
	}
	fmt.Println("======================================")

	var resultado interface{}
	for _, instr := range ctx.AllInstruccion() {
		res := v.Visit(instr)
		if ret, ok := res.(Retorno); ok {
			resultado = ret
			break
		}
	}

	if resultado == nil {
		for _, expr := range ctx.AllExpresion() {
			res := v.Visit(expr)
			if ret, ok := res.(Retorno); ok {
				resultado = ret
				break
			}
		}
	}

	return resultado
}

func (v *EvalVisitor) VisitRetorno(ctx *parser.RetornoContext) interface{} {
	if ctx.Expresion() != nil {
		val := v.Visit(ctx.Expresion())
		return Retorno{Valor: val}
	}
	return Retorno{Valor: nil}
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

// ========= GENERADOR DE ASM =========
func (v *EvalVisitor) GenerarASMFinal() string {
	// EmitirCodigoCompleto debe armar la sección .data + globales + funciones + _start con el punto de entrada correcto
	asm := EmitirCodigoCompleto(v.NombreEntorno)

	v.OutputASM.Reset()
	v.OutputASM.WriteString(asm)

	fmt.Println("=== Código ensamblador generado ===")
	fmt.Println(v.OutputASM.String())
	v.Console.WriteString(v.OutputASM.String())

	return asm
}

// ========= EXPRESIONES =========
func (v *EvalVisitor) VisitExpresion(ctx *parser.ExpresionContext) interface{} {
	if ctx.LlamadaFuncionesSinParametro() != nil {
		nombreFuncion := ctx.LlamadaFuncionesSinParametro().IDENTIFICADOR().GetText()

		if funciones, existe := v.Funciones[nombreFuncion]; existe {
			funcion := funciones[0]

			// Guardar entorno actual
			entornoAnterior := v.Tabla.EntornoActual
			nombreAnterior := v.NombreEntorno

			// Solo cambia el contexto, sin crear nuevo entorno visible
			v.Tabla.EntornoActual = symbols.NewEntorno(nil, "temp_ejecucion_"+nombreFuncion)
			v.NombreEntorno = "temp_ejecucion_" + nombreFuncion

			// Ejecutar cuerpo de la función
			var resultado interface{}
			if funcion.Bloque != nil {
				resultado = v.Visit(funcion.Bloque)
			}

			// Restaurar entorno anterior
			v.Tabla.EntornoActual = entornoAnterior
			v.NombreEntorno = nombreAnterior

			// Devolver el valor si hay retorno
			if ret, ok := resultado.(Retorno); ok {
				return ret.Valor
			}
			return nil
		}

		panic("Función no definida: " + nombreFuncion)
	}
	if ctx.LlamadaFuncionesConParametro() != nil {
		return v.VisitLlamadaFuncionesConParametro(ctx.LlamadaFuncionesConParametro().(*parser.LlamadaFuncionesConParametroContext))
	}

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

func GenerarASMFuncion(fn Funcion, v *EvalVisitor) {
	asm := fmt.Sprintf("\n # Función: %s\n", fn.Nombre)
	asm += fmt.Sprintf(".global fn_%s\n", fn.Nombre)
	asm += fmt.Sprintf("fn_%s:\n", fn.Nombre)
	asm += "    stp x29, x30, [sp, #-16]!\n"
	asm += "    mov x29, sp\n"

	if fn.Bloque != nil {
		if codigo, ok := v.Visit(fn.Bloque).(string); ok {
			asm += codigo
		}
	}

	asm += "    mov sp, x29\n"
	asm += "    ldp x29, x30, [sp], #16\n"
	asm += "    ret\n"

	traducciones.TextBuilder.WriteString(asm)
}

func EmitirCodigoCompleto(puntoEntrada string) string {
	var startCode strings.Builder

	startCode.WriteString(".global _start\n")
	startCode.WriteString(".text\n")
	startCode.WriteString("_start:\n")

	// Primero inserta TODO el código global fuera de funciones
	startCode.WriteString(traducciones.TextBuilder.String())

	// Luego, salto a la función principal si existe
	if puntoEntrada != "" && puntoEntrada != "global" {
		startCode.WriteString(fmt.Sprintf("    bl fn_%s\n", puntoEntrada))
	}

	// Código para salir del programa
	startCode.WriteString("    mov x0, #0\n")
	startCode.WriteString("    mov x8, #93\n")
	startCode.WriteString("    svc #0\n\n")

	// Finalmente agrega todas las funciones juntas
	startCode.WriteString(traducciones.FuncionesBuilder.String())

	return ".data\n" + traducciones.DataBuilder.String() + "\n" + startCode.String()
}

func ExtraerParametros(ctx parser.IListaParContext) []Parametro {
	var parametros []Parametro
	if ctx == nil {
		return parametros
	}
	for _, p := range ctx.AllParametro() {
		nombre := p.IDENTIFICADOR().GetText()
		tipo := p.Tipos().GetText()
		parametros = append(parametros, Parametro{Nombre: nombre, Tipo: tipo})
	}
	return parametros
}
