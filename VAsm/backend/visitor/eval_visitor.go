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
type BreakSignal struct{}
type ContinueSignal struct{}

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
	IfCtrl        *If
}

var funcionesGeneradas map[string]bool

func NewEvalVisitor(tabla *symbols.TablaSimbolos) *EvalVisitor {
	return &EvalVisitor{
		BasegramaticaVisitor: &parser.BasegramaticaVisitor{},
		Tabla:                tabla,
		Console:              &strings.Builder{},
		Entorno:              make(map[string]interface{}),
		Funciones:            make(map[string][]Funcion),
		IfCtrl:               NewIf(),
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
	if ctx.AsigIncre() != nil {
		v.Visit(ctx.AsigIncre())
		return nil
	}
	if ctx.AsigDecre() != nil {
		v.Visit(ctx.AsigDecre())
		return nil
	}
	if ctx.Asignacion() != nil {
		return v.Visit(ctx.Asignacion())
	}

	if ctx.Sif() != nil {
		return v.Visit(ctx.Sif())
	}
	if ctx.SSwitch() != nil {
		return v.Visit(ctx.SSwitch())
	}
	if ctx.Sfor() != nil {
		return v.Visit(ctx.Sfor())
	}
	if ctx.Break_() != nil {
		return BreakSignal{}
	}
	if ctx.Continue_() != nil {
		return ContinueSignal{}
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

// ========= INCRE Y DECRE =========
func (v *EvalVisitor) VisitAsigIncre(ctx *parser.AsigIncreContext) interface{} {
	id := ctx.IDENTIFICADOR().GetText()
	valorExpr := ctx.Expresion()
	valor := v.Visit(valorExpr)

	simbolo := v.Tabla.EntornoActual.BuscarSimbolo(id)
	if simbolo == nil {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Variable '%s' no declarada", id))
		return nil
	}

	tipoSimbolo := simbolo.TipoDato
	tipoValor := inferirTipo(valor)

	var tipoFinal string
	if tipoSimbolo == "float" || tipoValor == "float" {
		tipoFinal = "float"
	} else {
		tipoFinal = "int"
	}

	traducciones.ReservarVariableSiNoExiste(id, tipoFinal)

	v.OutputASM.WriteString(fmt.Sprintf("// %s += %v\n", id, valor))

	if tipoFinal == "float" {
		v.OutputASM.WriteString(fmt.Sprintf("adr x10, %s\n", id))
		v.OutputASM.WriteString("ldr s0, [x10]\n")

		if tipoValor == "int" {
			v.OutputASM.WriteString(fmt.Sprintf("mov x11, #%d\n", valor.(int)))
			v.OutputASM.WriteString("scvtf s1, x11\n")
		} else {
			v.OutputASM.WriteString(fmt.Sprintf("ldr s1, =%g\n", valor.(float64)))
		}

		v.OutputASM.WriteString("fadd s2, s0, s1\n")

		v.OutputASM.WriteString(fmt.Sprintf("adr x12, %s\n", id))
		v.OutputASM.WriteString("str s2, [x12]\n\n")

		simbolo.Valor = toFloat(simbolo.Valor) + toFloat(valor)

	} else {
		v.OutputASM.WriteString(fmt.Sprintf("adr x0, %s\n", id))
		v.OutputASM.WriteString("ldr x0, [x0]\n")
		v.OutputASM.WriteString(fmt.Sprintf("add x0, x0, #%d\n", valor.(int)))
		v.OutputASM.WriteString(fmt.Sprintf("adr x1, %s\n", id))
		v.OutputASM.WriteString("str x0, [x1]\n\n")

		simbolo.Valor = simbolo.Valor.(int) + valor.(int)
	}

	return nil
}

func (v *EvalVisitor) VisitAsigDecre(ctx *parser.AsigDecreContext) interface{} {
	id := ctx.IDENTIFICADOR().GetText()
	valorExpr := ctx.Expresion()
	valor := v.Visit(valorExpr)

	simbolo := v.Tabla.EntornoActual.BuscarSimbolo(id)
	if simbolo == nil {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Variable '%s' no declarada", id))
		return nil
	}

	tipoSimbolo := simbolo.TipoDato
	tipoValor := inferirTipo(valor)

	var tipoFinal string
	if tipoSimbolo == "float" || tipoValor == "float" {
		tipoFinal = "float"
	} else {
		tipoFinal = "int"
	}

	traducciones.ReservarVariableSiNoExiste(id, tipoFinal)
	v.OutputASM.WriteString(fmt.Sprintf("// %s -= %v\n", id, valor))

	if tipoFinal == "float" {
		v.OutputASM.WriteString(fmt.Sprintf("adr x10, %s\n", id))
		v.OutputASM.WriteString("ldr s0, [x10]\n")

		if tipoValor == "int" {
			v.OutputASM.WriteString(fmt.Sprintf("mov x11, #%d\n", valor.(int)))
			v.OutputASM.WriteString("scvtf s1, x11\n")
		} else {
			v.OutputASM.WriteString(fmt.Sprintf("ldr s1, =%g\n", valor.(float64)))
		}

		v.OutputASM.WriteString("fsub s2, s0, s1\n")
		v.OutputASM.WriteString(fmt.Sprintf("adr x12, %s\n", id))
		v.OutputASM.WriteString("str s2, [x12]\n\n")

		simbolo.Valor = toFloat(simbolo.Valor) - toFloat(valor)

	} else {
		v.OutputASM.WriteString(fmt.Sprintf("adr x0, %s\n", id))
		v.OutputASM.WriteString("ldr x0, [x0]\n")
		v.OutputASM.WriteString(fmt.Sprintf("sub x0, x0, #%d\n", valor.(int)))
		v.OutputASM.WriteString(fmt.Sprintf("adr x1, %s\n", id))
		v.OutputASM.WriteString("str x0, [x1]\n\n")

		simbolo.Valor = simbolo.Valor.(int) - valor.(int)
	}

	return nil
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
	if ctx.MAS() != nil {
		left := ctx.Expresion(0)
		right := ctx.Expresion(1)

		lid := left.GetText()
		rid := right.GetText()

		simboloL := v.Tabla.EntornoActual.BuscarSimbolo(lid)
		simboloR := v.Tabla.EntornoActual.BuscarSimbolo(rid)

		if simboloL != nil && simboloR != nil {
			tipoL := simboloL.TipoDato
			tipoR := simboloR.TipoDato

			var tipoFinal string
			if tipoL == "string" || tipoR == "string" {
				tipoFinal = "string"
			} else if tipoL == "float" || tipoR == "float" {
				tipoFinal = "float"
			} else {
				tipoFinal = "int"
			}

			traducciones.ReservarVariableSiNoExiste(lid, tipoL)
			traducciones.ReservarVariableSiNoExiste(rid, tipoR)

			if tipoFinal == "string" {
				var strVal, numVal string
				if tipoL == "string" {
					strVal = simboloL.Valor.(string)
					numVal = fmt.Sprintf("%v", simboloR.Valor)
				} else {
					strVal = simboloR.Valor.(string)
					numVal = fmt.Sprintf("%v", simboloL.Valor)
					return numVal + strVal
				}
				return strVal + numVal
			}

			traducciones.GenerarSumaASM(lid, rid, &traducciones.FuncionesBuilder, tipoFinal)
			v.OutputASM.WriteString("    bl add\n")
			switch tipoFinal {
			case "float":
				return toFloat(simboloL.Valor) + toFloat(simboloR.Valor)
			case "int":
				return simboloL.Valor.(int) + simboloR.Valor.(int)
			}
		}
	}

	if ctx.MENOS() != nil {
		left := ctx.Expresion(0)
		right := ctx.Expresion(1)

		lid := left.GetText()
		rid := right.GetText()

		simboloL := v.Tabla.EntornoActual.BuscarSimbolo(lid)
		simboloR := v.Tabla.EntornoActual.BuscarSimbolo(rid)

		if simboloL != nil && simboloR != nil {
			tipoL := simboloL.TipoDato
			tipoR := simboloR.TipoDato

			var tipoFinal string
			if tipoL == "float" || tipoR == "float" {
				tipoFinal = "float"
			} else {
				tipoFinal = "int"
			}

			if tipoFinal != "int" && tipoFinal != "float" {
				fmt.Println("Error: tipo no compatible para resta")
				return nil
			}

			traducciones.ReservarVariableSiNoExiste(lid, tipoL)
			traducciones.ReservarVariableSiNoExiste(rid, tipoR)

			traducciones.GenerarRestaASM(lid, rid, &traducciones.FuncionesBuilder, tipoFinal)
			v.OutputASM.WriteString("    bl add\n")
			switch tipoFinal {
			case "float":
				return toFloat(simboloL.Valor) - toFloat(simboloR.Valor)
			default:
				return simboloL.Valor.(int) - simboloR.Valor.(int)
			}
		}
	}

	if ctx.POR() != nil {
		left := ctx.Expresion(0)
		right := ctx.Expresion(1)

		lid := left.GetText()
		rid := right.GetText()

		simboloL := v.Tabla.EntornoActual.BuscarSimbolo(lid)
		simboloR := v.Tabla.EntornoActual.BuscarSimbolo(rid)

		if simboloL != nil && simboloR != nil {
			tipoL := simboloL.TipoDato
			tipoR := simboloR.TipoDato

			var tipoFinal string
			if tipoL == "float" || tipoR == "float" {
				tipoFinal = "float"
			} else {
				tipoFinal = "int"
			}

			if tipoFinal != "int" && tipoFinal != "float" {
				fmt.Println("Error: tipos no compatibles para multiplicación")
				return nil
			}

			traducciones.ReservarVariableSiNoExiste(lid, tipoL)
			traducciones.ReservarVariableSiNoExiste(rid, tipoR)

			traducciones.GenerarMultiplicacionASM(lid, rid, &traducciones.FuncionesBuilder, tipoFinal)
			v.OutputASM.WriteString("    bl add\n")

			switch tipoFinal {
			case "float":
				return toFloat(simboloL.Valor) * toFloat(simboloR.Valor)
			default:
				return simboloL.Valor.(int) * simboloR.Valor.(int)
			}
		}
	}

	if ctx.DIV() != nil {
		left := ctx.Expresion(0)
		right := ctx.Expresion(1)

		lid := left.GetText()
		rid := right.GetText()

		simboloL := v.Tabla.EntornoActual.BuscarSimbolo(lid)
		simboloR := v.Tabla.EntornoActual.BuscarSimbolo(rid)

		if simboloL != nil && simboloR != nil {
			tipoL := simboloL.TipoDato
			tipoR := simboloR.TipoDato

			var tipoFinal string
			if tipoL == "float" || tipoR == "float" {
				tipoFinal = "float"
			} else {
				tipoFinal = "int"
			}

			if tipoFinal != "int" && tipoFinal != "float" {
				fmt.Println("Error: tipos no compatibles para división")
				return nil
			}

			if (tipoFinal == "float" && toFloat(simboloR.Valor) == 0.0) || (tipoFinal == "int" && toInt(simboloR.Valor) == 0) {
				msg := fmt.Sprintf("Error semántico: división por cero no permitida")
				v.Tabla.Errores.NewSemanticError(ctx.GetStart(), msg)
				return nil
			}

			traducciones.ReservarVariableSiNoExiste(lid, tipoL)
			traducciones.ReservarVariableSiNoExiste(rid, tipoR)

			traducciones.GenerarDivisionASM(lid, rid, &traducciones.FuncionesBuilder, tipoFinal)
			v.OutputASM.WriteString("    bl add\n")

			switch tipoFinal {
			case "float":
				return toFloat(simboloL.Valor) / toFloat(simboloR.Valor)
			default:
				return simboloL.Valor.(int) / simboloR.Valor.(int)
			}
		}
	}

	if ctx.MODULO() != nil {
		left := ctx.Expresion(0)
		right := ctx.Expresion(1)

		lid := left.GetText()
		rid := right.GetText()

		simboloL := v.Tabla.EntornoActual.BuscarSimbolo(lid)
		simboloR := v.Tabla.EntornoActual.BuscarSimbolo(rid)

		if simboloL != nil && simboloR != nil {
			tipoL := simboloL.TipoDato
			tipoR := simboloR.TipoDato

			if tipoL != "int" && tipoL != "float" {
				fmt.Println("Error: lado izquierdo no es entero ni convertible")
				return nil
			}
			if tipoR != "int" && tipoR != "float" {
				fmt.Println("Error: lado derecho no es entero ni convertible")
				return nil
			}

			traducciones.ReservarVariableSiNoExiste(lid, tipoL)
			traducciones.ReservarVariableSiNoExiste(rid, tipoR)

			traducciones.GenerarModuloASM(lid, rid, &traducciones.FuncionesBuilder, tipoL, tipoR)
			v.OutputASM.WriteString("    bl add\n")
			valL := toInt(simboloL.Valor)
			valR := toInt(simboloR.Valor)
			if valR == 0 {
				fmt.Println("Error: división por cero en módulo")
				return nil
			}
			return valL % valR
		}
	}

	/*	RELACIONALES	*/
	if ctx.DIFERENTE() != nil {
		left := ctx.Expresion(0)
		right := ctx.Expresion(1)

		lid := left.GetText()
		rid := right.GetText()

		simboloL := v.Tabla.EntornoActual.BuscarSimbolo(lid)
		simboloR := v.Tabla.EntornoActual.BuscarSimbolo(rid)

		if simboloL != nil && simboloR != nil {
			tipoL := simboloL.TipoDato
			tipoR := simboloR.TipoDato

			traducciones.ReservarVariableSiNoExiste(lid, tipoL)
			traducciones.ReservarVariableSiNoExiste(rid, tipoR)
			traducciones.GenerarDiferenteASM(lid, rid, &traducciones.FuncionesBuilder, tipoL, tipoR)
			v.OutputASM.WriteString("    bl add\n")

			return toFloat(simboloL.Valor) != toFloat(simboloR.Valor)
		}
	}

	if ctx.IGUALDAD() != nil {
		left := ctx.Expresion(0)
		right := ctx.Expresion(1)

		lid := left.GetText()
		rid := right.GetText()

		simboloL := v.Tabla.EntornoActual.BuscarSimbolo(lid)
		simboloR := v.Tabla.EntornoActual.BuscarSimbolo(rid)

		if simboloL != nil && simboloR != nil {
			tipoL := simboloL.TipoDato
			tipoR := simboloR.TipoDato

			traducciones.ReservarVariableSiNoExiste(lid, tipoL)
			traducciones.ReservarVariableSiNoExiste(rid, tipoR)
			traducciones.GenerarIgualASM(lid, rid, &traducciones.FuncionesBuilder, tipoL, tipoR)
			v.OutputASM.WriteString("    bl add\n")

			return toFloat(simboloL.Valor) == toFloat(simboloR.Valor)
		}
	}

	if ctx.MENIGUAL() != nil {
		left := ctx.Expresion(0)
		right := ctx.Expresion(1)

		lid := left.GetText()
		rid := right.GetText()

		simboloL := v.Tabla.EntornoActual.BuscarSimbolo(lid)
		simboloR := v.Tabla.EntornoActual.BuscarSimbolo(rid)

		if simboloL != nil && simboloR != nil {
			tipoL := simboloL.TipoDato
			tipoR := simboloR.TipoDato

			traducciones.ReservarVariableSiNoExiste(lid, tipoL)
			traducciones.ReservarVariableSiNoExiste(rid, tipoR)
			traducciones.GenerarMenorIgualASM(lid, rid, &traducciones.FuncionesBuilder, tipoL, tipoR)
			v.OutputASM.WriteString("    bl add\n")

			return toFloat(simboloL.Valor) <= toFloat(simboloR.Valor)
		}
	}

	if ctx.MAYIGUAL() != nil {
		left := ctx.Expresion(0)
		right := ctx.Expresion(1)

		lid := left.GetText()
		rid := right.GetText()

		simboloL := v.Tabla.EntornoActual.BuscarSimbolo(lid)
		simboloR := v.Tabla.EntornoActual.BuscarSimbolo(rid)

		if simboloL != nil && simboloR != nil {
			tipoL := simboloL.TipoDato
			tipoR := simboloR.TipoDato

			traducciones.ReservarVariableSiNoExiste(lid, tipoL)
			traducciones.ReservarVariableSiNoExiste(rid, tipoR)

			traducciones.GenerarMayorIgualASM(lid, rid, &traducciones.FuncionesBuilder, tipoL, tipoR)
			v.OutputASM.WriteString("    bl add\n")

			return toFloat(simboloL.Valor) >= toFloat(simboloR.Valor)
		}
	}

	if ctx.MENOR() != nil {
		left := ctx.Expresion(0)
		right := ctx.Expresion(1)

		lid := left.GetText()
		rid := right.GetText()

		simboloL := v.Tabla.EntornoActual.BuscarSimbolo(lid)
		simboloR := v.Tabla.EntornoActual.BuscarSimbolo(rid)

		if simboloL != nil && simboloR != nil {
			tipoL := simboloL.TipoDato
			tipoR := simboloR.TipoDato

			traducciones.ReservarVariableSiNoExiste(lid, tipoL)
			traducciones.ReservarVariableSiNoExiste(rid, tipoR)
			traducciones.GenerarMenorASM(lid, rid, &traducciones.FuncionesBuilder, tipoL, tipoR)
			v.OutputASM.WriteString("    bl add\n")

			return toFloat(simboloL.Valor) < toFloat(simboloR.Valor)
		}
	}

	if ctx.MAYOR() != nil {
		left := ctx.Expresion(0)
		right := ctx.Expresion(1)

		lid := left.GetText()
		rid := right.GetText()

		simboloL := v.Tabla.EntornoActual.BuscarSimbolo(lid)
		simboloR := v.Tabla.EntornoActual.BuscarSimbolo(rid)

		if simboloL != nil && simboloR != nil {
			tipoL := simboloL.TipoDato
			tipoR := simboloR.TipoDato

			traducciones.ReservarVariableSiNoExiste(lid, tipoL)
			traducciones.ReservarVariableSiNoExiste(rid, tipoR)
			traducciones.GenerarMayorASM(lid, rid, &traducciones.FuncionesBuilder, tipoL, tipoR)
			v.OutputASM.WriteString("    bl add\n")

			return toFloat(simboloL.Valor) > toFloat(simboloR.Valor)
		}
	}
	//
	/*	BOOLEANOS	*/
	if ctx.DIFER() != nil {
		expr := ctx.Expresion(0)

		id := expr.GetText()
		simbolo := v.Tabla.EntornoActual.BuscarSimbolo(id)

		if simbolo != nil {
			if simbolo.TipoDato != "bool" {
				return false
			}
			traducciones.ReservarVariableSiNoExiste(id, "bool")
			traducciones.GenerarNotASM(id, &traducciones.FuncionesBuilder)
			v.OutputASM.WriteString("    bl add\n")
			val := simbolo.Valor.(bool)
			return !val
		}
		return false
	}
	if ctx.OR() != nil {
		left := ctx.Expresion(0)
		right := ctx.Expresion(1)

		lid := left.GetText()
		rid := right.GetText()

		simboloL := v.Tabla.EntornoActual.BuscarSimbolo(lid)
		simboloR := v.Tabla.EntornoActual.BuscarSimbolo(rid)

		if simboloL != nil && simboloR != nil {
			if simboloL.TipoDato != "bool" || simboloR.TipoDato != "bool" {
				return false
			}

			traducciones.ReservarVariableSiNoExiste(lid, "bool")
			traducciones.ReservarVariableSiNoExiste(rid, "bool")
			traducciones.GenerarOrASM(lid, rid, &traducciones.FuncionesBuilder)
			v.OutputASM.WriteString("    bl add\n")

			leftVal := simboloL.Valor.(bool)
			rightVal := simboloR.Valor.(bool)
			return leftVal || rightVal
		}
		return false
	}
	if ctx.AND() != nil {
		left := ctx.Expresion(0)
		right := ctx.Expresion(1)

		lid := left.GetText()
		rid := right.GetText()

		simboloL := v.Tabla.EntornoActual.BuscarSimbolo(lid)
		simboloR := v.Tabla.EntornoActual.BuscarSimbolo(rid)

		if simboloL != nil && simboloR != nil {
			if simboloL.TipoDato != "bool" || simboloR.TipoDato != "bool" {
				return false
			}

			traducciones.ReservarVariableSiNoExiste(lid, "bool")
			traducciones.ReservarVariableSiNoExiste(rid, "bool")
			traducciones.GenerarAndASM(lid, rid, &traducciones.FuncionesBuilder)
			v.OutputASM.WriteString("    bl add\n")

			leftVal := simboloL.Valor.(bool)
			rightVal := simboloR.Valor.(bool)
			return leftVal && rightVal
		}
		return false
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
func toFloat(val interface{}) float64 {
	switch v := val.(type) {
	case int:
		return float64(v)
	case float64:
		return v
	default:
		return 0.0
	}
}
func toInt(v interface{}) int {
	switch val := v.(type) {
	case int:
		return val
	case float64:
		return int(val)
	default:
		return 0
	}
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

// ASIGNACIONES
func (v *EvalVisitor) VisitAsignacionNormal(ctx *parser.AsignacionNormalContext) interface{} {
	id := ctx.IDENTIFICADOR().GetText()
	valor := v.Visit(ctx.Expresion())

	fmt.Printf("[DEBUG] Asignación normal: %s = %v\n", id, valor)

	simbolo := v.Tabla.EntornoActual.BuscarSimbolo(id)
	if simbolo == nil {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Variable '%s' no declarada", id))
		return nil
	}

	tipoSimbolo := simbolo.TipoDato
	tipoValor := inferirTipo(valor)

	var tipoFinal string
	if tipoSimbolo == "float" || tipoValor == "float" {
		tipoFinal = "float"
	} else {
		tipoFinal = "int"
	}

	traducciones.ReservarVariableSiNoExiste(id, tipoFinal)
	v.OutputASM.WriteString(fmt.Sprintf("// %s = (expresion)\n", id))

	if tipoFinal == "float" {
		// para float, cargar el valor en s registro y almacenar
		v.OutputASM.WriteString(fmt.Sprintf("adr x10, %s\n", id))

		if tipoValor == "int" {
			v.OutputASM.WriteString(fmt.Sprintf("mov x11, #%d\n", valor.(int)))
			v.OutputASM.WriteString("scvtf s0, x11\n")
		} else {
			v.OutputASM.WriteString(fmt.Sprintf("ldr s0, =%g\n", valor.(float64)))
		}

		v.OutputASM.WriteString("str s0, [x10]\n\n")

		simbolo.Valor = toFloat(valor)

	} else {
		// para int, cargar el valor en x registro y almacenar
		v.OutputASM.WriteString(fmt.Sprintf("adr x0, %s\n", id))

		if tipoValor == "int" {
			v.OutputASM.WriteString(fmt.Sprintf("mov x1, #%d\n", valor.(int)))
		} else {
			// si el valor es float, hacer cast a int (si tiene sentido)
			v.OutputASM.WriteString(fmt.Sprintf("// Casting float a int pendiente\n"))
		}

		v.OutputASM.WriteString("str x1, [x0]\n\n")

		if tipoValor == "int" {
			simbolo.Valor = valor.(int)
		} else {
			// si tienes que hacer cast, implementa aquí
		}
	}

	return nil
}

// incremento
func (v *EvalVisitor) VisitAsignacionIncremento(ctx *parser.AsignacionIncrementoContext) interface{} {
	id := ctx.IDENTIFICADOR().GetText()
	sim := v.Tabla.BuscarSimbolo(id)
	if sim == nil {
		panic("Variable no declarada: " + id)
	}
	if val, ok := sim.Valor.(int); ok {
		nuevo := val + 1
		v.Tabla.ActualizarVariable(id, nuevo)

		// Reservar y generar código ASM
		traducciones.ReservarVariableSiNoExiste(id, "int")
		traducciones.TextBuilder.WriteString(fmt.Sprintf("    adr x10, %s\n", id))
		traducciones.TextBuilder.WriteString("    ldr x11, [x10]\n")
		traducciones.TextBuilder.WriteString("    add x11, x11, #1\n")
		traducciones.TextBuilder.WriteString("    str x11, [x10]\n")

		return nil
	}
	panic("Incremento solo válido para enteros")
}

// Decremento
func (v *EvalVisitor) VisitAsignacionDecremento(ctx *parser.AsignacionDecrementoContext) interface{} {
	id := ctx.IDENTIFICADOR().GetText()
	sim := v.Tabla.BuscarSimbolo(id)
	if sim == nil {
		panic("Variable no declarada: " + id)
	}
	if val, ok := sim.Valor.(int); ok {
		nuevo := val - 1
		v.Tabla.ActualizarVariable(id, nuevo)

		traducciones.ReservarVariableSiNoExiste(id, "int")
		traducciones.TextBuilder.WriteString(fmt.Sprintf("    adr x10, %s\n", id))
		traducciones.TextBuilder.WriteString("    ldr x11, [x10]\n")
		traducciones.TextBuilder.WriteString("    sub x11, x11, #1\n")
		traducciones.TextBuilder.WriteString("    str x11, [x10]\n")

		return nil
	}
	panic("Decremento solo válido para enteros")
}

// IF
func (v *EvalVisitor) GenerarCondicion(ctx parser.IExpresionContext, etiquetaFallo string) {
	valor := v.Visit(ctx)

	if b, ok := valor.(bool); ok {
		if !b {
			fmt.Println("Generando condición con etiqueta:", etiquetaFallo)
			traducciones.TextBuilder.WriteString(fmt.Sprintf("    b %s\n", etiquetaFallo))
		}
		return
	}

	// Si es una variable booleana u operación relacional
	id := ctx.GetText()
	simbolo := v.Tabla.EntornoActual.BuscarSimbolo(id)
	if simbolo != nil && simbolo.TipoDato == "bool" {
		traducciones.ReservarVariableSiNoExiste(id, "bool")
		traducciones.TextBuilder.WriteString(fmt.Sprintf("    adr x10, %s\n", id))
		traducciones.TextBuilder.WriteString("    ldr x10, [x10]\n")
		traducciones.TextBuilder.WriteString("    cmp x10, #0\n")
		traducciones.TextBuilder.WriteString(fmt.Sprintf("    beq %s\n", etiquetaFallo))
		return
	}

	// Si no es variable booleana, asumimos que se generó el resultado booleano en x0
	traducciones.TextBuilder.WriteString("    cmp x0, #0\n")
	traducciones.TextBuilder.WriteString(fmt.Sprintf("    beq %s\n", etiquetaFallo))
}

func (v *EvalVisitor) VisitBloque(ctx *parser.BloqueContext) interface{} {
	instrucciones := ctx.AllInstruccion()
	var resultados []interface{}

	for _, instruccionCtx := range instrucciones {
		val := v.Visit(instruccionCtx)

		if _, esBreak := val.(BreakSignal); esBreak {
			return BreakSignal{}
		}
		if _, esContinue := val.(ContinueSignal); esContinue {
			return ContinueSignal{}
		}

		if list, ok := val.([]interface{}); ok {
			if len(list) > 0 {
				switch list[len(list)-1].(type) {
				case BreakSignal:
					return BreakSignal{}
				case ContinueSignal:
					return ContinueSignal{}
				}
			}
			resultados = append(resultados, list...)
		} else if val != nil {
			resultados = append(resultados, val)
		}
	}

	return resultados
}

func (v *EvalVisitor) VisitSif(ctx *parser.SifContext) interface{} {
	return v.IfCtrl.Ejecutar(
		func(t antlr.Tree) interface{} { return v.Visit(t.(antlr.ParseTree)) },
		v.GenerarCondicion,
		ctx,
	)
}
func (v *EvalVisitor) VisitForCondicional(ctx *parser.ForCondicionalContext) interface{} {
	fmt.Println("[DEBUG] Iniciando ejecución de for condicional")
	traducciones.TextBuilder.WriteString("\n// ====== FOR CONDICIONAL ======\n")
	for {
		// 1. Evaluar la condición
		condVal := v.Visit(ctx.Expresion())
		fmt.Printf("[DEBUG] Condición del for: %v\n", condVal)

		boolVal, ok := condVal.(bool)
		if !ok {
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "La condición del 'for' no es booleana")
			break
		}

		if !boolVal {
			fmt.Println("[DEBUG] Condición falsa. Saliendo del for.")
			break
		}

		// 2. Ejecutar el bloque
		val := v.Visit(ctx.Bloque())
		fmt.Println("[DEBUG] Bloque del for ejecutado.")

		if list, ok := val.([]interface{}); ok {
			if len(list) > 0 {
				switch list[len(list)-1].(type) {
				case BreakSignal:
					fmt.Println("[DEBUG] Break encontrado en bloque del for.")
					return list[:len(list)-1]
				case ContinueSignal:
					fmt.Println("[DEBUG] Continue encontrado en bloque del for.")
					continue
				}
			}
		} else if val != nil {
			switch val.(type) {
			case BreakSignal:
				fmt.Println("[DEBUG] Break encontrado (individual) en bloque del for.")
				break
			case ContinueSignal:
				fmt.Println("[DEBUG] Continue encontrado (individual) en bloque del for.")
				continue
			}
		}
	}

	fmt.Println("[DEBUG] Finalizó for condicional")
	return nil
}

func (v *EvalVisitor) VisitForClasico(ctx *parser.ForClasicoContext) interface{} {
	return v.ejecutarForClasico(ctx)
}

func (v *EvalVisitor) ejecutarForClasico(ctx *parser.ForClasicoContext) interface{} {
	fmt.Println("[DEBUG] Iniciando ejecución de for clásico")
	traducciones.TextBuilder.WriteString("\n// ====== FOR CLÁSICO ======\n")

	v.Visit(ctx.Asignacion(0))
	var resultado []interface{}

	for {
		condVal := v.Visit(ctx.Expresion())
		fmt.Printf("[DEBUG] Condición del for clásico: %v\n", condVal)

		boolVal, ok := condVal.(bool)
		if !ok {
			panic("Condición del for clásico no es booleana")
		}
		if !boolVal {
			fmt.Println("[DEBUG] Condición falsa. Saliendo del for clásico.")
			break
		}

		val := v.Visit(ctx.Bloque())
		fmt.Println("[DEBUG] Bloque del for clásico ejecutado.")

		if list, ok := val.([]interface{}); ok && len(list) > 0 {
			if _, isBreak := list[len(list)-1].(BreakSignal); isBreak {
				fmt.Println("[DEBUG] Break encontrado en bloque del for clásico.")
				resultado = append(resultado, list[:len(list)-1]...)
				break
			}
			if _, isContinue := list[len(list)-1].(ContinueSignal); isContinue {
				fmt.Println("[DEBUG] Continue encontrado en bloque del for clásico.")
				resultado = append(resultado, list[:len(list)-1]...)
				v.Visit(ctx.Asignacion(1))
				continue
			}
			resultado = append(resultado, list...)
		} else if _, isBreak := val.(BreakSignal); isBreak {
			fmt.Println("[DEBUG] Break encontrado (individual) en bloque del for clásico.")
			break
		} else if _, isContinue := val.(ContinueSignal); isContinue {
			fmt.Println("[DEBUG] Continue encontrado (individual) en bloque del for clásico.")
			v.Visit(ctx.Asignacion(1))
			continue
		} else if val != nil {
			resultado = append(resultado, val)
		}

		v.Visit(ctx.Asignacion(1))
	}

	fmt.Println("[DEBUG] Finalizó for clásico")
	return resultado
}

func (v *EvalVisitor) VisitSSwitch(ctx *parser.SSwitchContext) interface{} {
	valorSwitch := v.Visit(ctx.Expresion())
	traducciones.TextBuilder.WriteString("\n// ====== SWITCH ======\n")
	for _, caseCtx := range ctx.AllCaseBlock() {
		valorCase := v.Visit(caseCtx.Expresion())

		if valorSwitch == valorCase {
			res := v.Visit(caseCtx.Instrucciones())

			// Detectar break
			if list, ok := res.([]interface{}); ok {
				// Si el último elemento es BreakSignal, lo quita
				if len(list) > 0 {
					if _, isBreak := list[len(list)-1].(BreakSignal); isBreak {
						return list[:len(list)-1]
					}
				}
				return list
			}

			return res
		}
	}

	if ctx.DefaultBlock() != nil {
		res := v.Visit(ctx.DefaultBlock().Instrucciones())

		if list, ok := res.([]interface{}); ok {
			if len(list) > 0 {
				if _, isBreak := list[len(list)-1].(BreakSignal); isBreak {
					return list[:len(list)-1]
				}
			}
			return list
		}

		return res
	}

	return nil
}

func (v *EvalVisitor) VisitBreak(ctx *parser.Break_Context) interface{} {
	return BreakSignal{}
}

func (v *EvalVisitor) VisitContinue(ctx *parser.Continue_Context) interface{} {
	fmt.Println(">>> [DEBUG] Se ejecutó CONTINUE")
	return ContinueSignal{}
}
