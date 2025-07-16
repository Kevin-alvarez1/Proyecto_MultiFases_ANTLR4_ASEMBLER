package operaciones

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
