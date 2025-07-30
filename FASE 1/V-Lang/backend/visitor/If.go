package visitor

import (
	parser "vlang/backend/analizador/parser"

	"github.com/antlr4-go/antlr/v4"
)

type If struct{}

func NewIf() *If {
	return &If{}
}

func (i *If) Ejecutar(visitFn func(antlr.ParseTree) interface{}, ctx *parser.SifContext) interface{} {
	condicion := visitFn(ctx.Expresion())
	condBool, ok := condicion.(bool)
	if !ok {
		panic("Condición del IF no es booleana")
	}

	if condBool {
		result := i.procesarBloqueConRetorno(visitFn, ctx.Bloque(0))
		// Verificar si el resultado es un Retorno directamente
		if ret, ok := result.(Retorno); ok {
			return ret
		}
		return result
	}

	for j := 0; j < len(ctx.AllElseifPart()); j++ {
		elseifCtx := ctx.ElseifPart(j)
		condElseif := visitFn(elseifCtx.Expresion())
		condElseifBool, ok := condElseif.(bool)
		if !ok {
			panic("Condición del ELSE IF no es booleana")
		}
		if condElseifBool {
			result := i.procesarBloqueConRetorno(visitFn, elseifCtx.Bloque())
			if ret, ok := result.(Retorno); ok {
				return ret
			}
			return result
		}
	}

	if ctx.ELSE() != nil {
		result := i.procesarBloqueConRetorno(visitFn, ctx.Bloque(1))
		if ret, ok := result.(Retorno); ok {
			return ret
		}
		return result
	}

	return nil
}

func (i *If) procesarBloqueConRetorno(visitFn func(antlr.ParseTree) interface{}, ctx parser.IBloqueContext) interface{} {
	result := i.ejecutarBloque(visitFn, ctx)

	// Manejo mejorado del retorno
	switch v := result.(type) {
	case Retorno:
		return v // Devuelve tanto return con valor como sin valor
	case []interface{}:
		// Busca recursivamente cualquier Retorno en la lista
		for _, item := range v {
			if ret, ok := item.(Retorno); ok {
				return ret
			}
		}
	}

	return result
}

func (i *If) ejecutarBloque(visitFn func(antlr.ParseTree) interface{}, ctx parser.IBloqueContext) interface{} {
	instrucciones := ctx.AllInstrucciones()
	var resultados []interface{}

	for _, instruccionCtx := range instrucciones {
		res := visitFn(instruccionCtx)

		// Manejar señales de control inmediatamente
		switch v := res.(type) {
		case Retorno:
			return v // Aquí se captura el Retorno correctamente
		case BreakSignal:
			return v
		case ContinueSignal:
			return v
		}

		if list, ok := res.([]interface{}); ok {
			// Buscar señales de control en listas anidadas
			for _, item := range list {
				switch v := item.(type) {
				case Retorno:
					return v
				case BreakSignal:
					return v
				case ContinueSignal:
					return v
				}
			}
			resultados = append(resultados, list...)
		} else if res != nil {
			resultados = append(resultados, res)
		}
	}

	return resultados
}

// buscarRetornoEnResultados busca recursivamente un Retorno en los resultados
func buscarRetornoEnResultados(result interface{}) *Retorno {
	switch v := result.(type) {
	case Retorno:
		return &v
	case []interface{}:
		for _, item := range v {
			if ret := buscarRetornoEnResultados(item); ret != nil {
				return ret
			}
		}
	}
	return nil
}
