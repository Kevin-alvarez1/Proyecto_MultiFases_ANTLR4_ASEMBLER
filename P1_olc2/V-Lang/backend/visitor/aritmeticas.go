package visitor

import (
	"fmt"
	"math"
	"vlang/frontend/errors"
)

func normalizarValor(val interface{}) interface{} {
	// Si el valor es un Retorno, desempaqueta el valor real
	if ret, ok := val.(Retorno); ok {
		return ret.Valor
	} else if ret, ok := val.(*Retorno); ok {
		return ret.Valor
	}
	return val
}

// funcion para sumar
func Sumar(left, right interface{}) interface{} {
	// Si alguno es string, concatenar como string
	if _, ok := left.(string); ok {
		return fmt.Sprintf("%v%v", left, right)
	}
	if _, ok := right.(string); ok {
		return fmt.Sprintf("%v%v", left, right)
	}

	// Si ambos son int
	if l, ok := left.(int); ok {
		if r, ok := right.(int); ok {
			return l + r
		}
	}

	// Suma numérica como fallback
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
	// Primero intenta como enteros
	if l, ok := left.(int); ok {
		if r, ok := right.(int); ok {
			if r == 0 {
				return errors.SemanticError
			}
			return l % r
		}
	}

	// Si no son enteros, usar float64 y math.Mod
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
