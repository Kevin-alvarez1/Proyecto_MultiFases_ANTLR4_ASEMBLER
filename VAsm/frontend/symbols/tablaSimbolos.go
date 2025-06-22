package symbols

import (
	"VAsm/frontend/errors"
	"fmt"
	"sort"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// Tipos de símbolos
const (
	Variable = "Variable"
	Funcion  = "Función"
	Metodo   = "Método"
)

// Tipos de datos
const (
	Entero      = "int"
	Decimal     = "float"
	Texto       = "string"
	Booleano    = "bool"
	Void        = "void"
	Desconocido = "desconocido"
)

// Entorno representa un ámbito de declaración
type Entorno struct {
	Nombre   string
	Padre    *Entorno
	Hijos    []*Entorno
	Simbolos map[string]*Simbolo
	Anterior *Entorno
}

// Simbolo representa un elemento en la tabla de símbolos
type Simbolo struct {
	ID          string
	TipoSimbolo string
	TipoDato    string
	Ambito      string
	Linea       int
	Columna     int
	Valor       interface{}
}

// TablaSimbolos maneja todos los símbolos del programa
type TablaSimbolos struct {
	EntornoGlobal     *Entorno
	EntornoActual     *Entorno
	Errores           *errors.ErrorTable
	EntornosFunciones map[string]*Entorno
}

func NewEntorno(padre *Entorno, nombre string) *Entorno {
	return &Entorno{
		Padre:    padre,
		Nombre:   nombre,
		Simbolos: make(map[string]*Simbolo),
		Hijos:    []*Entorno{},
	}
}

func (ts *TablaSimbolos) Obtener(nombre string) *Simbolo {
	entorno := ts.EntornoActual
	for entorno != nil {
		if simbolo, ok := entorno.Simbolos[nombre]; ok {
			return simbolo
		}
		entorno = entorno.Padre
	}
	return nil
}

func NewTablaSimbolos() *TablaSimbolos {
	global := NewEntorno(nil, "global")
	return &TablaSimbolos{
		EntornoGlobal: global,
		EntornoActual: global,
	}
}

// NuevoEntorno crea un nuevo ámbito
func (ts *TablaSimbolos) NuevoEntorno(nombre string) {
	nuevo := NewEntorno(ts.EntornoActual, nombre)
	ts.EntornoActual.Hijos = append(ts.EntornoActual.Hijos, nuevo)
	ts.EntornoActual = nuevo
}

// SalirEntorno regresa al ámbito padre
func (ts *TablaSimbolos) SalirEntorno() {
	if ts.EntornoActual.Padre != nil {
		ts.EntornoActual = ts.EntornoActual.Padre
	}
}

func (e *Entorno) BuscarSimbolo(nombre string) *Simbolo {
	if e == nil || e.Simbolos == nil {
		return nil
	}

	if simbolo, existe := e.Simbolos[nombre]; existe {
		return simbolo
	}

	if e.Padre != nil {
		return e.Padre.BuscarSimbolo(nombre)
	}
	return nil
}

func (ts *TablaSimbolos) DeclararVariable(id, tipo string, valor interface{}, ctx antlr.ParserRuleContext, nombreEntorno string) {
	for _, s := range ts.EntornoActual.Simbolos {
		if s.ID == id {
			// Ya existe en este entorno → reasignar valor
			s.Valor = valor
			return
		}
	}

	// Si no existe → declarar
	simbolo := &Simbolo{
		ID:          id,
		TipoSimbolo: "variable",
		TipoDato:    tipo,
		Valor:       valor,
		Ambito:      ts.EntornoActual.Nombre,
		Linea:       ctx.GetStart().GetLine(),
		Columna:     ctx.GetStart().GetColumn(),
	}
	ts.Insertar(simbolo)
}

func (ts *TablaSimbolos) DeclararVariableSimple(id, tipo string, valor interface{}, ctx antlr.ParserRuleContext, nombreEntorno string) {
	for _, s := range ts.EntornoActual.Simbolos {
		if s.ID == id {
			s.Valor = valor
			return
		}
	}
	if ts.EntornoActual != nil {
	} else {
		fmt.Println("DEBUG: ts.EntornoActual is NIL, cannot access its fields.")
	}
	simbolo := &Simbolo{
		ID:          id,
		TipoSimbolo: Variable,
		TipoDato:    tipo,
		Ambito:      ts.EntornoActual.Nombre,
		Linea:       ctx.GetStart().GetLine(),
		Columna:     ctx.GetStart().GetColumn(),
		Valor:       valor,
	}

	ts.Insertar(simbolo)
}

func (ts *TablaSimbolos) Insertar(simbolo *Simbolo) {

	ts.EntornoActual.Simbolos[simbolo.ID] = simbolo
}

// DeclararFuncion añade una función a la tabla
func (ts *TablaSimbolos) DeclararFuncion(id string, tipoRetorno string, ctx antlr.ParserRuleContext) bool {
	token := ctx.GetStart()
	for _, s := range ts.EntornoGlobal.Simbolos {
		if s.ID == id && (s.TipoSimbolo == Funcion || s.TipoSimbolo == Metodo) {
			msg := fmt.Sprintf("función '%s' ya declarada (línea %d:%d)", id, s.Linea, s.Columna)
			ts.Errores.NewSemanticError(token, msg)
			return false
		}
	}

	simbolo := &Simbolo{
		ID:          id,
		TipoSimbolo: Funcion,
		TipoDato:    tipoRetorno,
		Ambito:      "Global",
		Linea:       token.GetLine(),
		Columna:     token.GetColumn(),
	}

	ts.EntornoGlobal.Simbolos[id] = simbolo
	return true
}

// BuscarSimbolo busca un símbolo en todos los ámbitos
func (ts *TablaSimbolos) BuscarSimbolo(id string) *Simbolo {
	actual := ts.EntornoActual
	for actual != nil {
		for _, s := range actual.Simbolos {
			if s.ID == id {
				return s
			}
		}
		actual = actual.Padre
	}
	return nil
}

func (ts *TablaSimbolos) GenerarReporte() []*Simbolo {
	var todosSimbolos []*Simbolo

	var recolectar func(e *Entorno)
	recolectar = func(e *Entorno) {
		for _, simbolo := range e.Simbolos {
			todosSimbolos = append(todosSimbolos, simbolo)
		}

		// Recorrer entornos hijos para recolectar sus símbolos
		for _, hijo := range e.Hijos {
			recolectar(hijo)
		}
	}

	recolectar(ts.EntornoGlobal)

	// Ordenar por Ámbito, Línea y Columna
	sort.Slice(todosSimbolos, func(i, j int) bool {
		if todosSimbolos[i].Ambito == todosSimbolos[j].Ambito {
			if todosSimbolos[i].Linea == todosSimbolos[j].Linea {
				return todosSimbolos[i].Columna < todosSimbolos[j].Columna
			}
			return todosSimbolos[i].Linea < todosSimbolos[j].Linea
		}
		return todosSimbolos[i].Ambito < todosSimbolos[j].Ambito
	})

	return todosSimbolos
}
func (ts *TablaSimbolos) ActualizarVariable(id string, valor interface{}) {
	if ts.EntornoActual.Simbolos == nil {
		panic("Entorno no inicializado")
	}
	if s, existe := ts.EntornoActual.Simbolos[id]; existe {
		s.Valor = valor
	} else {
		panic("Variable no encontrada: " + id)
	}
}

// ToString convierte la tabla de símbolos a formato de texto
func (ts *TablaSimbolos) ToString() string {
	sb := strings.Builder{}
	sb.WriteString("TABLA DE SÍMBOLOS\n")
	sb.WriteString("| ID | Tipo Símbolo | Tipo Dato | Ámbito | Línea | Columna |\n")
	sb.WriteString("|----|--------------|-----------|---------------------|-------|---------|\n")

	for _, s := range ts.GenerarReporte() {
		sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %d | %d |\n",
			s.ID, s.TipoSimbolo, s.TipoDato, s.Ambito, s.Linea, s.Columna))
	}

	return sb.String()
}
func (e *Entorno) ExisteSimbolo(nombre string) bool {
	_, existe := e.Simbolos[nombre] // Asumiendo que `Simbolos` es map[string]*Simbolo
	return existe
}
func (e *Entorno) AgregarSimbolo(id, tipoSimbolo, tipoDato, ambito string, linea, columna int, valor interface{}) {
	if e.Simbolos == nil {
		e.Simbolos = make(map[string]*Simbolo)
	}
	e.Simbolos[id] = &Simbolo{
		ID:          id,
		TipoSimbolo: tipoSimbolo,
		TipoDato:    tipoDato,
		Ambito:      ambito,
		Linea:       linea,
		Columna:     columna,
		Valor:       valor,
	}
}

func (ts *TablaSimbolos) Imprimir() {
	var recorrer func(e *Entorno, nivel int)
	recorrer = func(e *Entorno, nivel int) {
		fmt.Printf("%sEntorno: %s\n", strings.Repeat("  ", nivel), e.Nombre)
		for id, sim := range e.Simbolos {
			fmt.Printf("%s  %s (%s) = %v\n", strings.Repeat("  ", nivel), id, sim.TipoDato, sim.Valor)
		}
		for _, hijo := range e.Hijos {
			recorrer(hijo, nivel+1)
		}
	}
	recorrer(ts.EntornoGlobal, 0)
}
