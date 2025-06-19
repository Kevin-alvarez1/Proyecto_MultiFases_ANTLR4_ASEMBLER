package errors

import (
	"sort"

	"github.com/antlr4-go/antlr/v4"
)

const (
	LexicalError  = "Error léxico"
	SyntaxError   = "Error sintáctico"
	SemanticError = "Error semántico"
	RuntimeError  = "Error runtime"
)

// Definimos una tabla de errores
type Error struct {
	Line   int
	Column int
	Msg    string
	Type   string
}

// La tabla de Err
type ErrorTable struct {
	Errors []Error
}

type SyntaxErrorListener struct {
	*antlr.DefaultErrorListener
	ErrorTable *ErrorTable
}

type LexicalErrorListener struct {
	*antlr.DefaultErrorListener
	ErrorTable *ErrorTable
}

type SemanticErrorListener struct {
	*antlr.DefaultErrorListener
	ErrorTable *ErrorTable
}

func (et *ErrorTable) AddError(line int, column int, msg string, errorType string) {
	et.Errors = append(et.Errors, Error{line, column, msg, errorType})
}

func (et *ErrorTable) NewLexicalError(line int, column int, msg string) {
	et.AddError(line, column, msg, LexicalError)
}

func (et *ErrorTable) NewSyntaxError(line int, column int, msg string) {
	et.AddError(line, column, msg, SyntaxError)
}

func (et *ErrorTable) NewSemanticError(token antlr.Token, msg string) {
	et.AddError(token.GetLine(), token.GetColumn(), msg, SemanticError)
}
func NewSemanticErrorListener(errorTable *ErrorTable) *SemanticErrorListener {
	return &SemanticErrorListener{
		ErrorTable: errorTable,
	}
}

func (et *ErrorTable) NewRuntimeError(line int, column int, msg string) {
	et.AddError(line, column, msg, RuntimeError)
}

func NewErrorTable() *ErrorTable {
	return &ErrorTable{}
}

func NewSyntaxErrorListener(errorTable *ErrorTable) *SyntaxErrorListener {
	return &SyntaxErrorListener{
		ErrorTable: errorTable,
	}
}

func (e *SyntaxErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {

	e.ErrorTable.AddError(
		line,
		column,
		msg,
		SyntaxError,
	)

}

func NewLexicalErrorListener() *LexicalErrorListener {
	return &LexicalErrorListener{
		ErrorTable: NewErrorTable(),
	}
}
func NewSemanticError() *SemanticErrorListener {
	return &SemanticErrorListener{
		ErrorTable: NewErrorTable(),
	}
}
func (e *SemanticErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {

	e.ErrorTable.AddError(
		line,
		column,
		msg,
		LexicalError,
	)

}
func (e *LexicalErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {

	e.ErrorTable.AddError(
		line,
		column,
		msg,
		LexicalError,
	)

}

func (et *ErrorTable) GetSortedErrors() []Error {
	errors := append([]Error{}, et.Errors...) // Copia
	sort.Slice(errors, func(i, j int) bool {
		if errors[i].Line == errors[j].Line {
			return errors[i].Column < errors[j].Column
		}
		return errors[i].Line < errors[j].Line
	})
	return errors
}
