package traducciones

import (
	"VAsm/frontend/symbols"
	"errors"
	"fmt"
	"strings"
)

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
		codigo := fmt.Sprintf("// Declaración %s de tipo %s\nmov x0, #%d\n\n", id, tipo, val)
		outputASM.WriteString(codigo)

	case float64:
		// AArch64 no soporta mover float inmediato directamente, se necesita cargar desde memoria,
		codigo := fmt.Sprintf("// Declaración %s de tipo %s (float64) valor: %f\n", id, tipo, val)
		outputASM.WriteString(codigo)

	case bool:
		var boolVal int
		if val {
			boolVal = 1
		} else {
			boolVal = 0
		}
		codigo := fmt.Sprintf("// Declaración %s de tipo %s\nmov x0, #%d // bool\n\n", id, tipo, boolVal)
		outputASM.WriteString(codigo)

	case string:
		// Generar impresión usando la función para prints
		GenerarCodigoPrint(val, false)

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
