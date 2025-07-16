package traducciones

import (
	"VAsm/frontend/errors"
	"fmt"
	"math"
	"strings"
)

/*	ARITMETICAS 	*/
// funcion para sumar
func Sumar(left, right interface{}) interface{} {
	if _, ok := left.(string); ok {
		return fmt.Sprintf("%v%v", left, right)
	}
	if _, ok := right.(string); ok {
		return fmt.Sprintf("%v%v", left, right)
	}

	if l, ok := left.(int); ok {
		if r, ok := right.(int); ok {
			return l + r
		}
	}
	return toFloat64(left) + toFloat64(right)
}
// funcion para restar
func Restar(left, right interface{}) interface{} {
	fmt.Printf("[DEBUG Restar] left: %#v (%T), right: %#v (%T)\n", left, left, right, right)

	if l, ok := left.(int); ok {
		if r, ok := right.(int); ok {
			return l - r
		}
	}
	return toFloat64(left) - toFloat64(right)
}
// funcion para multiplicar
func Multiplicar(left, right interface{}) interface{} {

	if l, ok := left.(int); ok {
		if r, ok := right.(int); ok {
			return l * r
		}
	}
	return toFloat64(left) * toFloat64(right)
}
// funcion para division
func Division(left, right interface{}) interface{} {
	den := toFloat64(right)
	if den == 0 {
		return errors.SemanticError
	}
	return toFloat64(left) / den
}
// funcion para potencia
func Modulo(left, right interface{}) interface{} {
	if l, ok := left.(int); ok {
		if r, ok := right.(int); ok {
			if r == 0 {
				return errors.SemanticError
			}
			return l % r
		}
	}
	den := toFloat64(right)
	if den == 0 {
		return errors.SemanticError
	}
	return math.Mod(toFloat64(left), den)
}
// funcion negativa
func Negativo(valor interface{}) interface{} {
	switch v := valor.(type) {
	case int:
		return -v
	case float64:
		return -v
	default:
		panic("Error: Tipo no soportado para unario negativo")
	}
}

func toFloat64(val interface{}) float64 {
	fmt.Println("Valor recibido para conversión a float64:", val)
	switch v := val.(type) {
	case int:
		return float64(v)
	case float64:
		return v
	case nil:
		return 0.0
	default:
		panic(fmt.Errorf("Error: Tipo no soportado para operación aritmética: %v", val))
	}
}

	/*	RELACIONALES	*/
// funcion igualdad
func Igualdad(left, right interface{}) interface{} {
	if l, ok := left.(int); ok {
		if r, ok := right.(int); ok {
			return l == r
		}
	}
	return (left) == (right)
}
// funcion negacion
func NoIgualdad(left, right interface{}) interface{} {
	if l, ok := left.(int); ok {
		if r, ok := right.(int); ok {
			return l != r
		}
	}
	return (left) != (right)
}
// funcion menor igual
func MenorIgual(left, right interface{}) interface{} {
	if l, ok := left.(int); ok {
		if r, ok := right.(int); ok {
			return l <= r
		}
	}
	if l, ok := left.(float64); ok {
		if r, ok := right.(float64); ok {
			return l <= r
		}
	}
	if l, ok := left.(int); ok {
		if r, ok := right.(float64); ok {
			return float64(l) <= r
		}
	}
	if l, ok := left.(float64); ok {
		if r, ok := right.(int); ok {
			return l <= float64(r)
		}
	}
	return false
}
// funcion mayor igual
func MayorIgual(left, right interface{}) interface{} {
	if l, ok := left.(int); ok {
		if r, ok := right.(int); ok {
			return l >= r
		}
	}
	if l, ok := left.(float64); ok {
		if r, ok := right.(float64); ok {
			return l >= r
		}
	}
	if l, ok := left.(int); ok {
		if r, ok := right.(float64); ok {
			return float64(l) >= r
		}
	}
	if l, ok := left.(float64); ok {
		if r, ok := right.(int); ok {
			return l >= float64(r)
		}
	}
	return false
}
// funcion menor
func Menor(left, right interface{}) interface{} {
	if l, ok := left.(int); ok {
		if r, ok := right.(int); ok {
			return l < r
		}
	}
	if l, ok := left.(float64); ok {
		if r, ok := right.(float64); ok {
			return l < r
		}
	}
	if l, ok := left.(int); ok {
		if r, ok := right.(float64); ok {
			return float64(l) < r
		}
	}
	if l, ok := left.(float64); ok {
		if r, ok := right.(int); ok {
			return l < float64(r)
		}
	}
	return false
}
// funcion mayor
func Mayor(left, right interface{}) interface{} {
	if l, ok := left.(int); ok {
		if r, ok := right.(int); ok {
			return l > r
		}
	}
	if l, ok := left.(float64); ok {
		if r, ok := right.(float64); ok {
			return l > r
		}
	}
	if l, ok := left.(int); ok {
		if r, ok := right.(float64); ok {
			return float64(l) > r
		}
	}
	if l, ok := left.(float64); ok {
		if r, ok := right.(int); ok {
			return l > float64(r)
		}
	}
	return false
}

	/* 	BOOLEANAS 		*/
// funcion or
func Or(left, right interface{}) interface{} {
	return toBool(left) || toBool(right)
}

// funcion and
func And(left, right interface{}) interface{} {
	return toBool(left) && toBool(right)
}

// funcion not
func Not(val interface{}) interface{} {
	return !toBool(val)
}

// convetir valor a bool
func toBool(val interface{}) bool {
	switch v := val.(type) {
	case bool:
		return v
	case int:
		return v != 0
	case float64:
		return v != 0.0
	case string:
		return v != ""
	default:
		return false
	}
}

	/*	ESCAPES DE PRINTS	*/
func SecuenciaDeEscape(s string) string {
	processedString := s
	processedString = strings.ReplaceAll(processedString, "\\n", "\n")  // salto de línea
	processedString = strings.ReplaceAll(processedString, "\\t", "\t")  // tabulación
	processedString = strings.ReplaceAll(processedString, "\\\"", "\"") // comilla doble
	processedString = strings.ReplaceAll(processedString, "\\r", "\r")  // retorno de carro
	processedString = strings.ReplaceAll(processedString, "\\\\", "\\") // barra invertida
	processedString = strings.ReplaceAll(processedString, "\\'", "'")   // comilla simple
	return processedString
}