// Code generated from backend/analizador/parser/gramatica.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // gramatica
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type gramaticaParser struct {
	*antlr.BaseParser
}

var GramaticaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gramaticaParserInit() {
	staticData := &GramaticaParserStaticData
	staticData.LiteralNames = []string{
		"", "'.'", "'+='", "'-='", "'print'", "'println'", "'fn'", "'mut'",
		"'slice'", "'if'", "'else'", "'switch'", "'case'", "'default'", "'for'",
		"'break'", "'continue'", "'return'", "'int'", "'float64'", "'string'",
		"'bool'", "'nil'", "'true'", "'false'", "'indexOf'", "'join'", "'len'",
		"'append'", "'struct'", "'func'", "'range'", "'Atoi'", "'parseFloat'",
		"'typeOf'", "'+'", "'-'", "'*'", "'/'", "'%'", "'('", "')'", "'{'",
		"'}'", "'['", "']'", "'!='", "'=='", "'<='", "'>='", "", "'<'", "'>'",
		"'||'", "'&&'", "'!'", "':'", "';'", "','", "'++'", "'--'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "PRINT", "PRINTLN", "FUNCIONES", "MUT", "SLICE", "IF",
		"ELSE", "SWITCH", "CASE", "DEFAULT", "FOR", "BREAK", "CONTINUE", "RETURN",
		"INT", "FLOAT", "STRING", "BOOL", "NIL", "TRUE", "FALSE", "INDEXOF",
		"JOIN", "LEN", "APPEND", "STRUCT", "FUNC", "RANGE", "ATOI", "PARSEFLOAT",
		"TYPEOF", "MAS", "MENOS", "POR", "DIV", "MODULO", "PAR1", "PAR2", "LLAV1",
		"LLAV2", "COR1", "COR2", "DIFERENTE", "IGUALDAD", "MENIGUAL", "MAYIGUAL",
		"IGUAL", "MENOR", "MAYOR", "OR", "AND", "DIFER", "DOSPTS", "PTCOMA",
		"COMA", "FINCREMENTO", "FDECREMENTO", "IDENTIFICADOR", "ENTERO", "DECIMAL",
		"CADENA", "RUNE", "COMENTARIO", "MULTICOMENTARIO", "WS",
	}
	staticData.RuleNames = []string{
		"init", "instrucciones", "instruccion", "print", "println", "declaracionMultiple",
		"declaracionMultipleSimple", "declaracionMultipleSinTipo", "asignacionMultiple",
		"structDef", "atributos", "structInst", "structInit", "listaStructs",
		"accesoStruct", "asigStruct", "funStruct", "caja", "valores", "actStruct",
		"bloqueFuncion", "llamadaFuncionesSinParametro", "llamadaFuncionesConParametro",
		"fnSinParametro", "fnConParametro", "sliceDef", "sliceLiteral", "accesoElementoSlice",
		"modificacionElementoSlice", "fnIndexOf", "fnJoin", "fnLen", "fnAppend",
		"declaracionSliceVacio", "tipoRetorno", "retorno", "fnTypeOf", "fnAtoi",
		"fnParseToFloat", "asigIncre", "asigDecre", "asignacion", "listaIDS",
		"listaPar", "parametro", "listaExpr", "listaExprList", "expresion",
		"sif", "elseifPart", "bloque", "sfor", "sSwitch", "break_", "continue_",
		"caseBlock", "defaultBlock", "tipos",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 68, 647, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		1, 0, 5, 0, 118, 8, 0, 10, 0, 12, 0, 121, 9, 0, 1, 1, 4, 1, 124, 8, 1,
		11, 1, 12, 1, 125, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 156, 8, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 4, 9, 191, 8, 9, 11, 9, 12, 9,
		192, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 5, 13, 213, 8,
		13, 10, 13, 12, 13, 216, 9, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 4, 17, 242, 8, 17,
		11, 17, 12, 17, 243, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 5, 20,
		263, 8, 20, 10, 20, 12, 20, 266, 9, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23,
		282, 8, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 3, 24, 294, 8, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 320, 8, 26,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 3, 27, 335, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 3, 28, 354, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 388, 8, 33, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 35, 3, 35, 395, 8, 35, 1, 35, 1, 35, 3, 35, 399, 8, 35, 3,
		35, 401, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 41, 3, 41, 433, 8, 41, 1, 42, 1, 42, 1, 42, 5, 42, 438, 8, 42, 10, 42,
		12, 42, 441, 9, 42, 1, 43, 1, 43, 1, 43, 5, 43, 446, 8, 43, 10, 43, 12,
		43, 449, 9, 43, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 5, 45, 457, 8,
		45, 10, 45, 12, 45, 460, 9, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46,
		1, 46, 1, 46, 5, 46, 470, 8, 46, 10, 46, 12, 46, 473, 9, 46, 1, 46, 3,
		46, 476, 8, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 3, 47, 492, 8, 47, 1, 47, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 3, 47, 510, 8, 47, 1, 47, 1, 47, 1, 47, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 5, 47, 551, 8, 47, 10, 47, 12, 47, 554, 9,
		47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 3, 48, 562, 8, 48, 1, 48,
		1, 48, 5, 48, 566, 8, 48, 10, 48, 12, 48, 569, 9, 48, 1, 48, 1, 48, 3,
		48, 573, 8, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 3, 49,
		582, 8, 49, 1, 49, 1, 49, 1, 50, 1, 50, 5, 50, 588, 8, 50, 10, 50, 12,
		50, 591, 9, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51,
		1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1,
		51, 1, 51, 1, 51, 1, 51, 1, 51, 3, 51, 616, 8, 51, 1, 52, 1, 52, 1, 52,
		1, 52, 5, 52, 622, 8, 52, 10, 52, 12, 52, 625, 9, 52, 1, 52, 3, 52, 628,
		8, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1,
		55, 1, 55, 1, 56, 1, 56, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 0, 1, 94, 58,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
		38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72,
		74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106,
		108, 110, 112, 114, 0, 2, 3, 0, 18, 21, 61, 61, 65, 65, 2, 0, 18, 21, 61,
		61, 682, 0, 119, 1, 0, 0, 0, 2, 123, 1, 0, 0, 0, 4, 155, 1, 0, 0, 0, 6,
		157, 1, 0, 0, 0, 8, 162, 1, 0, 0, 0, 10, 167, 1, 0, 0, 0, 12, 173, 1, 0,
		0, 0, 14, 177, 1, 0, 0, 0, 16, 182, 1, 0, 0, 0, 18, 186, 1, 0, 0, 0, 20,
		196, 1, 0, 0, 0, 22, 200, 1, 0, 0, 0, 24, 205, 1, 0, 0, 0, 26, 209, 1,
		0, 0, 0, 28, 217, 1, 0, 0, 0, 30, 221, 1, 0, 0, 0, 32, 227, 1, 0, 0, 0,
		34, 239, 1, 0, 0, 0, 36, 247, 1, 0, 0, 0, 38, 253, 1, 0, 0, 0, 40, 264,
		1, 0, 0, 0, 42, 267, 1, 0, 0, 0, 44, 271, 1, 0, 0, 0, 46, 276, 1, 0, 0,
		0, 48, 287, 1, 0, 0, 0, 50, 299, 1, 0, 0, 0, 52, 319, 1, 0, 0, 0, 54, 334,
		1, 0, 0, 0, 56, 353, 1, 0, 0, 0, 58, 355, 1, 0, 0, 0, 60, 360, 1, 0, 0,
		0, 62, 365, 1, 0, 0, 0, 64, 370, 1, 0, 0, 0, 66, 387, 1, 0, 0, 0, 68, 389,
		1, 0, 0, 0, 70, 400, 1, 0, 0, 0, 72, 402, 1, 0, 0, 0, 74, 407, 1, 0, 0,
		0, 76, 412, 1, 0, 0, 0, 78, 417, 1, 0, 0, 0, 80, 421, 1, 0, 0, 0, 82, 432,
		1, 0, 0, 0, 84, 434, 1, 0, 0, 0, 86, 442, 1, 0, 0, 0, 88, 450, 1, 0, 0,
		0, 90, 453, 1, 0, 0, 0, 92, 461, 1, 0, 0, 0, 94, 509, 1, 0, 0, 0, 96, 555,
		1, 0, 0, 0, 98, 574, 1, 0, 0, 0, 100, 585, 1, 0, 0, 0, 102, 615, 1, 0,
		0, 0, 104, 617, 1, 0, 0, 0, 106, 631, 1, 0, 0, 0, 108, 633, 1, 0, 0, 0,
		110, 635, 1, 0, 0, 0, 112, 640, 1, 0, 0, 0, 114, 644, 1, 0, 0, 0, 116,
		118, 3, 2, 1, 0, 117, 116, 1, 0, 0, 0, 118, 121, 1, 0, 0, 0, 119, 117,
		1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 1, 1, 0, 0, 0, 121, 119, 1, 0, 0,
		0, 122, 124, 3, 4, 2, 0, 123, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125,
		123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 3, 1, 0, 0, 0, 127, 156, 3,
		6, 3, 0, 128, 156, 3, 8, 4, 0, 129, 156, 3, 10, 5, 0, 130, 156, 3, 12,
		6, 0, 131, 156, 3, 14, 7, 0, 132, 156, 3, 16, 8, 0, 133, 156, 3, 18, 9,
		0, 134, 156, 3, 22, 11, 0, 135, 156, 3, 28, 14, 0, 136, 156, 3, 32, 16,
		0, 137, 156, 3, 38, 19, 0, 138, 156, 3, 30, 15, 0, 139, 156, 3, 46, 23,
		0, 140, 156, 3, 48, 24, 0, 141, 156, 3, 42, 21, 0, 142, 156, 3, 44, 22,
		0, 143, 156, 3, 78, 39, 0, 144, 156, 3, 80, 40, 0, 145, 156, 3, 82, 41,
		0, 146, 156, 3, 96, 48, 0, 147, 156, 3, 102, 51, 0, 148, 156, 3, 104, 52,
		0, 149, 156, 3, 106, 53, 0, 150, 156, 3, 108, 54, 0, 151, 156, 3, 70, 35,
		0, 152, 156, 3, 50, 25, 0, 153, 156, 3, 66, 33, 0, 154, 156, 3, 56, 28,
		0, 155, 127, 1, 0, 0, 0, 155, 128, 1, 0, 0, 0, 155, 129, 1, 0, 0, 0, 155,
		130, 1, 0, 0, 0, 155, 131, 1, 0, 0, 0, 155, 132, 1, 0, 0, 0, 155, 133,
		1, 0, 0, 0, 155, 134, 1, 0, 0, 0, 155, 135, 1, 0, 0, 0, 155, 136, 1, 0,
		0, 0, 155, 137, 1, 0, 0, 0, 155, 138, 1, 0, 0, 0, 155, 139, 1, 0, 0, 0,
		155, 140, 1, 0, 0, 0, 155, 141, 1, 0, 0, 0, 155, 142, 1, 0, 0, 0, 155,
		143, 1, 0, 0, 0, 155, 144, 1, 0, 0, 0, 155, 145, 1, 0, 0, 0, 155, 146,
		1, 0, 0, 0, 155, 147, 1, 0, 0, 0, 155, 148, 1, 0, 0, 0, 155, 149, 1, 0,
		0, 0, 155, 150, 1, 0, 0, 0, 155, 151, 1, 0, 0, 0, 155, 152, 1, 0, 0, 0,
		155, 153, 1, 0, 0, 0, 155, 154, 1, 0, 0, 0, 156, 5, 1, 0, 0, 0, 157, 158,
		5, 4, 0, 0, 158, 159, 5, 40, 0, 0, 159, 160, 3, 90, 45, 0, 160, 161, 5,
		41, 0, 0, 161, 7, 1, 0, 0, 0, 162, 163, 5, 5, 0, 0, 163, 164, 5, 40, 0,
		0, 164, 165, 3, 90, 45, 0, 165, 166, 5, 41, 0, 0, 166, 9, 1, 0, 0, 0, 167,
		168, 5, 7, 0, 0, 168, 169, 3, 84, 42, 0, 169, 170, 3, 114, 57, 0, 170,
		171, 5, 50, 0, 0, 171, 172, 3, 90, 45, 0, 172, 11, 1, 0, 0, 0, 173, 174,
		5, 7, 0, 0, 174, 175, 3, 84, 42, 0, 175, 176, 3, 114, 57, 0, 176, 13, 1,
		0, 0, 0, 177, 178, 5, 7, 0, 0, 178, 179, 3, 84, 42, 0, 179, 180, 5, 50,
		0, 0, 180, 181, 3, 90, 45, 0, 181, 15, 1, 0, 0, 0, 182, 183, 3, 84, 42,
		0, 183, 184, 5, 50, 0, 0, 184, 185, 3, 90, 45, 0, 185, 17, 1, 0, 0, 0,
		186, 187, 5, 29, 0, 0, 187, 188, 5, 61, 0, 0, 188, 190, 5, 42, 0, 0, 189,
		191, 3, 20, 10, 0, 190, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 190,
		1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 195, 5, 43,
		0, 0, 195, 19, 1, 0, 0, 0, 196, 197, 3, 114, 57, 0, 197, 198, 5, 61, 0,
		0, 198, 199, 5, 57, 0, 0, 199, 21, 1, 0, 0, 0, 200, 201, 5, 61, 0, 0, 201,
		202, 5, 50, 0, 0, 202, 203, 5, 61, 0, 0, 203, 204, 3, 24, 12, 0, 204, 23,
		1, 0, 0, 0, 205, 206, 5, 42, 0, 0, 206, 207, 3, 26, 13, 0, 207, 208, 5,
		43, 0, 0, 208, 25, 1, 0, 0, 0, 209, 214, 3, 94, 47, 0, 210, 211, 5, 58,
		0, 0, 211, 213, 3, 94, 47, 0, 212, 210, 1, 0, 0, 0, 213, 216, 1, 0, 0,
		0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 27, 1, 0, 0, 0, 216,
		214, 1, 0, 0, 0, 217, 218, 5, 61, 0, 0, 218, 219, 5, 1, 0, 0, 219, 220,
		5, 61, 0, 0, 220, 29, 1, 0, 0, 0, 221, 222, 5, 61, 0, 0, 222, 223, 5, 1,
		0, 0, 223, 224, 5, 61, 0, 0, 224, 225, 5, 50, 0, 0, 225, 226, 3, 94, 47,
		0, 226, 31, 1, 0, 0, 0, 227, 228, 5, 6, 0, 0, 228, 229, 5, 40, 0, 0, 229,
		230, 5, 61, 0, 0, 230, 231, 5, 37, 0, 0, 231, 232, 5, 61, 0, 0, 232, 233,
		5, 41, 0, 0, 233, 234, 5, 61, 0, 0, 234, 235, 5, 40, 0, 0, 235, 236, 3,
		86, 43, 0, 236, 237, 5, 41, 0, 0, 237, 238, 3, 34, 17, 0, 238, 33, 1, 0,
		0, 0, 239, 241, 5, 42, 0, 0, 240, 242, 3, 36, 18, 0, 241, 240, 1, 0, 0,
		0, 242, 243, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244,
		245, 1, 0, 0, 0, 245, 246, 5, 43, 0, 0, 246, 35, 1, 0, 0, 0, 247, 248,
		5, 61, 0, 0, 248, 249, 5, 1, 0, 0, 249, 250, 5, 61, 0, 0, 250, 251, 5,
		50, 0, 0, 251, 252, 5, 61, 0, 0, 252, 37, 1, 0, 0, 0, 253, 254, 5, 61,
		0, 0, 254, 255, 5, 1, 0, 0, 255, 256, 5, 61, 0, 0, 256, 257, 5, 40, 0,
		0, 257, 258, 3, 26, 13, 0, 258, 259, 5, 41, 0, 0, 259, 39, 1, 0, 0, 0,
		260, 263, 3, 4, 2, 0, 261, 263, 3, 94, 47, 0, 262, 260, 1, 0, 0, 0, 262,
		261, 1, 0, 0, 0, 263, 266, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 264, 265,
		1, 0, 0, 0, 265, 41, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 267, 268, 5, 61,
		0, 0, 268, 269, 5, 40, 0, 0, 269, 270, 5, 41, 0, 0, 270, 43, 1, 0, 0, 0,
		271, 272, 5, 61, 0, 0, 272, 273, 5, 40, 0, 0, 273, 274, 3, 90, 45, 0, 274,
		275, 5, 41, 0, 0, 275, 45, 1, 0, 0, 0, 276, 277, 5, 6, 0, 0, 277, 278,
		5, 61, 0, 0, 278, 279, 5, 40, 0, 0, 279, 281, 5, 41, 0, 0, 280, 282, 3,
		68, 34, 0, 281, 280, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 283, 1, 0,
		0, 0, 283, 284, 5, 42, 0, 0, 284, 285, 3, 40, 20, 0, 285, 286, 5, 43, 0,
		0, 286, 47, 1, 0, 0, 0, 287, 288, 5, 6, 0, 0, 288, 289, 5, 61, 0, 0, 289,
		290, 5, 40, 0, 0, 290, 291, 3, 86, 43, 0, 291, 293, 5, 41, 0, 0, 292, 294,
		3, 68, 34, 0, 293, 292, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294, 295, 1,
		0, 0, 0, 295, 296, 5, 42, 0, 0, 296, 297, 3, 40, 20, 0, 297, 298, 5, 43,
		0, 0, 298, 49, 1, 0, 0, 0, 299, 300, 5, 61, 0, 0, 300, 301, 5, 50, 0, 0,
		301, 302, 3, 52, 26, 0, 302, 51, 1, 0, 0, 0, 303, 304, 5, 44, 0, 0, 304,
		305, 5, 45, 0, 0, 305, 306, 3, 114, 57, 0, 306, 307, 5, 42, 0, 0, 307,
		308, 3, 90, 45, 0, 308, 309, 5, 43, 0, 0, 309, 320, 1, 0, 0, 0, 310, 311,
		5, 44, 0, 0, 311, 312, 5, 45, 0, 0, 312, 313, 5, 44, 0, 0, 313, 314, 5,
		45, 0, 0, 314, 315, 3, 114, 57, 0, 315, 316, 5, 42, 0, 0, 316, 317, 3,
		92, 46, 0, 317, 318, 5, 43, 0, 0, 318, 320, 1, 0, 0, 0, 319, 303, 1, 0,
		0, 0, 319, 310, 1, 0, 0, 0, 320, 53, 1, 0, 0, 0, 321, 322, 5, 61, 0, 0,
		322, 323, 5, 44, 0, 0, 323, 324, 3, 94, 47, 0, 324, 325, 5, 45, 0, 0, 325,
		335, 1, 0, 0, 0, 326, 327, 5, 61, 0, 0, 327, 328, 5, 44, 0, 0, 328, 329,
		3, 94, 47, 0, 329, 330, 5, 45, 0, 0, 330, 331, 5, 44, 0, 0, 331, 332, 3,
		94, 47, 0, 332, 333, 5, 45, 0, 0, 333, 335, 1, 0, 0, 0, 334, 321, 1, 0,
		0, 0, 334, 326, 1, 0, 0, 0, 335, 55, 1, 0, 0, 0, 336, 337, 5, 61, 0, 0,
		337, 338, 5, 44, 0, 0, 338, 339, 3, 94, 47, 0, 339, 340, 5, 45, 0, 0, 340,
		341, 5, 50, 0, 0, 341, 342, 3, 94, 47, 0, 342, 354, 1, 0, 0, 0, 343, 344,
		5, 61, 0, 0, 344, 345, 5, 44, 0, 0, 345, 346, 3, 94, 47, 0, 346, 347, 5,
		45, 0, 0, 347, 348, 5, 44, 0, 0, 348, 349, 3, 94, 47, 0, 349, 350, 5, 45,
		0, 0, 350, 351, 5, 50, 0, 0, 351, 352, 3, 94, 47, 0, 352, 354, 1, 0, 0,
		0, 353, 336, 1, 0, 0, 0, 353, 343, 1, 0, 0, 0, 354, 57, 1, 0, 0, 0, 355,
		356, 5, 25, 0, 0, 356, 357, 5, 40, 0, 0, 357, 358, 3, 90, 45, 0, 358, 359,
		5, 41, 0, 0, 359, 59, 1, 0, 0, 0, 360, 361, 5, 26, 0, 0, 361, 362, 5, 40,
		0, 0, 362, 363, 3, 90, 45, 0, 363, 364, 5, 41, 0, 0, 364, 61, 1, 0, 0,
		0, 365, 366, 5, 27, 0, 0, 366, 367, 5, 40, 0, 0, 367, 368, 3, 90, 45, 0,
		368, 369, 5, 41, 0, 0, 369, 63, 1, 0, 0, 0, 370, 371, 5, 28, 0, 0, 371,
		372, 5, 40, 0, 0, 372, 373, 3, 90, 45, 0, 373, 374, 5, 41, 0, 0, 374, 65,
		1, 0, 0, 0, 375, 376, 5, 7, 0, 0, 376, 377, 5, 61, 0, 0, 377, 378, 5, 44,
		0, 0, 378, 379, 5, 45, 0, 0, 379, 388, 3, 114, 57, 0, 380, 381, 5, 7, 0,
		0, 381, 382, 5, 61, 0, 0, 382, 383, 5, 44, 0, 0, 383, 384, 5, 45, 0, 0,
		384, 385, 5, 44, 0, 0, 385, 386, 5, 45, 0, 0, 386, 388, 3, 114, 57, 0,
		387, 375, 1, 0, 0, 0, 387, 380, 1, 0, 0, 0, 388, 67, 1, 0, 0, 0, 389, 390,
		7, 0, 0, 0, 390, 69, 1, 0, 0, 0, 391, 392, 5, 17, 0, 0, 392, 394, 3, 94,
		47, 0, 393, 395, 5, 57, 0, 0, 394, 393, 1, 0, 0, 0, 394, 395, 1, 0, 0,
		0, 395, 401, 1, 0, 0, 0, 396, 398, 5, 17, 0, 0, 397, 399, 5, 57, 0, 0,
		398, 397, 1, 0, 0, 0, 398, 399, 1, 0, 0, 0, 399, 401, 1, 0, 0, 0, 400,
		391, 1, 0, 0, 0, 400, 396, 1, 0, 0, 0, 401, 71, 1, 0, 0, 0, 402, 403, 5,
		34, 0, 0, 403, 404, 5, 40, 0, 0, 404, 405, 3, 90, 45, 0, 405, 406, 5, 41,
		0, 0, 406, 73, 1, 0, 0, 0, 407, 408, 5, 32, 0, 0, 408, 409, 5, 40, 0, 0,
		409, 410, 3, 90, 45, 0, 410, 411, 5, 41, 0, 0, 411, 75, 1, 0, 0, 0, 412,
		413, 5, 33, 0, 0, 413, 414, 5, 40, 0, 0, 414, 415, 3, 90, 45, 0, 415, 416,
		5, 41, 0, 0, 416, 77, 1, 0, 0, 0, 417, 418, 5, 61, 0, 0, 418, 419, 5, 2,
		0, 0, 419, 420, 3, 94, 47, 0, 420, 79, 1, 0, 0, 0, 421, 422, 5, 61, 0,
		0, 422, 423, 5, 3, 0, 0, 423, 424, 3, 94, 47, 0, 424, 81, 1, 0, 0, 0, 425,
		426, 5, 61, 0, 0, 426, 427, 5, 50, 0, 0, 427, 433, 3, 94, 47, 0, 428, 429,
		5, 61, 0, 0, 429, 433, 5, 59, 0, 0, 430, 431, 5, 61, 0, 0, 431, 433, 5,
		60, 0, 0, 432, 425, 1, 0, 0, 0, 432, 428, 1, 0, 0, 0, 432, 430, 1, 0, 0,
		0, 433, 83, 1, 0, 0, 0, 434, 439, 5, 61, 0, 0, 435, 436, 5, 58, 0, 0, 436,
		438, 5, 61, 0, 0, 437, 435, 1, 0, 0, 0, 438, 441, 1, 0, 0, 0, 439, 437,
		1, 0, 0, 0, 439, 440, 1, 0, 0, 0, 440, 85, 1, 0, 0, 0, 441, 439, 1, 0,
		0, 0, 442, 447, 3, 88, 44, 0, 443, 444, 5, 58, 0, 0, 444, 446, 3, 88, 44,
		0, 445, 443, 1, 0, 0, 0, 446, 449, 1, 0, 0, 0, 447, 445, 1, 0, 0, 0, 447,
		448, 1, 0, 0, 0, 448, 87, 1, 0, 0, 0, 449, 447, 1, 0, 0, 0, 450, 451, 5,
		61, 0, 0, 451, 452, 3, 114, 57, 0, 452, 89, 1, 0, 0, 0, 453, 458, 3, 94,
		47, 0, 454, 455, 5, 58, 0, 0, 455, 457, 3, 94, 47, 0, 456, 454, 1, 0, 0,
		0, 457, 460, 1, 0, 0, 0, 458, 456, 1, 0, 0, 0, 458, 459, 1, 0, 0, 0, 459,
		91, 1, 0, 0, 0, 460, 458, 1, 0, 0, 0, 461, 462, 5, 42, 0, 0, 462, 463,
		3, 90, 45, 0, 463, 471, 5, 43, 0, 0, 464, 465, 5, 58, 0, 0, 465, 466, 5,
		42, 0, 0, 466, 467, 3, 90, 45, 0, 467, 468, 5, 43, 0, 0, 468, 470, 1, 0,
		0, 0, 469, 464, 1, 0, 0, 0, 470, 473, 1, 0, 0, 0, 471, 469, 1, 0, 0, 0,
		471, 472, 1, 0, 0, 0, 472, 475, 1, 0, 0, 0, 473, 471, 1, 0, 0, 0, 474,
		476, 5, 58, 0, 0, 475, 474, 1, 0, 0, 0, 475, 476, 1, 0, 0, 0, 476, 93,
		1, 0, 0, 0, 477, 478, 6, 47, -1, 0, 478, 479, 5, 36, 0, 0, 479, 510, 3,
		94, 47, 35, 480, 481, 5, 55, 0, 0, 481, 510, 3, 94, 47, 34, 482, 510, 5,
		62, 0, 0, 483, 510, 5, 63, 0, 0, 484, 510, 5, 64, 0, 0, 485, 510, 5, 65,
		0, 0, 486, 510, 5, 23, 0, 0, 487, 510, 5, 24, 0, 0, 488, 510, 3, 54, 27,
		0, 489, 491, 5, 44, 0, 0, 490, 492, 3, 90, 45, 0, 491, 490, 1, 0, 0, 0,
		491, 492, 1, 0, 0, 0, 492, 493, 1, 0, 0, 0, 493, 510, 5, 45, 0, 0, 494,
		510, 3, 42, 21, 0, 495, 510, 3, 44, 22, 0, 496, 510, 5, 61, 0, 0, 497,
		498, 5, 40, 0, 0, 498, 499, 3, 94, 47, 0, 499, 500, 5, 41, 0, 0, 500, 510,
		1, 0, 0, 0, 501, 510, 3, 74, 37, 0, 502, 510, 3, 76, 38, 0, 503, 510, 3,
		72, 36, 0, 504, 510, 3, 28, 14, 0, 505, 510, 3, 64, 32, 0, 506, 510, 3,
		58, 29, 0, 507, 510, 3, 60, 30, 0, 508, 510, 3, 62, 31, 0, 509, 477, 1,
		0, 0, 0, 509, 480, 1, 0, 0, 0, 509, 482, 1, 0, 0, 0, 509, 483, 1, 0, 0,
		0, 509, 484, 1, 0, 0, 0, 509, 485, 1, 0, 0, 0, 509, 486, 1, 0, 0, 0, 509,
		487, 1, 0, 0, 0, 509, 488, 1, 0, 0, 0, 509, 489, 1, 0, 0, 0, 509, 494,
		1, 0, 0, 0, 509, 495, 1, 0, 0, 0, 509, 496, 1, 0, 0, 0, 509, 497, 1, 0,
		0, 0, 509, 501, 1, 0, 0, 0, 509, 502, 1, 0, 0, 0, 509, 503, 1, 0, 0, 0,
		509, 504, 1, 0, 0, 0, 509, 505, 1, 0, 0, 0, 509, 506, 1, 0, 0, 0, 509,
		507, 1, 0, 0, 0, 509, 508, 1, 0, 0, 0, 510, 552, 1, 0, 0, 0, 511, 512,
		10, 33, 0, 0, 512, 513, 5, 39, 0, 0, 513, 551, 3, 94, 47, 34, 514, 515,
		10, 32, 0, 0, 515, 516, 5, 38, 0, 0, 516, 551, 3, 94, 47, 33, 517, 518,
		10, 31, 0, 0, 518, 519, 5, 37, 0, 0, 519, 551, 3, 94, 47, 32, 520, 521,
		10, 30, 0, 0, 521, 522, 5, 36, 0, 0, 522, 551, 3, 94, 47, 31, 523, 524,
		10, 29, 0, 0, 524, 525, 5, 35, 0, 0, 525, 551, 3, 94, 47, 30, 526, 527,
		10, 28, 0, 0, 527, 528, 5, 46, 0, 0, 528, 551, 3, 94, 47, 29, 529, 530,
		10, 27, 0, 0, 530, 531, 5, 47, 0, 0, 531, 551, 3, 94, 47, 28, 532, 533,
		10, 26, 0, 0, 533, 534, 5, 48, 0, 0, 534, 551, 3, 94, 47, 27, 535, 536,
		10, 25, 0, 0, 536, 537, 5, 49, 0, 0, 537, 551, 3, 94, 47, 26, 538, 539,
		10, 24, 0, 0, 539, 540, 5, 51, 0, 0, 540, 551, 3, 94, 47, 25, 541, 542,
		10, 23, 0, 0, 542, 543, 5, 52, 0, 0, 543, 551, 3, 94, 47, 24, 544, 545,
		10, 22, 0, 0, 545, 546, 5, 53, 0, 0, 546, 551, 3, 94, 47, 23, 547, 548,
		10, 21, 0, 0, 548, 549, 5, 54, 0, 0, 549, 551, 3, 94, 47, 22, 550, 511,
		1, 0, 0, 0, 550, 514, 1, 0, 0, 0, 550, 517, 1, 0, 0, 0, 550, 520, 1, 0,
		0, 0, 550, 523, 1, 0, 0, 0, 550, 526, 1, 0, 0, 0, 550, 529, 1, 0, 0, 0,
		550, 532, 1, 0, 0, 0, 550, 535, 1, 0, 0, 0, 550, 538, 1, 0, 0, 0, 550,
		541, 1, 0, 0, 0, 550, 544, 1, 0, 0, 0, 550, 547, 1, 0, 0, 0, 551, 554,
		1, 0, 0, 0, 552, 550, 1, 0, 0, 0, 552, 553, 1, 0, 0, 0, 553, 95, 1, 0,
		0, 0, 554, 552, 1, 0, 0, 0, 555, 561, 5, 9, 0, 0, 556, 557, 5, 40, 0, 0,
		557, 558, 3, 94, 47, 0, 558, 559, 5, 41, 0, 0, 559, 562, 1, 0, 0, 0, 560,
		562, 3, 94, 47, 0, 561, 556, 1, 0, 0, 0, 561, 560, 1, 0, 0, 0, 562, 563,
		1, 0, 0, 0, 563, 567, 3, 100, 50, 0, 564, 566, 3, 98, 49, 0, 565, 564,
		1, 0, 0, 0, 566, 569, 1, 0, 0, 0, 567, 565, 1, 0, 0, 0, 567, 568, 1, 0,
		0, 0, 568, 572, 1, 0, 0, 0, 569, 567, 1, 0, 0, 0, 570, 571, 5, 10, 0, 0,
		571, 573, 3, 100, 50, 0, 572, 570, 1, 0, 0, 0, 572, 573, 1, 0, 0, 0, 573,
		97, 1, 0, 0, 0, 574, 575, 5, 10, 0, 0, 575, 581, 5, 9, 0, 0, 576, 577,
		5, 40, 0, 0, 577, 578, 3, 94, 47, 0, 578, 579, 5, 41, 0, 0, 579, 582, 1,
		0, 0, 0, 580, 582, 3, 94, 47, 0, 581, 576, 1, 0, 0, 0, 581, 580, 1, 0,
		0, 0, 582, 583, 1, 0, 0, 0, 583, 584, 3, 100, 50, 0, 584, 99, 1, 0, 0,
		0, 585, 589, 5, 42, 0, 0, 586, 588, 3, 2, 1, 0, 587, 586, 1, 0, 0, 0, 588,
		591, 1, 0, 0, 0, 589, 587, 1, 0, 0, 0, 589, 590, 1, 0, 0, 0, 590, 592,
		1, 0, 0, 0, 591, 589, 1, 0, 0, 0, 592, 593, 5, 43, 0, 0, 593, 101, 1, 0,
		0, 0, 594, 595, 5, 14, 0, 0, 595, 596, 3, 94, 47, 0, 596, 597, 3, 100,
		50, 0, 597, 616, 1, 0, 0, 0, 598, 599, 5, 14, 0, 0, 599, 600, 3, 82, 41,
		0, 600, 601, 5, 57, 0, 0, 601, 602, 3, 94, 47, 0, 602, 603, 5, 57, 0, 0,
		603, 604, 3, 82, 41, 0, 604, 605, 3, 100, 50, 0, 605, 616, 1, 0, 0, 0,
		606, 607, 5, 14, 0, 0, 607, 608, 5, 61, 0, 0, 608, 609, 5, 58, 0, 0, 609,
		610, 5, 61, 0, 0, 610, 611, 5, 50, 0, 0, 611, 612, 5, 31, 0, 0, 612, 613,
		3, 94, 47, 0, 613, 614, 3, 100, 50, 0, 614, 616, 1, 0, 0, 0, 615, 594,
		1, 0, 0, 0, 615, 598, 1, 0, 0, 0, 615, 606, 1, 0, 0, 0, 616, 103, 1, 0,
		0, 0, 617, 618, 5, 11, 0, 0, 618, 619, 3, 94, 47, 0, 619, 623, 5, 42, 0,
		0, 620, 622, 3, 110, 55, 0, 621, 620, 1, 0, 0, 0, 622, 625, 1, 0, 0, 0,
		623, 621, 1, 0, 0, 0, 623, 624, 1, 0, 0, 0, 624, 627, 1, 0, 0, 0, 625,
		623, 1, 0, 0, 0, 626, 628, 3, 112, 56, 0, 627, 626, 1, 0, 0, 0, 627, 628,
		1, 0, 0, 0, 628, 629, 1, 0, 0, 0, 629, 630, 5, 43, 0, 0, 630, 105, 1, 0,
		0, 0, 631, 632, 5, 15, 0, 0, 632, 107, 1, 0, 0, 0, 633, 634, 5, 16, 0,
		0, 634, 109, 1, 0, 0, 0, 635, 636, 5, 12, 0, 0, 636, 637, 3, 94, 47, 0,
		637, 638, 5, 56, 0, 0, 638, 639, 3, 2, 1, 0, 639, 111, 1, 0, 0, 0, 640,
		641, 5, 13, 0, 0, 641, 642, 5, 56, 0, 0, 642, 643, 3, 2, 1, 0, 643, 113,
		1, 0, 0, 0, 644, 645, 7, 1, 0, 0, 645, 115, 1, 0, 0, 0, 35, 119, 125, 155,
		192, 214, 243, 262, 264, 281, 293, 319, 334, 353, 387, 394, 398, 400, 432,
		439, 447, 458, 471, 475, 491, 509, 550, 552, 561, 567, 572, 581, 589, 615,
		623, 627,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// gramaticaParserInit initializes any static state used to implement gramaticaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewgramaticaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GramaticaParserInit() {
	staticData := &GramaticaParserStaticData
	staticData.once.Do(gramaticaParserInit)
}

// NewgramaticaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewgramaticaParser(input antlr.TokenStream) *gramaticaParser {
	GramaticaParserInit()
	this := new(gramaticaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GramaticaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "gramatica.g4"

	return this
}

// gramaticaParser tokens.
const (
	gramaticaParserEOF             = antlr.TokenEOF
	gramaticaParserT__0            = 1
	gramaticaParserT__1            = 2
	gramaticaParserT__2            = 3
	gramaticaParserPRINT           = 4
	gramaticaParserPRINTLN         = 5
	gramaticaParserFUNCIONES       = 6
	gramaticaParserMUT             = 7
	gramaticaParserSLICE           = 8
	gramaticaParserIF              = 9
	gramaticaParserELSE            = 10
	gramaticaParserSWITCH          = 11
	gramaticaParserCASE            = 12
	gramaticaParserDEFAULT         = 13
	gramaticaParserFOR             = 14
	gramaticaParserBREAK           = 15
	gramaticaParserCONTINUE        = 16
	gramaticaParserRETURN          = 17
	gramaticaParserINT             = 18
	gramaticaParserFLOAT           = 19
	gramaticaParserSTRING          = 20
	gramaticaParserBOOL            = 21
	gramaticaParserNIL             = 22
	gramaticaParserTRUE            = 23
	gramaticaParserFALSE           = 24
	gramaticaParserINDEXOF         = 25
	gramaticaParserJOIN            = 26
	gramaticaParserLEN             = 27
	gramaticaParserAPPEND          = 28
	gramaticaParserSTRUCT          = 29
	gramaticaParserFUNC            = 30
	gramaticaParserRANGE           = 31
	gramaticaParserATOI            = 32
	gramaticaParserPARSEFLOAT      = 33
	gramaticaParserTYPEOF          = 34
	gramaticaParserMAS             = 35
	gramaticaParserMENOS           = 36
	gramaticaParserPOR             = 37
	gramaticaParserDIV             = 38
	gramaticaParserMODULO          = 39
	gramaticaParserPAR1            = 40
	gramaticaParserPAR2            = 41
	gramaticaParserLLAV1           = 42
	gramaticaParserLLAV2           = 43
	gramaticaParserCOR1            = 44
	gramaticaParserCOR2            = 45
	gramaticaParserDIFERENTE       = 46
	gramaticaParserIGUALDAD        = 47
	gramaticaParserMENIGUAL        = 48
	gramaticaParserMAYIGUAL        = 49
	gramaticaParserIGUAL           = 50
	gramaticaParserMENOR           = 51
	gramaticaParserMAYOR           = 52
	gramaticaParserOR              = 53
	gramaticaParserAND             = 54
	gramaticaParserDIFER           = 55
	gramaticaParserDOSPTS          = 56
	gramaticaParserPTCOMA          = 57
	gramaticaParserCOMA            = 58
	gramaticaParserFINCREMENTO     = 59
	gramaticaParserFDECREMENTO     = 60
	gramaticaParserIDENTIFICADOR   = 61
	gramaticaParserENTERO          = 62
	gramaticaParserDECIMAL         = 63
	gramaticaParserCADENA          = 64
	gramaticaParserRUNE            = 65
	gramaticaParserCOMENTARIO      = 66
	gramaticaParserMULTICOMENTARIO = 67
	gramaticaParserWS              = 68
)

// gramaticaParser rules.
const (
	gramaticaParserRULE_init                         = 0
	gramaticaParserRULE_instrucciones                = 1
	gramaticaParserRULE_instruccion                  = 2
	gramaticaParserRULE_print                        = 3
	gramaticaParserRULE_println                      = 4
	gramaticaParserRULE_declaracionMultiple          = 5
	gramaticaParserRULE_declaracionMultipleSimple    = 6
	gramaticaParserRULE_declaracionMultipleSinTipo   = 7
	gramaticaParserRULE_asignacionMultiple           = 8
	gramaticaParserRULE_structDef                    = 9
	gramaticaParserRULE_atributos                    = 10
	gramaticaParserRULE_structInst                   = 11
	gramaticaParserRULE_structInit                   = 12
	gramaticaParserRULE_listaStructs                 = 13
	gramaticaParserRULE_accesoStruct                 = 14
	gramaticaParserRULE_asigStruct                   = 15
	gramaticaParserRULE_funStruct                    = 16
	gramaticaParserRULE_caja                         = 17
	gramaticaParserRULE_valores                      = 18
	gramaticaParserRULE_actStruct                    = 19
	gramaticaParserRULE_bloqueFuncion                = 20
	gramaticaParserRULE_llamadaFuncionesSinParametro = 21
	gramaticaParserRULE_llamadaFuncionesConParametro = 22
	gramaticaParserRULE_fnSinParametro               = 23
	gramaticaParserRULE_fnConParametro               = 24
	gramaticaParserRULE_sliceDef                     = 25
	gramaticaParserRULE_sliceLiteral                 = 26
	gramaticaParserRULE_accesoElementoSlice          = 27
	gramaticaParserRULE_modificacionElementoSlice    = 28
	gramaticaParserRULE_fnIndexOf                    = 29
	gramaticaParserRULE_fnJoin                       = 30
	gramaticaParserRULE_fnLen                        = 31
	gramaticaParserRULE_fnAppend                     = 32
	gramaticaParserRULE_declaracionSliceVacio        = 33
	gramaticaParserRULE_tipoRetorno                  = 34
	gramaticaParserRULE_retorno                      = 35
	gramaticaParserRULE_fnTypeOf                     = 36
	gramaticaParserRULE_fnAtoi                       = 37
	gramaticaParserRULE_fnParseToFloat               = 38
	gramaticaParserRULE_asigIncre                    = 39
	gramaticaParserRULE_asigDecre                    = 40
	gramaticaParserRULE_asignacion                   = 41
	gramaticaParserRULE_listaIDS                     = 42
	gramaticaParserRULE_listaPar                     = 43
	gramaticaParserRULE_parametro                    = 44
	gramaticaParserRULE_listaExpr                    = 45
	gramaticaParserRULE_listaExprList                = 46
	gramaticaParserRULE_expresion                    = 47
	gramaticaParserRULE_sif                          = 48
	gramaticaParserRULE_elseifPart                   = 49
	gramaticaParserRULE_bloque                       = 50
	gramaticaParserRULE_sfor                         = 51
	gramaticaParserRULE_sSwitch                      = 52
	gramaticaParserRULE_break_                       = 53
	gramaticaParserRULE_continue_                    = 54
	gramaticaParserRULE_caseBlock                    = 55
	gramaticaParserRULE_defaultBlock                 = 56
	gramaticaParserRULE_tipos                        = 57
)

// IInitContext is an interface to support dynamic dispatch.
type IInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInstrucciones() []IInstruccionesContext
	Instrucciones(i int) IInstruccionesContext

	// IsInitContext differentiates from other interfaces.
	IsInitContext()
}

type InitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitContext() *InitContext {
	var p = new(InitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_init
	return p
}

func InitEmptyInitContext(p *InitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_init
}

func (*InitContext) IsInitContext() {}

func NewInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitContext {
	var p = new(InitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_init

	return p
}

func (s *InitContext) GetParser() antlr.Parser { return s.parser }

func (s *InitContext) AllInstrucciones() []IInstruccionesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstruccionesContext); ok {
			len++
		}
	}

	tst := make([]IInstruccionesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstruccionesContext); ok {
			tst[i] = t.(IInstruccionesContext)
			i++
		}
	}

	return tst
}

func (s *InitContext) Instrucciones(i int) IInstruccionesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *InitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterInit(s)
	}
}

func (s *InitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitInit(s)
	}
}

func (s *InitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Init() (localctx IInitContext) {
	localctx = NewInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, gramaticaParserRULE_init)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2305843009750813424) != 0 {
		{
			p.SetState(116)
			p.Instrucciones()
		}

		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInstruccion() []IInstruccionContext
	Instruccion(i int) IInstruccionContext

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_instrucciones
	return p
}

func InitEmptyInstruccionesContext(p *InstruccionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_instrucciones
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstruccionContext); ok {
			len++
		}
	}

	tst := make([]IInstruccionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstruccionContext); ok {
			tst[i] = t.(IInstruccionContext)
			i++
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (s *InstruccionesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitInstrucciones(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Instrucciones() (localctx IInstruccionesContext) {
	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, gramaticaParserRULE_instrucciones)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(122)
				p.Instruccion()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Print_() IPrintContext
	Println_() IPrintlnContext
	DeclaracionMultiple() IDeclaracionMultipleContext
	DeclaracionMultipleSimple() IDeclaracionMultipleSimpleContext
	DeclaracionMultipleSinTipo() IDeclaracionMultipleSinTipoContext
	AsignacionMultiple() IAsignacionMultipleContext
	StructDef() IStructDefContext
	StructInst() IStructInstContext
	AccesoStruct() IAccesoStructContext
	FunStruct() IFunStructContext
	ActStruct() IActStructContext
	AsigStruct() IAsigStructContext
	FnSinParametro() IFnSinParametroContext
	FnConParametro() IFnConParametroContext
	LlamadaFuncionesSinParametro() ILlamadaFuncionesSinParametroContext
	LlamadaFuncionesConParametro() ILlamadaFuncionesConParametroContext
	AsigIncre() IAsigIncreContext
	AsigDecre() IAsigDecreContext
	Asignacion() IAsignacionContext
	Sif() ISifContext
	Sfor() ISforContext
	SSwitch() ISSwitchContext
	Break_() IBreak_Context
	Continue_() IContinue_Context
	Retorno() IRetornoContext
	SliceDef() ISliceDefContext
	DeclaracionSliceVacio() IDeclaracionSliceVacioContext
	ModificacionElementoSlice() IModificacionElementoSliceContext

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_instruccion
	return p
}

func InitEmptyInstruccionContext(p *InstruccionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_instruccion
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Print_() IPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintContext)
}

func (s *InstruccionContext) Println_() IPrintlnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintlnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintlnContext)
}

func (s *InstruccionContext) DeclaracionMultiple() IDeclaracionMultipleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionMultipleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracionMultipleContext)
}

func (s *InstruccionContext) DeclaracionMultipleSimple() IDeclaracionMultipleSimpleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionMultipleSimpleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracionMultipleSimpleContext)
}

func (s *InstruccionContext) DeclaracionMultipleSinTipo() IDeclaracionMultipleSinTipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionMultipleSinTipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracionMultipleSinTipoContext)
}

func (s *InstruccionContext) AsignacionMultiple() IAsignacionMultipleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionMultipleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionMultipleContext)
}

func (s *InstruccionContext) StructDef() IStructDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDefContext)
}

func (s *InstruccionContext) StructInst() IStructInstContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructInstContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructInstContext)
}

func (s *InstruccionContext) AccesoStruct() IAccesoStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccesoStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccesoStructContext)
}

func (s *InstruccionContext) FunStruct() IFunStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunStructContext)
}

func (s *InstruccionContext) ActStruct() IActStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActStructContext)
}

func (s *InstruccionContext) AsigStruct() IAsigStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsigStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsigStructContext)
}

func (s *InstruccionContext) FnSinParametro() IFnSinParametroContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnSinParametroContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnSinParametroContext)
}

func (s *InstruccionContext) FnConParametro() IFnConParametroContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnConParametroContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnConParametroContext)
}

func (s *InstruccionContext) LlamadaFuncionesSinParametro() ILlamadaFuncionesSinParametroContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadaFuncionesSinParametroContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadaFuncionesSinParametroContext)
}

func (s *InstruccionContext) LlamadaFuncionesConParametro() ILlamadaFuncionesConParametroContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadaFuncionesConParametroContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadaFuncionesConParametroContext)
}

func (s *InstruccionContext) AsigIncre() IAsigIncreContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsigIncreContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsigIncreContext)
}

func (s *InstruccionContext) AsigDecre() IAsigDecreContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsigDecreContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsigDecreContext)
}

func (s *InstruccionContext) Asignacion() IAsignacionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *InstruccionContext) Sif() ISifContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISifContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISifContext)
}

func (s *InstruccionContext) Sfor() ISforContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISforContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISforContext)
}

func (s *InstruccionContext) SSwitch() ISSwitchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISSwitchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISSwitchContext)
}

func (s *InstruccionContext) Break_() IBreak_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreak_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreak_Context)
}

func (s *InstruccionContext) Continue_() IContinue_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinue_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinue_Context)
}

func (s *InstruccionContext) Retorno() IRetornoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRetornoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRetornoContext)
}

func (s *InstruccionContext) SliceDef() ISliceDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceDefContext)
}

func (s *InstruccionContext) DeclaracionSliceVacio() IDeclaracionSliceVacioContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionSliceVacioContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracionSliceVacioContext)
}

func (s *InstruccionContext) ModificacionElementoSlice() IModificacionElementoSliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModificacionElementoSliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModificacionElementoSliceContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (s *InstruccionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitInstruccion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Instruccion() (localctx IInstruccionContext) {
	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, gramaticaParserRULE_instruccion)
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Print_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.Println_()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(129)
			p.DeclaracionMultiple()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(130)
			p.DeclaracionMultipleSimple()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(131)
			p.DeclaracionMultipleSinTipo()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(132)
			p.AsignacionMultiple()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(133)
			p.StructDef()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(134)
			p.StructInst()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(135)
			p.AccesoStruct()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(136)
			p.FunStruct()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(137)
			p.ActStruct()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(138)
			p.AsigStruct()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(139)
			p.FnSinParametro()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(140)
			p.FnConParametro()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(141)
			p.LlamadaFuncionesSinParametro()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(142)
			p.LlamadaFuncionesConParametro()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(143)
			p.AsigIncre()
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(144)
			p.AsigDecre()
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(145)
			p.Asignacion()
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(146)
			p.Sif()
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(147)
			p.Sfor()
		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(148)
			p.SSwitch()
		}

	case 23:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(149)
			p.Break_()
		}

	case 24:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(150)
			p.Continue_()
		}

	case 25:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(151)
			p.Retorno()
		}

	case 26:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(152)
			p.SliceDef()
		}

	case 27:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(153)
			p.DeclaracionSliceVacio()
		}

	case 28:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(154)
			p.ModificacionElementoSlice()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintContext is an interface to support dynamic dispatch.
type IPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT() antlr.TerminalNode
	PAR1() antlr.TerminalNode
	ListaExpr() IListaExprContext
	PAR2() antlr.TerminalNode

	// IsPrintContext differentiates from other interfaces.
	IsPrintContext()
}

type PrintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintContext() *PrintContext {
	var p = new(PrintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_print
	return p
}

func InitEmptyPrintContext(p *PrintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_print
}

func (*PrintContext) IsPrintContext() {}

func NewPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintContext {
	var p = new(PrintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_print

	return p
}

func (s *PrintContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintContext) PRINT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPRINT, 0)
}

func (s *PrintContext) PAR1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR1, 0)
}

func (s *PrintContext) ListaExpr() IListaExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaExprContext)
}

func (s *PrintContext) PAR2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR2, 0)
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitPrint(s)
	}
}

func (s *PrintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitPrint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Print_() (localctx IPrintContext) {
	localctx = NewPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, gramaticaParserRULE_print)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(gramaticaParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.ListaExpr()
	}
	{
		p.SetState(160)
		p.Match(gramaticaParserPAR2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintlnContext is an interface to support dynamic dispatch.
type IPrintlnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINTLN() antlr.TerminalNode
	PAR1() antlr.TerminalNode
	ListaExpr() IListaExprContext
	PAR2() antlr.TerminalNode

	// IsPrintlnContext differentiates from other interfaces.
	IsPrintlnContext()
}

type PrintlnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintlnContext() *PrintlnContext {
	var p = new(PrintlnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_println
	return p
}

func InitEmptyPrintlnContext(p *PrintlnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_println
}

func (*PrintlnContext) IsPrintlnContext() {}

func NewPrintlnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintlnContext {
	var p = new(PrintlnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_println

	return p
}

func (s *PrintlnContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintlnContext) PRINTLN() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPRINTLN, 0)
}

func (s *PrintlnContext) PAR1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR1, 0)
}

func (s *PrintlnContext) ListaExpr() IListaExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaExprContext)
}

func (s *PrintlnContext) PAR2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR2, 0)
}

func (s *PrintlnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintlnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintlnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterPrintln(s)
	}
}

func (s *PrintlnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitPrintln(s)
	}
}

func (s *PrintlnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitPrintln(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Println_() (localctx IPrintlnContext) {
	localctx = NewPrintlnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, gramaticaParserRULE_println)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(gramaticaParserPRINTLN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(164)
		p.ListaExpr()
	}
	{
		p.SetState(165)
		p.Match(gramaticaParserPAR2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracionMultipleContext is an interface to support dynamic dispatch.
type IDeclaracionMultipleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MUT() antlr.TerminalNode
	ListaIDS() IListaIDSContext
	Tipos() ITiposContext
	IGUAL() antlr.TerminalNode
	ListaExpr() IListaExprContext

	// IsDeclaracionMultipleContext differentiates from other interfaces.
	IsDeclaracionMultipleContext()
}

type DeclaracionMultipleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracionMultipleContext() *DeclaracionMultipleContext {
	var p = new(DeclaracionMultipleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_declaracionMultiple
	return p
}

func InitEmptyDeclaracionMultipleContext(p *DeclaracionMultipleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_declaracionMultiple
}

func (*DeclaracionMultipleContext) IsDeclaracionMultipleContext() {}

func NewDeclaracionMultipleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionMultipleContext {
	var p = new(DeclaracionMultipleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_declaracionMultiple

	return p
}

func (s *DeclaracionMultipleContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionMultipleContext) MUT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserMUT, 0)
}

func (s *DeclaracionMultipleContext) ListaIDS() IListaIDSContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaIDSContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaIDSContext)
}

func (s *DeclaracionMultipleContext) Tipos() ITiposContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITiposContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITiposContext)
}

func (s *DeclaracionMultipleContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIGUAL, 0)
}

func (s *DeclaracionMultipleContext) ListaExpr() IListaExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaExprContext)
}

func (s *DeclaracionMultipleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionMultipleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionMultipleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterDeclaracionMultiple(s)
	}
}

func (s *DeclaracionMultipleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitDeclaracionMultiple(s)
	}
}

func (s *DeclaracionMultipleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitDeclaracionMultiple(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) DeclaracionMultiple() (localctx IDeclaracionMultipleContext) {
	localctx = NewDeclaracionMultipleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, gramaticaParserRULE_declaracionMultiple)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(gramaticaParserMUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.ListaIDS()
	}
	{
		p.SetState(169)
		p.Tipos()
	}
	{
		p.SetState(170)
		p.Match(gramaticaParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.ListaExpr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracionMultipleSimpleContext is an interface to support dynamic dispatch.
type IDeclaracionMultipleSimpleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MUT() antlr.TerminalNode
	ListaIDS() IListaIDSContext
	Tipos() ITiposContext

	// IsDeclaracionMultipleSimpleContext differentiates from other interfaces.
	IsDeclaracionMultipleSimpleContext()
}

type DeclaracionMultipleSimpleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracionMultipleSimpleContext() *DeclaracionMultipleSimpleContext {
	var p = new(DeclaracionMultipleSimpleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_declaracionMultipleSimple
	return p
}

func InitEmptyDeclaracionMultipleSimpleContext(p *DeclaracionMultipleSimpleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_declaracionMultipleSimple
}

func (*DeclaracionMultipleSimpleContext) IsDeclaracionMultipleSimpleContext() {}

func NewDeclaracionMultipleSimpleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionMultipleSimpleContext {
	var p = new(DeclaracionMultipleSimpleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_declaracionMultipleSimple

	return p
}

func (s *DeclaracionMultipleSimpleContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionMultipleSimpleContext) MUT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserMUT, 0)
}

func (s *DeclaracionMultipleSimpleContext) ListaIDS() IListaIDSContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaIDSContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaIDSContext)
}

func (s *DeclaracionMultipleSimpleContext) Tipos() ITiposContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITiposContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITiposContext)
}

func (s *DeclaracionMultipleSimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionMultipleSimpleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionMultipleSimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterDeclaracionMultipleSimple(s)
	}
}

func (s *DeclaracionMultipleSimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitDeclaracionMultipleSimple(s)
	}
}

func (s *DeclaracionMultipleSimpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitDeclaracionMultipleSimple(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) DeclaracionMultipleSimple() (localctx IDeclaracionMultipleSimpleContext) {
	localctx = NewDeclaracionMultipleSimpleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, gramaticaParserRULE_declaracionMultipleSimple)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Match(gramaticaParserMUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
		p.ListaIDS()
	}
	{
		p.SetState(175)
		p.Tipos()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracionMultipleSinTipoContext is an interface to support dynamic dispatch.
type IDeclaracionMultipleSinTipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MUT() antlr.TerminalNode
	ListaIDS() IListaIDSContext
	IGUAL() antlr.TerminalNode
	ListaExpr() IListaExprContext

	// IsDeclaracionMultipleSinTipoContext differentiates from other interfaces.
	IsDeclaracionMultipleSinTipoContext()
}

type DeclaracionMultipleSinTipoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracionMultipleSinTipoContext() *DeclaracionMultipleSinTipoContext {
	var p = new(DeclaracionMultipleSinTipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_declaracionMultipleSinTipo
	return p
}

func InitEmptyDeclaracionMultipleSinTipoContext(p *DeclaracionMultipleSinTipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_declaracionMultipleSinTipo
}

func (*DeclaracionMultipleSinTipoContext) IsDeclaracionMultipleSinTipoContext() {}

func NewDeclaracionMultipleSinTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionMultipleSinTipoContext {
	var p = new(DeclaracionMultipleSinTipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_declaracionMultipleSinTipo

	return p
}

func (s *DeclaracionMultipleSinTipoContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionMultipleSinTipoContext) MUT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserMUT, 0)
}

func (s *DeclaracionMultipleSinTipoContext) ListaIDS() IListaIDSContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaIDSContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaIDSContext)
}

func (s *DeclaracionMultipleSinTipoContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIGUAL, 0)
}

func (s *DeclaracionMultipleSinTipoContext) ListaExpr() IListaExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaExprContext)
}

func (s *DeclaracionMultipleSinTipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionMultipleSinTipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionMultipleSinTipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterDeclaracionMultipleSinTipo(s)
	}
}

func (s *DeclaracionMultipleSinTipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitDeclaracionMultipleSinTipo(s)
	}
}

func (s *DeclaracionMultipleSinTipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitDeclaracionMultipleSinTipo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) DeclaracionMultipleSinTipo() (localctx IDeclaracionMultipleSinTipoContext) {
	localctx = NewDeclaracionMultipleSinTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, gramaticaParserRULE_declaracionMultipleSinTipo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Match(gramaticaParserMUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)
		p.ListaIDS()
	}
	{
		p.SetState(179)
		p.Match(gramaticaParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
		p.ListaExpr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignacionMultipleContext is an interface to support dynamic dispatch.
type IAsignacionMultipleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ListaIDS() IListaIDSContext
	IGUAL() antlr.TerminalNode
	ListaExpr() IListaExprContext

	// IsAsignacionMultipleContext differentiates from other interfaces.
	IsAsignacionMultipleContext()
}

type AsignacionMultipleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignacionMultipleContext() *AsignacionMultipleContext {
	var p = new(AsignacionMultipleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_asignacionMultiple
	return p
}

func InitEmptyAsignacionMultipleContext(p *AsignacionMultipleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_asignacionMultiple
}

func (*AsignacionMultipleContext) IsAsignacionMultipleContext() {}

func NewAsignacionMultipleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionMultipleContext {
	var p = new(AsignacionMultipleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_asignacionMultiple

	return p
}

func (s *AsignacionMultipleContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionMultipleContext) ListaIDS() IListaIDSContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaIDSContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaIDSContext)
}

func (s *AsignacionMultipleContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIGUAL, 0)
}

func (s *AsignacionMultipleContext) ListaExpr() IListaExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaExprContext)
}

func (s *AsignacionMultipleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionMultipleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionMultipleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAsignacionMultiple(s)
	}
}

func (s *AsignacionMultipleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAsignacionMultiple(s)
	}
}

func (s *AsignacionMultipleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAsignacionMultiple(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) AsignacionMultiple() (localctx IAsignacionMultipleContext) {
	localctx = NewAsignacionMultipleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, gramaticaParserRULE_asignacionMultiple)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.ListaIDS()
	}
	{
		p.SetState(183)
		p.Match(gramaticaParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.ListaExpr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructDefContext is an interface to support dynamic dispatch.
type IStructDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	IDENTIFICADOR() antlr.TerminalNode
	LLAV1() antlr.TerminalNode
	LLAV2() antlr.TerminalNode
	AllAtributos() []IAtributosContext
	Atributos(i int) IAtributosContext

	// IsStructDefContext differentiates from other interfaces.
	IsStructDefContext()
}

type StructDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDefContext() *StructDefContext {
	var p = new(StructDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_structDef
	return p
}

func InitEmptyStructDefContext(p *StructDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_structDef
}

func (*StructDefContext) IsStructDefContext() {}

func NewStructDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDefContext {
	var p = new(StructDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_structDef

	return p
}

func (s *StructDefContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDefContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSTRUCT, 0)
}

func (s *StructDefContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *StructDefContext) LLAV1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAV1, 0)
}

func (s *StructDefContext) LLAV2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAV2, 0)
}

func (s *StructDefContext) AllAtributos() []IAtributosContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAtributosContext); ok {
			len++
		}
	}

	tst := make([]IAtributosContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAtributosContext); ok {
			tst[i] = t.(IAtributosContext)
			i++
		}
	}

	return tst
}

func (s *StructDefContext) Atributos(i int) IAtributosContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtributosContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtributosContext)
}

func (s *StructDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterStructDef(s)
	}
}

func (s *StructDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitStructDef(s)
	}
}

func (s *StructDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitStructDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) StructDef() (localctx IStructDefContext) {
	localctx = NewStructDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, gramaticaParserRULE_structDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Match(gramaticaParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(188)
		p.Match(gramaticaParserLLAV1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2305843009217626112) != 0) {
		{
			p.SetState(189)
			p.Atributos()
		}

		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(194)
		p.Match(gramaticaParserLLAV2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtributosContext is an interface to support dynamic dispatch.
type IAtributosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tipos() ITiposContext
	IDENTIFICADOR() antlr.TerminalNode
	PTCOMA() antlr.TerminalNode

	// IsAtributosContext differentiates from other interfaces.
	IsAtributosContext()
}

type AtributosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtributosContext() *AtributosContext {
	var p = new(AtributosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_atributos
	return p
}

func InitEmptyAtributosContext(p *AtributosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_atributos
}

func (*AtributosContext) IsAtributosContext() {}

func NewAtributosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtributosContext {
	var p = new(AtributosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_atributos

	return p
}

func (s *AtributosContext) GetParser() antlr.Parser { return s.parser }

func (s *AtributosContext) Tipos() ITiposContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITiposContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITiposContext)
}

func (s *AtributosContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *AtributosContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPTCOMA, 0)
}

func (s *AtributosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtributosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtributosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAtributos(s)
	}
}

func (s *AtributosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAtributos(s)
	}
}

func (s *AtributosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAtributos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Atributos() (localctx IAtributosContext) {
	localctx = NewAtributosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, gramaticaParserRULE_atributos)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Tipos()
	}
	{
		p.SetState(197)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)
		p.Match(gramaticaParserPTCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructInstContext is an interface to support dynamic dispatch.
type IStructInstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFICADOR() []antlr.TerminalNode
	IDENTIFICADOR(i int) antlr.TerminalNode
	IGUAL() antlr.TerminalNode
	StructInit() IStructInitContext

	// IsStructInstContext differentiates from other interfaces.
	IsStructInstContext()
}

type StructInstContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructInstContext() *StructInstContext {
	var p = new(StructInstContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_structInst
	return p
}

func InitEmptyStructInstContext(p *StructInstContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_structInst
}

func (*StructInstContext) IsStructInstContext() {}

func NewStructInstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructInstContext {
	var p = new(StructInstContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_structInst

	return p
}

func (s *StructInstContext) GetParser() antlr.Parser { return s.parser }

func (s *StructInstContext) AllIDENTIFICADOR() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserIDENTIFICADOR)
}

func (s *StructInstContext) IDENTIFICADOR(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, i)
}

func (s *StructInstContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIGUAL, 0)
}

func (s *StructInstContext) StructInit() IStructInitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructInitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructInitContext)
}

func (s *StructInstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructInstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructInstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterStructInst(s)
	}
}

func (s *StructInstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitStructInst(s)
	}
}

func (s *StructInstContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitStructInst(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) StructInst() (localctx IStructInstContext) {
	localctx = NewStructInstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, gramaticaParserRULE_structInst)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.Match(gramaticaParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
		p.StructInit()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructInitContext is an interface to support dynamic dispatch.
type IStructInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LLAV1() antlr.TerminalNode
	ListaStructs() IListaStructsContext
	LLAV2() antlr.TerminalNode

	// IsStructInitContext differentiates from other interfaces.
	IsStructInitContext()
}

type StructInitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructInitContext() *StructInitContext {
	var p = new(StructInitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_structInit
	return p
}

func InitEmptyStructInitContext(p *StructInitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_structInit
}

func (*StructInitContext) IsStructInitContext() {}

func NewStructInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructInitContext {
	var p = new(StructInitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_structInit

	return p
}

func (s *StructInitContext) GetParser() antlr.Parser { return s.parser }

func (s *StructInitContext) LLAV1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAV1, 0)
}

func (s *StructInitContext) ListaStructs() IListaStructsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaStructsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaStructsContext)
}

func (s *StructInitContext) LLAV2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAV2, 0)
}

func (s *StructInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterStructInit(s)
	}
}

func (s *StructInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitStructInit(s)
	}
}

func (s *StructInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitStructInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) StructInit() (localctx IStructInitContext) {
	localctx = NewStructInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, gramaticaParserRULE_structInit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(gramaticaParserLLAV1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.ListaStructs()
	}
	{
		p.SetState(207)
		p.Match(gramaticaParserLLAV2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListaStructsContext is an interface to support dynamic dispatch.
type IListaStructsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	AllCOMA() []antlr.TerminalNode
	COMA(i int) antlr.TerminalNode

	// IsListaStructsContext differentiates from other interfaces.
	IsListaStructsContext()
}

type ListaStructsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaStructsContext() *ListaStructsContext {
	var p = new(ListaStructsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_listaStructs
	return p
}

func InitEmptyListaStructsContext(p *ListaStructsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_listaStructs
}

func (*ListaStructsContext) IsListaStructsContext() {}

func NewListaStructsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaStructsContext {
	var p = new(ListaStructsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_listaStructs

	return p
}

func (s *ListaStructsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaStructsContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *ListaStructsContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ListaStructsContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCOMA)
}

func (s *ListaStructsContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOMA, i)
}

func (s *ListaStructsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaStructsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaStructsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterListaStructs(s)
	}
}

func (s *ListaStructsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitListaStructs(s)
	}
}

func (s *ListaStructsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitListaStructs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) ListaStructs() (localctx IListaStructsContext) {
	localctx = NewListaStructsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, gramaticaParserRULE_listaStructs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.expresion(0)
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserCOMA {
		{
			p.SetState(210)
			p.Match(gramaticaParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
			p.expresion(0)
		}

		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAccesoStructContext is an interface to support dynamic dispatch.
type IAccesoStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFICADOR() []antlr.TerminalNode
	IDENTIFICADOR(i int) antlr.TerminalNode

	// IsAccesoStructContext differentiates from other interfaces.
	IsAccesoStructContext()
}

type AccesoStructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccesoStructContext() *AccesoStructContext {
	var p = new(AccesoStructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_accesoStruct
	return p
}

func InitEmptyAccesoStructContext(p *AccesoStructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_accesoStruct
}

func (*AccesoStructContext) IsAccesoStructContext() {}

func NewAccesoStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccesoStructContext {
	var p = new(AccesoStructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_accesoStruct

	return p
}

func (s *AccesoStructContext) GetParser() antlr.Parser { return s.parser }

func (s *AccesoStructContext) AllIDENTIFICADOR() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserIDENTIFICADOR)
}

func (s *AccesoStructContext) IDENTIFICADOR(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, i)
}

func (s *AccesoStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccesoStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccesoStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAccesoStruct(s)
	}
}

func (s *AccesoStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAccesoStruct(s)
	}
}

func (s *AccesoStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAccesoStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) AccesoStruct() (localctx IAccesoStructContext) {
	localctx = NewAccesoStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, gramaticaParserRULE_accesoStruct)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(218)
		p.Match(gramaticaParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(219)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsigStructContext is an interface to support dynamic dispatch.
type IAsigStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFICADOR() []antlr.TerminalNode
	IDENTIFICADOR(i int) antlr.TerminalNode
	IGUAL() antlr.TerminalNode
	Expresion() IExpresionContext

	// IsAsigStructContext differentiates from other interfaces.
	IsAsigStructContext()
}

type AsigStructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsigStructContext() *AsigStructContext {
	var p = new(AsigStructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_asigStruct
	return p
}

func InitEmptyAsigStructContext(p *AsigStructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_asigStruct
}

func (*AsigStructContext) IsAsigStructContext() {}

func NewAsigStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsigStructContext {
	var p = new(AsigStructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_asigStruct

	return p
}

func (s *AsigStructContext) GetParser() antlr.Parser { return s.parser }

func (s *AsigStructContext) AllIDENTIFICADOR() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserIDENTIFICADOR)
}

func (s *AsigStructContext) IDENTIFICADOR(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, i)
}

func (s *AsigStructContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIGUAL, 0)
}

func (s *AsigStructContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsigStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsigStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsigStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAsigStruct(s)
	}
}

func (s *AsigStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAsigStruct(s)
	}
}

func (s *AsigStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAsigStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) AsigStruct() (localctx IAsigStructContext) {
	localctx = NewAsigStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, gramaticaParserRULE_asigStruct)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.Match(gramaticaParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Match(gramaticaParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.expresion(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunStructContext is an interface to support dynamic dispatch.
type IFunStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCIONES() antlr.TerminalNode
	AllPAR1() []antlr.TerminalNode
	PAR1(i int) antlr.TerminalNode
	AllIDENTIFICADOR() []antlr.TerminalNode
	IDENTIFICADOR(i int) antlr.TerminalNode
	POR() antlr.TerminalNode
	AllPAR2() []antlr.TerminalNode
	PAR2(i int) antlr.TerminalNode
	ListaPar() IListaParContext
	Caja() ICajaContext

	// IsFunStructContext differentiates from other interfaces.
	IsFunStructContext()
}

type FunStructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunStructContext() *FunStructContext {
	var p = new(FunStructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_funStruct
	return p
}

func InitEmptyFunStructContext(p *FunStructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_funStruct
}

func (*FunStructContext) IsFunStructContext() {}

func NewFunStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunStructContext {
	var p = new(FunStructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_funStruct

	return p
}

func (s *FunStructContext) GetParser() antlr.Parser { return s.parser }

func (s *FunStructContext) FUNCIONES() antlr.TerminalNode {
	return s.GetToken(gramaticaParserFUNCIONES, 0)
}

func (s *FunStructContext) AllPAR1() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserPAR1)
}

func (s *FunStructContext) PAR1(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR1, i)
}

func (s *FunStructContext) AllIDENTIFICADOR() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserIDENTIFICADOR)
}

func (s *FunStructContext) IDENTIFICADOR(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, i)
}

func (s *FunStructContext) POR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPOR, 0)
}

func (s *FunStructContext) AllPAR2() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserPAR2)
}

func (s *FunStructContext) PAR2(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR2, i)
}

func (s *FunStructContext) ListaPar() IListaParContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaParContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaParContext)
}

func (s *FunStructContext) Caja() ICajaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICajaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICajaContext)
}

func (s *FunStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFunStruct(s)
	}
}

func (s *FunStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFunStruct(s)
	}
}

func (s *FunStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitFunStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) FunStruct() (localctx IFunStructContext) {
	localctx = NewFunStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, gramaticaParserRULE_funStruct)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Match(gramaticaParserFUNCIONES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(228)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.Match(gramaticaParserPOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
		p.Match(gramaticaParserPAR2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(233)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(234)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(235)
		p.ListaPar()
	}
	{
		p.SetState(236)
		p.Match(gramaticaParserPAR2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.Caja()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICajaContext is an interface to support dynamic dispatch.
type ICajaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LLAV1() antlr.TerminalNode
	LLAV2() antlr.TerminalNode
	AllValores() []IValoresContext
	Valores(i int) IValoresContext

	// IsCajaContext differentiates from other interfaces.
	IsCajaContext()
}

type CajaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCajaContext() *CajaContext {
	var p = new(CajaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_caja
	return p
}

func InitEmptyCajaContext(p *CajaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_caja
}

func (*CajaContext) IsCajaContext() {}

func NewCajaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CajaContext {
	var p = new(CajaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_caja

	return p
}

func (s *CajaContext) GetParser() antlr.Parser { return s.parser }

func (s *CajaContext) LLAV1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAV1, 0)
}

func (s *CajaContext) LLAV2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAV2, 0)
}

func (s *CajaContext) AllValores() []IValoresContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValoresContext); ok {
			len++
		}
	}

	tst := make([]IValoresContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValoresContext); ok {
			tst[i] = t.(IValoresContext)
			i++
		}
	}

	return tst
}

func (s *CajaContext) Valores(i int) IValoresContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValoresContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValoresContext)
}

func (s *CajaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CajaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CajaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterCaja(s)
	}
}

func (s *CajaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitCaja(s)
	}
}

func (s *CajaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitCaja(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Caja() (localctx ICajaContext) {
	localctx = NewCajaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, gramaticaParserRULE_caja)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(gramaticaParserLLAV1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == gramaticaParserIDENTIFICADOR {
		{
			p.SetState(240)
			p.Valores()
		}

		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(245)
		p.Match(gramaticaParserLLAV2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValoresContext is an interface to support dynamic dispatch.
type IValoresContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFICADOR() []antlr.TerminalNode
	IDENTIFICADOR(i int) antlr.TerminalNode
	IGUAL() antlr.TerminalNode

	// IsValoresContext differentiates from other interfaces.
	IsValoresContext()
}

type ValoresContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValoresContext() *ValoresContext {
	var p = new(ValoresContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_valores
	return p
}

func InitEmptyValoresContext(p *ValoresContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_valores
}

func (*ValoresContext) IsValoresContext() {}

func NewValoresContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValoresContext {
	var p = new(ValoresContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_valores

	return p
}

func (s *ValoresContext) GetParser() antlr.Parser { return s.parser }

func (s *ValoresContext) AllIDENTIFICADOR() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserIDENTIFICADOR)
}

func (s *ValoresContext) IDENTIFICADOR(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, i)
}

func (s *ValoresContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIGUAL, 0)
}

func (s *ValoresContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValoresContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValoresContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterValores(s)
	}
}

func (s *ValoresContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitValores(s)
	}
}

func (s *ValoresContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitValores(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Valores() (localctx IValoresContext) {
	localctx = NewValoresContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, gramaticaParserRULE_valores)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)
		p.Match(gramaticaParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
		p.Match(gramaticaParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(251)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IActStructContext is an interface to support dynamic dispatch.
type IActStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFICADOR() []antlr.TerminalNode
	IDENTIFICADOR(i int) antlr.TerminalNode
	PAR1() antlr.TerminalNode
	ListaStructs() IListaStructsContext
	PAR2() antlr.TerminalNode

	// IsActStructContext differentiates from other interfaces.
	IsActStructContext()
}

type ActStructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActStructContext() *ActStructContext {
	var p = new(ActStructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_actStruct
	return p
}

func InitEmptyActStructContext(p *ActStructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_actStruct
}

func (*ActStructContext) IsActStructContext() {}

func NewActStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActStructContext {
	var p = new(ActStructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_actStruct

	return p
}

func (s *ActStructContext) GetParser() antlr.Parser { return s.parser }

func (s *ActStructContext) AllIDENTIFICADOR() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserIDENTIFICADOR)
}

func (s *ActStructContext) IDENTIFICADOR(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, i)
}

func (s *ActStructContext) PAR1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR1, 0)
}

func (s *ActStructContext) ListaStructs() IListaStructsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaStructsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaStructsContext)
}

func (s *ActStructContext) PAR2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR2, 0)
}

func (s *ActStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterActStruct(s)
	}
}

func (s *ActStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitActStruct(s)
	}
}

func (s *ActStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitActStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) ActStruct() (localctx IActStructContext) {
	localctx = NewActStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, gramaticaParserRULE_actStruct)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(254)
		p.Match(gramaticaParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(255)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(256)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(257)
		p.ListaStructs()
	}
	{
		p.SetState(258)
		p.Match(gramaticaParserPAR2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBloqueFuncionContext is an interface to support dynamic dispatch.
type IBloqueFuncionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInstruccion() []IInstruccionContext
	Instruccion(i int) IInstruccionContext
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext

	// IsBloqueFuncionContext differentiates from other interfaces.
	IsBloqueFuncionContext()
}

type BloqueFuncionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloqueFuncionContext() *BloqueFuncionContext {
	var p = new(BloqueFuncionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_bloqueFuncion
	return p
}

func InitEmptyBloqueFuncionContext(p *BloqueFuncionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_bloqueFuncion
}

func (*BloqueFuncionContext) IsBloqueFuncionContext() {}

func NewBloqueFuncionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueFuncionContext {
	var p = new(BloqueFuncionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_bloqueFuncion

	return p
}

func (s *BloqueFuncionContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueFuncionContext) AllInstruccion() []IInstruccionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstruccionContext); ok {
			len++
		}
	}

	tst := make([]IInstruccionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstruccionContext); ok {
			tst[i] = t.(IInstruccionContext)
			i++
		}
	}

	return tst
}

func (s *BloqueFuncionContext) Instruccion(i int) IInstruccionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *BloqueFuncionContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *BloqueFuncionContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *BloqueFuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueFuncionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloqueFuncionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterBloqueFuncion(s)
	}
}

func (s *BloqueFuncionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitBloqueFuncion(s)
	}
}

func (s *BloqueFuncionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitBloqueFuncion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) BloqueFuncion() (localctx IBloqueFuncionContext) {
	localctx = NewBloqueFuncionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, gramaticaParserRULE_bloqueFuncion)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-4)) & ^0x3f) == 0 && ((int64(1)<<(_la-4))&4469823804636937391) != 0 {
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(260)
				p.Instruccion()
			}

		case 2:
			{
				p.SetState(261)
				p.expresion(0)
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILlamadaFuncionesSinParametroContext is an interface to support dynamic dispatch.
type ILlamadaFuncionesSinParametroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFICADOR() antlr.TerminalNode
	PAR1() antlr.TerminalNode
	PAR2() antlr.TerminalNode

	// IsLlamadaFuncionesSinParametroContext differentiates from other interfaces.
	IsLlamadaFuncionesSinParametroContext()
}

type LlamadaFuncionesSinParametroContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLlamadaFuncionesSinParametroContext() *LlamadaFuncionesSinParametroContext {
	var p = new(LlamadaFuncionesSinParametroContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamadaFuncionesSinParametro
	return p
}

func InitEmptyLlamadaFuncionesSinParametroContext(p *LlamadaFuncionesSinParametroContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamadaFuncionesSinParametro
}

func (*LlamadaFuncionesSinParametroContext) IsLlamadaFuncionesSinParametroContext() {}

func NewLlamadaFuncionesSinParametroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LlamadaFuncionesSinParametroContext {
	var p = new(LlamadaFuncionesSinParametroContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_llamadaFuncionesSinParametro

	return p
}

func (s *LlamadaFuncionesSinParametroContext) GetParser() antlr.Parser { return s.parser }

func (s *LlamadaFuncionesSinParametroContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *LlamadaFuncionesSinParametroContext) PAR1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR1, 0)
}

func (s *LlamadaFuncionesSinParametroContext) PAR2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR2, 0)
}

func (s *LlamadaFuncionesSinParametroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadaFuncionesSinParametroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LlamadaFuncionesSinParametroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterLlamadaFuncionesSinParametro(s)
	}
}

func (s *LlamadaFuncionesSinParametroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitLlamadaFuncionesSinParametro(s)
	}
}

func (s *LlamadaFuncionesSinParametroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitLlamadaFuncionesSinParametro(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) LlamadaFuncionesSinParametro() (localctx ILlamadaFuncionesSinParametroContext) {
	localctx = NewLlamadaFuncionesSinParametroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, gramaticaParserRULE_llamadaFuncionesSinParametro)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
		p.Match(gramaticaParserPAR2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILlamadaFuncionesConParametroContext is an interface to support dynamic dispatch.
type ILlamadaFuncionesConParametroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFICADOR() antlr.TerminalNode
	PAR1() antlr.TerminalNode
	ListaExpr() IListaExprContext
	PAR2() antlr.TerminalNode

	// IsLlamadaFuncionesConParametroContext differentiates from other interfaces.
	IsLlamadaFuncionesConParametroContext()
}

type LlamadaFuncionesConParametroContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLlamadaFuncionesConParametroContext() *LlamadaFuncionesConParametroContext {
	var p = new(LlamadaFuncionesConParametroContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamadaFuncionesConParametro
	return p
}

func InitEmptyLlamadaFuncionesConParametroContext(p *LlamadaFuncionesConParametroContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamadaFuncionesConParametro
}

func (*LlamadaFuncionesConParametroContext) IsLlamadaFuncionesConParametroContext() {}

func NewLlamadaFuncionesConParametroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LlamadaFuncionesConParametroContext {
	var p = new(LlamadaFuncionesConParametroContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_llamadaFuncionesConParametro

	return p
}

func (s *LlamadaFuncionesConParametroContext) GetParser() antlr.Parser { return s.parser }

func (s *LlamadaFuncionesConParametroContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *LlamadaFuncionesConParametroContext) PAR1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR1, 0)
}

func (s *LlamadaFuncionesConParametroContext) ListaExpr() IListaExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaExprContext)
}

func (s *LlamadaFuncionesConParametroContext) PAR2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR2, 0)
}

func (s *LlamadaFuncionesConParametroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadaFuncionesConParametroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LlamadaFuncionesConParametroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterLlamadaFuncionesConParametro(s)
	}
}

func (s *LlamadaFuncionesConParametroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitLlamadaFuncionesConParametro(s)
	}
}

func (s *LlamadaFuncionesConParametroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitLlamadaFuncionesConParametro(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) LlamadaFuncionesConParametro() (localctx ILlamadaFuncionesConParametroContext) {
	localctx = NewLlamadaFuncionesConParametroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, gramaticaParserRULE_llamadaFuncionesConParametro)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(272)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(273)
		p.ListaExpr()
	}
	{
		p.SetState(274)
		p.Match(gramaticaParserPAR2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFnSinParametroContext is an interface to support dynamic dispatch.
type IFnSinParametroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCIONES() antlr.TerminalNode
	IDENTIFICADOR() antlr.TerminalNode
	PAR1() antlr.TerminalNode
	PAR2() antlr.TerminalNode
	LLAV1() antlr.TerminalNode
	BloqueFuncion() IBloqueFuncionContext
	LLAV2() antlr.TerminalNode
	TipoRetorno() ITipoRetornoContext

	// IsFnSinParametroContext differentiates from other interfaces.
	IsFnSinParametroContext()
}

type FnSinParametroContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnSinParametroContext() *FnSinParametroContext {
	var p = new(FnSinParametroContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_fnSinParametro
	return p
}

func InitEmptyFnSinParametroContext(p *FnSinParametroContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_fnSinParametro
}

func (*FnSinParametroContext) IsFnSinParametroContext() {}

func NewFnSinParametroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnSinParametroContext {
	var p = new(FnSinParametroContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_fnSinParametro

	return p
}

func (s *FnSinParametroContext) GetParser() antlr.Parser { return s.parser }

func (s *FnSinParametroContext) FUNCIONES() antlr.TerminalNode {
	return s.GetToken(gramaticaParserFUNCIONES, 0)
}

func (s *FnSinParametroContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *FnSinParametroContext) PAR1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR1, 0)
}

func (s *FnSinParametroContext) PAR2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR2, 0)
}

func (s *FnSinParametroContext) LLAV1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAV1, 0)
}

func (s *FnSinParametroContext) BloqueFuncion() IBloqueFuncionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueFuncionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueFuncionContext)
}

func (s *FnSinParametroContext) LLAV2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAV2, 0)
}

func (s *FnSinParametroContext) TipoRetorno() ITipoRetornoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoRetornoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoRetornoContext)
}

func (s *FnSinParametroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnSinParametroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnSinParametroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFnSinParametro(s)
	}
}

func (s *FnSinParametroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFnSinParametro(s)
	}
}

func (s *FnSinParametroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitFnSinParametro(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) FnSinParametro() (localctx IFnSinParametroContext) {
	localctx = NewFnSinParametroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, gramaticaParserRULE_fnSinParametro)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Match(gramaticaParserFUNCIONES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(277)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(278)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(279)
		p.Match(gramaticaParserPAR2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-18)) & ^0x3f) == 0 && ((int64(1)<<(_la-18))&149533581377551) != 0 {
		{
			p.SetState(280)
			p.TipoRetorno()
		}

	}
	{
		p.SetState(283)
		p.Match(gramaticaParserLLAV1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(284)
		p.BloqueFuncion()
	}
	{
		p.SetState(285)
		p.Match(gramaticaParserLLAV2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFnConParametroContext is an interface to support dynamic dispatch.
type IFnConParametroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCIONES() antlr.TerminalNode
	IDENTIFICADOR() antlr.TerminalNode
	PAR1() antlr.TerminalNode
	ListaPar() IListaParContext
	PAR2() antlr.TerminalNode
	LLAV1() antlr.TerminalNode
	BloqueFuncion() IBloqueFuncionContext
	LLAV2() antlr.TerminalNode
	TipoRetorno() ITipoRetornoContext

	// IsFnConParametroContext differentiates from other interfaces.
	IsFnConParametroContext()
}

type FnConParametroContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnConParametroContext() *FnConParametroContext {
	var p = new(FnConParametroContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_fnConParametro
	return p
}

func InitEmptyFnConParametroContext(p *FnConParametroContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_fnConParametro
}

func (*FnConParametroContext) IsFnConParametroContext() {}

func NewFnConParametroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnConParametroContext {
	var p = new(FnConParametroContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_fnConParametro

	return p
}

func (s *FnConParametroContext) GetParser() antlr.Parser { return s.parser }

func (s *FnConParametroContext) FUNCIONES() antlr.TerminalNode {
	return s.GetToken(gramaticaParserFUNCIONES, 0)
}

func (s *FnConParametroContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *FnConParametroContext) PAR1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR1, 0)
}

func (s *FnConParametroContext) ListaPar() IListaParContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaParContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaParContext)
}

func (s *FnConParametroContext) PAR2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR2, 0)
}

func (s *FnConParametroContext) LLAV1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAV1, 0)
}

func (s *FnConParametroContext) BloqueFuncion() IBloqueFuncionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueFuncionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueFuncionContext)
}

func (s *FnConParametroContext) LLAV2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAV2, 0)
}

func (s *FnConParametroContext) TipoRetorno() ITipoRetornoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoRetornoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoRetornoContext)
}

func (s *FnConParametroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnConParametroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnConParametroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFnConParametro(s)
	}
}

func (s *FnConParametroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFnConParametro(s)
	}
}

func (s *FnConParametroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitFnConParametro(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) FnConParametro() (localctx IFnConParametroContext) {
	localctx = NewFnConParametroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, gramaticaParserRULE_fnConParametro)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
		p.Match(gramaticaParserFUNCIONES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(288)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(289)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(290)
		p.ListaPar()
	}
	{
		p.SetState(291)
		p.Match(gramaticaParserPAR2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-18)) & ^0x3f) == 0 && ((int64(1)<<(_la-18))&149533581377551) != 0 {
		{
			p.SetState(292)
			p.TipoRetorno()
		}

	}
	{
		p.SetState(295)
		p.Match(gramaticaParserLLAV1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(296)
		p.BloqueFuncion()
	}
	{
		p.SetState(297)
		p.Match(gramaticaParserLLAV2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISliceDefContext is an interface to support dynamic dispatch.
type ISliceDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFICADOR() antlr.TerminalNode
	IGUAL() antlr.TerminalNode
	SliceLiteral() ISliceLiteralContext

	// IsSliceDefContext differentiates from other interfaces.
	IsSliceDefContext()
}

type SliceDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceDefContext() *SliceDefContext {
	var p = new(SliceDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sliceDef
	return p
}

func InitEmptySliceDefContext(p *SliceDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sliceDef
}

func (*SliceDefContext) IsSliceDefContext() {}

func NewSliceDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceDefContext {
	var p = new(SliceDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_sliceDef

	return p
}

func (s *SliceDefContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceDefContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *SliceDefContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIGUAL, 0)
}

func (s *SliceDefContext) SliceLiteral() ISliceLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceLiteralContext)
}

func (s *SliceDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSliceDef(s)
	}
}

func (s *SliceDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSliceDef(s)
	}
}

func (s *SliceDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSliceDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) SliceDef() (localctx ISliceDefContext) {
	localctx = NewSliceDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, gramaticaParserRULE_sliceDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(300)
		p.Match(gramaticaParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(301)
		p.SliceLiteral()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISliceLiteralContext is an interface to support dynamic dispatch.
type ISliceLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCOR1() []antlr.TerminalNode
	COR1(i int) antlr.TerminalNode
	AllCOR2() []antlr.TerminalNode
	COR2(i int) antlr.TerminalNode
	Tipos() ITiposContext
	LLAV1() antlr.TerminalNode
	ListaExpr() IListaExprContext
	LLAV2() antlr.TerminalNode
	ListaExprList() IListaExprListContext

	// IsSliceLiteralContext differentiates from other interfaces.
	IsSliceLiteralContext()
}

type SliceLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceLiteralContext() *SliceLiteralContext {
	var p = new(SliceLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sliceLiteral
	return p
}

func InitEmptySliceLiteralContext(p *SliceLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sliceLiteral
}

func (*SliceLiteralContext) IsSliceLiteralContext() {}

func NewSliceLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceLiteralContext {
	var p = new(SliceLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_sliceLiteral

	return p
}

func (s *SliceLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceLiteralContext) AllCOR1() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCOR1)
}

func (s *SliceLiteralContext) COR1(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOR1, i)
}

func (s *SliceLiteralContext) AllCOR2() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCOR2)
}

func (s *SliceLiteralContext) COR2(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOR2, i)
}

func (s *SliceLiteralContext) Tipos() ITiposContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITiposContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITiposContext)
}

func (s *SliceLiteralContext) LLAV1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAV1, 0)
}

func (s *SliceLiteralContext) ListaExpr() IListaExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaExprContext)
}

func (s *SliceLiteralContext) LLAV2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAV2, 0)
}

func (s *SliceLiteralContext) ListaExprList() IListaExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaExprListContext)
}

func (s *SliceLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSliceLiteral(s)
	}
}

func (s *SliceLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSliceLiteral(s)
	}
}

func (s *SliceLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSliceLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) SliceLiteral() (localctx ISliceLiteralContext) {
	localctx = NewSliceLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, gramaticaParserRULE_sliceLiteral)
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(303)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(304)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)
			p.Tipos()
		}
		{
			p.SetState(306)
			p.Match(gramaticaParserLLAV1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(307)
			p.ListaExpr()
		}
		{
			p.SetState(308)
			p.Match(gramaticaParserLLAV2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(310)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(311)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(313)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(314)
			p.Tipos()
		}
		{
			p.SetState(315)
			p.Match(gramaticaParserLLAV1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)
			p.ListaExprList()
		}
		{
			p.SetState(317)
			p.Match(gramaticaParserLLAV2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAccesoElementoSliceContext is an interface to support dynamic dispatch.
type IAccesoElementoSliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFICADOR() antlr.TerminalNode
	AllCOR1() []antlr.TerminalNode
	COR1(i int) antlr.TerminalNode
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	AllCOR2() []antlr.TerminalNode
	COR2(i int) antlr.TerminalNode

	// IsAccesoElementoSliceContext differentiates from other interfaces.
	IsAccesoElementoSliceContext()
}

type AccesoElementoSliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccesoElementoSliceContext() *AccesoElementoSliceContext {
	var p = new(AccesoElementoSliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_accesoElementoSlice
	return p
}

func InitEmptyAccesoElementoSliceContext(p *AccesoElementoSliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_accesoElementoSlice
}

func (*AccesoElementoSliceContext) IsAccesoElementoSliceContext() {}

func NewAccesoElementoSliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccesoElementoSliceContext {
	var p = new(AccesoElementoSliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_accesoElementoSlice

	return p
}

func (s *AccesoElementoSliceContext) GetParser() antlr.Parser { return s.parser }

func (s *AccesoElementoSliceContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *AccesoElementoSliceContext) AllCOR1() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCOR1)
}

func (s *AccesoElementoSliceContext) COR1(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOR1, i)
}

func (s *AccesoElementoSliceContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *AccesoElementoSliceContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AccesoElementoSliceContext) AllCOR2() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCOR2)
}

func (s *AccesoElementoSliceContext) COR2(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOR2, i)
}

func (s *AccesoElementoSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccesoElementoSliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccesoElementoSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAccesoElementoSlice(s)
	}
}

func (s *AccesoElementoSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAccesoElementoSlice(s)
	}
}

func (s *AccesoElementoSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAccesoElementoSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) AccesoElementoSlice() (localctx IAccesoElementoSliceContext) {
	localctx = NewAccesoElementoSliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, gramaticaParserRULE_accesoElementoSlice)
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(321)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(322)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.expresion(0)
		}
		{
			p.SetState(324)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(326)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(327)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.expresion(0)
		}
		{
			p.SetState(329)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(330)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.expresion(0)
		}
		{
			p.SetState(332)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IModificacionElementoSliceContext is an interface to support dynamic dispatch.
type IModificacionElementoSliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFICADOR() antlr.TerminalNode
	AllCOR1() []antlr.TerminalNode
	COR1(i int) antlr.TerminalNode
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	AllCOR2() []antlr.TerminalNode
	COR2(i int) antlr.TerminalNode
	IGUAL() antlr.TerminalNode

	// IsModificacionElementoSliceContext differentiates from other interfaces.
	IsModificacionElementoSliceContext()
}

type ModificacionElementoSliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModificacionElementoSliceContext() *ModificacionElementoSliceContext {
	var p = new(ModificacionElementoSliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_modificacionElementoSlice
	return p
}

func InitEmptyModificacionElementoSliceContext(p *ModificacionElementoSliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_modificacionElementoSlice
}

func (*ModificacionElementoSliceContext) IsModificacionElementoSliceContext() {}

func NewModificacionElementoSliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModificacionElementoSliceContext {
	var p = new(ModificacionElementoSliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_modificacionElementoSlice

	return p
}

func (s *ModificacionElementoSliceContext) GetParser() antlr.Parser { return s.parser }

func (s *ModificacionElementoSliceContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *ModificacionElementoSliceContext) AllCOR1() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCOR1)
}

func (s *ModificacionElementoSliceContext) COR1(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOR1, i)
}

func (s *ModificacionElementoSliceContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *ModificacionElementoSliceContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ModificacionElementoSliceContext) AllCOR2() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCOR2)
}

func (s *ModificacionElementoSliceContext) COR2(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOR2, i)
}

func (s *ModificacionElementoSliceContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIGUAL, 0)
}

func (s *ModificacionElementoSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModificacionElementoSliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModificacionElementoSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterModificacionElementoSlice(s)
	}
}

func (s *ModificacionElementoSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitModificacionElementoSlice(s)
	}
}

func (s *ModificacionElementoSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitModificacionElementoSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) ModificacionElementoSlice() (localctx IModificacionElementoSliceContext) {
	localctx = NewModificacionElementoSliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, gramaticaParserRULE_modificacionElementoSlice)
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(336)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
			p.expresion(0)
		}
		{
			p.SetState(339)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(340)
			p.Match(gramaticaParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.expresion(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(343)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(344)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)
			p.expresion(0)
		}
		{
			p.SetState(346)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(347)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(348)
			p.expresion(0)
		}
		{
			p.SetState(349)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(350)
			p.Match(gramaticaParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)
			p.expresion(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFnIndexOfContext is an interface to support dynamic dispatch.
type IFnIndexOfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INDEXOF() antlr.TerminalNode
	PAR1() antlr.TerminalNode
	ListaExpr() IListaExprContext
	PAR2() antlr.TerminalNode

	// IsFnIndexOfContext differentiates from other interfaces.
	IsFnIndexOfContext()
}

type FnIndexOfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnIndexOfContext() *FnIndexOfContext {
	var p = new(FnIndexOfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_fnIndexOf
	return p
}

func InitEmptyFnIndexOfContext(p *FnIndexOfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_fnIndexOf
}

func (*FnIndexOfContext) IsFnIndexOfContext() {}

func NewFnIndexOfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnIndexOfContext {
	var p = new(FnIndexOfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_fnIndexOf

	return p
}

func (s *FnIndexOfContext) GetParser() antlr.Parser { return s.parser }

func (s *FnIndexOfContext) INDEXOF() antlr.TerminalNode {
	return s.GetToken(gramaticaParserINDEXOF, 0)
}

func (s *FnIndexOfContext) PAR1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR1, 0)
}

func (s *FnIndexOfContext) ListaExpr() IListaExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaExprContext)
}

func (s *FnIndexOfContext) PAR2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR2, 0)
}

func (s *FnIndexOfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnIndexOfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnIndexOfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFnIndexOf(s)
	}
}

func (s *FnIndexOfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFnIndexOf(s)
	}
}

func (s *FnIndexOfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitFnIndexOf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) FnIndexOf() (localctx IFnIndexOfContext) {
	localctx = NewFnIndexOfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, gramaticaParserRULE_fnIndexOf)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Match(gramaticaParserINDEXOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(356)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(357)
		p.ListaExpr()
	}
	{
		p.SetState(358)
		p.Match(gramaticaParserPAR2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFnJoinContext is an interface to support dynamic dispatch.
type IFnJoinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JOIN() antlr.TerminalNode
	PAR1() antlr.TerminalNode
	ListaExpr() IListaExprContext
	PAR2() antlr.TerminalNode

	// IsFnJoinContext differentiates from other interfaces.
	IsFnJoinContext()
}

type FnJoinContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnJoinContext() *FnJoinContext {
	var p = new(FnJoinContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_fnJoin
	return p
}

func InitEmptyFnJoinContext(p *FnJoinContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_fnJoin
}

func (*FnJoinContext) IsFnJoinContext() {}

func NewFnJoinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnJoinContext {
	var p = new(FnJoinContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_fnJoin

	return p
}

func (s *FnJoinContext) GetParser() antlr.Parser { return s.parser }

func (s *FnJoinContext) JOIN() antlr.TerminalNode {
	return s.GetToken(gramaticaParserJOIN, 0)
}

func (s *FnJoinContext) PAR1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR1, 0)
}

func (s *FnJoinContext) ListaExpr() IListaExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaExprContext)
}

func (s *FnJoinContext) PAR2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR2, 0)
}

func (s *FnJoinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnJoinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnJoinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFnJoin(s)
	}
}

func (s *FnJoinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFnJoin(s)
	}
}

func (s *FnJoinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitFnJoin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) FnJoin() (localctx IFnJoinContext) {
	localctx = NewFnJoinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, gramaticaParserRULE_fnJoin)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(360)
		p.Match(gramaticaParserJOIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(361)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(362)
		p.ListaExpr()
	}
	{
		p.SetState(363)
		p.Match(gramaticaParserPAR2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFnLenContext is an interface to support dynamic dispatch.
type IFnLenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEN() antlr.TerminalNode
	PAR1() antlr.TerminalNode
	ListaExpr() IListaExprContext
	PAR2() antlr.TerminalNode

	// IsFnLenContext differentiates from other interfaces.
	IsFnLenContext()
}

type FnLenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnLenContext() *FnLenContext {
	var p = new(FnLenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_fnLen
	return p
}

func InitEmptyFnLenContext(p *FnLenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_fnLen
}

func (*FnLenContext) IsFnLenContext() {}

func NewFnLenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnLenContext {
	var p = new(FnLenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_fnLen

	return p
}

func (s *FnLenContext) GetParser() antlr.Parser { return s.parser }

func (s *FnLenContext) LEN() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLEN, 0)
}

func (s *FnLenContext) PAR1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR1, 0)
}

func (s *FnLenContext) ListaExpr() IListaExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaExprContext)
}

func (s *FnLenContext) PAR2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR2, 0)
}

func (s *FnLenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnLenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnLenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFnLen(s)
	}
}

func (s *FnLenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFnLen(s)
	}
}

func (s *FnLenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitFnLen(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) FnLen() (localctx IFnLenContext) {
	localctx = NewFnLenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, gramaticaParserRULE_fnLen)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.Match(gramaticaParserLEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(366)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(367)
		p.ListaExpr()
	}
	{
		p.SetState(368)
		p.Match(gramaticaParserPAR2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFnAppendContext is an interface to support dynamic dispatch.
type IFnAppendContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	APPEND() antlr.TerminalNode
	PAR1() antlr.TerminalNode
	ListaExpr() IListaExprContext
	PAR2() antlr.TerminalNode

	// IsFnAppendContext differentiates from other interfaces.
	IsFnAppendContext()
}

type FnAppendContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnAppendContext() *FnAppendContext {
	var p = new(FnAppendContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_fnAppend
	return p
}

func InitEmptyFnAppendContext(p *FnAppendContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_fnAppend
}

func (*FnAppendContext) IsFnAppendContext() {}

func NewFnAppendContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnAppendContext {
	var p = new(FnAppendContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_fnAppend

	return p
}

func (s *FnAppendContext) GetParser() antlr.Parser { return s.parser }

func (s *FnAppendContext) APPEND() antlr.TerminalNode {
	return s.GetToken(gramaticaParserAPPEND, 0)
}

func (s *FnAppendContext) PAR1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR1, 0)
}

func (s *FnAppendContext) ListaExpr() IListaExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaExprContext)
}

func (s *FnAppendContext) PAR2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR2, 0)
}

func (s *FnAppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnAppendContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnAppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFnAppend(s)
	}
}

func (s *FnAppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFnAppend(s)
	}
}

func (s *FnAppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitFnAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) FnAppend() (localctx IFnAppendContext) {
	localctx = NewFnAppendContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, gramaticaParserRULE_fnAppend)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.Match(gramaticaParserAPPEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(371)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(372)
		p.ListaExpr()
	}
	{
		p.SetState(373)
		p.Match(gramaticaParserPAR2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracionSliceVacioContext is an interface to support dynamic dispatch.
type IDeclaracionSliceVacioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MUT() antlr.TerminalNode
	IDENTIFICADOR() antlr.TerminalNode
	AllCOR1() []antlr.TerminalNode
	COR1(i int) antlr.TerminalNode
	AllCOR2() []antlr.TerminalNode
	COR2(i int) antlr.TerminalNode
	Tipos() ITiposContext

	// IsDeclaracionSliceVacioContext differentiates from other interfaces.
	IsDeclaracionSliceVacioContext()
}

type DeclaracionSliceVacioContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracionSliceVacioContext() *DeclaracionSliceVacioContext {
	var p = new(DeclaracionSliceVacioContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_declaracionSliceVacio
	return p
}

func InitEmptyDeclaracionSliceVacioContext(p *DeclaracionSliceVacioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_declaracionSliceVacio
}

func (*DeclaracionSliceVacioContext) IsDeclaracionSliceVacioContext() {}

func NewDeclaracionSliceVacioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionSliceVacioContext {
	var p = new(DeclaracionSliceVacioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_declaracionSliceVacio

	return p
}

func (s *DeclaracionSliceVacioContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionSliceVacioContext) MUT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserMUT, 0)
}

func (s *DeclaracionSliceVacioContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *DeclaracionSliceVacioContext) AllCOR1() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCOR1)
}

func (s *DeclaracionSliceVacioContext) COR1(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOR1, i)
}

func (s *DeclaracionSliceVacioContext) AllCOR2() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCOR2)
}

func (s *DeclaracionSliceVacioContext) COR2(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOR2, i)
}

func (s *DeclaracionSliceVacioContext) Tipos() ITiposContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITiposContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITiposContext)
}

func (s *DeclaracionSliceVacioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionSliceVacioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionSliceVacioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterDeclaracionSliceVacio(s)
	}
}

func (s *DeclaracionSliceVacioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitDeclaracionSliceVacio(s)
	}
}

func (s *DeclaracionSliceVacioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitDeclaracionSliceVacio(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) DeclaracionSliceVacio() (localctx IDeclaracionSliceVacioContext) {
	localctx = NewDeclaracionSliceVacioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, gramaticaParserRULE_declaracionSliceVacio)
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(375)
			p.Match(gramaticaParserMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(376)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(377)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(378)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(379)
			p.Tipos()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(380)
			p.Match(gramaticaParserMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(381)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(382)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(383)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(384)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(385)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)
			p.Tipos()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipoRetornoContext is an interface to support dynamic dispatch.
type ITipoRetornoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode
	RUNE() antlr.TerminalNode
	IDENTIFICADOR() antlr.TerminalNode

	// IsTipoRetornoContext differentiates from other interfaces.
	IsTipoRetornoContext()
}

type TipoRetornoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipoRetornoContext() *TipoRetornoContext {
	var p = new(TipoRetornoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_tipoRetorno
	return p
}

func InitEmptyTipoRetornoContext(p *TipoRetornoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_tipoRetorno
}

func (*TipoRetornoContext) IsTipoRetornoContext() {}

func NewTipoRetornoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoRetornoContext {
	var p = new(TipoRetornoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_tipoRetorno

	return p
}

func (s *TipoRetornoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoRetornoContext) INT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserINT, 0)
}

func (s *TipoRetornoContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserFLOAT, 0)
}

func (s *TipoRetornoContext) BOOL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserBOOL, 0)
}

func (s *TipoRetornoContext) STRING() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSTRING, 0)
}

func (s *TipoRetornoContext) RUNE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRUNE, 0)
}

func (s *TipoRetornoContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *TipoRetornoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoRetornoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoRetornoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterTipoRetorno(s)
	}
}

func (s *TipoRetornoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitTipoRetorno(s)
	}
}

func (s *TipoRetornoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitTipoRetorno(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) TipoRetorno() (localctx ITipoRetornoContext) {
	localctx = NewTipoRetornoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, gramaticaParserRULE_tipoRetorno)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(389)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-18)) & ^0x3f) == 0 && ((int64(1)<<(_la-18))&149533581377551) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRetornoContext is an interface to support dynamic dispatch.
type IRetornoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expresion() IExpresionContext
	PTCOMA() antlr.TerminalNode

	// IsRetornoContext differentiates from other interfaces.
	IsRetornoContext()
}

type RetornoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetornoContext() *RetornoContext {
	var p = new(RetornoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_retorno
	return p
}

func InitEmptyRetornoContext(p *RetornoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_retorno
}

func (*RetornoContext) IsRetornoContext() {}

func NewRetornoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetornoContext {
	var p = new(RetornoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_retorno

	return p
}

func (s *RetornoContext) GetParser() antlr.Parser { return s.parser }

func (s *RetornoContext) RETURN() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRETURN, 0)
}

func (s *RetornoContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *RetornoContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPTCOMA, 0)
}

func (s *RetornoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetornoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RetornoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterRetorno(s)
	}
}

func (s *RetornoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitRetorno(s)
	}
}

func (s *RetornoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitRetorno(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Retorno() (localctx IRetornoContext) {
	localctx = NewRetornoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, gramaticaParserRULE_retorno)
	var _la int

	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(391)
			p.Match(gramaticaParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(392)
			p.expresion(0)
		}
		p.SetState(394)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gramaticaParserPTCOMA {
			{
				p.SetState(393)
				p.Match(gramaticaParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(396)
			p.Match(gramaticaParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == gramaticaParserPTCOMA {
			{
				p.SetState(397)
				p.Match(gramaticaParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFnTypeOfContext is an interface to support dynamic dispatch.
type IFnTypeOfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPEOF() antlr.TerminalNode
	PAR1() antlr.TerminalNode
	ListaExpr() IListaExprContext
	PAR2() antlr.TerminalNode

	// IsFnTypeOfContext differentiates from other interfaces.
	IsFnTypeOfContext()
}

type FnTypeOfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnTypeOfContext() *FnTypeOfContext {
	var p = new(FnTypeOfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_fnTypeOf
	return p
}

func InitEmptyFnTypeOfContext(p *FnTypeOfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_fnTypeOf
}

func (*FnTypeOfContext) IsFnTypeOfContext() {}

func NewFnTypeOfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnTypeOfContext {
	var p = new(FnTypeOfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_fnTypeOf

	return p
}

func (s *FnTypeOfContext) GetParser() antlr.Parser { return s.parser }

func (s *FnTypeOfContext) TYPEOF() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTYPEOF, 0)
}

func (s *FnTypeOfContext) PAR1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR1, 0)
}

func (s *FnTypeOfContext) ListaExpr() IListaExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaExprContext)
}

func (s *FnTypeOfContext) PAR2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR2, 0)
}

func (s *FnTypeOfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnTypeOfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnTypeOfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFnTypeOf(s)
	}
}

func (s *FnTypeOfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFnTypeOf(s)
	}
}

func (s *FnTypeOfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitFnTypeOf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) FnTypeOf() (localctx IFnTypeOfContext) {
	localctx = NewFnTypeOfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, gramaticaParserRULE_fnTypeOf)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(402)
		p.Match(gramaticaParserTYPEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(403)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(404)
		p.ListaExpr()
	}
	{
		p.SetState(405)
		p.Match(gramaticaParserPAR2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFnAtoiContext is an interface to support dynamic dispatch.
type IFnAtoiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ATOI() antlr.TerminalNode
	PAR1() antlr.TerminalNode
	ListaExpr() IListaExprContext
	PAR2() antlr.TerminalNode

	// IsFnAtoiContext differentiates from other interfaces.
	IsFnAtoiContext()
}

type FnAtoiContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnAtoiContext() *FnAtoiContext {
	var p = new(FnAtoiContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_fnAtoi
	return p
}

func InitEmptyFnAtoiContext(p *FnAtoiContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_fnAtoi
}

func (*FnAtoiContext) IsFnAtoiContext() {}

func NewFnAtoiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnAtoiContext {
	var p = new(FnAtoiContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_fnAtoi

	return p
}

func (s *FnAtoiContext) GetParser() antlr.Parser { return s.parser }

func (s *FnAtoiContext) ATOI() antlr.TerminalNode {
	return s.GetToken(gramaticaParserATOI, 0)
}

func (s *FnAtoiContext) PAR1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR1, 0)
}

func (s *FnAtoiContext) ListaExpr() IListaExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaExprContext)
}

func (s *FnAtoiContext) PAR2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR2, 0)
}

func (s *FnAtoiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnAtoiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnAtoiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFnAtoi(s)
	}
}

func (s *FnAtoiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFnAtoi(s)
	}
}

func (s *FnAtoiContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitFnAtoi(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) FnAtoi() (localctx IFnAtoiContext) {
	localctx = NewFnAtoiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, gramaticaParserRULE_fnAtoi)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(407)
		p.Match(gramaticaParserATOI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(408)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(409)
		p.ListaExpr()
	}
	{
		p.SetState(410)
		p.Match(gramaticaParserPAR2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFnParseToFloatContext is an interface to support dynamic dispatch.
type IFnParseToFloatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PARSEFLOAT() antlr.TerminalNode
	PAR1() antlr.TerminalNode
	ListaExpr() IListaExprContext
	PAR2() antlr.TerminalNode

	// IsFnParseToFloatContext differentiates from other interfaces.
	IsFnParseToFloatContext()
}

type FnParseToFloatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnParseToFloatContext() *FnParseToFloatContext {
	var p = new(FnParseToFloatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_fnParseToFloat
	return p
}

func InitEmptyFnParseToFloatContext(p *FnParseToFloatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_fnParseToFloat
}

func (*FnParseToFloatContext) IsFnParseToFloatContext() {}

func NewFnParseToFloatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnParseToFloatContext {
	var p = new(FnParseToFloatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_fnParseToFloat

	return p
}

func (s *FnParseToFloatContext) GetParser() antlr.Parser { return s.parser }

func (s *FnParseToFloatContext) PARSEFLOAT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPARSEFLOAT, 0)
}

func (s *FnParseToFloatContext) PAR1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR1, 0)
}

func (s *FnParseToFloatContext) ListaExpr() IListaExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaExprContext)
}

func (s *FnParseToFloatContext) PAR2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR2, 0)
}

func (s *FnParseToFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnParseToFloatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnParseToFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFnParseToFloat(s)
	}
}

func (s *FnParseToFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFnParseToFloat(s)
	}
}

func (s *FnParseToFloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitFnParseToFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) FnParseToFloat() (localctx IFnParseToFloatContext) {
	localctx = NewFnParseToFloatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, gramaticaParserRULE_fnParseToFloat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(412)
		p.Match(gramaticaParserPARSEFLOAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(413)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(414)
		p.ListaExpr()
	}
	{
		p.SetState(415)
		p.Match(gramaticaParserPAR2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsigIncreContext is an interface to support dynamic dispatch.
type IAsigIncreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFICADOR() antlr.TerminalNode
	Expresion() IExpresionContext

	// IsAsigIncreContext differentiates from other interfaces.
	IsAsigIncreContext()
}

type AsigIncreContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsigIncreContext() *AsigIncreContext {
	var p = new(AsigIncreContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_asigIncre
	return p
}

func InitEmptyAsigIncreContext(p *AsigIncreContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_asigIncre
}

func (*AsigIncreContext) IsAsigIncreContext() {}

func NewAsigIncreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsigIncreContext {
	var p = new(AsigIncreContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_asigIncre

	return p
}

func (s *AsigIncreContext) GetParser() antlr.Parser { return s.parser }

func (s *AsigIncreContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *AsigIncreContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsigIncreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsigIncreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsigIncreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAsigIncre(s)
	}
}

func (s *AsigIncreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAsigIncre(s)
	}
}

func (s *AsigIncreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAsigIncre(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) AsigIncre() (localctx IAsigIncreContext) {
	localctx = NewAsigIncreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, gramaticaParserRULE_asigIncre)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(417)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(418)
		p.Match(gramaticaParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(419)
		p.expresion(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsigDecreContext is an interface to support dynamic dispatch.
type IAsigDecreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFICADOR() antlr.TerminalNode
	Expresion() IExpresionContext

	// IsAsigDecreContext differentiates from other interfaces.
	IsAsigDecreContext()
}

type AsigDecreContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsigDecreContext() *AsigDecreContext {
	var p = new(AsigDecreContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_asigDecre
	return p
}

func InitEmptyAsigDecreContext(p *AsigDecreContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_asigDecre
}

func (*AsigDecreContext) IsAsigDecreContext() {}

func NewAsigDecreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsigDecreContext {
	var p = new(AsigDecreContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_asigDecre

	return p
}

func (s *AsigDecreContext) GetParser() antlr.Parser { return s.parser }

func (s *AsigDecreContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *AsigDecreContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsigDecreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsigDecreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsigDecreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAsigDecre(s)
	}
}

func (s *AsigDecreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAsigDecre(s)
	}
}

func (s *AsigDecreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAsigDecre(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) AsigDecre() (localctx IAsigDecreContext) {
	localctx = NewAsigDecreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, gramaticaParserRULE_asigDecre)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(421)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(422)
		p.Match(gramaticaParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(423)
		p.expresion(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_asignacion
	return p
}

func InitEmptyAsignacionContext(p *AsignacionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_asignacion
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) CopyAll(ctx *AsignacionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AsignacionIncrementoContext struct {
	AsignacionContext
}

func NewAsignacionIncrementoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionIncrementoContext {
	var p = new(AsignacionIncrementoContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *AsignacionIncrementoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionIncrementoContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *AsignacionIncrementoContext) FINCREMENTO() antlr.TerminalNode {
	return s.GetToken(gramaticaParserFINCREMENTO, 0)
}

func (s *AsignacionIncrementoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAsignacionIncremento(s)
	}
}

func (s *AsignacionIncrementoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAsignacionIncremento(s)
	}
}

func (s *AsignacionIncrementoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAsignacionIncremento(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionNormalContext struct {
	AsignacionContext
}

func NewAsignacionNormalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionNormalContext {
	var p = new(AsignacionNormalContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *AsignacionNormalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionNormalContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *AsignacionNormalContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIGUAL, 0)
}

func (s *AsignacionNormalContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignacionNormalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAsignacionNormal(s)
	}
}

func (s *AsignacionNormalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAsignacionNormal(s)
	}
}

func (s *AsignacionNormalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAsignacionNormal(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionDecrementoContext struct {
	AsignacionContext
}

func NewAsignacionDecrementoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionDecrementoContext {
	var p = new(AsignacionDecrementoContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *AsignacionDecrementoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionDecrementoContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *AsignacionDecrementoContext) FDECREMENTO() antlr.TerminalNode {
	return s.GetToken(gramaticaParserFDECREMENTO, 0)
}

func (s *AsignacionDecrementoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAsignacionDecremento(s)
	}
}

func (s *AsignacionDecrementoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAsignacionDecremento(s)
	}
}

func (s *AsignacionDecrementoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitAsignacionDecremento(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Asignacion() (localctx IAsignacionContext) {
	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, gramaticaParserRULE_asignacion)
	p.SetState(432)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAsignacionNormalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(425)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(426)
			p.Match(gramaticaParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(427)
			p.expresion(0)
		}

	case 2:
		localctx = NewAsignacionIncrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(428)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(429)
			p.Match(gramaticaParserFINCREMENTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewAsignacionDecrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(430)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(431)
			p.Match(gramaticaParserFDECREMENTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListaIDSContext is an interface to support dynamic dispatch.
type IListaIDSContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFICADOR() []antlr.TerminalNode
	IDENTIFICADOR(i int) antlr.TerminalNode
	AllCOMA() []antlr.TerminalNode
	COMA(i int) antlr.TerminalNode

	// IsListaIDSContext differentiates from other interfaces.
	IsListaIDSContext()
}

type ListaIDSContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaIDSContext() *ListaIDSContext {
	var p = new(ListaIDSContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_listaIDS
	return p
}

func InitEmptyListaIDSContext(p *ListaIDSContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_listaIDS
}

func (*ListaIDSContext) IsListaIDSContext() {}

func NewListaIDSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaIDSContext {
	var p = new(ListaIDSContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_listaIDS

	return p
}

func (s *ListaIDSContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaIDSContext) AllIDENTIFICADOR() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserIDENTIFICADOR)
}

func (s *ListaIDSContext) IDENTIFICADOR(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, i)
}

func (s *ListaIDSContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCOMA)
}

func (s *ListaIDSContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOMA, i)
}

func (s *ListaIDSContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaIDSContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaIDSContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterListaIDS(s)
	}
}

func (s *ListaIDSContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitListaIDS(s)
	}
}

func (s *ListaIDSContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitListaIDS(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) ListaIDS() (localctx IListaIDSContext) {
	localctx = NewListaIDSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, gramaticaParserRULE_listaIDS)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(434)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(439)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserCOMA {
		{
			p.SetState(435)
			p.Match(gramaticaParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(436)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(441)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListaParContext is an interface to support dynamic dispatch.
type IListaParContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParametro() []IParametroContext
	Parametro(i int) IParametroContext
	AllCOMA() []antlr.TerminalNode
	COMA(i int) antlr.TerminalNode

	// IsListaParContext differentiates from other interfaces.
	IsListaParContext()
}

type ListaParContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaParContext() *ListaParContext {
	var p = new(ListaParContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_listaPar
	return p
}

func InitEmptyListaParContext(p *ListaParContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_listaPar
}

func (*ListaParContext) IsListaParContext() {}

func NewListaParContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaParContext {
	var p = new(ListaParContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_listaPar

	return p
}

func (s *ListaParContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaParContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *ListaParContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *ListaParContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCOMA)
}

func (s *ListaParContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOMA, i)
}

func (s *ListaParContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaParContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaParContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterListaPar(s)
	}
}

func (s *ListaParContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitListaPar(s)
	}
}

func (s *ListaParContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitListaPar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) ListaPar() (localctx IListaParContext) {
	localctx = NewListaParContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, gramaticaParserRULE_listaPar)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(442)
		p.Parametro()
	}
	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserCOMA {
		{
			p.SetState(443)
			p.Match(gramaticaParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(444)
			p.Parametro()
		}

		p.SetState(449)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametroContext is an interface to support dynamic dispatch.
type IParametroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFICADOR() antlr.TerminalNode
	Tipos() ITiposContext

	// IsParametroContext differentiates from other interfaces.
	IsParametroContext()
}

type ParametroContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametroContext() *ParametroContext {
	var p = new(ParametroContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_parametro
	return p
}

func InitEmptyParametroContext(p *ParametroContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_parametro
}

func (*ParametroContext) IsParametroContext() {}

func NewParametroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametroContext {
	var p = new(ParametroContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_parametro

	return p
}

func (s *ParametroContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametroContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *ParametroContext) Tipos() ITiposContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITiposContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITiposContext)
}

func (s *ParametroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterParametro(s)
	}
}

func (s *ParametroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitParametro(s)
	}
}

func (s *ParametroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitParametro(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Parametro() (localctx IParametroContext) {
	localctx = NewParametroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, gramaticaParserRULE_parametro)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(450)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(451)
		p.Tipos()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListaExprContext is an interface to support dynamic dispatch.
type IListaExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	AllCOMA() []antlr.TerminalNode
	COMA(i int) antlr.TerminalNode

	// IsListaExprContext differentiates from other interfaces.
	IsListaExprContext()
}

type ListaExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaExprContext() *ListaExprContext {
	var p = new(ListaExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_listaExpr
	return p
}

func InitEmptyListaExprContext(p *ListaExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_listaExpr
}

func (*ListaExprContext) IsListaExprContext() {}

func NewListaExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaExprContext {
	var p = new(ListaExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_listaExpr

	return p
}

func (s *ListaExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaExprContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *ListaExprContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ListaExprContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCOMA)
}

func (s *ListaExprContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOMA, i)
}

func (s *ListaExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterListaExpr(s)
	}
}

func (s *ListaExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitListaExpr(s)
	}
}

func (s *ListaExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitListaExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) ListaExpr() (localctx IListaExprContext) {
	localctx = NewListaExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, gramaticaParserRULE_listaExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(453)
		p.expresion(0)
	}
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserCOMA {
		{
			p.SetState(454)
			p.Match(gramaticaParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(455)
			p.expresion(0)
		}

		p.SetState(460)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListaExprListContext is an interface to support dynamic dispatch.
type IListaExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLLAV1() []antlr.TerminalNode
	LLAV1(i int) antlr.TerminalNode
	AllListaExpr() []IListaExprContext
	ListaExpr(i int) IListaExprContext
	AllLLAV2() []antlr.TerminalNode
	LLAV2(i int) antlr.TerminalNode
	AllCOMA() []antlr.TerminalNode
	COMA(i int) antlr.TerminalNode

	// IsListaExprListContext differentiates from other interfaces.
	IsListaExprListContext()
}

type ListaExprListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaExprListContext() *ListaExprListContext {
	var p = new(ListaExprListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_listaExprList
	return p
}

func InitEmptyListaExprListContext(p *ListaExprListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_listaExprList
}

func (*ListaExprListContext) IsListaExprListContext() {}

func NewListaExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaExprListContext {
	var p = new(ListaExprListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_listaExprList

	return p
}

func (s *ListaExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaExprListContext) AllLLAV1() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserLLAV1)
}

func (s *ListaExprListContext) LLAV1(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAV1, i)
}

func (s *ListaExprListContext) AllListaExpr() []IListaExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IListaExprContext); ok {
			len++
		}
	}

	tst := make([]IListaExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IListaExprContext); ok {
			tst[i] = t.(IListaExprContext)
			i++
		}
	}

	return tst
}

func (s *ListaExprListContext) ListaExpr(i int) IListaExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaExprContext)
}

func (s *ListaExprListContext) AllLLAV2() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserLLAV2)
}

func (s *ListaExprListContext) LLAV2(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAV2, i)
}

func (s *ListaExprListContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserCOMA)
}

func (s *ListaExprListContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOMA, i)
}

func (s *ListaExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterListaExprList(s)
	}
}

func (s *ListaExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitListaExprList(s)
	}
}

func (s *ListaExprListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitListaExprList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) ListaExprList() (localctx IListaExprListContext) {
	localctx = NewListaExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, gramaticaParserRULE_listaExprList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(461)
		p.Match(gramaticaParserLLAV1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(462)
		p.ListaExpr()
	}
	{
		p.SetState(463)
		p.Match(gramaticaParserLLAV2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(471)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(464)
				p.Match(gramaticaParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(465)
				p.Match(gramaticaParserLLAV1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(466)
				p.ListaExpr()
			}
			{
				p.SetState(467)
				p.Match(gramaticaParserLLAV2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(473)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(475)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserCOMA {
		{
			p.SetState(474)
			p.Match(gramaticaParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MENOS() antlr.TerminalNode
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	DIFER() antlr.TerminalNode
	ENTERO() antlr.TerminalNode
	DECIMAL() antlr.TerminalNode
	CADENA() antlr.TerminalNode
	RUNE() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	AccesoElementoSlice() IAccesoElementoSliceContext
	COR1() antlr.TerminalNode
	COR2() antlr.TerminalNode
	ListaExpr() IListaExprContext
	LlamadaFuncionesSinParametro() ILlamadaFuncionesSinParametroContext
	LlamadaFuncionesConParametro() ILlamadaFuncionesConParametroContext
	IDENTIFICADOR() antlr.TerminalNode
	PAR1() antlr.TerminalNode
	PAR2() antlr.TerminalNode
	FnAtoi() IFnAtoiContext
	FnParseToFloat() IFnParseToFloatContext
	FnTypeOf() IFnTypeOfContext
	AccesoStruct() IAccesoStructContext
	FnAppend() IFnAppendContext
	FnIndexOf() IFnIndexOfContext
	FnJoin() IFnJoinContext
	FnLen() IFnLenContext
	MODULO() antlr.TerminalNode
	DIV() antlr.TerminalNode
	POR() antlr.TerminalNode
	MAS() antlr.TerminalNode
	DIFERENTE() antlr.TerminalNode
	IGUALDAD() antlr.TerminalNode
	MENIGUAL() antlr.TerminalNode
	MAYIGUAL() antlr.TerminalNode
	MENOR() antlr.TerminalNode
	MAYOR() antlr.TerminalNode
	OR() antlr.TerminalNode
	AND() antlr.TerminalNode

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_expresion
	return p
}

func InitEmptyExpresionContext(p *ExpresionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_expresion
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) MENOS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserMENOS, 0)
}

func (s *ExpresionContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *ExpresionContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpresionContext) DIFER() antlr.TerminalNode {
	return s.GetToken(gramaticaParserDIFER, 0)
}

func (s *ExpresionContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(gramaticaParserENTERO, 0)
}

func (s *ExpresionContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserDECIMAL, 0)
}

func (s *ExpresionContext) CADENA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCADENA, 0)
}

func (s *ExpresionContext) RUNE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRUNE, 0)
}

func (s *ExpresionContext) TRUE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserTRUE, 0)
}

func (s *ExpresionContext) FALSE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserFALSE, 0)
}

func (s *ExpresionContext) AccesoElementoSlice() IAccesoElementoSliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccesoElementoSliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccesoElementoSliceContext)
}

func (s *ExpresionContext) COR1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOR1, 0)
}

func (s *ExpresionContext) COR2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOR2, 0)
}

func (s *ExpresionContext) ListaExpr() IListaExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaExprContext)
}

func (s *ExpresionContext) LlamadaFuncionesSinParametro() ILlamadaFuncionesSinParametroContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadaFuncionesSinParametroContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadaFuncionesSinParametroContext)
}

func (s *ExpresionContext) LlamadaFuncionesConParametro() ILlamadaFuncionesConParametroContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadaFuncionesConParametroContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadaFuncionesConParametroContext)
}

func (s *ExpresionContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *ExpresionContext) PAR1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR1, 0)
}

func (s *ExpresionContext) PAR2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR2, 0)
}

func (s *ExpresionContext) FnAtoi() IFnAtoiContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnAtoiContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnAtoiContext)
}

func (s *ExpresionContext) FnParseToFloat() IFnParseToFloatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnParseToFloatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnParseToFloatContext)
}

func (s *ExpresionContext) FnTypeOf() IFnTypeOfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnTypeOfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnTypeOfContext)
}

func (s *ExpresionContext) AccesoStruct() IAccesoStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccesoStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccesoStructContext)
}

func (s *ExpresionContext) FnAppend() IFnAppendContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnAppendContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnAppendContext)
}

func (s *ExpresionContext) FnIndexOf() IFnIndexOfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnIndexOfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnIndexOfContext)
}

func (s *ExpresionContext) FnJoin() IFnJoinContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnJoinContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnJoinContext)
}

func (s *ExpresionContext) FnLen() IFnLenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnLenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnLenContext)
}

func (s *ExpresionContext) MODULO() antlr.TerminalNode {
	return s.GetToken(gramaticaParserMODULO, 0)
}

func (s *ExpresionContext) DIV() antlr.TerminalNode {
	return s.GetToken(gramaticaParserDIV, 0)
}

func (s *ExpresionContext) POR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPOR, 0)
}

func (s *ExpresionContext) MAS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserMAS, 0)
}

func (s *ExpresionContext) DIFERENTE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserDIFERENTE, 0)
}

func (s *ExpresionContext) IGUALDAD() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIGUALDAD, 0)
}

func (s *ExpresionContext) MENIGUAL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserMENIGUAL, 0)
}

func (s *ExpresionContext) MAYIGUAL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserMAYIGUAL, 0)
}

func (s *ExpresionContext) MENOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserMENOR, 0)
}

func (s *ExpresionContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserMAYOR, 0)
}

func (s *ExpresionContext) OR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserOR, 0)
}

func (s *ExpresionContext) AND() antlr.TerminalNode {
	return s.GetToken(gramaticaParserAND, 0)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExpresion(s)
	}
}

func (s *ExpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExpresion(s)
	}
}

func (s *ExpresionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitExpresion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *gramaticaParser) expresion(_p int) (localctx IExpresionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 94
	p.EnterRecursionRule(localctx, 94, gramaticaParserRULE_expresion, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(509)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(478)
			p.Match(gramaticaParserMENOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(479)
			p.expresion(35)
		}

	case 2:
		{
			p.SetState(480)
			p.Match(gramaticaParserDIFER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(481)
			p.expresion(34)
		}

	case 3:
		{
			p.SetState(482)
			p.Match(gramaticaParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(483)
			p.Match(gramaticaParserDECIMAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		{
			p.SetState(484)
			p.Match(gramaticaParserCADENA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		{
			p.SetState(485)
			p.Match(gramaticaParserRUNE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		{
			p.SetState(486)
			p.Match(gramaticaParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		{
			p.SetState(487)
			p.Match(gramaticaParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		{
			p.SetState(488)
			p.AccesoElementoSlice()
		}

	case 10:
		{
			p.SetState(489)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(491)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-23)) & ^0x3f) == 0 && ((int64(1)<<(_la-23))&8525512322623) != 0 {
			{
				p.SetState(490)
				p.ListaExpr()
			}

		}
		{
			p.SetState(493)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		{
			p.SetState(494)
			p.LlamadaFuncionesSinParametro()
		}

	case 12:
		{
			p.SetState(495)
			p.LlamadaFuncionesConParametro()
		}

	case 13:
		{
			p.SetState(496)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		{
			p.SetState(497)
			p.Match(gramaticaParserPAR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(498)
			p.expresion(0)
		}
		{
			p.SetState(499)
			p.Match(gramaticaParserPAR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		{
			p.SetState(501)
			p.FnAtoi()
		}

	case 16:
		{
			p.SetState(502)
			p.FnParseToFloat()
		}

	case 17:
		{
			p.SetState(503)
			p.FnTypeOf()
		}

	case 18:
		{
			p.SetState(504)
			p.AccesoStruct()
		}

	case 19:
		{
			p.SetState(505)
			p.FnAppend()
		}

	case 20:
		{
			p.SetState(506)
			p.FnIndexOf()
		}

	case 21:
		{
			p.SetState(507)
			p.FnJoin()
		}

	case 22:
		{
			p.SetState(508)
			p.FnLen()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(552)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(550)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(511)

				if !(p.Precpred(p.GetParserRuleContext(), 33)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 33)", ""))
					goto errorExit
				}
				{
					p.SetState(512)
					p.Match(gramaticaParserMODULO)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(513)
					p.expresion(34)
				}

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(514)

				if !(p.Precpred(p.GetParserRuleContext(), 32)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 32)", ""))
					goto errorExit
				}
				{
					p.SetState(515)
					p.Match(gramaticaParserDIV)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(516)
					p.expresion(33)
				}

			case 3:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(517)

				if !(p.Precpred(p.GetParserRuleContext(), 31)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 31)", ""))
					goto errorExit
				}
				{
					p.SetState(518)
					p.Match(gramaticaParserPOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(519)
					p.expresion(32)
				}

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(520)

				if !(p.Precpred(p.GetParserRuleContext(), 30)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 30)", ""))
					goto errorExit
				}
				{
					p.SetState(521)
					p.Match(gramaticaParserMENOS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(522)
					p.expresion(31)
				}

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(523)

				if !(p.Precpred(p.GetParserRuleContext(), 29)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 29)", ""))
					goto errorExit
				}
				{
					p.SetState(524)
					p.Match(gramaticaParserMAS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(525)
					p.expresion(30)
				}

			case 6:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(526)

				if !(p.Precpred(p.GetParserRuleContext(), 28)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 28)", ""))
					goto errorExit
				}
				{
					p.SetState(527)
					p.Match(gramaticaParserDIFERENTE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(528)
					p.expresion(29)
				}

			case 7:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(529)

				if !(p.Precpred(p.GetParserRuleContext(), 27)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 27)", ""))
					goto errorExit
				}
				{
					p.SetState(530)
					p.Match(gramaticaParserIGUALDAD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(531)
					p.expresion(28)
				}

			case 8:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(532)

				if !(p.Precpred(p.GetParserRuleContext(), 26)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 26)", ""))
					goto errorExit
				}
				{
					p.SetState(533)
					p.Match(gramaticaParserMENIGUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(534)
					p.expresion(27)
				}

			case 9:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(535)

				if !(p.Precpred(p.GetParserRuleContext(), 25)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 25)", ""))
					goto errorExit
				}
				{
					p.SetState(536)
					p.Match(gramaticaParserMAYIGUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(537)
					p.expresion(26)
				}

			case 10:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(538)

				if !(p.Precpred(p.GetParserRuleContext(), 24)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 24)", ""))
					goto errorExit
				}
				{
					p.SetState(539)
					p.Match(gramaticaParserMENOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(540)
					p.expresion(25)
				}

			case 11:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(541)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
					goto errorExit
				}
				{
					p.SetState(542)
					p.Match(gramaticaParserMAYOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(543)
					p.expresion(24)
				}

			case 12:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(544)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(545)
					p.Match(gramaticaParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(546)
					p.expresion(23)
				}

			case 13:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(547)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(548)
					p.Match(gramaticaParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(549)
					p.expresion(22)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(554)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISifContext is an interface to support dynamic dispatch.
type ISifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	AllBloque() []IBloqueContext
	Bloque(i int) IBloqueContext
	PAR1() antlr.TerminalNode
	Expresion() IExpresionContext
	PAR2() antlr.TerminalNode
	AllElseifPart() []IElseifPartContext
	ElseifPart(i int) IElseifPartContext
	ELSE() antlr.TerminalNode

	// IsSifContext differentiates from other interfaces.
	IsSifContext()
}

type SifContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySifContext() *SifContext {
	var p = new(SifContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sif
	return p
}

func InitEmptySifContext(p *SifContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sif
}

func (*SifContext) IsSifContext() {}

func NewSifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SifContext {
	var p = new(SifContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_sif

	return p
}

func (s *SifContext) GetParser() antlr.Parser { return s.parser }

func (s *SifContext) IF() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIF, 0)
}

func (s *SifContext) AllBloque() []IBloqueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBloqueContext); ok {
			len++
		}
	}

	tst := make([]IBloqueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBloqueContext); ok {
			tst[i] = t.(IBloqueContext)
			i++
		}
	}

	return tst
}

func (s *SifContext) Bloque(i int) IBloqueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *SifContext) PAR1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR1, 0)
}

func (s *SifContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *SifContext) PAR2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR2, 0)
}

func (s *SifContext) AllElseifPart() []IElseifPartContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseifPartContext); ok {
			len++
		}
	}

	tst := make([]IElseifPartContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseifPartContext); ok {
			tst[i] = t.(IElseifPartContext)
			i++
		}
	}

	return tst
}

func (s *SifContext) ElseifPart(i int) IElseifPartContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseifPartContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseifPartContext)
}

func (s *SifContext) ELSE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserELSE, 0)
}

func (s *SifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSif(s)
	}
}

func (s *SifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSif(s)
	}
}

func (s *SifContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSif(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Sif() (localctx ISifContext) {
	localctx = NewSifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, gramaticaParserRULE_sif)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(555)
		p.Match(gramaticaParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(561)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(556)
			p.Match(gramaticaParserPAR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(557)
			p.expresion(0)
		}
		{
			p.SetState(558)
			p.Match(gramaticaParserPAR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(560)
			p.expresion(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(563)
		p.Bloque()
	}
	p.SetState(567)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(564)
				p.ElseifPart()
			}

		}
		p.SetState(569)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(572)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserELSE {
		{
			p.SetState(570)
			p.Match(gramaticaParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(571)
			p.Bloque()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseifPartContext is an interface to support dynamic dispatch.
type IElseifPartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	IF() antlr.TerminalNode
	Bloque() IBloqueContext
	PAR1() antlr.TerminalNode
	Expresion() IExpresionContext
	PAR2() antlr.TerminalNode

	// IsElseifPartContext differentiates from other interfaces.
	IsElseifPartContext()
}

type ElseifPartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseifPartContext() *ElseifPartContext {
	var p = new(ElseifPartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_elseifPart
	return p
}

func InitEmptyElseifPartContext(p *ElseifPartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_elseifPart
}

func (*ElseifPartContext) IsElseifPartContext() {}

func NewElseifPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseifPartContext {
	var p = new(ElseifPartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_elseifPart

	return p
}

func (s *ElseifPartContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseifPartContext) ELSE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserELSE, 0)
}

func (s *ElseifPartContext) IF() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIF, 0)
}

func (s *ElseifPartContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *ElseifPartContext) PAR1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR1, 0)
}

func (s *ElseifPartContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ElseifPartContext) PAR2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPAR2, 0)
}

func (s *ElseifPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseifPartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseifPartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterElseifPart(s)
	}
}

func (s *ElseifPartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitElseifPart(s)
	}
}

func (s *ElseifPartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitElseifPart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) ElseifPart() (localctx IElseifPartContext) {
	localctx = NewElseifPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, gramaticaParserRULE_elseifPart)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(574)
		p.Match(gramaticaParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(575)
		p.Match(gramaticaParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(581)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(576)
			p.Match(gramaticaParserPAR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(577)
			p.expresion(0)
		}
		{
			p.SetState(578)
			p.Match(gramaticaParserPAR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(580)
			p.expresion(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(583)
		p.Bloque()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBloqueContext is an interface to support dynamic dispatch.
type IBloqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LLAV1() antlr.TerminalNode
	LLAV2() antlr.TerminalNode
	AllInstrucciones() []IInstruccionesContext
	Instrucciones(i int) IInstruccionesContext

	// IsBloqueContext differentiates from other interfaces.
	IsBloqueContext()
}

type BloqueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloqueContext() *BloqueContext {
	var p = new(BloqueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_bloque
	return p
}

func InitEmptyBloqueContext(p *BloqueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_bloque
}

func (*BloqueContext) IsBloqueContext() {}

func NewBloqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueContext {
	var p = new(BloqueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_bloque

	return p
}

func (s *BloqueContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueContext) LLAV1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAV1, 0)
}

func (s *BloqueContext) LLAV2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAV2, 0)
}

func (s *BloqueContext) AllInstrucciones() []IInstruccionesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstruccionesContext); ok {
			len++
		}
	}

	tst := make([]IInstruccionesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstruccionesContext); ok {
			tst[i] = t.(IInstruccionesContext)
			i++
		}
	}

	return tst
}

func (s *BloqueContext) Instrucciones(i int) IInstruccionesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *BloqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloqueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterBloque(s)
	}
}

func (s *BloqueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitBloque(s)
	}
}

func (s *BloqueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitBloque(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Bloque() (localctx IBloqueContext) {
	localctx = NewBloqueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, gramaticaParserRULE_bloque)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(585)
		p.Match(gramaticaParserLLAV1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(589)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2305843009750813424) != 0 {
		{
			p.SetState(586)
			p.Instrucciones()
		}

		p.SetState(591)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(592)
		p.Match(gramaticaParserLLAV2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISforContext is an interface to support dynamic dispatch.
type ISforContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSforContext differentiates from other interfaces.
	IsSforContext()
}

type SforContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySforContext() *SforContext {
	var p = new(SforContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sfor
	return p
}

func InitEmptySforContext(p *SforContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sfor
}

func (*SforContext) IsSforContext() {}

func NewSforContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SforContext {
	var p = new(SforContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_sfor

	return p
}

func (s *SforContext) GetParser() antlr.Parser { return s.parser }

func (s *SforContext) CopyAll(ctx *SforContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SforContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SforContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForClasicoContext struct {
	SforContext
}

func NewForClasicoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForClasicoContext {
	var p = new(ForClasicoContext)

	InitEmptySforContext(&p.SforContext)
	p.parser = parser
	p.CopyAll(ctx.(*SforContext))

	return p
}

func (s *ForClasicoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForClasicoContext) FOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserFOR, 0)
}

func (s *ForClasicoContext) AllAsignacion() []IAsignacionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAsignacionContext); ok {
			len++
		}
	}

	tst := make([]IAsignacionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAsignacionContext); ok {
			tst[i] = t.(IAsignacionContext)
			i++
		}
	}

	return tst
}

func (s *ForClasicoContext) Asignacion(i int) IAsignacionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *ForClasicoContext) AllPTCOMA() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserPTCOMA)
}

func (s *ForClasicoContext) PTCOMA(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserPTCOMA, i)
}

func (s *ForClasicoContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ForClasicoContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *ForClasicoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterForClasico(s)
	}
}

func (s *ForClasicoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitForClasico(s)
	}
}

func (s *ForClasicoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitForClasico(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForCondicionalContext struct {
	SforContext
}

func NewForCondicionalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForCondicionalContext {
	var p = new(ForCondicionalContext)

	InitEmptySforContext(&p.SforContext)
	p.parser = parser
	p.CopyAll(ctx.(*SforContext))

	return p
}

func (s *ForCondicionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForCondicionalContext) FOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserFOR, 0)
}

func (s *ForCondicionalContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ForCondicionalContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *ForCondicionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterForCondicional(s)
	}
}

func (s *ForCondicionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitForCondicional(s)
	}
}

func (s *ForCondicionalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitForCondicional(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForRangeContext struct {
	SforContext
}

func NewForRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForRangeContext {
	var p = new(ForRangeContext)

	InitEmptySforContext(&p.SforContext)
	p.parser = parser
	p.CopyAll(ctx.(*SforContext))

	return p
}

func (s *ForRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForRangeContext) FOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserFOR, 0)
}

func (s *ForRangeContext) AllIDENTIFICADOR() []antlr.TerminalNode {
	return s.GetTokens(gramaticaParserIDENTIFICADOR)
}

func (s *ForRangeContext) IDENTIFICADOR(i int) antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, i)
}

func (s *ForRangeContext) COMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOMA, 0)
}

func (s *ForRangeContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIGUAL, 0)
}

func (s *ForRangeContext) RANGE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRANGE, 0)
}

func (s *ForRangeContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ForRangeContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *ForRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterForRange(s)
	}
}

func (s *ForRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitForRange(s)
	}
}

func (s *ForRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitForRange(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Sfor() (localctx ISforContext) {
	localctx = NewSforContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, gramaticaParserRULE_sfor)
	p.SetState(615)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForCondicionalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(594)
			p.Match(gramaticaParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(595)
			p.expresion(0)
		}
		{
			p.SetState(596)
			p.Bloque()
		}

	case 2:
		localctx = NewForClasicoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(598)
			p.Match(gramaticaParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(599)
			p.Asignacion()
		}
		{
			p.SetState(600)
			p.Match(gramaticaParserPTCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(601)
			p.expresion(0)
		}
		{
			p.SetState(602)
			p.Match(gramaticaParserPTCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(603)
			p.Asignacion()
		}
		{
			p.SetState(604)
			p.Bloque()
		}

	case 3:
		localctx = NewForRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(606)
			p.Match(gramaticaParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(607)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(608)
			p.Match(gramaticaParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(609)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(610)
			p.Match(gramaticaParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(611)
			p.Match(gramaticaParserRANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(612)
			p.expresion(0)
		}
		{
			p.SetState(613)
			p.Bloque()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISSwitchContext is an interface to support dynamic dispatch.
type ISSwitchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expresion() IExpresionContext
	LLAV1() antlr.TerminalNode
	LLAV2() antlr.TerminalNode
	AllCaseBlock() []ICaseBlockContext
	CaseBlock(i int) ICaseBlockContext
	DefaultBlock() IDefaultBlockContext

	// IsSSwitchContext differentiates from other interfaces.
	IsSSwitchContext()
}

type SSwitchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySSwitchContext() *SSwitchContext {
	var p = new(SSwitchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sSwitch
	return p
}

func InitEmptySSwitchContext(p *SSwitchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sSwitch
}

func (*SSwitchContext) IsSSwitchContext() {}

func NewSSwitchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SSwitchContext {
	var p = new(SSwitchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_sSwitch

	return p
}

func (s *SSwitchContext) GetParser() antlr.Parser { return s.parser }

func (s *SSwitchContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSWITCH, 0)
}

func (s *SSwitchContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *SSwitchContext) LLAV1() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAV1, 0)
}

func (s *SSwitchContext) LLAV2() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLLAV2, 0)
}

func (s *SSwitchContext) AllCaseBlock() []ICaseBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICaseBlockContext); ok {
			len++
		}
	}

	tst := make([]ICaseBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICaseBlockContext); ok {
			tst[i] = t.(ICaseBlockContext)
			i++
		}
	}

	return tst
}

func (s *SSwitchContext) CaseBlock(i int) ICaseBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseBlockContext)
}

func (s *SSwitchContext) DefaultBlock() IDefaultBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultBlockContext)
}

func (s *SSwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SSwitchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SSwitchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSSwitch(s)
	}
}

func (s *SSwitchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSSwitch(s)
	}
}

func (s *SSwitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitSSwitch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) SSwitch() (localctx ISSwitchContext) {
	localctx = NewSSwitchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, gramaticaParserRULE_sSwitch)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(617)
		p.Match(gramaticaParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(618)
		p.expresion(0)
	}
	{
		p.SetState(619)
		p.Match(gramaticaParserLLAV1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(623)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserCASE {
		{
			p.SetState(620)
			p.CaseBlock()
		}

		p.SetState(625)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(627)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserDEFAULT {
		{
			p.SetState(626)
			p.DefaultBlock()
		}

	}
	{
		p.SetState(629)
		p.Match(gramaticaParserLLAV2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreak_Context is an interface to support dynamic dispatch.
type IBreak_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode

	// IsBreak_Context differentiates from other interfaces.
	IsBreak_Context()
}

type Break_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreak_Context() *Break_Context {
	var p = new(Break_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_break_
	return p
}

func InitEmptyBreak_Context(p *Break_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_break_
}

func (*Break_Context) IsBreak_Context() {}

func NewBreak_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Break_Context {
	var p = new(Break_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_break_

	return p
}

func (s *Break_Context) GetParser() antlr.Parser { return s.parser }

func (s *Break_Context) BREAK() antlr.TerminalNode {
	return s.GetToken(gramaticaParserBREAK, 0)
}

func (s *Break_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Break_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Break_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterBreak_(s)
	}
}

func (s *Break_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitBreak_(s)
	}
}

func (s *Break_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitBreak_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Break_() (localctx IBreak_Context) {
	localctx = NewBreak_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, gramaticaParserRULE_break_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(631)
		p.Match(gramaticaParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinue_Context is an interface to support dynamic dispatch.
type IContinue_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinue_Context differentiates from other interfaces.
	IsContinue_Context()
}

type Continue_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinue_Context() *Continue_Context {
	var p = new(Continue_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_continue_
	return p
}

func InitEmptyContinue_Context(p *Continue_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_continue_
}

func (*Continue_Context) IsContinue_Context() {}

func NewContinue_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Continue_Context {
	var p = new(Continue_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_continue_

	return p
}

func (s *Continue_Context) GetParser() antlr.Parser { return s.parser }

func (s *Continue_Context) CONTINUE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCONTINUE, 0)
}

func (s *Continue_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Continue_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Continue_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterContinue_(s)
	}
}

func (s *Continue_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitContinue_(s)
	}
}

func (s *Continue_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitContinue_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Continue_() (localctx IContinue_Context) {
	localctx = NewContinue_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, gramaticaParserRULE_continue_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(633)
		p.Match(gramaticaParserCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICaseBlockContext is an interface to support dynamic dispatch.
type ICaseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	Expresion() IExpresionContext
	DOSPTS() antlr.TerminalNode
	Instrucciones() IInstruccionesContext

	// IsCaseBlockContext differentiates from other interfaces.
	IsCaseBlockContext()
}

type CaseBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseBlockContext() *CaseBlockContext {
	var p = new(CaseBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_caseBlock
	return p
}

func InitEmptyCaseBlockContext(p *CaseBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_caseBlock
}

func (*CaseBlockContext) IsCaseBlockContext() {}

func NewCaseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseBlockContext {
	var p = new(CaseBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_caseBlock

	return p
}

func (s *CaseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseBlockContext) CASE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCASE, 0)
}

func (s *CaseBlockContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *CaseBlockContext) DOSPTS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserDOSPTS, 0)
}

func (s *CaseBlockContext) Instrucciones() IInstruccionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *CaseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterCaseBlock(s)
	}
}

func (s *CaseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitCaseBlock(s)
	}
}

func (s *CaseBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitCaseBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) CaseBlock() (localctx ICaseBlockContext) {
	localctx = NewCaseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, gramaticaParserRULE_caseBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(635)
		p.Match(gramaticaParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(636)
		p.expresion(0)
	}
	{
		p.SetState(637)
		p.Match(gramaticaParserDOSPTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(638)
		p.Instrucciones()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefaultBlockContext is an interface to support dynamic dispatch.
type IDefaultBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	DOSPTS() antlr.TerminalNode
	Instrucciones() IInstruccionesContext

	// IsDefaultBlockContext differentiates from other interfaces.
	IsDefaultBlockContext()
}

type DefaultBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultBlockContext() *DefaultBlockContext {
	var p = new(DefaultBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_defaultBlock
	return p
}

func InitEmptyDefaultBlockContext(p *DefaultBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_defaultBlock
}

func (*DefaultBlockContext) IsDefaultBlockContext() {}

func NewDefaultBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultBlockContext {
	var p = new(DefaultBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_defaultBlock

	return p
}

func (s *DefaultBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultBlockContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserDEFAULT, 0)
}

func (s *DefaultBlockContext) DOSPTS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserDOSPTS, 0)
}

func (s *DefaultBlockContext) Instrucciones() IInstruccionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *DefaultBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterDefaultBlock(s)
	}
}

func (s *DefaultBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitDefaultBlock(s)
	}
}

func (s *DefaultBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitDefaultBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) DefaultBlock() (localctx IDefaultBlockContext) {
	localctx = NewDefaultBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, gramaticaParserRULE_defaultBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(640)
		p.Match(gramaticaParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(641)
		p.Match(gramaticaParserDOSPTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(642)
		p.Instrucciones()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITiposContext is an interface to support dynamic dispatch.
type ITiposContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	IDENTIFICADOR() antlr.TerminalNode

	// IsTiposContext differentiates from other interfaces.
	IsTiposContext()
}

type TiposContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTiposContext() *TiposContext {
	var p = new(TiposContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_tipos
	return p
}

func InitEmptyTiposContext(p *TiposContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_tipos
}

func (*TiposContext) IsTiposContext() {}

func NewTiposContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TiposContext {
	var p = new(TiposContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_tipos

	return p
}

func (s *TiposContext) GetParser() antlr.Parser { return s.parser }

func (s *TiposContext) INT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserINT, 0)
}

func (s *TiposContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserFLOAT, 0)
}

func (s *TiposContext) STRING() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSTRING, 0)
}

func (s *TiposContext) BOOL() antlr.TerminalNode {
	return s.GetToken(gramaticaParserBOOL, 0)
}

func (s *TiposContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserIDENTIFICADOR, 0)
}

func (s *TiposContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TiposContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TiposContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterTipos(s)
	}
}

func (s *TiposContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitTipos(s)
	}
}

func (s *TiposContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gramaticaVisitor:
		return t.VisitTipos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gramaticaParser) Tipos() (localctx ITiposContext) {
	localctx = NewTiposContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, gramaticaParserRULE_tipos)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(644)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2305843009217626112) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *gramaticaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 47:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *gramaticaParser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 33)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 32)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 31)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 30)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 29)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 28)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 27)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 26)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 25)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 24)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 21)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
