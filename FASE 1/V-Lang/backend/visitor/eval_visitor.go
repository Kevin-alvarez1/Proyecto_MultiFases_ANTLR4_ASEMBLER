package visitor

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"vlang/backend/analizador/operaciones"
	parser "vlang/backend/analizador/parser"

	"vlang/frontend/symbols"

	"github.com/antlr4-go/antlr/v4"
)

type StructAtributo struct {
	Tipo string
	Id   string
}

type StructMethod struct {
	Nombre    string
	Receptor  string
	Parametro []Parametro
	Cuerpo    antlr.ParserRuleContext
}

type MetodoKey struct {
	Struct string
	Nombre string
}

type Funcion struct {
	Nombre               string
	TipoRetorno          string
	Parametros           []Parametro
	Bloque               antlr.ParseTree
	ValorRetorno         interface{}
	EntornoDeDeclaracion *symbols.Entorno
}

type Retorno struct {
	Valor interface{}
}
type Parametro struct {
	Nombre string
	Tipo   string
}
type BreakSignal struct{}
type ContinueSignal struct{}
type EvalVisitor struct {
	*parser.BasegramaticaVisitor
	Entorno                map[string]interface{}
	Tabla                  *symbols.TablaSimbolos
	Console                *strings.Builder
	StructDef              map[string][]StructAtributo
	Metodos                map[MetodoKey]StructMethod
	Funciones              map[string][]Funcion
	NombreEntorno          string
	EntornoEjecucionActivo *symbols.Entorno
	IfCtrl                 *If
}

func NewEvalVisitor(tabla *symbols.TablaSimbolos) *EvalVisitor {
	return &EvalVisitor{
		BasegramaticaVisitor: &parser.BasegramaticaVisitor{},
		Entorno:              make(map[string]interface{}),
		Tabla:                tabla,
		Console:              &strings.Builder{},
		StructDef:            make(map[string][]StructAtributo),
		Metodos:              make(map[MetodoKey]StructMethod),
		Funciones:            make(map[string][]Funcion),
		IfCtrl:               NewIf(),
	}
}

func (v *EvalVisitor) Visit(tree antlr.ParseTree) interface{} {
	fmt.Println("Visit: ", reflect.TypeOf(tree))
	return tree.Accept(v)
}

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

func (v *EvalVisitor) VisitInstruccion(ctx *parser.InstruccionContext) interface{} {
	if ctx.Print_() != nil {
		val := v.Visit(ctx.Print_())
		return val
	}
	if ctx.Println_() != nil {
		val := v.Visit(ctx.Println_())
		return val
	}
	if ctx.DeclaracionMultiple() != nil {
		v.Visit(ctx.DeclaracionMultiple())
		return ""
	}
	if ctx.DeclaracionMultipleSimple() != nil {
		v.Visit(ctx.DeclaracionMultipleSimple())
		return ""
	}
	if ctx.DeclaracionMultipleSinTipo() != nil {
		v.Visit(ctx.DeclaracionMultipleSinTipo())
		return ""
	}
	if ctx.AsignacionMultiple() != nil {
		v.Visit(ctx.AsignacionMultiple())
		return ""
	}

	if ctx.StructDef() != nil {
		v.Visit(ctx.StructDef())
		return nil
	}

	if ctx.StructInst() != nil {
		v.Visit(ctx.StructInst())
		return nil
	}

	if ctx.AccesoStruct() != nil {
		v.Visit(ctx.AccesoStruct())
		return nil
	}

	if ctx.AsigStruct() != nil {
		v.Visit(ctx.AsigStruct())
		return nil
	}
	if ctx.FunStruct() != nil {
		v.Visit(ctx.FunStruct())
		return nil
	}
	if ctx.ActStruct() != nil {
		v.Visit(ctx.ActStruct())
		return nil
	}
	if ctx.FnSinParametro() != nil {
		v.Visit(ctx.FnSinParametro())
		return nil
	}
	if ctx.FnConParametro() != nil {
		v.Visit(ctx.FnConParametro())
		return nil
	}
	if ctx.LlamadaFuncionesSinParametro() != nil {
		v.Visit(ctx.LlamadaFuncionesSinParametro())
		return nil
	}
	if ctx.LlamadaFuncionesConParametro() != nil {
		v.Visit(ctx.LlamadaFuncionesConParametro())
		return nil
	}
	if ctx.SliceDef() != nil {
		v.Visit(ctx.SliceDef())
		return nil
	}
	if ctx.ModificacionElementoSlice() != nil {
		v.Visit(ctx.ModificacionElementoSlice())
		return nil
	}
	if ctx.DeclaracionSliceVacio() != nil {
		v.Visit(ctx.DeclaracionSliceVacio())
		return nil
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
	if ctx.Sif() != nil {
		return v.Visit(ctx.Sif())
	}
	if ctx.Asignacion() != nil {
		return v.Visit(ctx.Asignacion())
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

func (v *EvalVisitor) VisitInstrucciones(ctx *parser.InstruccionesContext) interface{} {
	var results []interface{}

	for _, instr := range ctx.AllInstruccion() {
		val := v.Visit(instr)

		if _, ok := val.(BreakSignal); ok {
			return BreakSignal{}
		}
		if _, ok := val.(ContinueSignal); ok {
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
			results = append(results, list...)
		} else if val != nil {
			results = append(results, val)
		}
	}

	return results
}

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

func (v *EvalVisitor) VisitPrintln(ctx *parser.PrintlnContext) interface{} {
	var output strings.Builder
	expresiones := ctx.ListaExpr().AllExpresion()

	for i, expr := range expresiones {
		val := v.Visit(expr)
		fmt.Printf("[DEBUG PRINTLN] Argumento %d: Valor %#v (Tipo: %T)\n", i, val, val)
		output.WriteString(valorAString(val))
		if i < len(expresiones)-1 {
			output.WriteString(" ")
		}
	}

	output.WriteString("\n")
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

	for i, id := range ids {
		valor := valores[i]
		valorTipo := inferirTipo(valor)
		if !tiposCompatibles(tipo, valorTipo) {
			msg := fmt.Sprintf("Error: no se puede asignar %s a variable %s de tipo %s",
				valorTipo, id, tipo)
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), msg)
			continue
		}

		simbolo := &symbols.Simbolo{
			ID:          id,
			TipoDato:    tipo,
			Valor:       valor,
			Ambito:      v.Tabla.EntornoActual.Nombre,
			TipoSimbolo: "variable",
			Linea:       ctx.GetStart().GetLine(),
			Columna:     ctx.GetStart().GetColumn(),
		}
		v.Tabla.EntornoActual.Simbolos[id] = simbolo
	}
	return nil
}

func (v *EvalVisitor) VisitDeclaracionMultipleSinTipo(ctx *parser.DeclaracionMultipleSinTipoContext) interface{} {
	listaExpr := ctx.ListaExpr()
	valores := obtenerValores(listaExpr, v)
	// Obtener todos los IDs
	ids := obtenerIDs(ctx.ListaIDS())
	if len(ids) != len(valores) {
		msg := fmt.Sprintf("Error semántico: cantidad de variables (%d) no coincide con valores (%d)",
			len(ids), len(valores))
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), msg)
		return nil
	}
	// Procesar cada asignación
	for i, id := range ids {
		valor := valores[i]
		tipo := inferirTipo(valor)
		v.Entorno[id] = valor
		fmt.Printf("Variable declarada: %s tipo: %s con valor: %v\n", id, tipo, valor)
		v.Tabla.DeclararVariable(id, tipo, valor, ctx, v.Tabla.EntornoActual.Nombre)
	}
	return nil
}

func (v *EvalVisitor) VisitDeclaracionMultipleSimple(ctx *parser.DeclaracionMultipleSimpleContext) interface{} {
	if ctx.Tipos() == nil || ctx.ListaIDS() == nil {
		return fmt.Errorf("Error: declaración incompleta")
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
		valorDefault = "\"\""
	case "bool":
		valorDefault = false
	default:
		valorDefault = nil
	}
	for _, id := range ids {
		v.Tabla.DeclararVariableSimple(id, tipo, valorDefault, ctx, v.Tabla.EntornoActual.Nombre)
		v.Entorno[id] = valorDefault
	}
	return nil
}

func (v *EvalVisitor) VisitAsignacionMultiple(ctx *parser.AsignacionMultipleContext) interface{} {
	ids := obtenerIDs(ctx.ListaIDS())
	valores := obtenerValores(ctx.ListaExpr(), v)

	if len(ids) != len(valores) {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "Número de variables y valores no coincide en declaración múltiple")
		return nil
	}

	for i, id := range ids {
		valor := valores[i]
		simboloExistente := v.Tabla.EntornoActual.BuscarSimbolo(id)
		if simboloExistente != nil {
			simboloExistente.Valor = valor
		} else {
			simbolo := &symbols.Simbolo{
				ID:       id,
				TipoDato: inferirTipo(valor),
				Valor:    valor,
				Ambito:   v.Tabla.EntornoActual.Nombre,
				Linea:    ctx.GetStart().GetLine(),
				Columna:  ctx.GetStart().GetColumn(),
			}
			v.Tabla.EntornoActual.Simbolos[id] = simbolo
		}
		// Copia profunda para evitar aliasing accidental
		switch val := valor.(type) {
		case []interface{}:
			copia := make([]interface{}, len(val))
			copy(copia, val)
			valor = copia
		case [][]interface{}:
			copia := make([][]interface{}, len(val))
			for j, fila := range val {
				if fila == nil {
					copia[j] = nil
					continue
				}
				subcopia := make([]interface{}, len(fila))
				copy(subcopia, fila)
				copia[j] = subcopia
			}
			valor = copia
		}

		tipoInferido := inferirTipo(valor)

		fmt.Printf("DEBUG: Buscando símbolo '%s' en entorno '%s'\n", id, v.Tabla.EntornoActual.Nombre)

		simbolo := v.Tabla.EntornoActual.BuscarSimbolo(id)
		if simbolo != nil {
			if !tiposCompatibles(simbolo.TipoDato, tipoInferido) {
				msg := fmt.Sprintf("No se puede asignar tipo '%s' a '%s' de tipo '%s'", tipoInferido, id, simbolo.TipoDato)
				v.Tabla.Errores.NewSemanticError(ctx.GetStart(), msg)
				continue
			}
			simbolo.Valor = valor
			v.Entorno[id] = valor
			fmt.Printf("Actualización: %s = %v (tipo: %s)\n", id, valor, tipoInferido)
		} else {
			v.Tabla.DeclararVariableSimple(id, tipoInferido, valor, ctx, v.Tabla.EntornoActual.Nombre)
			v.Entorno[id] = valor
			fmt.Printf("Declaración: %s = %v (tipo: %s)\n", id, valor, tipoInferido)
		}
	}

	return nil
}

// Funciones sin parametros
func (v *EvalVisitor) VisitFnSinParametro(ctx *parser.FnSinParametroContext) interface{} {
	nombre := ctx.IDENTIFICADOR().GetText()
	fmt.Printf("[VisitFnSinParametro] Definiendo función '%s'\n", nombre)

	if _, existe := v.Funciones[nombre]; existe {
		panic(fmt.Sprintf("[VisitFnSinParametro] ERROR: Función '%s' ya declarada", nombre))
	}

	tipoRetorno := "void"
	if ctx.TipoRetorno() != nil {
		tipoRetorno = ctx.TipoRetorno().GetText()
		fmt.Printf("[VisitFnSinParametro] Tipo de retorno especificado: %s\n", tipoRetorno)
	} else {
		fmt.Println("[VisitFnSinParametro] Tipo de retorno por defecto: void")
	}

	fn := Funcion{
		TipoRetorno:  tipoRetorno,
		Nombre:       nombre,
		Bloque:       ctx.BloqueFuncion(),
		ValorRetorno: nil,
	}

	v.Funciones[nombre] = []Funcion{fn}
	fmt.Printf("[VisitFnSinParametro] Función '%s' registrada correctamente con tipo de retorno '%s'\n", nombre, tipoRetorno)

	return nil
}

/*
metodo comentado porque da fallo los structs por el entorno temporal de bloque

	func (v *EvalVisitor) VisitFnSinParametro(ctx *parser.FnSinParametroContext) interface{} {
		nombre := ctx.IDENTIFICADOR().GetText()
		fmt.Printf("[VisitFnSinParametro] Definiendo función '%s'\n", nombre)

		if _, existe := v.Funciones[nombre]; existe {
			panic(fmt.Sprintf("[VisitFnSinParametro] ERROR: Función '%s' ya declarada", nombre))
		}

		tipoRetorno := "void"
		if ctx.TipoRetorno() != nil {
			tipoRetorno = ctx.TipoRetorno().GetText()
			fmt.Printf("[VisitFnSinParametro] Tipo de retorno especificado: %s\n", tipoRetorno)
		} else {
			fmt.Println("[VisitFnSinParametro] Tipo de retorno por defecto: void")
		}

		fn := Funcion{
			TipoRetorno:  tipoRetorno,
			Nombre:       nombre,
			Bloque:       ctx.BloqueFuncion(),
			ValorRetorno: nil,
		}

		v.Funciones[nombre] = []Funcion{fn}
		fmt.Printf("[VisitFnSinParametro] Función '%s' registrada correctamente con tipo de retorno '%s'\n", nombre, tipoRetorno)

		if ctx.BloqueFuncion() != nil {
			entornoAnterior := v.Tabla.EntornoActual
			nombreEntornoAnterior := v.NombreEntorno

			nuevoEntornoDef := symbols.NewEntorno(entornoAnterior, "def_"+nombre)
			v.Tabla.EntornoActual.Hijos = append(v.Tabla.EntornoActual.Hijos, nuevoEntornoDef)

			v.Tabla.EntornoActual = nuevoEntornoDef
			v.NombreEntorno = "def_" + nombre

			for _, child := range ctx.BloqueFuncion().GetChildren() {
				if instrCtx, ok := child.(*parser.InstruccionContext); ok {
					switch {
					case instrCtx.DeclaracionMultiple() != nil:
						fmt.Printf("[DEBUG DEF] DeclaracionMultiple encontrada en %s. Procesando.\n", v.Tabla.EntornoActual.Nombre)
						v.Visit(instrCtx.DeclaracionMultiple())
					case instrCtx.DeclaracionMultipleSimple() != nil:
						fmt.Printf("[DEBUG DEF] DeclaracionMultipleSimple encontrada en %s. Procesando.\n", v.Tabla.EntornoActual.Nombre)
						v.Visit(instrCtx.DeclaracionMultipleSimple())
					case instrCtx.DeclaracionMultipleSinTipo() != nil:
						fmt.Printf("[DEBUG DEF] DeclaracionMultipleSinTipo encontrada en %s. Procesando.\n", v.Tabla.EntornoActual.Nombre)
						v.Visit(instrCtx.DeclaracionMultipleSinTipo())
					default:
						fmt.Printf("[DEBUG DEF] **SALTANDO** Instrucción no-declaración durante la definición de función: %T\n", instrCtx.GetChild(0))
					}
				} else {
					fmt.Printf("[DEBUG DEF] **SALTANDO** Elemento no-instrucción durante la definición de función: %T\n", child)
				}
			}

			v.Tabla.EntornoActual = entornoAnterior
			v.NombreEntorno = nombreEntornoAnterior
		}

		return nil
	}
*/
func (v *EvalVisitor) VisitLlamadaFuncionesSinParametro(ctx *parser.LlamadaFuncionesSinParametroContext) interface{} {
	nombre := ctx.IDENTIFICADOR().GetText()

	if funcs, existe := v.Funciones[nombre]; existe {
		funcion := funcs[0]

		entornoAnterior := v.Tabla.EntornoActual
		nombreEntornoAnterior := v.NombreEntorno

		entornoEjecucionFn := symbols.NewEntorno(entornoAnterior, "ejecucion_"+nombre)
		entornoAnterior.Hijos = append(entornoAnterior.Hijos, entornoEjecucionFn)

		v.Tabla.EntornoActual = entornoEjecucionFn
		v.NombreEntorno = "ejecucion_" + nombre

		var resultado interface{}
		if funcion.Bloque != nil {
			resultado = v.Visit(funcion.Bloque)
		}

		v.Tabla.EntornoActual = entornoAnterior
		v.NombreEntorno = nombreEntornoAnterior

		if ret, ok := resultado.(Retorno); ok {
			return ret.Valor
		}

		return nil
	}

	panic("Función no definida: " + nombre)
}

// funciones con parametros
func (v *EvalVisitor) VisitFnConParametro(ctx *parser.FnConParametroContext) interface{} {
	nombre := ctx.IDENTIFICADOR().GetText()
	fmt.Printf("[VisitFnConParametro] Definiendo función '%s'\n", nombre)
	parametros := ExtraerParametros(ctx.ListaPar())

	if _, existe := v.Funciones[nombre]; existe {
		panic(fmt.Sprintf("[VisitFnConParametro] ERROR: Función '%s' ya declarada", nombre))
	}

	tipoRetorno := "void"
	if ctx.TipoRetorno() != nil {
		tipoRetorno = ctx.TipoRetorno().GetText()
	}

	entornoDeDeclaracion := v.Tabla.EntornoActual

	fn := Funcion{
		TipoRetorno: tipoRetorno,
		Nombre:      nombre,
		Bloque:      ctx.BloqueFuncion(),
		Parametros:  parametros,

		EntornoDeDeclaracion: entornoDeDeclaracion,
	}
	v.Funciones[nombre] = []Funcion{fn}

	fmt.Printf("[VisitFnConParametro] Función '%s' registrada correctamente con tipo de retorno '%s'\n", nombre, tipoRetorno)
	return nil
}
func (v *EvalVisitor) VisitLlamadaFuncionesConParametro(ctx *parser.LlamadaFuncionesConParametroContext) interface{} {
	nombre := ctx.IDENTIFICADOR().GetText()

	funcs, existe := v.Funciones[nombre]
	if !existe {
		panic("Función no definida: " + nombre)
	}

	funcion := funcs[0]

	// Guardar el entorno actual antes de entrar a la función
	entornoAnterior := v.Tabla.EntornoActual
	nombreEntornoAnterior := v.NombreEntorno

	entornoEjecucionFn := symbols.NewEntorno(entornoAnterior, "ejecucion_"+nombre)
	v.Tabla.EntornoActual = entornoEjecucionFn
	v.NombreEntorno = "ejecucion_" + nombre

	defer func() {
		v.Tabla.EntornoActual = entornoAnterior
		v.NombreEntorno = nombreEntornoAnterior
	}()

	paramCtxs := funcion.Parametros

	valoresExpr := ctx.ListaExpr().AllExpresion()

	// Validar número de argumentos
	if len(paramCtxs) != len(valoresExpr) {
		panic(fmt.Sprintf("[VisitLlamadaFuncionesConParametro] ERROR: Número de argumentos incorrecto para función '%s', esperado %d, recibido %d",
			nombre, len(paramCtxs), len(valoresExpr)))
	}

	for i, par := range paramCtxs {
		id := par.Nombre
		tipo := par.Tipo
		valorEvaluado := v.Visit(valoresExpr[i])
		fmt.Printf("[VisitLlamadaFuncionesConParametro] Parámetro '%s' (tipo %s) recibe valor '%v' (tipo %T)\n", id, tipo, valorEvaluado, valorEvaluado)

		v.Tabla.DeclararVariable(id, tipo, valorEvaluado, ctx, v.Tabla.EntornoActual.Nombre)
	}

	resultado := v.Visit(funcion.Bloque)

	// El defer se encargará de restaurar el entorno.
	if ret, ok := resultado.(Retorno); ok {
		return ret.Valor
	}

	return nil
}

func (v *EvalVisitor) VisitBloqueFuncion(ctx *parser.BloqueFuncionContext) interface{} {
	entornoAnterior := v.Tabla.EntornoActual
	nombreAnterior := v.NombreEntorno

	entornoBloque := symbols.NewEntorno(v.Tabla.EntornoActual, "bloque_funcion_interno")

	v.Tabla.EntornoActual = entornoBloque
	v.NombreEntorno = "bloque_funcion_interno"

	defer func() {
		v.Tabla.EntornoActual = entornoAnterior
		v.NombreEntorno = nombreAnterior
	}()

	var resultado interface{} = nil

	for _, instrCtx := range ctx.AllInstruccion() {
		res := v.Visit(instrCtx)
		if ret, ok := res.(Retorno); ok {
			resultado = ret
			break
		}
	}

	return resultado
}

// return
func (v *EvalVisitor) VisitRetorno(ctx *parser.RetornoContext) interface{} {
	fmt.Printf("DEBUG Retorno - Expresión: %v\n", ctx.Expresion())
	fmt.Printf("DEBUG Retorno - Texto completo: %s\n", ctx.GetText())

	if ctx.Expresion() != nil {
		valor := v.Visit(ctx.Expresion())
		fmt.Printf("DEBUG Retorno con valor: %v\n", valor)
		return Retorno{Valor: valor}
	}

	fmt.Println("DEBUG Retorno sin valor")
	return Retorno{Valor: nil}
}

// seccion slices
func (v *EvalVisitor) VisitSliceDef(ctx *parser.SliceDefContext) interface{} {
	id := ctx.IDENTIFICADOR().GetText()
	literal := ctx.SliceLiteral()

	tipo := literal.Tipos().GetText()

	if literal.ListaExpr() != nil {
		exprs := literal.ListaExpr()
		var elementos []interface{}

		for _, exprCtx := range exprs.AllExpresion() {
			valor := v.Visit(exprCtx)

			if !v.checkTipo(valor, tipo) {
				v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "tipo incompatible en slice inválida")
				return nil
			}
			elementos = append(elementos, valor)
		}

		v.Tabla.DeclararVariableSimple(id, "slice("+tipo+")", elementos, ctx, v.Tabla.EntornoActual.Nombre)

	} else if literal.ListaExprList() != nil {
		// 3. Bidimensional
		lista := literal.ListaExprList()
		var matriz [][]interface{}

		for _, filaCtx := range lista.AllListaExpr() {
			var fila []interface{}
			for _, exprCtx := range filaCtx.AllExpresion() {
				valor := v.Visit(exprCtx)
				if !v.checkTipo(valor, tipo) {
					v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "tipo incompatible en slice bidimensional")
					return nil
				}
				fila = append(fila, valor)
			}
			matriz = append(matriz, fila)
		}

		v.Tabla.DeclararVariableSimple(id, "slice(slice("+tipo+"))", matriz, ctx, v.Tabla.EntornoActual.Nombre)

	} else {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "Definición de slice inválida")
	}

	return nil
}

func (v *EvalVisitor) VisitDeclaracionSliceVacio(ctx *parser.DeclaracionSliceVacioContext) interface{} {
	id := ctx.IDENTIFICADOR().GetText()
	tipo := ctx.Tipos().GetText()

	esBidimensional := false
	if tn3, ok := ctx.GetChild(3).(antlr.TerminalNode); ok && tn3.GetText() == "[" {
		if tn4, ok := ctx.GetChild(4).(antlr.TerminalNode); ok && tn4.GetText() == "]" {
			esBidimensional = true
		}
	}

	var valorInicial interface{}
	var tipoDeclarado string

	if esBidimensional {
		var sliceBidimensionalVacio [][]interface{}
		valorInicial = sliceBidimensionalVacio
		tipoDeclarado = "slice(slice(" + tipo + "))"
	} else {
		var sliceUnidimensionalVacio []interface{}
		valorInicial = sliceUnidimensionalVacio
		tipoDeclarado = "slice(" + tipo + ")"
	}

	v.Tabla.DeclararVariable(
		id,
		tipoDeclarado,
		valorInicial,
		ctx,
		v.Tabla.EntornoActual.Nombre,
	)

	return nil
}

func (v *EvalVisitor) VisitAccesoElementoSlice(ctx *parser.AccesoElementoSliceContext) interface{} {
	id := ctx.IDENTIFICADOR().GetText()
	simbolo := v.Tabla.EntornoActual.BuscarSimbolo(id)

	if simbolo == nil {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), fmt.Sprintf("La variable '%s' no está definida", id))
		return nil
	}

	expresiones := ctx.AllExpresion()
	exprCount := len(expresiones)

	if exprCount == 1 {
		iRaw := v.Visit(expresiones[0])
		i, ok := iRaw.(int)
		if !ok {
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "El índice debe ser un entero")
			return nil
		}

		slice, ok := simbolo.Valor.([]interface{})
		if !ok {
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "La variable no es un slice")
			return nil
		}

		if i < 0 || i >= len(slice) {
			v.Tabla.Errores.NewRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "Índice fuera de rango")
			return nil
		}

		valor := slice[i]
		fmt.Printf("[DEBUG VisitAccesoElementoSlice] Retornando slice[%d]: valor=%v tipo=%T\n", i, valor, valor)

		return valor

	} else if exprCount == 2 {
		iRaw := v.Visit(expresiones[0])
		jRaw := v.Visit(expresiones[1])

		i, ok1 := iRaw.(int)
		j, ok2 := jRaw.(int)
		if !ok1 || !ok2 {
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "Ambos índices deben ser enteros")
			return nil
		}

		matrizBidimensional, ok := simbolo.Valor.([][]interface{})
		if !ok {
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "La variable no es un slice bidimensional")
			return nil
		}

		if i < 0 || i >= len(matrizBidimensional) {
			v.Tabla.Errores.NewRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "Índice de fila fuera de rango")
			return nil
		}

		fila := matrizBidimensional[i]

		if j < 0 || j >= len(fila) {
			v.Tabla.Errores.NewRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "Índice de columna fuera de rango")
			return nil
		}

		return fila[j]
	}

	v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "Número inválido de índices")
	return nil
}
func (v *EvalVisitor) VisitModificacionElementoSlice(ctx *parser.ModificacionElementoSliceContext) interface{} {
	id := ctx.IDENTIFICADOR().GetText()
	simbolo := v.Tabla.EntornoActual.BuscarSimbolo(id)

	if simbolo == nil {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), fmt.Sprintf("La variable '%s' no está definida", id))
		return nil
	}

	exprCount := len(ctx.AllExpresion())
	valorNuevo := v.Visit(ctx.Expresion(exprCount - 1))

	tipoElemento := inferirTipo(valorNuevo)
	tipoSlice := simbolo.TipoDato

	if exprCount == 2 {
		// Un índice → slice[i] = val
		iRaw := v.Visit(ctx.Expresion(0))
		i, ok := iRaw.(int)
		if !ok {
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "Índice debe ser un entero")
			return nil
		}

		slice, ok := simbolo.Valor.([]interface{})
		if !ok {
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "La variable no es un slice")
			return nil
		}

		if i < 0 || i >= len(slice) {
			v.Tabla.Errores.NewRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "Índice fuera de rango")
			return nil
		}

		if !tiposCompatibles(tipoElementoDeSlice(tipoSlice), tipoElemento) {
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "Tipo no compatible con elemento del slice")
			return nil
		}

		slice[i] = valorNuevo
		fmt.Printf("[MODIFICACIÓN 1D] %s[%d] = %v\n", id, i, valorNuevo)
		return nil
	} else if exprCount == 3 {
		// Dos índices → slice[i][j] = val
		iRaw := v.Visit(ctx.Expresion(0))
		jRaw := v.Visit(ctx.Expresion(1))
		i, ok1 := iRaw.(int)
		j, ok2 := jRaw.(int)

		if !ok1 || !ok2 {
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "Índices deben ser enteros")
			return nil
		}

		matrizBidimensional, ok := simbolo.Valor.([][]interface{})
		if !ok {
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "La variable no es un slice bidimensional")
			return nil
		}

		if i < 0 || i >= len(matrizBidimensional) { // Usar matrizBidimensional aquí
			v.Tabla.Errores.NewRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "Índice de fila fuera de rango")
			return nil
		}

		fila := matrizBidimensional[i]

		if j < 0 || j >= len(fila) {
			v.Tabla.Errores.NewRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "Índice de columna fuera de rango")
			return nil
		}

		if !tiposCompatibles(tipoElementoDeSlice(tipoElementoDeSlice(tipoSlice)), tipoElemento) {
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "Tipo no compatible con elemento de la matriz")
			return nil
		}

		// Asignar el nuevo valor
		fila[j] = valorNuevo
		fmt.Printf("[MODIFICACIÓN 2D] %s[%d][%d] = %v\n", id, i, j, valorNuevo)
		return nil
	}

	v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "Número inválido de índices")
	return nil
}

// funciones embebida visitor:
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

	return intVal
}
func tipoElementoDeSlice(tipoSlice string) string {
	if strings.HasPrefix(tipoSlice, "slice(") && strings.HasSuffix(tipoSlice, ")") {
		return tipoSlice[6 : len(tipoSlice)-1]
	}
	return ""
}

func (v *EvalVisitor) VisitFnParseToFloat(ctx *parser.FnParseToFloatContext) interface{} {
	text := ctx.ListaExpr().GetText()
	text = strings.Trim(text, `"`)

	// Convertir a float
	floatVal, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return 0.0
	}

	return floatVal
}
func (v *EvalVisitor) VisitTipos(ctx *parser.TiposContext) interface{} {
	return ctx.GetText()
}
func (v *EvalVisitor) VisitFnTypeOf(ctx *parser.FnTypeOfContext) interface{} {
	if ctx.ListaExpr() == nil {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "TypeOF requiere un argumento")
		return "undefined"
	}

	expresiones := ctx.ListaExpr().AllExpresion()
	if len(expresiones) != 1 {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "TypeOF requiere exactamente un argumento")
		return "undefined"
	}

	// Obtener el valor del argumento
	val := v.Visit(expresiones[0])

	// Determinar el tipo del valor
	switch val.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case bool:
		return "bool"
	case string:
		return "string"
	case rune:
		return "rune"
	case nil:
		return "nil"
	default:
		// Para tipos compuestos
		rt := reflect.TypeOf(val)
		switch rt.Kind() {
		case reflect.Struct:
			return rt.Name()
		case reflect.Map:
			return "struct"
		case reflect.Slice:
			return "slice"
		default:
			return rt.String()
		}
	}
}

// funciones para los slices
func (v *EvalVisitor) VisitFnIndexOf(ctx *parser.FnIndexOfContext) interface{} {
	if ctx.ListaExpr() == nil || len(ctx.ListaExpr().AllExpresion()) != 2 {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "IndexOf requiere dos argumentos")
		return -1
	}

	sliceExpr := ctx.ListaExpr().AllExpresion()[0]
	valorExpr := ctx.ListaExpr().AllExpresion()[1]

	slice := v.Visit(sliceExpr)
	valor := v.Visit(valorExpr)

	if slice == nil || valor == nil {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "Argumentos no pueden ser nil")
		return -1
	}

	sliceVal, ok := slice.([]interface{})
	if !ok {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "El primer argumento debe ser un slice")
		return -1
	}

	for i, elem := range sliceVal {
		if reflect.DeepEqual(elem, valor) {
			return i
		}
	}

	return -1
}
func (v *EvalVisitor) VisitFnJoin(ctx *parser.FnJoinContext) interface{} {
	if ctx.ListaExpr() == nil || len(ctx.ListaExpr().AllExpresion()) != 2 {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "Join requiere dos argumentos")
		return ""
	}
	sliceExpr := ctx.ListaExpr().AllExpresion()[0]
	valorExpr := ctx.ListaExpr().AllExpresion()[1]
	slice := v.Visit(sliceExpr)

	if slice == nil {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "El primer argumento no puede ser nil")
		return ""
	}
	var sliceVal []interface{}
	switch s := slice.(type) {
	case []interface{}:
		sliceVal = s
	case []string:
		for _, v := range s {
			sliceVal = append(sliceVal, v)
		}
	case []int:
		for _, v := range s {
			sliceVal = append(sliceVal, v)
		}
	default:
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "El primer argumento de join debe ser un arreglo")
		return ""
	}

	valor := v.Visit(valorExpr)
	if valor == nil {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "El segundo argumento no puede ser nil")
		return ""
	}
	var strElems []string
	for _, elem := range sliceVal {
		strElem := fmt.Sprintf("%v", elem)
		strElems = append(strElems, strElem)
	}
	output := strings.Join(strElems, fmt.Sprintf("%v", valor))
	fmt.Printf("[DEBUG JOIN] Resultado de Join: %s\n", output)

	return output
}

func (v *EvalVisitor) VisitFnLen(ctx *parser.FnLenContext) interface{} {
	if ctx.ListaExpr() == nil || len(ctx.ListaExpr().AllExpresion()) != 1 {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "Len requiere un argumento")
		return -1
	}
	sliceExpr := ctx.ListaExpr().AllExpresion()[0]
	slice := v.Visit(sliceExpr)
	if slice == nil {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "El argumento no puede ser nil")
		return -1
	}
	sliceVal, ok := slice.([]interface{})
	if !ok {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "El argumento debe ser un slice")
		return -1
	}
	return len(sliceVal)
}

func (v *EvalVisitor) VisitFnAppend(ctx *parser.FnAppendContext) interface{} {
	if ctx.ListaExpr() == nil || len(ctx.ListaExpr().AllExpresion()) < 2 {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "Append requiere al menos dos argumentos")
		return nil
	}

	sliceExpr := ctx.ListaExpr().AllExpresion()[0]
	valorExprs := ctx.ListaExpr().AllExpresion()[1:] // Arguments to add

	primerArgValor := v.Visit(sliceExpr) // Evaluate the slice being appended to (matriz)
	if primerArgValor == nil {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "The first argument cannot be nil")
		return nil
	}

	if sliceUnidimensional, ok := primerArgValor.([]interface{}); ok {

		for _, valorExpr := range valorExprs {
			valorAAgregar := v.Visit(valorExpr)
			if valorAAgregar == nil {
				v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "A value to add cannot be nil")
				continue
			}

			sliceUnidimensional = append(sliceUnidimensional, valorAAgregar)
		}
		fmt.Printf("[DEBUG APPEND 1D] Result of Append (1D): %v\n", sliceUnidimensional)
		return sliceUnidimensional

	} else if sliceBidimensional, ok := primerArgValor.([][]interface{}); ok {
		if len(valorExprs) != 1 {
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "Appending to a bidimensional slice only accepts one row at a time")
			return nil
		}

		valorAAgregarRaw := v.Visit(valorExprs[0])
		valorAAgregarFila, ok := valorAAgregarRaw.([]interface{})
		if !ok {
			v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "The value to add to a bidimensional slice must be another slice (a row)")
			return nil
		}

		sliceBidimensional = append(sliceBidimensional, valorAAgregarFila)
		fmt.Printf("[DEBUG APPEND 2D] Result of Append (2D): %v\n", sliceBidimensional)
		return sliceBidimensional

	} else {
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), "The first argument for append is not a valid slice type")
		return nil
	}
}
func (v *EvalVisitor) VisitListaIDS(ctx *parser.ListaIDSContext) interface{} {
	var resultados []interface{}

	for _, idToken := range ctx.AllIDENTIFICADOR() {
		id := idToken.GetText()
		val, ok := v.Entorno[id]
		if !ok {
			v.Tabla.Errores.NewSemanticError(idToken.GetSymbol(),
				fmt.Sprintf("Variable '%s' no declarada", id))
			return nil
		}
		resultados = append(resultados, val)
	}

	if len(resultados) == 1 {
		return resultados[0]
	}

	return resultados
}

/* ASIGNAR INCREMENTO Y DECREMENTO */
func (v *EvalVisitor) VisitAsigIncre(ctx *parser.AsigIncreContext) interface{} {
	id := ctx.IDENTIFICADOR().GetText()
	simbolo := v.Tabla.Obtener(id)
	if simbolo == nil {
		panic(fmt.Sprintf("Variable '%s' no definida", id))
	}

	valorActual := simbolo.Valor
	valorNuevo := v.Visit(ctx.Expresion())

	resultado := Sumar(valorActual, valorNuevo)
	simbolo.Valor = resultado

	return resultado
}

func (v *EvalVisitor) VisitAsigDecre(ctx *parser.AsigDecreContext) interface{} {
	id := ctx.IDENTIFICADOR().GetText()
	simbolo := v.Tabla.Obtener(id)
	if simbolo == nil {
		panic(fmt.Sprintf("Variable '%s' no definida", id))
	}

	valorActual := simbolo.Valor
	valorNuevo := v.Visit(ctx.Expresion())

	resultado := Restar(valorActual, valorNuevo)
	simbolo.Valor = resultado

	return resultado
}

func (v *EvalVisitor) VisitStructDef(ctx *parser.StructDefContext) interface{} {
	nombreStruct := ctx.IDENTIFICADOR().GetText()
	var atributos []StructAtributo

	fmt.Printf("[StructDef] Definiendo struct: %s\n", nombreStruct)

	for _, atrCtx := range ctx.AllAtributos() {
		tipo := atrCtx.Tipos().GetText()
		id := atrCtx.IDENTIFICADOR().GetText()
		fmt.Printf("  - Atributo: %s %s\n", tipo, id)
		atributos = append(atributos, StructAtributo{Tipo: tipo, Id: id})
	}

	v.StructDef[nombreStruct] = atributos
	fmt.Printf("[StructDef] Struct '%s' registrado con atributos: %+v\n", nombreStruct, atributos)
	return nil
}

func (v *EvalVisitor) VisitAtributos(ctx *parser.AtributosContext) interface{} {
	fmt.Println("[Atributos] Visitando atributos")
	return v.VisitChildren(ctx)
}

func (v *EvalVisitor) VisitStructInst(ctx *parser.StructInstContext) interface{} {
	varName := ctx.IDENTIFICADOR(0).GetText()
	structName := ctx.IDENTIFICADOR(1).GetText()
	fmt.Printf("[StructInst] Instanciando '%s' con struct '%s'\n", varName, structName)

	// Verificar si el struct existe
	atributos, ok := v.StructDef[structName]
	if !ok {
		panic(fmt.Sprintf("[StructInst] Struct '%s' no definido", structName))
	}

	// Obtener valores desde la inicialización
	valores := v.Visit(ctx.StructInit()).([]interface{})
	fmt.Printf("[StructInst] Valores recibidos para '%s': %v\n", structName, valores)

	if len(valores) != len(atributos) {
		panic(fmt.Sprintf("[StructInst] Cantidad de valores (%d) no coincide con los atributos (%d) del struct '%s'", len(valores), len(atributos), structName))
	}

	// Crear mapa con los valores
	instancia := make(map[string]interface{})
	for i, atr := range atributos {
		instancia[atr.Id] = valores[i]
		fmt.Printf("  - Atributo '%s' = %v\n", atr.Id, valores[i])
	}

	// Guardar en entorno
	v.Entorno[varName] = instancia
	simbolo := &symbols.Simbolo{
		ID:          varName,
		TipoDato:    structName,
		Valor:       instancia,
		Ambito:      v.Tabla.EntornoActual.Nombre,
		TipoSimbolo: "struct",
		Linea:       ctx.GetStart().GetLine(),
		Columna:     ctx.GetStart().GetColumn(),
	}
	v.Tabla.EntornoActual.Simbolos[varName] = simbolo
	fmt.Printf("[StructInst] Instancia creada: %s = %+v\n", varName, instancia)

	return nil
}

func (v *EvalVisitor) VisitStructInit(ctx *parser.StructInitContext) interface{} {
	fmt.Println("[StructInit] Inicializando struct...")
	return v.Visit(ctx.ListaStructs())
}

func (v *EvalVisitor) VisitListaStructs(ctx *parser.ListaStructsContext) interface{} {
	fmt.Println("[ListaStructs] Visitando lista de valores para struct...")

	var valores []interface{}
	for _, expr := range ctx.AllExpresion() {
		val := v.Visit(expr)
		valores = append(valores, val)
		fmt.Printf("  - Valor agregado: %v\n", val)
	}

	return valores
}

func (v *EvalVisitor) VisitAccesoStruct(ctx *parser.AccesoStructContext) interface{} {
	structName := ctx.IDENTIFICADOR(0).GetText()
	atr := ctx.IDENTIFICADOR(1).GetText()

	fmt.Printf("[AccesoStruct] Accediendo a '%s.%s'\n", structName, atr)

	val, ok := v.Entorno[structName]
	if !ok {
		panic(fmt.Sprintf("[AccesoStruct] Variable '%s' no definida", structName))
	}

	instancia, ok := val.(map[string]interface{})
	if !ok {
		panic(fmt.Sprintf("[AccesoStruct] '%s' no es una instancia de struct", structName))
	}

	res, ok := instancia[atr]
	if !ok {
		panic(fmt.Sprintf("[AccesoStruct] Atributo '%s' no existe en struct '%s'", atr, structName))
	}

	fmt.Printf("[AccesoStruct] Valor obtenido: %v\n", res)
	return res
}

func (v *EvalVisitor) VisitAsigStruct(ctx *parser.AsigStructContext) interface{} {
	structName := ctx.IDENTIFICADOR(0).GetText()
	atr := ctx.IDENTIFICADOR(1).GetText()
	valor := v.Visit(ctx.Expresion())

	fmt.Printf("[AsigStruct] Asignando '%s.%s' = %v\n", structName, atr, valor)

	val, ok := v.Entorno[structName]
	if !ok {
		panic(fmt.Sprintf("[AsigStruct] Variable '%s' no definida", structName))
	}

	instancia, ok := val.(map[string]interface{})
	if !ok {
		panic(fmt.Sprintf("[AsigStruct] '%s' no es una instancia de struct", structName))
	}

	if _, ok := instancia[atr]; !ok {
		panic(fmt.Sprintf("[AsigStruct] Atributo '%s' no existe en struct '%s'", atr, structName))
	}

	instancia[atr] = valor
	fmt.Printf("[AsigStruct] Modificado: %s.%s = %v\n", structName, atr, valor)
	return nil
}

/* FUNCIONES STRUCTS */
func (v *EvalVisitor) VisitFunStruct(ctx *parser.FunStructContext) interface{} {
	receptor := ctx.IDENTIFICADOR(0).GetText()
	structReceiver := ctx.IDENTIFICADOR(1).GetText()
	nombreMetodo := ctx.IDENTIFICADOR(2).GetText()

	fmt.Printf("[FunStruct] Registrando método '%s' para struct '%s'\n", nombreMetodo, structReceiver)

	var parametros []Parametro
	if ctx.ListaPar() != nil {
		parametros = v.Visit(ctx.ListaPar()).([]Parametro)
	}

	key := MetodoKey{
		Struct: structReceiver,
		Nombre: nombreMetodo,
	}

	v.Metodos[key] = StructMethod{
		Nombre:    nombreMetodo,
		Receptor:  receptor,
		Parametro: parametros,
		Cuerpo:    ctx.Caja().(*parser.CajaContext),
	}

	return nil
}

func (v *EvalVisitor) VisitCaja(ctx *parser.CajaContext) interface{} {
	fmt.Println("[Caja] Ejecutando cuerpo del método struct")
	for _, val := range ctx.AllValores() {
		v.Visit(val)
	}
	return nil
}

func (v *EvalVisitor) VisitValores(ctx *parser.ValoresContext) interface{} {
	instanciaID := ctx.IDENTIFICADOR(0).GetText()
	atributoID := ctx.IDENTIFICADOR(1).GetText()
	nuevoValorID := ctx.IDENTIFICADOR(2).GetText()

	valorInstancia, ok := v.Entorno[instanciaID]
	if !ok {
		panic(fmt.Sprintf("[Valores] Instancia '%s' no encontrada", instanciaID))
	}

	instancia, ok := valorInstancia.(map[string]interface{})
	if !ok {
		panic(fmt.Sprintf("[Valores] '%s' no es una instancia válida de struct", instanciaID))
	}

	nuevoValor, ok := v.Entorno[nuevoValorID]
	if !ok {
		panic(fmt.Sprintf("[Valores] Variable '%s' no encontrada", nuevoValorID))
	}

	instancia[atributoID] = nuevoValor
	fmt.Printf("[Valores] %s.%s = %v\n", instanciaID, atributoID, nuevoValor)

	return nil
}

func (v *EvalVisitor) VisitActStruct(ctx *parser.ActStructContext) interface{} {
	nombreInstancia := ctx.IDENTIFICADOR(0).GetText()
	metodoNombre := ctx.IDENTIFICADOR(1).GetText()

	fmt.Printf("[ActStruct] Llamando método '%s' de '%s'\n", metodoNombre, nombreInstancia)

	instancia, ok := v.Entorno[nombreInstancia].(map[string]interface{})
	if !ok {
		panic(fmt.Sprintf("[ActStruct] '%s' no es una instancia de struct", nombreInstancia))
	}

	var tipoStruct string
	for tipo, atts := range v.StructDef {
		coinciden := true
		for _, attr := range atts {
			if _, existe := instancia[attr.Id]; !existe {
				coinciden = false
				break
			}
		}
		if coinciden {
			tipoStruct = tipo
			break
		}
	}
	if tipoStruct == "" {
		panic(fmt.Sprintf("[ActStruct] No se pudo determinar el tipo de struct de '%s'", nombreInstancia))
	}

	key := MetodoKey{Struct: tipoStruct, Nombre: metodoNombre}
	metodo, ok := v.Metodos[key]
	if !ok {
		panic(fmt.Sprintf("[ActStruct] Método '%s' no encontrado para struct '%s'", metodoNombre, tipoStruct))
	}

	var valores []interface{}
	if ctx.ListaStructs() != nil {
		valores = v.Visit(ctx.ListaStructs()).([]interface{})
	}
	if len(valores) != len(metodo.Parametro) {
		panic("[ActStruct] Cantidad de argumentos no coincide con los parámetros")
	}

	entornoMetodo := make(map[string]interface{})

	entornoMetodo[metodo.Receptor] = instancia

	for i, param := range metodo.Parametro {
		entornoMetodo[param.Nombre] = valores[i]
	}

	entornoAnterior := v.Entorno
	v.Entorno = entornoMetodo

	v.Visit(metodo.Cuerpo)

	v.Entorno = entornoAnterior

	return nil
}
func (v *EvalVisitor) VisitListaPar(ctx *parser.ListaParContext) interface{} {
	var parametros []Parametro
	for _, par := range ctx.AllParametro() {
		param := v.Visit(par).(Parametro)
		parametros = append(parametros, param)
	}
	return parametros
}

func (v *EvalVisitor) VisitParametro(ctx *parser.ParametroContext) interface{} {
	nombre := ctx.IDENTIFICADOR().GetText()
	tipo := ctx.Tipos().GetText()
	return Parametro{Nombre: nombre, Tipo: tipo}
}

/* EXRESIONES */

func (v *EvalVisitor) VisitExpresion(ctx *parser.ExpresionContext) interface{} {
	if ctx.AccesoStruct() != nil {
		return v.Visit(ctx.AccesoStruct())
	}

	if ctx.GetChildCount() == 3 {
		if left, ok := ctx.GetChild(0).(antlr.TerminalNode); ok {
			if right, ok := ctx.GetChild(2).(antlr.TerminalNode); ok {
				if left.GetText() == "[" && right.GetText() == "]" {
					// Caso arreglo literal
					if listaCtx, ok := ctx.GetChild(1).(*parser.ListaExprContext); ok {
						var resultado []interface{}
						for _, expr := range listaCtx.AllExpresion() {
							val := v.Visit(expr)
							resultado = append(resultado, val)
						}
						return resultado
					}
					return []interface{}{} // arreglo vacío
				}
			}
		}
	}

	if ctx.LlamadaFuncionesSinParametro() != nil {
		// Obtener el nombre de la función
		nombreFuncion := ctx.LlamadaFuncionesSinParametro().IDENTIFICADOR().GetText()

		// Verificar si la función existe
		if funciones, existe := v.Funciones[nombreFuncion]; existe {
			resultado := v.Visit(funciones[0].Bloque)

			// Manejar el valor de retorno
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

	if ctx.IDENTIFICADOR() != nil {
		id := ctx.IDENTIFICADOR().GetText()
		fmt.Printf("[DEBUG EXPR] Buscando símbolo '%s' en entorno '%s'\n", id, v.Tabla.EntornoActual.Nombre)

		// Primero buscar como variable
		simbolo := buscarSimboloRecursivo(v.Tabla.EntornoActual, id)
		if simbolo != nil {
			fmt.Printf("[DEBUG EXPR] Símbolo '%s' ENCONTRADO. Valor: %#v (Tipo: %T)\n", id, simbolo.Valor, simbolo.Valor)
			return simbolo.Valor
		}

		if funciones, ok := v.Funciones[id]; ok {
			fmt.Printf("[DEBUG EXPR] Identificador '%s' es una función. Retornando función.\n", id)
			return funciones[0]
		}

		// No es ni variable ni función
		fmt.Printf("[DEBUG EXPR] Símbolo '%s' NO encontrado. Devolviendo nil.\n", id)
		v.Tabla.Errores.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Identificador '%s' no definido", id))
		return nil
	}

	if ctx.MAS() != nil {
		return Sumar(v.Visit(ctx.Expresion(0)), v.Visit(ctx.Expresion(1)))
	}

	// para resta
	if ctx.MENOS() != nil {
		if len(ctx.AllExpresion()) == 1 {
			return Negativo(v.Visit(ctx.Expresion(0)))
		} else {
			return Restar(v.Visit(ctx.Expresion(0)), v.Visit(ctx.Expresion(1)))
		}
	}

	// para multiplicacion
	if ctx.POR() != nil {
		return Multiplicar(v.Visit(ctx.Expresion(0)), v.Visit(ctx.Expresion(1)))
	}

	// para division
	if ctx.DIV() != nil {
		return Division(v.Visit(ctx.Expresion(0)), v.Visit(ctx.Expresion(1)))
	}

	// para potencia
	if ctx.MODULO() != nil {
		return Modulo(v.Visit(ctx.Expresion(0)), v.Visit(ctx.Expresion(1)))
	}
	/* OPERACIONES RELACIONALES */
	if ctx.DIFERENTE() != nil {
		fmt.Println(ctx.Expresion(0).GetText())
		fmt.Println(ctx.Expresion(1).GetText())

		return operaciones.NoIgualdad(v.Visit(ctx.Expresion(0)), v.Visit(ctx.Expresion(1)))
	}

	// para igualdad
	if ctx.IGUALDAD() != nil {
		return operaciones.Igualdad(v.Visit(ctx.Expresion(0)), v.Visit(ctx.Expresion(1)))
	}

	// para menor igual
	if ctx.MENIGUAL() != nil {
		return operaciones.MenorIgual(v.Visit(ctx.Expresion(0)), v.Visit(ctx.Expresion(1)))
	}

	// para mayor igual
	if ctx.MAYIGUAL() != nil {
		return operaciones.MayorIgual(v.Visit(ctx.Expresion(0)), v.Visit(ctx.Expresion(1)))
	}

	// para menor
	if ctx.MENOR() != nil {
		return operaciones.Menor(v.Visit(ctx.Expresion(0)), v.Visit(ctx.Expresion(1)))
	}

	// para mayor
	if ctx.MAYOR() != nil {
		return operaciones.Mayor(v.Visit(ctx.Expresion(0)), v.Visit(ctx.Expresion(1)))
	}
	/* EXPRESIONES BOOLEANAS */
	// para not
	if ctx.DIFER() != nil {
		return operaciones.Not(v.Visit(ctx.Expresion(0)))
	}

	// para or
	if ctx.OR() != nil {
		return operaciones.Or(v.Visit(ctx.Expresion(0)), v.Visit(ctx.Expresion(1)))
	}

	// para and
	if ctx.AND() != nil {
		return operaciones.And(v.Visit(ctx.Expresion(0)), v.Visit(ctx.Expresion(1)))
	}

	/* EXPRESIONES NATIVAS */
	if ctx.CADENA() != nil {
		text := ctx.CADENA().GetText()
		return operaciones.SecuenciaDeEscape(text[1 : len(text)-1])
	}
	if ctx.RUNE() != nil {
		text := ctx.RUNE().GetText()
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
	if ctx.AccesoElementoSlice() != nil {
		return v.Visit(ctx.AccesoElementoSlice())
	}
	// Funciones embebida

	if ctx.FnAtoi() != nil {
		return v.Visit(ctx.FnAtoi())
	}
	if ctx.FnParseToFloat() != nil {
		return v.Visit(ctx.FnParseToFloat())
	}
	if ctx.FnTypeOf() != nil {
		return v.Visit(ctx.FnTypeOf())
	}
	if ctx.FnAppend() != nil {
		return v.Visit(ctx.FnAppend())
	}
	if ctx.FnIndexOf() != nil {
		return v.Visit(ctx.FnIndexOf())
	}
	if ctx.FnJoin() != nil {
		return v.Visit(ctx.FnJoin())
	}
	if ctx.FnLen() != nil {
		return v.Visit(ctx.FnLen())
	}

	return 0
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
func obtenerValores(ctx parser.IListaExprContext, visitor *EvalVisitor) []interface{} {
	valores := []interface{}{}

	if ctx == nil {
		return valores
	}

	listaCtx, ok := ctx.(*parser.ListaExprContext)
	if !ok {
		return valores
	}

	for _, expr := range listaCtx.AllExpresion() {
		valores = append(valores, visitor.Visit(expr))
	}

	return valores
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

func tiposCompatibles(declarado string, real string) bool {
	declarado = strings.ToLower(declarado)
	real = strings.ToLower(real)
	fmt.Printf("[DEBUG tiposCompatibles] declarado: %s, real: %s\n", declarado, real)

	if declarado == real {
		return true
	}
	if declarado == "float64" && real == "int" {
		return true
	}
	if strings.HasPrefix(declarado, "slice(") && strings.HasPrefix(real, "slice(") {
		inner1 := extraerTipoInterno(declarado)
		inner2 := extraerTipoInterno(real)
		return tiposCompatibles(inner1, inner2)
	}
	return false
}
func extraerTipoInterno(tipo string) string {
	if strings.HasPrefix(tipo, "slice(") && strings.HasSuffix(tipo, ")") {
		return tipo[6 : len(tipo)-1]
	}
	return tipo
}

// buscarSimboloRecursivo busca un símbolo en el entorno actual y sus padres
func buscarSimboloRecursivo(entorno *symbols.Entorno, id string) *symbols.Simbolo {
	currentEntorno := entorno
	for currentEntorno != nil {
		if simbolo, ok := currentEntorno.Simbolos[id]; ok {
			return simbolo
		}
		currentEntorno = currentEntorno.Padre
	}
	return nil
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

func (v *EvalVisitor) VisitAsignacionNormal(ctx *parser.AsignacionNormalContext) interface{} {
	id := ctx.IDENTIFICADOR().GetText()
	val := v.Visit(ctx.Expresion())

	sim := v.Tabla.BuscarSimbolo(id)
	if sim == nil {
		panic("Variable no declarada: " + id)
	}

	v.Tabla.ActualizarVariable(id, val)
	return nil
}

func (v *EvalVisitor) VisitAsignacionIncremento(ctx *parser.AsignacionIncrementoContext) interface{} {
	id := ctx.IDENTIFICADOR().GetText()
	sim := v.Tabla.BuscarSimbolo(id)
	if sim == nil {
		panic("Variable no declarada: " + id)
	}
	if val, ok := sim.Valor.(int); ok {
		v.Tabla.ActualizarVariable(id, val+1)
		fmt.Printf("Incremento: %s ahora vale %d\n", id, val+1)
	} else {
		panic("Incremento solo válido para enteros")
	}
	return nil
}

func (v *EvalVisitor) VisitAsignacionDecremento(ctx *parser.AsignacionDecrementoContext) interface{} {
	id := ctx.IDENTIFICADOR().GetText()
	sim := v.Tabla.BuscarSimbolo(id)
	if sim == nil {
		panic("Variable no declarada: " + id)
	}
	if val, ok := sim.Valor.(int); ok {
		v.Tabla.ActualizarVariable(id, val-1)
	} else {
		panic("Decremento solo válido para enteros")
	}
	return nil
}

// Bloque
func (v *EvalVisitor) VisitBloque(ctx *parser.BloqueContext) interface{} {
	instrucciones := ctx.AllInstrucciones()
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
	return v.IfCtrl.Ejecutar(v.Visit, ctx)
}

func (v *EvalVisitor) VisitForCondicional(ctx *parser.ForCondicionalContext) interface{} {
	return v.ejecutarForCondicional(ctx.Expresion(), ctx.Bloque())
}

func (v *EvalVisitor) VisitForClasico(ctx *parser.ForClasicoContext) interface{} {
	return v.ejecutarForClasico(ctx)
}

func (v *EvalVisitor) VisitForRange(ctx *parser.ForRangeContext) interface{} {
	return v.ejecutarForRange(ctx)
}

func (v *EvalVisitor) ejecutarForCondicional(cond parser.IExpresionContext, bloque parser.IBloqueContext) interface{} {
	var resultados []interface{}

	for {
		condVal := v.Visit(cond)
		boolVal, ok := condVal.(bool)
		if !ok {
			panic("Condición del for no es booleana")
		}
		if !boolVal {
			break
		}

		val := v.Visit(bloque)
		if list, ok := val.([]interface{}); ok {
			if len(list) > 0 {

				if _, isBreak := list[len(list)-1].(BreakSignal); isBreak {
					resultados = append(resultados, list[:len(list)-1]...)
					break
				}

				if _, isContinue := list[len(list)-1].(ContinueSignal); isContinue {
					resultados = append(resultados, list[:len(list)-1]...)
					continue
				}
			}
			resultados = append(resultados, list...)
		} else if _, isBreak := val.(BreakSignal); isBreak {
			break
		} else if _, isContinue := val.(ContinueSignal); isContinue {
			continue
		} else if val != nil {
			resultados = append(resultados, val)
		}
	}
	return resultados
}

func (v *EvalVisitor) ejecutarForClasico(ctx *parser.ForClasicoContext) interface{} {

	v.Visit(ctx.Asignacion(0))
	var resultado []interface{}

	for {
		condVal := v.Visit(ctx.Expresion())
		boolVal, ok := condVal.(bool)
		if !ok {
			panic("Condición del for clásico no es booleana")
		}
		if !boolVal {
			break
		}

		val := v.Visit(ctx.Bloque())

		if list, ok := val.([]interface{}); ok {
			if len(list) > 0 {

				if _, isBreak := list[len(list)-1].(BreakSignal); isBreak {
					resultado = append(resultado, list[:len(list)-1]...)
					break
				}

				if _, isContinue := list[len(list)-1].(ContinueSignal); isContinue {
					resultado = append(resultado, list[:len(list)-1]...)

					v.Visit(ctx.Asignacion(1))
					continue
				}
			}
			resultado = append(resultado, list...)
		} else if _, isBreak := val.(BreakSignal); isBreak {
			break
		} else if _, isContinue := val.(ContinueSignal); isContinue {

			v.Visit(ctx.Asignacion(1))
			continue
		} else if val != nil {
			resultado = append(resultado, val)
		}

		v.Visit(ctx.Asignacion(1))
	}

	return resultado
}

func (v *EvalVisitor) ejecutarForRange(ctx *parser.ForRangeContext) interface{} {
	idIndice := ctx.IDENTIFICADOR(0).GetText()
	idValor := ctx.IDENTIFICADOR(1).GetText()

	val := v.Visit(ctx.Expresion())
	lista, ok := val.([]interface{})
	if !ok {
		panic("El for-range espera una lista")
	}

	var resultados []interface{}

	for i, elem := range lista {
		v.Tabla.ActualizarVariable(idIndice, i)
		v.Tabla.ActualizarVariable(idValor, elem)
		v.Entorno[idValor] = elem

		val := v.Visit(ctx.Bloque())
		if list, ok := val.([]interface{}); ok {
			if len(list) > 0 {
				if _, isBreak := list[len(list)-1].(BreakSignal); isBreak {
					resultados = append(resultados, list[:len(list)-1]...)
					break
				}
				if _, isContinue := list[len(list)-1].(ContinueSignal); isContinue {
					resultados = append(resultados, list[:len(list)-1]...)
					continue
				}
			}
			resultados = append(resultados, list...)
		} else if _, isBreak := val.(BreakSignal); isBreak {
			break
		} else if _, isContinue := val.(ContinueSignal); isContinue {
			continue
		} else if val != nil {
			resultados = append(resultados, val)
		}
	}
	return resultados
}

func (v *EvalVisitor) VisitSSwitch(ctx *parser.SSwitchContext) interface{} {
	valorSwitch := v.Visit(ctx.Expresion())

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

func (v *EvalVisitor) checkTipo(valor interface{}, esperado string) bool {
	switch esperado {
	case "int":
		_, ok := valor.(int)
		return ok
	case "float":
		_, ok := valor.(float64)
		return ok
	case "string":
		_, ok := valor.(string)
		return ok
	}
	return false
}
