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
		"", "'+='", "'-='", "'print'", "'println'", "'fn'", "'mut'", "'slice'",
		"'if'", "'else'", "'switch'", "'case'", "'default'", "'for'", "'break'",
		"'continue'", "'return'", "'int'", "'float64'", "'string'", "'bool'",
		"'nil'", "'true'", "'false'", "'indexOf'", "'join'", "'len'", "'append'",
		"'struct'", "'func'", "'range'", "'Atoi'", "'parseFloat'", "'typeOf'",
		"'+'", "'-'", "'*'", "'/'", "'%'", "'('", "')'", "'{'", "'}'", "'['",
		"']'", "'!='", "'=='", "'<='", "'>='", "", "'<'", "'>'", "'||'", "'&&'",
		"'!'", "':'", "';'", "','", "'++'", "'--'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "PRINT", "PRINTLN", "FUNCIONES", "MUT", "SLICE", "IF", "ELSE",
		"SWITCH", "CASE", "DEFAULT", "FOR", "BREAK", "CONTINUE", "RETURN", "INT",
		"FLOAT", "STRING", "BOOL", "NIL", "TRUE", "FALSE", "INDEXOF", "JOIN",
		"LEN", "APPEND", "STRUCT", "FUNC", "RANGE", "ATOI", "PARSEFLOAT", "TYPEOF",
		"MAS", "MENOS", "POR", "DIV", "MODULO", "PAR1", "PAR2", "LLAV1", "LLAV2",
		"COR1", "COR2", "DIFERENTE", "IGUALDAD", "MENIGUAL", "MAYIGUAL", "IGUAL",
		"MENOR", "MAYOR", "OR", "AND", "DIFER", "DOSPTS", "PTCOMA", "COMA",
		"FINCREMENTO", "FDECREMENTO", "IDENTIFICADOR", "ENTERO", "DECIMAL",
		"CADENA", "RUNE", "COMENTARIO", "MULTICOMENTARIO", "WS",
	}
	staticData.RuleNames = []string{
		"init", "instrucciones", "instruccion", "print", "println", "declaracionMultiple",
		"declaracionMultipleSimple", "declaracionMultipleSinTipo", "asignacionMultiple",
		"bloque", "llamadaFuncionesSinParametro", "llamadaFuncionesConParametro",
		"fnSinParametro", "fnConParametro", "sliceDef", "sliceLiteral", "accesoElementoSlice",
		"modificacionElementoSlice", "fnIndexOf", "fnJoin", "fnLen", "fnAppend",
		"declaracionSliceVacio", "tipoRetorno", "retorno", "fnTypeOf", "fnAtoi",
		"fnParseToFloat", "asigIncre", "asigDecre", "asignacion", "listaIDS",
		"listaPar", "parametro", "listaExpr", "listaExprList", "expresion",
		"sif", "elseifPart", "sfor", "sSwitch", "break_", "continue_", "caseBlock",
		"defaultBlock", "tipos",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 67, 529, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 1, 0, 5, 0, 94, 8,
		0, 10, 0, 12, 0, 97, 9, 0, 1, 1, 4, 1, 100, 8, 1, 11, 1, 12, 1, 101, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 127,
		8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 161, 8, 9,
		10, 9, 12, 9, 164, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 182,
		8, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 192,
		8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 3, 15, 216, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 231, 8, 16,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 250, 8, 17, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3,
		22, 284, 8, 22, 1, 23, 1, 23, 1, 24, 1, 24, 3, 24, 290, 8, 24, 1, 24, 3,
		24, 293, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 3, 30, 325, 8, 30, 1, 31, 1, 31, 1, 31, 5, 31, 330, 8, 31, 10, 31,
		12, 31, 333, 9, 31, 1, 32, 1, 32, 1, 32, 5, 32, 338, 8, 32, 10, 32, 12,
		32, 341, 9, 32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 5, 34, 349, 8,
		34, 10, 34, 12, 34, 352, 9, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 5, 35, 362, 8, 35, 10, 35, 12, 35, 365, 9, 35, 1, 35, 3,
		35, 368, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 384, 8, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 3, 36, 401, 8, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 5, 36, 442, 8, 36, 10, 36, 12, 36, 445, 9, 36, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 453, 8, 37, 1, 37, 1, 37,
		5, 37, 457, 8, 37, 10, 37, 12, 37, 460, 9, 37, 1, 37, 1, 37, 3, 37, 464,
		8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 473, 8,
		38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39,
		1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 3, 39, 498, 8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 5, 40,
		504, 8, 40, 10, 40, 12, 40, 507, 9, 40, 1, 40, 3, 40, 510, 8, 40, 1, 40,
		1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1,
		44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 0, 1, 72, 46, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
		80, 82, 84, 86, 88, 90, 0, 2, 3, 0, 17, 20, 60, 60, 64, 64, 2, 0, 17, 20,
		60, 60, 565, 0, 95, 1, 0, 0, 0, 2, 99, 1, 0, 0, 0, 4, 126, 1, 0, 0, 0,
		6, 128, 1, 0, 0, 0, 8, 133, 1, 0, 0, 0, 10, 138, 1, 0, 0, 0, 12, 144, 1,
		0, 0, 0, 14, 148, 1, 0, 0, 0, 16, 153, 1, 0, 0, 0, 18, 157, 1, 0, 0, 0,
		20, 167, 1, 0, 0, 0, 22, 171, 1, 0, 0, 0, 24, 176, 1, 0, 0, 0, 26, 185,
		1, 0, 0, 0, 28, 195, 1, 0, 0, 0, 30, 215, 1, 0, 0, 0, 32, 230, 1, 0, 0,
		0, 34, 249, 1, 0, 0, 0, 36, 251, 1, 0, 0, 0, 38, 256, 1, 0, 0, 0, 40, 261,
		1, 0, 0, 0, 42, 266, 1, 0, 0, 0, 44, 283, 1, 0, 0, 0, 46, 285, 1, 0, 0,
		0, 48, 287, 1, 0, 0, 0, 50, 294, 1, 0, 0, 0, 52, 299, 1, 0, 0, 0, 54, 304,
		1, 0, 0, 0, 56, 309, 1, 0, 0, 0, 58, 313, 1, 0, 0, 0, 60, 324, 1, 0, 0,
		0, 62, 326, 1, 0, 0, 0, 64, 334, 1, 0, 0, 0, 66, 342, 1, 0, 0, 0, 68, 345,
		1, 0, 0, 0, 70, 353, 1, 0, 0, 0, 72, 400, 1, 0, 0, 0, 74, 446, 1, 0, 0,
		0, 76, 465, 1, 0, 0, 0, 78, 497, 1, 0, 0, 0, 80, 499, 1, 0, 0, 0, 82, 513,
		1, 0, 0, 0, 84, 515, 1, 0, 0, 0, 86, 517, 1, 0, 0, 0, 88, 522, 1, 0, 0,
		0, 90, 526, 1, 0, 0, 0, 92, 94, 3, 2, 1, 0, 93, 92, 1, 0, 0, 0, 94, 97,
		1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 1, 1, 0, 0, 0,
		97, 95, 1, 0, 0, 0, 98, 100, 3, 4, 2, 0, 99, 98, 1, 0, 0, 0, 100, 101,
		1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 3, 1, 0, 0,
		0, 103, 127, 3, 6, 3, 0, 104, 127, 3, 8, 4, 0, 105, 127, 3, 10, 5, 0, 106,
		127, 3, 12, 6, 0, 107, 127, 3, 14, 7, 0, 108, 127, 3, 16, 8, 0, 109, 127,
		3, 24, 12, 0, 110, 127, 3, 26, 13, 0, 111, 127, 3, 20, 10, 0, 112, 127,
		3, 22, 11, 0, 113, 127, 3, 48, 24, 0, 114, 127, 3, 56, 28, 0, 115, 127,
		3, 58, 29, 0, 116, 127, 3, 60, 30, 0, 117, 127, 3, 74, 37, 0, 118, 127,
		3, 78, 39, 0, 119, 127, 3, 80, 40, 0, 120, 127, 3, 82, 41, 0, 121, 127,
		3, 84, 42, 0, 122, 127, 3, 28, 14, 0, 123, 127, 3, 44, 22, 0, 124, 127,
		3, 34, 17, 0, 125, 127, 3, 18, 9, 0, 126, 103, 1, 0, 0, 0, 126, 104, 1,
		0, 0, 0, 126, 105, 1, 0, 0, 0, 126, 106, 1, 0, 0, 0, 126, 107, 1, 0, 0,
		0, 126, 108, 1, 0, 0, 0, 126, 109, 1, 0, 0, 0, 126, 110, 1, 0, 0, 0, 126,
		111, 1, 0, 0, 0, 126, 112, 1, 0, 0, 0, 126, 113, 1, 0, 0, 0, 126, 114,
		1, 0, 0, 0, 126, 115, 1, 0, 0, 0, 126, 116, 1, 0, 0, 0, 126, 117, 1, 0,
		0, 0, 126, 118, 1, 0, 0, 0, 126, 119, 1, 0, 0, 0, 126, 120, 1, 0, 0, 0,
		126, 121, 1, 0, 0, 0, 126, 122, 1, 0, 0, 0, 126, 123, 1, 0, 0, 0, 126,
		124, 1, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127, 5, 1, 0, 0, 0, 128, 129, 5,
		3, 0, 0, 129, 130, 5, 39, 0, 0, 130, 131, 3, 68, 34, 0, 131, 132, 5, 40,
		0, 0, 132, 7, 1, 0, 0, 0, 133, 134, 5, 4, 0, 0, 134, 135, 5, 39, 0, 0,
		135, 136, 3, 68, 34, 0, 136, 137, 5, 40, 0, 0, 137, 9, 1, 0, 0, 0, 138,
		139, 5, 6, 0, 0, 139, 140, 3, 62, 31, 0, 140, 141, 3, 90, 45, 0, 141, 142,
		5, 49, 0, 0, 142, 143, 3, 68, 34, 0, 143, 11, 1, 0, 0, 0, 144, 145, 5,
		6, 0, 0, 145, 146, 3, 62, 31, 0, 146, 147, 3, 90, 45, 0, 147, 13, 1, 0,
		0, 0, 148, 149, 5, 6, 0, 0, 149, 150, 3, 62, 31, 0, 150, 151, 5, 49, 0,
		0, 151, 152, 3, 68, 34, 0, 152, 15, 1, 0, 0, 0, 153, 154, 3, 62, 31, 0,
		154, 155, 5, 49, 0, 0, 155, 156, 3, 68, 34, 0, 156, 17, 1, 0, 0, 0, 157,
		162, 5, 41, 0, 0, 158, 161, 3, 4, 2, 0, 159, 161, 3, 72, 36, 0, 160, 158,
		1, 0, 0, 0, 160, 159, 1, 0, 0, 0, 161, 164, 1, 0, 0, 0, 162, 160, 1, 0,
		0, 0, 162, 163, 1, 0, 0, 0, 163, 165, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0,
		165, 166, 5, 42, 0, 0, 166, 19, 1, 0, 0, 0, 167, 168, 5, 60, 0, 0, 168,
		169, 5, 39, 0, 0, 169, 170, 5, 40, 0, 0, 170, 21, 1, 0, 0, 0, 171, 172,
		5, 60, 0, 0, 172, 173, 5, 39, 0, 0, 173, 174, 3, 68, 34, 0, 174, 175, 5,
		40, 0, 0, 175, 23, 1, 0, 0, 0, 176, 177, 5, 5, 0, 0, 177, 178, 5, 60, 0,
		0, 178, 179, 5, 39, 0, 0, 179, 181, 5, 40, 0, 0, 180, 182, 3, 46, 23, 0,
		181, 180, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183,
		184, 3, 18, 9, 0, 184, 25, 1, 0, 0, 0, 185, 186, 5, 5, 0, 0, 186, 187,
		5, 60, 0, 0, 187, 188, 5, 39, 0, 0, 188, 189, 3, 64, 32, 0, 189, 191, 5,
		40, 0, 0, 190, 192, 3, 46, 23, 0, 191, 190, 1, 0, 0, 0, 191, 192, 1, 0,
		0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 3, 18, 9, 0, 194, 27, 1, 0, 0, 0,
		195, 196, 5, 60, 0, 0, 196, 197, 5, 49, 0, 0, 197, 198, 3, 30, 15, 0, 198,
		29, 1, 0, 0, 0, 199, 200, 5, 43, 0, 0, 200, 201, 5, 44, 0, 0, 201, 202,
		3, 90, 45, 0, 202, 203, 5, 41, 0, 0, 203, 204, 3, 68, 34, 0, 204, 205,
		5, 42, 0, 0, 205, 216, 1, 0, 0, 0, 206, 207, 5, 43, 0, 0, 207, 208, 5,
		44, 0, 0, 208, 209, 5, 43, 0, 0, 209, 210, 5, 44, 0, 0, 210, 211, 3, 90,
		45, 0, 211, 212, 5, 41, 0, 0, 212, 213, 3, 70, 35, 0, 213, 214, 5, 42,
		0, 0, 214, 216, 1, 0, 0, 0, 215, 199, 1, 0, 0, 0, 215, 206, 1, 0, 0, 0,
		216, 31, 1, 0, 0, 0, 217, 218, 5, 60, 0, 0, 218, 219, 5, 43, 0, 0, 219,
		220, 3, 72, 36, 0, 220, 221, 5, 44, 0, 0, 221, 231, 1, 0, 0, 0, 222, 223,
		5, 60, 0, 0, 223, 224, 5, 43, 0, 0, 224, 225, 3, 72, 36, 0, 225, 226, 5,
		44, 0, 0, 226, 227, 5, 43, 0, 0, 227, 228, 3, 72, 36, 0, 228, 229, 5, 44,
		0, 0, 229, 231, 1, 0, 0, 0, 230, 217, 1, 0, 0, 0, 230, 222, 1, 0, 0, 0,
		231, 33, 1, 0, 0, 0, 232, 233, 5, 60, 0, 0, 233, 234, 5, 43, 0, 0, 234,
		235, 3, 72, 36, 0, 235, 236, 5, 44, 0, 0, 236, 237, 5, 49, 0, 0, 237, 238,
		3, 72, 36, 0, 238, 250, 1, 0, 0, 0, 239, 240, 5, 60, 0, 0, 240, 241, 5,
		43, 0, 0, 241, 242, 3, 72, 36, 0, 242, 243, 5, 44, 0, 0, 243, 244, 5, 43,
		0, 0, 244, 245, 3, 72, 36, 0, 245, 246, 5, 44, 0, 0, 246, 247, 5, 49, 0,
		0, 247, 248, 3, 72, 36, 0, 248, 250, 1, 0, 0, 0, 249, 232, 1, 0, 0, 0,
		249, 239, 1, 0, 0, 0, 250, 35, 1, 0, 0, 0, 251, 252, 5, 24, 0, 0, 252,
		253, 5, 39, 0, 0, 253, 254, 3, 68, 34, 0, 254, 255, 5, 40, 0, 0, 255, 37,
		1, 0, 0, 0, 256, 257, 5, 25, 0, 0, 257, 258, 5, 39, 0, 0, 258, 259, 3,
		68, 34, 0, 259, 260, 5, 40, 0, 0, 260, 39, 1, 0, 0, 0, 261, 262, 5, 26,
		0, 0, 262, 263, 5, 39, 0, 0, 263, 264, 3, 68, 34, 0, 264, 265, 5, 40, 0,
		0, 265, 41, 1, 0, 0, 0, 266, 267, 5, 27, 0, 0, 267, 268, 5, 39, 0, 0, 268,
		269, 3, 68, 34, 0, 269, 270, 5, 40, 0, 0, 270, 43, 1, 0, 0, 0, 271, 272,
		5, 6, 0, 0, 272, 273, 5, 60, 0, 0, 273, 274, 5, 43, 0, 0, 274, 275, 5,
		44, 0, 0, 275, 284, 3, 90, 45, 0, 276, 277, 5, 6, 0, 0, 277, 278, 5, 60,
		0, 0, 278, 279, 5, 43, 0, 0, 279, 280, 5, 44, 0, 0, 280, 281, 5, 43, 0,
		0, 281, 282, 5, 44, 0, 0, 282, 284, 3, 90, 45, 0, 283, 271, 1, 0, 0, 0,
		283, 276, 1, 0, 0, 0, 284, 45, 1, 0, 0, 0, 285, 286, 7, 0, 0, 0, 286, 47,
		1, 0, 0, 0, 287, 289, 5, 16, 0, 0, 288, 290, 3, 72, 36, 0, 289, 288, 1,
		0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 292, 1, 0, 0, 0, 291, 293, 5, 56, 0,
		0, 292, 291, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 49, 1, 0, 0, 0, 294,
		295, 5, 33, 0, 0, 295, 296, 5, 39, 0, 0, 296, 297, 3, 68, 34, 0, 297, 298,
		5, 40, 0, 0, 298, 51, 1, 0, 0, 0, 299, 300, 5, 31, 0, 0, 300, 301, 5, 39,
		0, 0, 301, 302, 3, 68, 34, 0, 302, 303, 5, 40, 0, 0, 303, 53, 1, 0, 0,
		0, 304, 305, 5, 32, 0, 0, 305, 306, 5, 39, 0, 0, 306, 307, 3, 68, 34, 0,
		307, 308, 5, 40, 0, 0, 308, 55, 1, 0, 0, 0, 309, 310, 5, 60, 0, 0, 310,
		311, 5, 1, 0, 0, 311, 312, 3, 72, 36, 0, 312, 57, 1, 0, 0, 0, 313, 314,
		5, 60, 0, 0, 314, 315, 5, 2, 0, 0, 315, 316, 3, 72, 36, 0, 316, 59, 1,
		0, 0, 0, 317, 318, 5, 60, 0, 0, 318, 319, 5, 49, 0, 0, 319, 325, 3, 72,
		36, 0, 320, 321, 5, 60, 0, 0, 321, 325, 5, 58, 0, 0, 322, 323, 5, 60, 0,
		0, 323, 325, 5, 59, 0, 0, 324, 317, 1, 0, 0, 0, 324, 320, 1, 0, 0, 0, 324,
		322, 1, 0, 0, 0, 325, 61, 1, 0, 0, 0, 326, 331, 5, 60, 0, 0, 327, 328,
		5, 57, 0, 0, 328, 330, 5, 60, 0, 0, 329, 327, 1, 0, 0, 0, 330, 333, 1,
		0, 0, 0, 331, 329, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 63, 1, 0, 0,
		0, 333, 331, 1, 0, 0, 0, 334, 339, 3, 66, 33, 0, 335, 336, 5, 57, 0, 0,
		336, 338, 3, 66, 33, 0, 337, 335, 1, 0, 0, 0, 338, 341, 1, 0, 0, 0, 339,
		337, 1, 0, 0, 0, 339, 340, 1, 0, 0, 0, 340, 65, 1, 0, 0, 0, 341, 339, 1,
		0, 0, 0, 342, 343, 5, 60, 0, 0, 343, 344, 3, 90, 45, 0, 344, 67, 1, 0,
		0, 0, 345, 350, 3, 72, 36, 0, 346, 347, 5, 57, 0, 0, 347, 349, 3, 72, 36,
		0, 348, 346, 1, 0, 0, 0, 349, 352, 1, 0, 0, 0, 350, 348, 1, 0, 0, 0, 350,
		351, 1, 0, 0, 0, 351, 69, 1, 0, 0, 0, 352, 350, 1, 0, 0, 0, 353, 354, 5,
		41, 0, 0, 354, 355, 3, 68, 34, 0, 355, 363, 5, 42, 0, 0, 356, 357, 5, 57,
		0, 0, 357, 358, 5, 41, 0, 0, 358, 359, 3, 68, 34, 0, 359, 360, 5, 42, 0,
		0, 360, 362, 1, 0, 0, 0, 361, 356, 1, 0, 0, 0, 362, 365, 1, 0, 0, 0, 363,
		361, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0, 364, 367, 1, 0, 0, 0, 365, 363,
		1, 0, 0, 0, 366, 368, 5, 57, 0, 0, 367, 366, 1, 0, 0, 0, 367, 368, 1, 0,
		0, 0, 368, 71, 1, 0, 0, 0, 369, 370, 6, 36, -1, 0, 370, 371, 5, 35, 0,
		0, 371, 401, 3, 72, 36, 34, 372, 373, 5, 54, 0, 0, 373, 401, 3, 72, 36,
		33, 374, 401, 5, 61, 0, 0, 375, 401, 5, 62, 0, 0, 376, 401, 5, 63, 0, 0,
		377, 401, 5, 64, 0, 0, 378, 401, 5, 22, 0, 0, 379, 401, 5, 23, 0, 0, 380,
		401, 3, 32, 16, 0, 381, 383, 5, 43, 0, 0, 382, 384, 3, 68, 34, 0, 383,
		382, 1, 0, 0, 0, 383, 384, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 401,
		5, 44, 0, 0, 386, 401, 5, 60, 0, 0, 387, 388, 5, 39, 0, 0, 388, 389, 3,
		72, 36, 0, 389, 390, 5, 40, 0, 0, 390, 401, 1, 0, 0, 0, 391, 401, 3, 52,
		26, 0, 392, 401, 3, 54, 27, 0, 393, 401, 3, 50, 25, 0, 394, 401, 3, 20,
		10, 0, 395, 401, 3, 22, 11, 0, 396, 401, 3, 42, 21, 0, 397, 401, 3, 36,
		18, 0, 398, 401, 3, 38, 19, 0, 399, 401, 3, 40, 20, 0, 400, 369, 1, 0,
		0, 0, 400, 372, 1, 0, 0, 0, 400, 374, 1, 0, 0, 0, 400, 375, 1, 0, 0, 0,
		400, 376, 1, 0, 0, 0, 400, 377, 1, 0, 0, 0, 400, 378, 1, 0, 0, 0, 400,
		379, 1, 0, 0, 0, 400, 380, 1, 0, 0, 0, 400, 381, 1, 0, 0, 0, 400, 386,
		1, 0, 0, 0, 400, 387, 1, 0, 0, 0, 400, 391, 1, 0, 0, 0, 400, 392, 1, 0,
		0, 0, 400, 393, 1, 0, 0, 0, 400, 394, 1, 0, 0, 0, 400, 395, 1, 0, 0, 0,
		400, 396, 1, 0, 0, 0, 400, 397, 1, 0, 0, 0, 400, 398, 1, 0, 0, 0, 400,
		399, 1, 0, 0, 0, 401, 443, 1, 0, 0, 0, 402, 403, 10, 32, 0, 0, 403, 404,
		5, 38, 0, 0, 404, 442, 3, 72, 36, 33, 405, 406, 10, 31, 0, 0, 406, 407,
		5, 37, 0, 0, 407, 442, 3, 72, 36, 32, 408, 409, 10, 30, 0, 0, 409, 410,
		5, 36, 0, 0, 410, 442, 3, 72, 36, 31, 411, 412, 10, 29, 0, 0, 412, 413,
		5, 35, 0, 0, 413, 442, 3, 72, 36, 30, 414, 415, 10, 28, 0, 0, 415, 416,
		5, 34, 0, 0, 416, 442, 3, 72, 36, 29, 417, 418, 10, 27, 0, 0, 418, 419,
		5, 45, 0, 0, 419, 442, 3, 72, 36, 28, 420, 421, 10, 26, 0, 0, 421, 422,
		5, 46, 0, 0, 422, 442, 3, 72, 36, 27, 423, 424, 10, 25, 0, 0, 424, 425,
		5, 47, 0, 0, 425, 442, 3, 72, 36, 26, 426, 427, 10, 24, 0, 0, 427, 428,
		5, 48, 0, 0, 428, 442, 3, 72, 36, 25, 429, 430, 10, 23, 0, 0, 430, 431,
		5, 50, 0, 0, 431, 442, 3, 72, 36, 24, 432, 433, 10, 22, 0, 0, 433, 434,
		5, 51, 0, 0, 434, 442, 3, 72, 36, 23, 435, 436, 10, 21, 0, 0, 436, 437,
		5, 52, 0, 0, 437, 442, 3, 72, 36, 22, 438, 439, 10, 20, 0, 0, 439, 440,
		5, 53, 0, 0, 440, 442, 3, 72, 36, 21, 441, 402, 1, 0, 0, 0, 441, 405, 1,
		0, 0, 0, 441, 408, 1, 0, 0, 0, 441, 411, 1, 0, 0, 0, 441, 414, 1, 0, 0,
		0, 441, 417, 1, 0, 0, 0, 441, 420, 1, 0, 0, 0, 441, 423, 1, 0, 0, 0, 441,
		426, 1, 0, 0, 0, 441, 429, 1, 0, 0, 0, 441, 432, 1, 0, 0, 0, 441, 435,
		1, 0, 0, 0, 441, 438, 1, 0, 0, 0, 442, 445, 1, 0, 0, 0, 443, 441, 1, 0,
		0, 0, 443, 444, 1, 0, 0, 0, 444, 73, 1, 0, 0, 0, 445, 443, 1, 0, 0, 0,
		446, 452, 5, 8, 0, 0, 447, 448, 5, 39, 0, 0, 448, 449, 3, 72, 36, 0, 449,
		450, 5, 40, 0, 0, 450, 453, 1, 0, 0, 0, 451, 453, 3, 72, 36, 0, 452, 447,
		1, 0, 0, 0, 452, 451, 1, 0, 0, 0, 453, 454, 1, 0, 0, 0, 454, 458, 3, 18,
		9, 0, 455, 457, 3, 76, 38, 0, 456, 455, 1, 0, 0, 0, 457, 460, 1, 0, 0,
		0, 458, 456, 1, 0, 0, 0, 458, 459, 1, 0, 0, 0, 459, 463, 1, 0, 0, 0, 460,
		458, 1, 0, 0, 0, 461, 462, 5, 9, 0, 0, 462, 464, 3, 18, 9, 0, 463, 461,
		1, 0, 0, 0, 463, 464, 1, 0, 0, 0, 464, 75, 1, 0, 0, 0, 465, 466, 5, 9,
		0, 0, 466, 472, 5, 8, 0, 0, 467, 468, 5, 39, 0, 0, 468, 469, 3, 72, 36,
		0, 469, 470, 5, 40, 0, 0, 470, 473, 1, 0, 0, 0, 471, 473, 3, 72, 36, 0,
		472, 467, 1, 0, 0, 0, 472, 471, 1, 0, 0, 0, 473, 474, 1, 0, 0, 0, 474,
		475, 3, 18, 9, 0, 475, 77, 1, 0, 0, 0, 476, 477, 5, 13, 0, 0, 477, 478,
		3, 72, 36, 0, 478, 479, 3, 18, 9, 0, 479, 498, 1, 0, 0, 0, 480, 481, 5,
		13, 0, 0, 481, 482, 3, 60, 30, 0, 482, 483, 5, 56, 0, 0, 483, 484, 3, 72,
		36, 0, 484, 485, 5, 56, 0, 0, 485, 486, 3, 60, 30, 0, 486, 487, 3, 18,
		9, 0, 487, 498, 1, 0, 0, 0, 488, 489, 5, 13, 0, 0, 489, 490, 5, 60, 0,
		0, 490, 491, 5, 57, 0, 0, 491, 492, 5, 60, 0, 0, 492, 493, 5, 49, 0, 0,
		493, 494, 5, 30, 0, 0, 494, 495, 3, 72, 36, 0, 495, 496, 3, 18, 9, 0, 496,
		498, 1, 0, 0, 0, 497, 476, 1, 0, 0, 0, 497, 480, 1, 0, 0, 0, 497, 488,
		1, 0, 0, 0, 498, 79, 1, 0, 0, 0, 499, 500, 5, 10, 0, 0, 500, 501, 3, 72,
		36, 0, 501, 505, 5, 41, 0, 0, 502, 504, 3, 86, 43, 0, 503, 502, 1, 0, 0,
		0, 504, 507, 1, 0, 0, 0, 505, 503, 1, 0, 0, 0, 505, 506, 1, 0, 0, 0, 506,
		509, 1, 0, 0, 0, 507, 505, 1, 0, 0, 0, 508, 510, 3, 88, 44, 0, 509, 508,
		1, 0, 0, 0, 509, 510, 1, 0, 0, 0, 510, 511, 1, 0, 0, 0, 511, 512, 5, 42,
		0, 0, 512, 81, 1, 0, 0, 0, 513, 514, 5, 14, 0, 0, 514, 83, 1, 0, 0, 0,
		515, 516, 5, 15, 0, 0, 516, 85, 1, 0, 0, 0, 517, 518, 5, 11, 0, 0, 518,
		519, 3, 72, 36, 0, 519, 520, 5, 55, 0, 0, 520, 521, 3, 2, 1, 0, 521, 87,
		1, 0, 0, 0, 522, 523, 5, 12, 0, 0, 523, 524, 5, 55, 0, 0, 524, 525, 3,
		2, 1, 0, 525, 89, 1, 0, 0, 0, 526, 527, 7, 1, 0, 0, 527, 91, 1, 0, 0, 0,
		30, 95, 101, 126, 160, 162, 181, 191, 215, 230, 249, 283, 289, 292, 324,
		331, 339, 350, 363, 367, 383, 400, 441, 443, 452, 458, 463, 472, 497, 505,
		509,
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
	gramaticaParserPRINT           = 3
	gramaticaParserPRINTLN         = 4
	gramaticaParserFUNCIONES       = 5
	gramaticaParserMUT             = 6
	gramaticaParserSLICE           = 7
	gramaticaParserIF              = 8
	gramaticaParserELSE            = 9
	gramaticaParserSWITCH          = 10
	gramaticaParserCASE            = 11
	gramaticaParserDEFAULT         = 12
	gramaticaParserFOR             = 13
	gramaticaParserBREAK           = 14
	gramaticaParserCONTINUE        = 15
	gramaticaParserRETURN          = 16
	gramaticaParserINT             = 17
	gramaticaParserFLOAT           = 18
	gramaticaParserSTRING          = 19
	gramaticaParserBOOL            = 20
	gramaticaParserNIL             = 21
	gramaticaParserTRUE            = 22
	gramaticaParserFALSE           = 23
	gramaticaParserINDEXOF         = 24
	gramaticaParserJOIN            = 25
	gramaticaParserLEN             = 26
	gramaticaParserAPPEND          = 27
	gramaticaParserSTRUCT          = 28
	gramaticaParserFUNC            = 29
	gramaticaParserRANGE           = 30
	gramaticaParserATOI            = 31
	gramaticaParserPARSEFLOAT      = 32
	gramaticaParserTYPEOF          = 33
	gramaticaParserMAS             = 34
	gramaticaParserMENOS           = 35
	gramaticaParserPOR             = 36
	gramaticaParserDIV             = 37
	gramaticaParserMODULO          = 38
	gramaticaParserPAR1            = 39
	gramaticaParserPAR2            = 40
	gramaticaParserLLAV1           = 41
	gramaticaParserLLAV2           = 42
	gramaticaParserCOR1            = 43
	gramaticaParserCOR2            = 44
	gramaticaParserDIFERENTE       = 45
	gramaticaParserIGUALDAD        = 46
	gramaticaParserMENIGUAL        = 47
	gramaticaParserMAYIGUAL        = 48
	gramaticaParserIGUAL           = 49
	gramaticaParserMENOR           = 50
	gramaticaParserMAYOR           = 51
	gramaticaParserOR              = 52
	gramaticaParserAND             = 53
	gramaticaParserDIFER           = 54
	gramaticaParserDOSPTS          = 55
	gramaticaParserPTCOMA          = 56
	gramaticaParserCOMA            = 57
	gramaticaParserFINCREMENTO     = 58
	gramaticaParserFDECREMENTO     = 59
	gramaticaParserIDENTIFICADOR   = 60
	gramaticaParserENTERO          = 61
	gramaticaParserDECIMAL         = 62
	gramaticaParserCADENA          = 63
	gramaticaParserRUNE            = 64
	gramaticaParserCOMENTARIO      = 65
	gramaticaParserMULTICOMENTARIO = 66
	gramaticaParserWS              = 67
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
	gramaticaParserRULE_bloque                       = 9
	gramaticaParserRULE_llamadaFuncionesSinParametro = 10
	gramaticaParserRULE_llamadaFuncionesConParametro = 11
	gramaticaParserRULE_fnSinParametro               = 12
	gramaticaParserRULE_fnConParametro               = 13
	gramaticaParserRULE_sliceDef                     = 14
	gramaticaParserRULE_sliceLiteral                 = 15
	gramaticaParserRULE_accesoElementoSlice          = 16
	gramaticaParserRULE_modificacionElementoSlice    = 17
	gramaticaParserRULE_fnIndexOf                    = 18
	gramaticaParserRULE_fnJoin                       = 19
	gramaticaParserRULE_fnLen                        = 20
	gramaticaParserRULE_fnAppend                     = 21
	gramaticaParserRULE_declaracionSliceVacio        = 22
	gramaticaParserRULE_tipoRetorno                  = 23
	gramaticaParserRULE_retorno                      = 24
	gramaticaParserRULE_fnTypeOf                     = 25
	gramaticaParserRULE_fnAtoi                       = 26
	gramaticaParserRULE_fnParseToFloat               = 27
	gramaticaParserRULE_asigIncre                    = 28
	gramaticaParserRULE_asigDecre                    = 29
	gramaticaParserRULE_asignacion                   = 30
	gramaticaParserRULE_listaIDS                     = 31
	gramaticaParserRULE_listaPar                     = 32
	gramaticaParserRULE_parametro                    = 33
	gramaticaParserRULE_listaExpr                    = 34
	gramaticaParserRULE_listaExprList                = 35
	gramaticaParserRULE_expresion                    = 36
	gramaticaParserRULE_sif                          = 37
	gramaticaParserRULE_elseifPart                   = 38
	gramaticaParserRULE_sfor                         = 39
	gramaticaParserRULE_sSwitch                      = 40
	gramaticaParserRULE_break_                       = 41
	gramaticaParserRULE_continue_                    = 42
	gramaticaParserRULE_caseBlock                    = 43
	gramaticaParserRULE_defaultBlock                 = 44
	gramaticaParserRULE_tipos                        = 45
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
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1152923703630226808) != 0 {
		{
			p.SetState(92)
			p.Instrucciones()
		}

		p.SetState(97)
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
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(98)
				p.Instruccion()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(101)
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
	FnSinParametro() IFnSinParametroContext
	FnConParametro() IFnConParametroContext
	LlamadaFuncionesSinParametro() ILlamadaFuncionesSinParametroContext
	LlamadaFuncionesConParametro() ILlamadaFuncionesConParametroContext
	Retorno() IRetornoContext
	AsigIncre() IAsigIncreContext
	AsigDecre() IAsigDecreContext
	Asignacion() IAsignacionContext
	Sif() ISifContext
	Sfor() ISforContext
	SSwitch() ISSwitchContext
	Break_() IBreak_Context
	Continue_() IContinue_Context
	SliceDef() ISliceDefContext
	DeclaracionSliceVacio() IDeclaracionSliceVacioContext
	ModificacionElementoSlice() IModificacionElementoSliceContext
	Bloque() IBloqueContext

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

func (s *InstruccionContext) Bloque() IBloqueContext {
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
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.Print_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.Println_()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(105)
			p.DeclaracionMultiple()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(106)
			p.DeclaracionMultipleSimple()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(107)
			p.DeclaracionMultipleSinTipo()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(108)
			p.AsignacionMultiple()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(109)
			p.FnSinParametro()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(110)
			p.FnConParametro()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(111)
			p.LlamadaFuncionesSinParametro()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(112)
			p.LlamadaFuncionesConParametro()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(113)
			p.Retorno()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(114)
			p.AsigIncre()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(115)
			p.AsigDecre()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(116)
			p.Asignacion()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(117)
			p.Sif()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(118)
			p.Sfor()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(119)
			p.SSwitch()
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(120)
			p.Break_()
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(121)
			p.Continue_()
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(122)
			p.SliceDef()
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(123)
			p.DeclaracionSliceVacio()
		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(124)
			p.ModificacionElementoSlice()
		}

	case 23:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(125)
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
		p.SetState(128)
		p.Match(gramaticaParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.ListaExpr()
	}
	{
		p.SetState(131)
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
		p.SetState(133)
		p.Match(gramaticaParserPRINTLN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.ListaExpr()
	}
	{
		p.SetState(136)
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
		p.SetState(138)
		p.Match(gramaticaParserMUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.ListaIDS()
	}
	{
		p.SetState(140)
		p.Tipos()
	}
	{
		p.SetState(141)
		p.Match(gramaticaParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
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
		p.SetState(144)
		p.Match(gramaticaParserMUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.ListaIDS()
	}
	{
		p.SetState(146)
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
		p.SetState(148)
		p.Match(gramaticaParserMUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.ListaIDS()
	}
	{
		p.SetState(150)
		p.Match(gramaticaParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
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
		p.SetState(153)
		p.ListaIDS()
	}
	{
		p.SetState(154)
		p.Match(gramaticaParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
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

// IBloqueContext is an interface to support dynamic dispatch.
type IBloqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LLAV1() antlr.TerminalNode
	LLAV2() antlr.TerminalNode
	AllInstruccion() []IInstruccionContext
	Instruccion(i int) IInstruccionContext
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext

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

func (s *BloqueContext) AllInstruccion() []IInstruccionContext {
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

func (s *BloqueContext) Instruccion(i int) IInstruccionContext {
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

func (s *BloqueContext) AllExpresion() []IExpresionContext {
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

func (s *BloqueContext) Expresion(i int) IExpresionContext {
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
	p.EnterRule(localctx, 18, gramaticaParserRULE_bloque)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(gramaticaParserLLAV1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-3)) & ^0x3f) == 0 && ((int64(1)<<(_la-3))&4469824079481289903) != 0 {
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(158)
				p.Instruccion()
			}

		case 2:
			{
				p.SetState(159)
				p.expresion(0)
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(165)
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
	p.EnterRule(localctx, 20, gramaticaParserRULE_llamadaFuncionesSinParametro)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
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
	p.EnterRule(localctx, 22, gramaticaParserRULE_llamadaFuncionesConParametro)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.ListaExpr()
	}
	{
		p.SetState(174)
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
	Bloque() IBloqueContext
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

func (s *FnSinParametroContext) Bloque() IBloqueContext {
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
	p.EnterRule(localctx, 24, gramaticaParserRULE_fnSinParametro)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Match(gramaticaParserFUNCIONES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(179)
		p.Match(gramaticaParserPAR2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-17)) & ^0x3f) == 0 && ((int64(1)<<(_la-17))&149533581377551) != 0 {
		{
			p.SetState(180)
			p.TipoRetorno()
		}

	}
	{
		p.SetState(183)
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
	Bloque() IBloqueContext
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

func (s *FnConParametroContext) Bloque() IBloqueContext {
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
	p.EnterRule(localctx, 26, gramaticaParserRULE_fnConParametro)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Match(gramaticaParserFUNCIONES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(188)
		p.ListaPar()
	}
	{
		p.SetState(189)
		p.Match(gramaticaParserPAR2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-17)) & ^0x3f) == 0 && ((int64(1)<<(_la-17))&149533581377551) != 0 {
		{
			p.SetState(190)
			p.TipoRetorno()
		}

	}
	{
		p.SetState(193)
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
	p.EnterRule(localctx, 28, gramaticaParserRULE_sliceDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(196)
		p.Match(gramaticaParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
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
	p.EnterRule(localctx, 30, gramaticaParserRULE_sliceLiteral)
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.Tipos()
		}
		{
			p.SetState(202)
			p.Match(gramaticaParserLLAV1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.ListaExpr()
		}
		{
			p.SetState(204)
			p.Match(gramaticaParserLLAV2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
			p.Tipos()
		}
		{
			p.SetState(211)
			p.Match(gramaticaParserLLAV1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(212)
			p.ListaExprList()
		}
		{
			p.SetState(213)
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
	p.EnterRule(localctx, 32, gramaticaParserRULE_accesoElementoSlice)
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
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
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)
			p.expresion(0)
		}
		{
			p.SetState(220)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(222)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.expresion(0)
		}
		{
			p.SetState(225)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)
			p.expresion(0)
		}
		{
			p.SetState(228)
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
	p.EnterRule(localctx, 34, gramaticaParserRULE_modificacionElementoSlice)
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.expresion(0)
		}
		{
			p.SetState(235)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)
			p.Match(gramaticaParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.expresion(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(239)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(240)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.expresion(0)
		}
		{
			p.SetState(242)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.expresion(0)
		}
		{
			p.SetState(245)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)
			p.Match(gramaticaParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)
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
	p.EnterRule(localctx, 36, gramaticaParserRULE_fnIndexOf)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.Match(gramaticaParserINDEXOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(252)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(253)
		p.ListaExpr()
	}
	{
		p.SetState(254)
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
	p.EnterRule(localctx, 38, gramaticaParserRULE_fnJoin)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.Match(gramaticaParserJOIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(257)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.ListaExpr()
	}
	{
		p.SetState(259)
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
	p.EnterRule(localctx, 40, gramaticaParserRULE_fnLen)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Match(gramaticaParserLEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(263)
		p.ListaExpr()
	}
	{
		p.SetState(264)
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
	p.EnterRule(localctx, 42, gramaticaParserRULE_fnAppend)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Match(gramaticaParserAPPEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
		p.ListaExpr()
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
	p.EnterRule(localctx, 44, gramaticaParserRULE_declaracionSliceVacio)
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(271)
			p.Match(gramaticaParserMUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(272)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(273)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(274)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(275)
			p.Tipos()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(276)
			p.Match(gramaticaParserMUT)
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
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(280)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)
			p.Match(gramaticaParserCOR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
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
	p.EnterRule(localctx, 46, gramaticaParserRULE_tipoRetorno)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-17)) & ^0x3f) == 0 && ((int64(1)<<(_la-17))&149533581377551) != 0) {
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
	p.EnterRule(localctx, 48, gramaticaParserRULE_retorno)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
		p.Match(gramaticaParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(288)
			p.expresion(0)
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserPTCOMA {
		{
			p.SetState(291)
			p.Match(gramaticaParserPTCOMA)
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
	p.EnterRule(localctx, 50, gramaticaParserRULE_fnTypeOf)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)
		p.Match(gramaticaParserTYPEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(295)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(296)
		p.ListaExpr()
	}
	{
		p.SetState(297)
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
	p.EnterRule(localctx, 52, gramaticaParserRULE_fnAtoi)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.Match(gramaticaParserATOI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(300)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(301)
		p.ListaExpr()
	}
	{
		p.SetState(302)
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
	p.EnterRule(localctx, 54, gramaticaParserRULE_fnParseToFloat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)
		p.Match(gramaticaParserPARSEFLOAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(305)
		p.Match(gramaticaParserPAR1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
		p.ListaExpr()
	}
	{
		p.SetState(307)
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
	p.EnterRule(localctx, 56, gramaticaParserRULE_asigIncre)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(310)
		p.Match(gramaticaParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(311)
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
	p.EnterRule(localctx, 58, gramaticaParserRULE_asigDecre)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(314)
		p.Match(gramaticaParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(315)
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
	p.EnterRule(localctx, 60, gramaticaParserRULE_asignacion)
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAsignacionNormalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(317)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(318)
			p.Match(gramaticaParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(319)
			p.expresion(0)
		}

	case 2:
		localctx = NewAsignacionIncrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(320)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(321)
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
			p.SetState(322)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
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
	p.EnterRule(localctx, 62, gramaticaParserRULE_listaIDS)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserCOMA {
		{
			p.SetState(327)
			p.Match(gramaticaParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(333)
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
	p.EnterRule(localctx, 64, gramaticaParserRULE_listaPar)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
		p.Parametro()
	}
	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserCOMA {
		{
			p.SetState(335)
			p.Match(gramaticaParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)
			p.Parametro()
		}

		p.SetState(341)
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
	p.EnterRule(localctx, 66, gramaticaParserRULE_parametro)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(342)
		p.Match(gramaticaParserIDENTIFICADOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)
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
	p.EnterRule(localctx, 68, gramaticaParserRULE_listaExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.expresion(0)
	}
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserCOMA {
		{
			p.SetState(346)
			p.Match(gramaticaParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(347)
			p.expresion(0)
		}

		p.SetState(352)
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
	p.EnterRule(localctx, 70, gramaticaParserRULE_listaExprList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(353)
		p.Match(gramaticaParserLLAV1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(354)
		p.ListaExpr()
	}
	{
		p.SetState(355)
		p.Match(gramaticaParserLLAV2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(356)
				p.Match(gramaticaParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(357)
				p.Match(gramaticaParserLLAV1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(358)
				p.ListaExpr()
			}
			{
				p.SetState(359)
				p.Match(gramaticaParserLLAV2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(365)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserCOMA {
		{
			p.SetState(366)
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
	IDENTIFICADOR() antlr.TerminalNode
	PAR1() antlr.TerminalNode
	PAR2() antlr.TerminalNode
	FnAtoi() IFnAtoiContext
	FnParseToFloat() IFnParseToFloatContext
	FnTypeOf() IFnTypeOfContext
	LlamadaFuncionesSinParametro() ILlamadaFuncionesSinParametroContext
	LlamadaFuncionesConParametro() ILlamadaFuncionesConParametroContext
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
	_startState := 72
	p.EnterRecursionRule(localctx, 72, gramaticaParserRULE_expresion, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(370)
			p.Match(gramaticaParserMENOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)
			p.expresion(34)
		}

	case 2:
		{
			p.SetState(372)
			p.Match(gramaticaParserDIFER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)
			p.expresion(33)
		}

	case 3:
		{
			p.SetState(374)
			p.Match(gramaticaParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(375)
			p.Match(gramaticaParserDECIMAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		{
			p.SetState(376)
			p.Match(gramaticaParserCADENA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		{
			p.SetState(377)
			p.Match(gramaticaParserRUNE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		{
			p.SetState(378)
			p.Match(gramaticaParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		{
			p.SetState(379)
			p.Match(gramaticaParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		{
			p.SetState(380)
			p.AccesoElementoSlice()
		}

	case 10:
		{
			p.SetState(381)
			p.Match(gramaticaParserCOR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(383)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-22)) & ^0x3f) == 0 && ((int64(1)<<(_la-22))&8525512322623) != 0 {
			{
				p.SetState(382)
				p.ListaExpr()
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

	case 11:
		{
			p.SetState(386)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		{
			p.SetState(387)
			p.Match(gramaticaParserPAR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(388)
			p.expresion(0)
		}
		{
			p.SetState(389)
			p.Match(gramaticaParserPAR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		{
			p.SetState(391)
			p.FnAtoi()
		}

	case 14:
		{
			p.SetState(392)
			p.FnParseToFloat()
		}

	case 15:
		{
			p.SetState(393)
			p.FnTypeOf()
		}

	case 16:
		{
			p.SetState(394)
			p.LlamadaFuncionesSinParametro()
		}

	case 17:
		{
			p.SetState(395)
			p.LlamadaFuncionesConParametro()
		}

	case 18:
		{
			p.SetState(396)
			p.FnAppend()
		}

	case 19:
		{
			p.SetState(397)
			p.FnIndexOf()
		}

	case 20:
		{
			p.SetState(398)
			p.FnJoin()
		}

	case 21:
		{
			p.SetState(399)
			p.FnLen()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(441)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(402)

				if !(p.Precpred(p.GetParserRuleContext(), 32)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 32)", ""))
					goto errorExit
				}
				{
					p.SetState(403)
					p.Match(gramaticaParserMODULO)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(404)
					p.expresion(33)
				}

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(405)

				if !(p.Precpred(p.GetParserRuleContext(), 31)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 31)", ""))
					goto errorExit
				}
				{
					p.SetState(406)
					p.Match(gramaticaParserDIV)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(407)
					p.expresion(32)
				}

			case 3:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(408)

				if !(p.Precpred(p.GetParserRuleContext(), 30)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 30)", ""))
					goto errorExit
				}
				{
					p.SetState(409)
					p.Match(gramaticaParserPOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(410)
					p.expresion(31)
				}

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(411)

				if !(p.Precpred(p.GetParserRuleContext(), 29)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 29)", ""))
					goto errorExit
				}
				{
					p.SetState(412)
					p.Match(gramaticaParserMENOS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(413)
					p.expresion(30)
				}

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(414)

				if !(p.Precpred(p.GetParserRuleContext(), 28)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 28)", ""))
					goto errorExit
				}
				{
					p.SetState(415)
					p.Match(gramaticaParserMAS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(416)
					p.expresion(29)
				}

			case 6:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(417)

				if !(p.Precpred(p.GetParserRuleContext(), 27)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 27)", ""))
					goto errorExit
				}
				{
					p.SetState(418)
					p.Match(gramaticaParserDIFERENTE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(419)
					p.expresion(28)
				}

			case 7:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(420)

				if !(p.Precpred(p.GetParserRuleContext(), 26)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 26)", ""))
					goto errorExit
				}
				{
					p.SetState(421)
					p.Match(gramaticaParserIGUALDAD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(422)
					p.expresion(27)
				}

			case 8:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(423)

				if !(p.Precpred(p.GetParserRuleContext(), 25)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 25)", ""))
					goto errorExit
				}
				{
					p.SetState(424)
					p.Match(gramaticaParserMENIGUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(425)
					p.expresion(26)
				}

			case 9:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(426)

				if !(p.Precpred(p.GetParserRuleContext(), 24)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 24)", ""))
					goto errorExit
				}
				{
					p.SetState(427)
					p.Match(gramaticaParserMAYIGUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(428)
					p.expresion(25)
				}

			case 10:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(429)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
					goto errorExit
				}
				{
					p.SetState(430)
					p.Match(gramaticaParserMENOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(431)
					p.expresion(24)
				}

			case 11:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(432)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(433)
					p.Match(gramaticaParserMAYOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(434)
					p.expresion(23)
				}

			case 12:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(435)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(436)
					p.Match(gramaticaParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(437)
					p.expresion(22)
				}

			case 13:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gramaticaParserRULE_expresion)
				p.SetState(438)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(439)
					p.Match(gramaticaParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(440)
					p.expresion(21)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(445)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 74, gramaticaParserRULE_sif)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(446)
		p.Match(gramaticaParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(452)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(447)
			p.Match(gramaticaParserPAR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(448)
			p.expresion(0)
		}
		{
			p.SetState(449)
			p.Match(gramaticaParserPAR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(451)
			p.expresion(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(454)
		p.Bloque()
	}
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(455)
				p.ElseifPart()
			}

		}
		p.SetState(460)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(463)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserELSE {
		{
			p.SetState(461)
			p.Match(gramaticaParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(462)
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
	p.EnterRule(localctx, 76, gramaticaParserRULE_elseifPart)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(465)
		p.Match(gramaticaParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(466)
		p.Match(gramaticaParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(467)
			p.Match(gramaticaParserPAR1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(468)
			p.expresion(0)
		}
		{
			p.SetState(469)
			p.Match(gramaticaParserPAR2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(471)
			p.expresion(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(474)
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
	p.EnterRule(localctx, 78, gramaticaParserRULE_sfor)
	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForCondicionalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(476)
			p.Match(gramaticaParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(477)
			p.expresion(0)
		}
		{
			p.SetState(478)
			p.Bloque()
		}

	case 2:
		localctx = NewForClasicoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(480)
			p.Match(gramaticaParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(481)
			p.Asignacion()
		}
		{
			p.SetState(482)
			p.Match(gramaticaParserPTCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(483)
			p.expresion(0)
		}
		{
			p.SetState(484)
			p.Match(gramaticaParserPTCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(485)
			p.Asignacion()
		}
		{
			p.SetState(486)
			p.Bloque()
		}

	case 3:
		localctx = NewForRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(488)
			p.Match(gramaticaParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(489)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(490)
			p.Match(gramaticaParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(491)
			p.Match(gramaticaParserIDENTIFICADOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(492)
			p.Match(gramaticaParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(493)
			p.Match(gramaticaParserRANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(494)
			p.expresion(0)
		}
		{
			p.SetState(495)
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
	p.EnterRule(localctx, 80, gramaticaParserRULE_sSwitch)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(499)
		p.Match(gramaticaParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(500)
		p.expresion(0)
	}
	{
		p.SetState(501)
		p.Match(gramaticaParserLLAV1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(505)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == gramaticaParserCASE {
		{
			p.SetState(502)
			p.CaseBlock()
		}

		p.SetState(507)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(509)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == gramaticaParserDEFAULT {
		{
			p.SetState(508)
			p.DefaultBlock()
		}

	}
	{
		p.SetState(511)
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
	p.EnterRule(localctx, 82, gramaticaParserRULE_break_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(513)
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
	p.EnterRule(localctx, 84, gramaticaParserRULE_continue_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(515)
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
	p.EnterRule(localctx, 86, gramaticaParserRULE_caseBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(517)
		p.Match(gramaticaParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(518)
		p.expresion(0)
	}
	{
		p.SetState(519)
		p.Match(gramaticaParserDOSPTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(520)
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
	p.EnterRule(localctx, 88, gramaticaParserRULE_defaultBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(522)
		p.Match(gramaticaParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(523)
		p.Match(gramaticaParserDOSPTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(524)
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
	p.EnterRule(localctx, 90, gramaticaParserRULE_tipos)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(526)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1152921504608813056) != 0) {
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
	case 36:
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
		return p.Precpred(p.GetParserRuleContext(), 32)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 31)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 30)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 29)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 28)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 27)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 26)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 25)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 24)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 20)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
