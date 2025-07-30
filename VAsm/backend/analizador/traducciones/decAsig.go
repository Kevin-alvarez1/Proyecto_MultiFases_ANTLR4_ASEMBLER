package traducciones

import (
	"VAsm/frontend/symbols"
	"errors"
	"fmt"
	"strings"
)

var contadorFloat int
var contadorString int
var variablesReservadas = make(map[string]bool)

// ProcesarDeclaracionMultiple recibe el visitor como parámetro
func ProcesarDeclaracionMultiple(
	ids []string,
	valores []interface{},
	tipo string,
	tabla *symbols.TablaSimbolos,
	outputASM *strings.Builder,
) error {
	fmt.Println("GENERANDO DECLARACIÓN ASM DE:", ids)

	for i, id := range ids {
		valor := valores[i]
		valorTipo := inferirTipo(valor)

		if !tiposCompatibles(tipo, valorTipo) {
			msg := fmt.Sprintf("Error: no se puede asignar %s a variable %s de tipo %s",
				valorTipo, id, tipo)
			return errors.New(msg)
		}

		generarCodigoDeclaracion(id, tipo, valor, outputASM)
		reservarVariableEnData(id, tipo)

		simbolo := &symbols.Simbolo{
			ID:          id,
			TipoDato:    tipo,
			Valor:       valor,
			Ambito:      tabla.EntornoActual.Nombre,
			TipoSimbolo: "variable",
		}
		tabla.EntornoActual.Simbolos[id] = simbolo
	}
	return nil
}

func generarCodigoDeclaracion(id, tipo string, valor interface{}, outputASM *strings.Builder) {
	reservarVariableEnData(id, tipo)

	switch val := valor.(type) {
	case int:
		generarCodigoInt(id, val, outputASM)

	case float64:
		generarCodigoFloat(id, val, outputASM)

	case bool:
		generarCodigoBool(id, val, outputASM)

	case string:
		generarCodigoString(id, val, outputASM)

	case []interface{}:
		// Para slices, solo mostramos comentario con contenido
		codigo := fmt.Sprintf("// Declaración %s tipo slice con %d elementos: [", id, len(val))
		for i, elem := range val {
			if i > 0 {
				codigo += ", "
			}
			codigo += fmt.Sprintf("%v", elem)
		}
		codigo += "]\n"
		outputASM.WriteString(codigo)

	default:
		outputASM.WriteString(fmt.Sprintf("// %s: tipo no soportado para ASM (%T)\n", id, val))
	}
}

func GenerarCodigoDeclaracionSinTipo(id, tipo string, valor interface{}, outputASM *strings.Builder) error {
	reservarVariableEnData(id, tipo)

	switch tipo {
	case "int":
		val, ok := valor.(int)
		if !ok {
			return fmt.Errorf("Valor para variable %s no es int", id)
		}
		generarCodigoInt(id, val, outputASM)

	case "float", "float64":
		val, ok := valor.(float64)
		if !ok {
			return fmt.Errorf("Valor para variable %s no es float64", id)
		}
		generarCodigoFloat(id, val, outputASM)

	case "string":
		val, ok := valor.(string)
		if !ok {
			return fmt.Errorf("Valor para variable %s no es string", id)
		}
		generarCodigoString(id, val, outputASM)

	case "bool":
		val, ok := valor.(bool)
		if !ok {
			return fmt.Errorf("Valor para variable %s no es bool", id)
		}
		generarCodigoBool(id, val, outputASM)

	default:
		return fmt.Errorf("Tipo %s no soportado para generación ASM", tipo)
	}
	return nil
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
func generarCodigoInt(id string, valor int, outputASM *strings.Builder) {
	outputASM.WriteString(fmt.Sprintf("mov x10, #%d\n", valor))
	outputASM.WriteString(fmt.Sprintf("adr x11, %s\n", id))
	outputASM.WriteString("str x10, [x11]\n\n")
}

func generarCodigoFloat(id string, val float64, outputASM *strings.Builder) {
	etiqueta := fmt.Sprintf("float_val_%s_%d", id, contadorFloat)
	contadorFloat++
	DataBuilder.WriteString(fmt.Sprintf("%s: .float %f\n", etiqueta, val))
	outputASM.WriteString(fmt.Sprintf("ldr s0, =%s\n", etiqueta))
	outputASM.WriteString(fmt.Sprintf("adr x11, %s\n", id))
	outputASM.WriteString("str s0, [x11]\n\n")
}

func generarCodigoString(id string, valor string, outputASM *strings.Builder) string {
	etiqueta := fmt.Sprintf("String_val_%s_%d", id, contadorString)
	contadorString++
	escaped := escape(valor)
	DataBuilder.WriteString(fmt.Sprintf("%s: .ascii \"%s\"\n", etiqueta, escaped))
	outputASM.WriteString(fmt.Sprintf("adr x10, %s\n", etiqueta))
	outputASM.WriteString(fmt.Sprintf("adr x11, %s\n", id))
	outputASM.WriteString("str x10, [x11]\n\n")
	return etiqueta
}

func generarCodigoBool(id string, val bool, outputASM *strings.Builder) {
	bit := 0
	if val {
		bit = 1
	}
	outputASM.WriteString(fmt.Sprintf("mov x10, #%d\n", bit))
	outputASM.WriteString(fmt.Sprintf("adr x11, %s\n", id))
	outputASM.WriteString("str x10, [x11]\n\n")
}
func reservarVariableEnData(id string, tipo string) {
	if variablesReservadas[id] {
		return
	}
	fmt.Println("RESERVANDO EN .data:", id)

	switch tipo {
	case "int", "bool":
		DataBuilder.WriteString(fmt.Sprintf("%s: .quad 0\n", id))
	case "float":
		DataBuilder.WriteString(fmt.Sprintf("%s: .float 0.0\n", id))
	case "string":
		DataBuilder.WriteString(fmt.Sprintf("%s: .quad 0\n", id)) // puntero o vacío
	default:
		DataBuilder.WriteString(fmt.Sprintf("%s: .quad 0\n", id))
	}
	variablesReservadas[id] = true
}

func GenerarSumaASM(id1, id2 string, builder *strings.Builder, tipo string) string {
	labelResult := fmt.Sprintf("resultado_suma_%d", contadorSuma)
	contadorSuma++

	if !variablesReservadas[labelResult] {
		reservarVariableEnData(labelResult, tipo)
		variablesReservadas[labelResult] = true
	}

	builder.WriteString(fmt.Sprintf("\n// Suma %d\n", contadorSuma))

	if tipo == "int" {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
		builder.WriteString("ldr x10, [x10]\n")
		builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
		builder.WriteString("ldr x11, [x11]\n")
		builder.WriteString("add x12, x10, x11\n")
		builder.WriteString(fmt.Sprintf("adr x13, %s\n", labelResult))
		builder.WriteString("str x12, [x13]\n\n")
	}

	if tipo == "float" {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
		builder.WriteString("ldr s0, [x10]\n")
		builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
		builder.WriteString("ldr s1, [x11]\n")
		builder.WriteString("fadd s2, s0, s1\n")
		builder.WriteString(fmt.Sprintf("adr x12, %s\n", labelResult))
		builder.WriteString("str s2, [x12]\n\n")
	}

	return labelResult
}

func GenerarRestaASM(id1, id2 string, builder *strings.Builder, tipo string) string {
	labelResult := fmt.Sprintf("resultado_resta_%d", contadorResta)
	contadorResta++

	if !variablesReservadas[labelResult] {
		reservarVariableEnData(labelResult, tipo)
		variablesReservadas[labelResult] = true
	}

	builder.WriteString(fmt.Sprintf("\n// Resta %d\n", contadorResta))

	if tipo == "int" {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
		builder.WriteString("ldr x10, [x10]\n")
		builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
		builder.WriteString("ldr x11, [x11]\n")
		builder.WriteString("sub x12, x10, x11\n")
		builder.WriteString(fmt.Sprintf("adr x13, %s\n", labelResult))
		builder.WriteString("str x12, [x13]\n\n")
	}

	if tipo == "float" {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
		builder.WriteString("ldr s0, [x10]\n")
		builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
		builder.WriteString("ldr s1, [x11]\n")
		builder.WriteString("fsub s2, s0, s1\n")
		builder.WriteString(fmt.Sprintf("adr x12, %s\n", labelResult))
		builder.WriteString("str s2, [x12]\n\n")
	}

	return labelResult
}

func GenerarMultiplicacionASM(id1, id2 string, builder *strings.Builder, tipo string) string {
	labelResult := fmt.Sprintf("resultado_mul_%d", contadorMulti)
	contadorMulti++

	if !variablesReservadas[labelResult] {
		reservarVariableEnData(labelResult, tipo)
		variablesReservadas[labelResult] = true
	}

	builder.WriteString(fmt.Sprintf("\n// Multiplicación %d\n", contadorMulti))

	if tipo == "int" {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
		builder.WriteString("ldr x10, [x10]\n")
		builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
		builder.WriteString("ldr x11, [x11]\n")
		builder.WriteString("mul x12, x10, x11\n")
		builder.WriteString(fmt.Sprintf("adr x13, %s\n", labelResult))
		builder.WriteString("str x12, [x13]\n\n")
	}

	if tipo == "float" {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
		builder.WriteString("ldr s0, [x10]\n")
		builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
		builder.WriteString("ldr s1, [x11]\n")
		builder.WriteString("fmul s2, s0, s1\n")
		builder.WriteString(fmt.Sprintf("adr x12, %s\n", labelResult))
		builder.WriteString("str s2, [x12]\n\n")
	}

	return labelResult
}

func GenerarDivisionASM(id1, id2 string, builder *strings.Builder, tipo string) string {
	labelResult := fmt.Sprintf("resultado_div_%d", contadorDivision)
	contadorDivision++

	if !variablesReservadas[labelResult] {
		reservarVariableEnData(labelResult, tipo)
		variablesReservadas[labelResult] = true
	}

	builder.WriteString(fmt.Sprintf("\n// División %d\n", contadorDivision))

	if tipo == "int" {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
		builder.WriteString("ldr x10, [x10]\n")
		builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
		builder.WriteString("ldr x11, [x11]\n")
		builder.WriteString("sdiv x12, x10, x11\n")
		builder.WriteString(fmt.Sprintf("adr x13, %s\n", labelResult))
		builder.WriteString("str x12, [x13]\n\n")
	}

	if tipo == "float" {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
		builder.WriteString("ldr s0, [x10]\n")
		builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
		builder.WriteString("ldr s1, [x11]\n")
		builder.WriteString("fdiv s2, s0, s1\n")
		builder.WriteString(fmt.Sprintf("adr x12, %s\n", labelResult))
		builder.WriteString("str s2, [x12]\n\n")
	}

	return labelResult
}

func GenerarModuloASM(id1, id2 string, builder *strings.Builder, tipo1, tipo2 string) string {
	labelResult := fmt.Sprintf("resultado_modulo_%d", contadorMod)
	contadorMod++

	if !variablesReservadas[labelResult] {
		reservarVariableEnData(labelResult, "int")
		variablesReservadas[labelResult] = true
	}
	builder.WriteString(fmt.Sprintf("\n// Módulo %d\n", contadorMod))
	builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
	if tipo1 == "float" {
		builder.WriteString("ldr s0, [x10]\n")
		builder.WriteString("fcvtzs x10, s0\n")
	} else {
		builder.WriteString("ldr x10, [x10]\n")
	}
	builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
	if tipo2 == "float" {
		builder.WriteString("ldr s1, [x11]\n")
		builder.WriteString("fcvtzs x11, s1\n")
	} else {
		builder.WriteString("ldr x11, [x11]\n")
	}
	builder.WriteString("sdiv x12, x10, x11\n")
	builder.WriteString("mul x12, x12, x11\n")
	builder.WriteString("sub x12, x10, x12\n")
	builder.WriteString(fmt.Sprintf("adr x13, %s\n", labelResult))
	builder.WriteString("str x12, [x13]\n")

	return labelResult
}

func GenerarDiferenteASM(id1, id2 string, builder *strings.Builder, tipo1, tipo2 string) {
	builder.WriteString(fmt.Sprintf("// Comparación != entre %s y %s\n", id1, id2))

	if tipo1 == "float" || tipo2 == "float" {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
		if tipo1 == "int" {
			builder.WriteString("ldr x12, [x10]\n")
			builder.WriteString("scvtf s0, x12\n")
		} else {
			builder.WriteString("ldr s0, [x10]\n")
		}

		builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
		if tipo2 == "int" {
			builder.WriteString("ldr x13, [x11]\n")
			builder.WriteString("scvtf s1, x13\n")
		} else {
			builder.WriteString("ldr s1, [x11]\n")
		}

		builder.WriteString("fcmp s0, s1\n")
		builder.WriteString("cset x0, ne\n")
	} else {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
		builder.WriteString("ldr x10, [x10]\n")
		builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
		builder.WriteString("ldr x11, [x11]\n")
		builder.WriteString("cmp x10, x11\n")
		builder.WriteString("cset x0, ne\n")
	}
}

func GenerarIgualASM(id1, id2 string, builder *strings.Builder, tipo1, tipo2 string) {
	builder.WriteString(fmt.Sprintf("// Comparación == entre %s y %s\n", id1, id2))

	if tipo1 == "float" || tipo2 == "float" {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
		if tipo1 == "int" {
			builder.WriteString("ldr x12, [x10]\n")
			builder.WriteString("scvtf s0, x12\n")
		} else {
			builder.WriteString("ldr s0, [x10]\n")
		}

		builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
		if tipo2 == "int" {
			builder.WriteString("ldr x13, [x11]\n")
			builder.WriteString("scvtf s1, x13\n")
		} else {
			builder.WriteString("ldr s1, [x11]\n")
		}
		builder.WriteString("fcmp s0, s1\n")
		builder.WriteString("cset x0, eq\n")

	} else {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
		builder.WriteString("ldr x10, [x10]\n")
		builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
		builder.WriteString("ldr x11, [x11]\n")
		builder.WriteString("cmp x10, x11\n")
		builder.WriteString("cset x0, eq\n")
	}
}

func GenerarMenorIgualASM(id1, id2 string, builder *strings.Builder, tipo1, tipo2 string) {
	builder.WriteString(fmt.Sprintf("// Comparación <= entre %s y %s\n", id1, id2))

	if tipo1 == "float" || tipo2 == "float" {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
		if tipo1 == "int" {
			builder.WriteString("ldr x12, [x10]\n")
			builder.WriteString("scvtf s0, x12\n")
		} else {
			builder.WriteString("ldr s0, [x10]\n")
		}
		builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
		if tipo2 == "int" {
			builder.WriteString("ldr x13, [x11]\n")
			builder.WriteString("scvtf s1, x13\n")
		} else {
			builder.WriteString("ldr s1, [x11]\n")
		}

		builder.WriteString("fcmp s0, s1\n")
		builder.WriteString("cset x0, le\n")
	} else {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
		builder.WriteString("ldr x10, [x10]\n")
		builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
		builder.WriteString("ldr x11, [x11]\n")
		builder.WriteString("cmp x10, x11\n")
		builder.WriteString("cset x0, le\n")
	}
}

func GenerarMayorIgualASM(id1, id2 string, builder *strings.Builder, tipo1, tipo2 string) {
	builder.WriteString(fmt.Sprintf("// Comparación >= entre %s y %s\n", id1, id2))

	if tipo1 == "float" || tipo2 == "float" {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
		if tipo1 == "int" {
			builder.WriteString("ldr x12, [x10]\n")
			builder.WriteString("scvtf s0, x12\n")
		} else {
			builder.WriteString("ldr s0, [x10]\n")
		}

		builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
		if tipo2 == "int" {
			builder.WriteString("ldr x13, [x11]\n")
			builder.WriteString("scvtf s1, x13\n")
		} else {
			builder.WriteString("ldr s1, [x11]\n")
		}
		builder.WriteString("fcmp s0, s1\n")
		builder.WriteString("cset x0, ge\n")

	} else {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
		builder.WriteString("ldr x10, [x10]\n")
		builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
		builder.WriteString("ldr x11, [x11]\n")
		builder.WriteString("cmp x10, x11\n")
		builder.WriteString("cset x0, ge\n")
	}
}

func GenerarMenorASM(id1, id2 string, builder *strings.Builder, tipo1, tipo2 string) {
	builder.WriteString(fmt.Sprintf("// Comparación < entre %s y %s\n", id1, id2))

	if tipo1 == "float" || tipo2 == "float" {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
		if tipo1 == "int" {
			builder.WriteString("ldr x12, [x10]\n")
			builder.WriteString("scvtf s0, x12\n")
		} else {
			builder.WriteString("ldr s0, [x10]\n")
		}

		builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
		if tipo2 == "int" {
			builder.WriteString("ldr x13, [x11]\n")
			builder.WriteString("scvtf s1, x13\n")
		} else {
			builder.WriteString("ldr s1, [x11]\n")
		}

		builder.WriteString("fcmp s0, s1\n")
		builder.WriteString("cset x0, lt\n")

	} else {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
		builder.WriteString("ldr x10, [x10]\n")
		builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
		builder.WriteString("ldr x11, [x11]\n")
		builder.WriteString("cmp x10, x11\n")
		builder.WriteString("cset x0, lt\n")
	}
}

func GenerarMayorASM(id1, id2 string, builder *strings.Builder, tipo1, tipo2 string) {
	builder.WriteString(fmt.Sprintf("// Comparación > entre %s y %s\n", id1, id2))

	if tipo1 == "float" || tipo2 == "float" {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
		if tipo1 == "int" {
			builder.WriteString("ldr x12, [x10]\n")
			builder.WriteString("scvtf s0, x12\n")
		} else {
			builder.WriteString("ldr s0, [x10]\n")
		}
		builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
		if tipo2 == "int" {
			builder.WriteString("ldr x13, [x11]\n")
			builder.WriteString("scvtf s1, x13\n")
		} else {
			builder.WriteString("ldr s1, [x11]\n")
		}

		builder.WriteString("fcmp s0, s1\n")
		builder.WriteString("cset x0, gt\n")

	} else {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
		builder.WriteString("ldr x10, [x10]\n")
		builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
		builder.WriteString("ldr x11, [x11]\n")
		builder.WriteString("cmp x10, x11\n")
		builder.WriteString("cset x0, gt\n")
	}
}
func GenerarAndASM(id1, id2 string, builder *strings.Builder) {
	builder.WriteString(fmt.Sprintf("// Operación lógica AND entre %s && %s\n", id1, id2))

	builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
	builder.WriteString("ldr x10, [x10]\n")

	builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
	builder.WriteString("ldr x11, [x11]\n")

	builder.WriteString("and x12, x10, x11\n")
	builder.WriteString("cmp x12, #0\n")
	builder.WriteString("cset x0, ne\n")
}

func GenerarOrASM(id1, id2 string, builder *strings.Builder) {
	builder.WriteString(fmt.Sprintf("// Operación lógica OR entre %s || %s\n", id1, id2))
	builder.WriteString(fmt.Sprintf("adr x10, %s\n", id1))
	builder.WriteString("ldr x10, [x10]\n")
	builder.WriteString(fmt.Sprintf("adr x11, %s\n", id2))
	builder.WriteString("ldr x11, [x11]\n")
	builder.WriteString("orr x12, x10, x11\n")
	builder.WriteString("cmp x12, #0\n")
	builder.WriteString("cset x0, ne\n") 
}

func GenerarNotASM(id string, builder *strings.Builder) {
	builder.WriteString(fmt.Sprintf("// Operación lógica NOT !%s\n", id))
	builder.WriteString(fmt.Sprintf("adr x10, %s\n", id))
	builder.WriteString("ldr x10, [x10]\n")
	builder.WriteString("cmp x10, #0\n")
	builder.WriteString("cset x0, eq\n")
}

func GenerarIncrementoASM(id string, valor interface{}, tipo string, builder *strings.Builder) {
	builder.WriteString(fmt.Sprintf("// %s += %v\n", id, valor))

	if tipo == "float" {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id))
		builder.WriteString("ldr s0, [x10]\n")

		if intVal, ok := valor.(int); ok {
			builder.WriteString(fmt.Sprintf("mov x11, #%d\n", intVal))
			builder.WriteString("scvtf s1, x11\n")
		} else if floatVal, ok := valor.(float64); ok {
			builder.WriteString(fmt.Sprintf("ldr s1, =%g\n", floatVal))
		}

		builder.WriteString("fadd s2, s0, s1\n")
		builder.WriteString(fmt.Sprintf("adr x12, %s\n", id))
		builder.WriteString("str s2, [x12]\n\n")
	} else {
		builder.WriteString(fmt.Sprintf("adr x0, %s\n", id))
		builder.WriteString("ldr x0, [x0]\n")
		builder.WriteString(fmt.Sprintf("add x0, x0, #%d\n", valor.(int)))
		builder.WriteString(fmt.Sprintf("adr x1, %s\n", id))
		builder.WriteString("str x0, [x1]\n\n")
	}
}

func GenerarDecrementoASM(id string, valor interface{}, tipo string, builder *strings.Builder) {
	builder.WriteString(fmt.Sprintf("// %s -= %v\n", id, valor))

	if tipo == "float" {
		builder.WriteString(fmt.Sprintf("adr x10, %s\n", id))
		builder.WriteString("ldr s0, [x10]\n")

		if intVal, ok := valor.(int); ok {
			builder.WriteString(fmt.Sprintf("mov x11, #%d\n", intVal))
			builder.WriteString("scvtf s1, x11\n")
		} else if floatVal, ok := valor.(float64); ok {
			builder.WriteString(fmt.Sprintf("ldr s1, =%g\n", floatVal))
		}

		builder.WriteString("fsub s2, s0, s1\n")
		builder.WriteString(fmt.Sprintf("adr x12, %s\n", id))
		builder.WriteString("str s2, [x12]\n\n")
	} else {
		builder.WriteString(fmt.Sprintf("adr x0, %s\n", id))
		builder.WriteString("ldr x0, [x0]\n")
		builder.WriteString(fmt.Sprintf("sub x0, x0, #%d\n", valor.(int)))
		builder.WriteString(fmt.Sprintf("adr x1, %s\n", id))
		builder.WriteString("str x0, [x1]\n\n")
	}
}

func ReservarVariableSiNoExiste(id, tipo string) {
	if !variablesReservadas[id] {
		reservarVariableEnData(id, tipo)
	}
}

var contadorIf int

func GenerarEtiquetaUnica(prefijo string) string {
	contadorIf++
	return fmt.Sprintf("%s_%d", prefijo, contadorIf)
}