package traducciones

import (
	"VAsm/frontend/symbols"
	"errors"
	"fmt"
	"strings"
)

var contadorFloat int
var contadorString int

// ProcesarDeclaracionMultiple recibe el visitor como parámetro
func ProcesarDeclaracionMultiple(
	ids []string,
	valores []interface{},
	tipo string,
	tabla *symbols.TablaSimbolos,
	outputASM *strings.Builder,
) error {
	for i, id := range ids {
		valor := valores[i]
		valorTipo := inferirTipo(valor)

		if !tiposCompatibles(tipo, valorTipo) {
			msg := fmt.Sprintf("Error: no se puede asignar %s a variable %s de tipo %s",
				valorTipo, id, tipo)
			return errors.New(msg)
		}

		generarCodigoDeclaracion(id, tipo, valor, outputASM)

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
	// Comentario y código ASM que mueve el valor al registro x0 (puedes adaptar si quieres más lógica)
	codigo := fmt.Sprintf("mov x0, #%d\n\n", valor)
	outputASM.WriteString(codigo)
}

func generarCodigoFloat(id string, val float64, outputASM *strings.Builder) {
	etiqueta := fmt.Sprintf("float_val_%s_%d", id, contadorFloat)
	contadorFloat++

	dataBuilder.WriteString(fmt.Sprintf("%s: .float %f\n", etiqueta, val))
	textBuilder.WriteString(fmt.Sprintf("ldr s0, =%s\n", etiqueta))
}
func generarCodigoString(id string, valor string, outputASM *strings.Builder) string {
	etiqueta := fmt.Sprintf("String_val_%s_%d", id, contadorString)
	contadorString++
	escaped := escape(valor)

	// Cadena única en .data
	dataBuilder.WriteString(fmt.Sprintf("%s: .ascii \"%s\"\n", etiqueta, escaped))
	return etiqueta
}

func generarCodigoBool(id string, val bool, outputASM *strings.Builder) {
	bit := 0
	if val {
		bit = 1
	}
	codigo := fmt.Sprintf("mov x0, #%d\n\n", bit)
	outputASM.WriteString(codigo)
}
