package operaciones

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
