// Code generated from Gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Gramatica

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "Laboratorio5/gen"
import "Laboratorio5/entorno"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 30, 252,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 5, 3, 40, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 7, 3, 52, 10, 3, 12, 3, 14, 3, 55, 11, 3, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 67, 10, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 79, 10, 4,
	12, 4, 14, 4, 82, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5,
	103, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 123, 10, 5, 12, 5,
	14, 5, 126, 11, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 156, 10, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 191, 10, 7, 12, 7, 14,
	7, 194, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 5, 8, 208, 10, 8, 3, 9, 7, 9, 211, 10, 9, 12, 9, 14, 9,
	214, 11, 9, 3, 10, 3, 10, 5, 10, 218, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 232, 10,
	11, 12, 11, 14, 11, 235, 11, 11, 3, 11, 3, 11, 5, 11, 239, 10, 11, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 2, 6, 4, 6, 8, 12, 14, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	2, 5, 3, 2, 3, 4, 3, 2, 5, 6, 4, 2, 3, 4, 15, 15, 2, 276, 2, 26, 3, 2,
	2, 2, 4, 39, 3, 2, 2, 2, 6, 66, 3, 2, 2, 2, 8, 102, 3, 2, 2, 2, 10, 127,
	3, 2, 2, 2, 12, 155, 3, 2, 2, 2, 14, 207, 3, 2, 2, 2, 16, 212, 3, 2, 2,
	2, 18, 217, 3, 2, 2, 2, 20, 219, 3, 2, 2, 2, 22, 242, 3, 2, 2, 2, 24, 247,
	3, 2, 2, 2, 26, 27, 5, 16, 9, 2, 27, 28, 7, 2, 2, 3, 28, 3, 3, 2, 2, 2,
	29, 30, 8, 3, 1, 2, 30, 31, 7, 7, 2, 2, 31, 32, 5, 4, 3, 2, 32, 33, 7,
	8, 2, 2, 33, 34, 8, 3, 1, 2, 34, 40, 3, 2, 2, 2, 35, 36, 7, 27, 2, 2, 36,
	40, 8, 3, 1, 2, 37, 38, 7, 28, 2, 2, 38, 40, 8, 3, 1, 2, 39, 29, 3, 2,
	2, 2, 39, 35, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 40, 53, 3, 2, 2, 2, 41, 42,
	12, 7, 2, 2, 42, 43, 9, 2, 2, 2, 43, 44, 5, 4, 3, 8, 44, 45, 8, 3, 1, 2,
	45, 52, 3, 2, 2, 2, 46, 47, 12, 6, 2, 2, 47, 48, 9, 3, 2, 2, 48, 49, 5,
	4, 3, 7, 49, 50, 8, 3, 1, 2, 50, 52, 3, 2, 2, 2, 51, 41, 3, 2, 2, 2, 51,
	46, 3, 2, 2, 2, 52, 55, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2,
	2, 54, 5, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 56, 57, 8, 4, 1, 2, 57, 58, 7,
	7, 2, 2, 58, 59, 5, 6, 4, 2, 59, 60, 7, 8, 2, 2, 60, 61, 8, 4, 1, 2, 61,
	67, 3, 2, 2, 2, 62, 63, 7, 27, 2, 2, 63, 67, 8, 4, 1, 2, 64, 65, 7, 28,
	2, 2, 65, 67, 8, 4, 1, 2, 66, 56, 3, 2, 2, 2, 66, 62, 3, 2, 2, 2, 66, 64,
	3, 2, 2, 2, 67, 80, 3, 2, 2, 2, 68, 69, 12, 7, 2, 2, 69, 70, 9, 2, 2, 2,
	70, 71, 5, 6, 4, 8, 71, 72, 8, 4, 1, 2, 72, 79, 3, 2, 2, 2, 73, 74, 12,
	6, 2, 2, 74, 75, 9, 3, 2, 2, 75, 76, 5, 6, 4, 7, 76, 77, 8, 4, 1, 2, 77,
	79, 3, 2, 2, 2, 78, 68, 3, 2, 2, 2, 78, 73, 3, 2, 2, 2, 79, 82, 3, 2, 2,
	2, 80, 78, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 7, 3, 2, 2, 2, 82, 80, 3,
	2, 2, 2, 83, 84, 8, 5, 1, 2, 84, 85, 7, 9, 2, 2, 85, 86, 5, 8, 5, 10, 86,
	87, 8, 5, 1, 2, 87, 103, 3, 2, 2, 2, 88, 89, 5, 4, 3, 2, 89, 90, 5, 14,
	8, 2, 90, 91, 5, 4, 3, 2, 91, 92, 8, 5, 1, 2, 92, 103, 3, 2, 2, 2, 93,
	94, 7, 7, 2, 2, 94, 95, 5, 8, 5, 2, 95, 96, 7, 8, 2, 2, 96, 97, 8, 5, 1,
	2, 97, 103, 3, 2, 2, 2, 98, 99, 7, 13, 2, 2, 99, 103, 8, 5, 1, 2, 100,
	101, 7, 14, 2, 2, 101, 103, 8, 5, 1, 2, 102, 83, 3, 2, 2, 2, 102, 88, 3,
	2, 2, 2, 102, 93, 3, 2, 2, 2, 102, 98, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2,
	103, 124, 3, 2, 2, 2, 104, 105, 12, 8, 2, 2, 105, 106, 7, 10, 2, 2, 106,
	107, 5, 10, 6, 2, 107, 108, 5, 8, 5, 9, 108, 109, 8, 5, 1, 2, 109, 123,
	3, 2, 2, 2, 110, 111, 12, 7, 2, 2, 111, 112, 7, 11, 2, 2, 112, 113, 5,
	10, 6, 2, 113, 114, 5, 8, 5, 8, 114, 115, 8, 5, 1, 2, 115, 123, 3, 2, 2,
	2, 116, 117, 12, 6, 2, 2, 117, 118, 7, 12, 2, 2, 118, 119, 5, 10, 6, 2,
	119, 120, 5, 8, 5, 7, 120, 121, 8, 5, 1, 2, 121, 123, 3, 2, 2, 2, 122,
	104, 3, 2, 2, 2, 122, 110, 3, 2, 2, 2, 122, 116, 3, 2, 2, 2, 123, 126,
	3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 9, 3, 2, 2,
	2, 126, 124, 3, 2, 2, 2, 127, 128, 8, 6, 1, 2, 128, 11, 3, 2, 2, 2, 129,
	130, 8, 7, 1, 2, 130, 131, 7, 5, 2, 2, 131, 132, 5, 12, 7, 16, 132, 133,
	8, 7, 1, 2, 133, 156, 3, 2, 2, 2, 134, 135, 7, 6, 2, 2, 135, 136, 5, 12,
	7, 15, 136, 137, 8, 7, 1, 2, 137, 156, 3, 2, 2, 2, 138, 139, 7, 9, 2, 2,
	139, 140, 5, 12, 7, 14, 140, 141, 8, 7, 1, 2, 141, 156, 3, 2, 2, 2, 142,
	143, 7, 7, 2, 2, 143, 144, 5, 12, 7, 2, 144, 145, 7, 8, 2, 2, 145, 146,
	8, 7, 1, 2, 146, 156, 3, 2, 2, 2, 147, 148, 7, 27, 2, 2, 148, 156, 8, 7,
	1, 2, 149, 150, 7, 28, 2, 2, 150, 156, 8, 7, 1, 2, 151, 152, 7, 13, 2,
	2, 152, 156, 8, 7, 1, 2, 153, 154, 7, 14, 2, 2, 154, 156, 8, 7, 1, 2, 155,
	129, 3, 2, 2, 2, 155, 134, 3, 2, 2, 2, 155, 138, 3, 2, 2, 2, 155, 142,
	3, 2, 2, 2, 155, 147, 3, 2, 2, 2, 155, 149, 3, 2, 2, 2, 155, 151, 3, 2,
	2, 2, 155, 153, 3, 2, 2, 2, 156, 192, 3, 2, 2, 2, 157, 158, 12, 13, 2,
	2, 158, 159, 9, 4, 2, 2, 159, 160, 5, 12, 7, 14, 160, 161, 8, 7, 1, 2,
	161, 191, 3, 2, 2, 2, 162, 163, 12, 12, 2, 2, 163, 164, 9, 3, 2, 2, 164,
	165, 5, 12, 7, 13, 165, 166, 8, 7, 1, 2, 166, 191, 3, 2, 2, 2, 167, 168,
	12, 11, 2, 2, 168, 169, 5, 14, 8, 2, 169, 170, 5, 12, 7, 12, 170, 171,
	8, 7, 1, 2, 171, 191, 3, 2, 2, 2, 172, 173, 12, 10, 2, 2, 173, 174, 7,
	10, 2, 2, 174, 175, 5, 10, 6, 2, 175, 176, 5, 12, 7, 11, 176, 177, 8, 7,
	1, 2, 177, 191, 3, 2, 2, 2, 178, 179, 12, 9, 2, 2, 179, 180, 7, 11, 2,
	2, 180, 181, 5, 10, 6, 2, 181, 182, 5, 12, 7, 10, 182, 183, 8, 7, 1, 2,
	183, 191, 3, 2, 2, 2, 184, 185, 12, 8, 2, 2, 185, 186, 7, 12, 2, 2, 186,
	187, 5, 10, 6, 2, 187, 188, 5, 12, 7, 9, 188, 189, 8, 7, 1, 2, 189, 191,
	3, 2, 2, 2, 190, 157, 3, 2, 2, 2, 190, 162, 3, 2, 2, 2, 190, 167, 3, 2,
	2, 2, 190, 172, 3, 2, 2, 2, 190, 178, 3, 2, 2, 2, 190, 184, 3, 2, 2, 2,
	191, 194, 3, 2, 2, 2, 192, 190, 3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193,
	13, 3, 2, 2, 2, 194, 192, 3, 2, 2, 2, 195, 196, 7, 16, 2, 2, 196, 208,
	8, 8, 1, 2, 197, 198, 7, 17, 2, 2, 198, 208, 8, 8, 1, 2, 199, 200, 7, 18,
	2, 2, 200, 208, 8, 8, 1, 2, 201, 202, 7, 19, 2, 2, 202, 208, 8, 8, 1, 2,
	203, 204, 7, 20, 2, 2, 204, 208, 8, 8, 1, 2, 205, 206, 7, 21, 2, 2, 206,
	208, 8, 8, 1, 2, 207, 195, 3, 2, 2, 2, 207, 197, 3, 2, 2, 2, 207, 199,
	3, 2, 2, 2, 207, 201, 3, 2, 2, 2, 207, 203, 3, 2, 2, 2, 207, 205, 3, 2,
	2, 2, 208, 15, 3, 2, 2, 2, 209, 211, 5, 18, 10, 2, 210, 209, 3, 2, 2, 2,
	211, 214, 3, 2, 2, 2, 212, 210, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213,
	17, 3, 2, 2, 2, 214, 212, 3, 2, 2, 2, 215, 218, 5, 20, 11, 2, 216, 218,
	5, 22, 12, 2, 217, 215, 3, 2, 2, 2, 217, 216, 3, 2, 2, 2, 218, 19, 3, 2,
	2, 2, 219, 220, 7, 22, 2, 2, 220, 221, 5, 12, 7, 2, 221, 222, 8, 11, 1,
	2, 222, 223, 5, 24, 13, 2, 223, 233, 8, 11, 1, 2, 224, 225, 7, 23, 2, 2,
	225, 226, 7, 22, 2, 2, 226, 227, 5, 12, 7, 2, 227, 228, 8, 11, 1, 2, 228,
	229, 5, 24, 13, 2, 229, 230, 8, 11, 1, 2, 230, 232, 3, 2, 2, 2, 231, 224,
	3, 2, 2, 2, 232, 235, 3, 2, 2, 2, 233, 231, 3, 2, 2, 2, 233, 234, 3, 2,
	2, 2, 234, 238, 3, 2, 2, 2, 235, 233, 3, 2, 2, 2, 236, 237, 7, 23, 2, 2,
	237, 239, 5, 24, 13, 2, 238, 236, 3, 2, 2, 2, 238, 239, 3, 2, 2, 2, 239,
	240, 3, 2, 2, 2, 240, 241, 8, 11, 1, 2, 241, 21, 3, 2, 2, 2, 242, 243,
	7, 28, 2, 2, 243, 244, 7, 24, 2, 2, 244, 245, 5, 12, 7, 2, 245, 246, 8,
	12, 1, 2, 246, 23, 3, 2, 2, 2, 247, 248, 7, 25, 2, 2, 248, 249, 5, 16,
	9, 2, 249, 250, 7, 26, 2, 2, 250, 25, 3, 2, 2, 2, 19, 39, 51, 53, 66, 78,
	80, 102, 122, 124, 155, 190, 192, 207, 212, 217, 233, 238,
}
var literalNames = []string{
	"", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'!'", "'&&'", "'^'", "'||'",
	"'true'", "'false'", "'%'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='",
	"'if'", "'else'", "'='", "'{'", "'}'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "NUM", "ID", "COMMENT", "WHITESPACE",
}

var ruleNames = []string{
	"start", "exp", "exp2", "cond", "marcador", "expresion", "oprel", "instrucciones",
	"instruccion", "inst_if", "inst_asignacion", "bloque",
}

type GramaticaParser struct {
	*antlr.BaseParser
}

// NewGramaticaParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *GramaticaParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewGramaticaParser(input antlr.TokenStream) *GramaticaParser {
	this := new(GramaticaParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Gramatica.g4"

	return this
}

// posición relativa de un símbolo
var desp int = 0

// entorno actual
var tope *entorno.Entorno

// temporal
var tmp int = 1

// GramaticaParser tokens.
const (
	GramaticaParserEOF        = antlr.TokenEOF
	GramaticaParserT__0       = 1
	GramaticaParserT__1       = 2
	GramaticaParserT__2       = 3
	GramaticaParserT__3       = 4
	GramaticaParserT__4       = 5
	GramaticaParserT__5       = 6
	GramaticaParserT__6       = 7
	GramaticaParserT__7       = 8
	GramaticaParserT__8       = 9
	GramaticaParserT__9       = 10
	GramaticaParserT__10      = 11
	GramaticaParserT__11      = 12
	GramaticaParserT__12      = 13
	GramaticaParserT__13      = 14
	GramaticaParserT__14      = 15
	GramaticaParserT__15      = 16
	GramaticaParserT__16      = 17
	GramaticaParserT__17      = 18
	GramaticaParserT__18      = 19
	GramaticaParserT__19      = 20
	GramaticaParserT__20      = 21
	GramaticaParserT__21      = 22
	GramaticaParserT__22      = 23
	GramaticaParserT__23      = 24
	GramaticaParserNUM        = 25
	GramaticaParserID         = 26
	GramaticaParserCOMMENT    = 27
	GramaticaParserWHITESPACE = 28
)

// GramaticaParser rules.
const (
	GramaticaParserRULE_start           = 0
	GramaticaParserRULE_exp             = 1
	GramaticaParserRULE_exp2            = 2
	GramaticaParserRULE_cond            = 3
	GramaticaParserRULE_marcador        = 4
	GramaticaParserRULE_expresion       = 5
	GramaticaParserRULE_oprel           = 6
	GramaticaParserRULE_instrucciones   = 7
	GramaticaParserRULE_instruccion     = 8
	GramaticaParserRULE_inst_if         = 9
	GramaticaParserRULE_inst_asignacion = 10
	GramaticaParserRULE_bloque          = 11
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *GramaticaParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GramaticaParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Instrucciones()
	}
	{
		p.SetState(25)
		p.Match(GramaticaParserEOF)
	}

	return localctx
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetN returns the n token.
	GetN() antlr.Token

	// GetId returns the id token.
	GetId() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetN sets the n token.
	SetN(antlr.Token)

	// SetId sets the id token.
	SetId(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpContext

	// GetE returns the e rule contexts.
	GetE() IExpContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpContext)

	// SetE sets the e rule contexts.
	SetE(IExpContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpContext)

	// GetDir returns the dir attribute.
	GetDir() string

	// SetDir sets the dir attribute.
	SetDir(string)

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	dir    string
	e1     IExpContext
	e      IExpContext
	n      antlr.Token
	id     antlr.Token
	op     antlr.Token
	e2     IExpContext
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) GetN() antlr.Token { return s.n }

func (s *ExpContext) GetId() antlr.Token { return s.id }

func (s *ExpContext) GetOp() antlr.Token { return s.op }

func (s *ExpContext) SetN(v antlr.Token) { s.n = v }

func (s *ExpContext) SetId(v antlr.Token) { s.id = v }

func (s *ExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpContext) GetE1() IExpContext { return s.e1 }

func (s *ExpContext) GetE() IExpContext { return s.e }

func (s *ExpContext) GetE2() IExpContext { return s.e2 }

func (s *ExpContext) SetE1(v IExpContext) { s.e1 = v }

func (s *ExpContext) SetE(v IExpContext) { s.e = v }

func (s *ExpContext) SetE2(v IExpContext) { s.e2 = v }

func (s *ExpContext) GetDir() string { return s.dir }

func (s *ExpContext) SetDir(v string) { s.dir = v }

func (s *ExpContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *ExpContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpContext) NUM() antlr.TerminalNode {
	return s.GetToken(GramaticaParserNUM, 0)
}

func (s *ExpContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitExp(s)
	}
}

func (p *GramaticaParser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *GramaticaParser) exp(_p int) (localctx IExpContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, GramaticaParserRULE_exp, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserT__4:
		{
			p.SetState(28)
			p.Match(GramaticaParserT__4)
		}
		{
			p.SetState(29)

			var _x = p.exp(0)

			localctx.(*ExpContext).e = _x
		}
		{
			p.SetState(30)
			p.Match(GramaticaParserT__5)
		}

		localctx.(*ExpContext).dir = localctx.(*ExpContext).GetE().GetDir()

	case GramaticaParserNUM:
		{
			p.SetState(33)

			var _m = p.Match(GramaticaParserNUM)

			localctx.(*ExpContext).n = _m
		}
		localctx.(*ExpContext).dir = (func() string {
			if localctx.(*ExpContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*ExpContext).GetN().GetText()
			}
		}())

	case GramaticaParserID:
		{
			p.SetState(35)

			var _m = p.Match(GramaticaParserID)

			localctx.(*ExpContext).id = _m
		}
		localctx.(*ExpContext).dir = (func() string {
			if localctx.(*ExpContext).GetId() == nil {
				return ""
			} else {
				return localctx.(*ExpContext).GetId().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(49)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpContext(p, _parentctx, _parentState)
				localctx.(*ExpContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_exp)
				p.SetState(39)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(40)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserT__0 || _la == GramaticaParserT__1) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(41)

					var _x = p.exp(6)

					localctx.(*ExpContext).e2 = _x
				}

				localctx.(*ExpContext).dir = gen.NewTemp()
				gen.Gen(localctx.(*ExpContext).dir + " = " + localctx.(*ExpContext).GetE1().GetDir() + " " + (func() string {
					if localctx.(*ExpContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpContext).GetOp().GetText()
					}
				}()) + " " + localctx.(*ExpContext).GetE2().GetDir())

			case 2:
				localctx = NewExpContext(p, _parentctx, _parentState)
				localctx.(*ExpContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_exp)
				p.SetState(44)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(45)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserT__2 || _la == GramaticaParserT__3) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(46)

					var _x = p.exp(5)

					localctx.(*ExpContext).e2 = _x
				}

				localctx.(*ExpContext).dir = gen.NewTemp()
				gen.Gen(localctx.(*ExpContext).dir + " = " + localctx.(*ExpContext).GetE1().GetDir() + " " + (func() string {
					if localctx.(*ExpContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpContext).GetOp().GetText()
					}
				}()) + " " + localctx.(*ExpContext).GetE2().GetDir())

			}

		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IExp2Context is an interface to support dynamic dispatch.
type IExp2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetN returns the n token.
	GetN() antlr.Token

	// GetId returns the id token.
	GetId() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetN sets the n token.
	SetN(antlr.Token)

	// SetId sets the id token.
	SetId(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetE1 returns the e1 rule contexts.
	GetE1() IExp2Context

	// GetE returns the e rule contexts.
	GetE() IExp2Context

	// GetE2 returns the e2 rule contexts.
	GetE2() IExp2Context

	// SetE1 sets the e1 rule contexts.
	SetE1(IExp2Context)

	// SetE sets the e rule contexts.
	SetE(IExp2Context)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExp2Context)

	// GetDir returns the dir attribute.
	GetDir() string

	// GetNum returns the num attribute.
	GetNum() int

	// SetDir sets the dir attribute.
	SetDir(string)

	// SetNum sets the num attribute.
	SetNum(int)

	// IsExp2Context differentiates from other interfaces.
	IsExp2Context()
}

type Exp2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	dir    string
	num    int
	e1     IExp2Context
	e      IExp2Context
	n      antlr.Token
	id     antlr.Token
	op     antlr.Token
	e2     IExp2Context
}

func NewEmptyExp2Context() *Exp2Context {
	var p = new(Exp2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_exp2
	return p
}

func (*Exp2Context) IsExp2Context() {}

func NewExp2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp2Context {
	var p = new(Exp2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_exp2

	return p
}

func (s *Exp2Context) GetParser() antlr.Parser { return s.parser }

func (s *Exp2Context) GetN() antlr.Token { return s.n }

func (s *Exp2Context) GetId() antlr.Token { return s.id }

func (s *Exp2Context) GetOp() antlr.Token { return s.op }

func (s *Exp2Context) SetN(v antlr.Token) { s.n = v }

func (s *Exp2Context) SetId(v antlr.Token) { s.id = v }

func (s *Exp2Context) SetOp(v antlr.Token) { s.op = v }

func (s *Exp2Context) GetE1() IExp2Context { return s.e1 }

func (s *Exp2Context) GetE() IExp2Context { return s.e }

func (s *Exp2Context) GetE2() IExp2Context { return s.e2 }

func (s *Exp2Context) SetE1(v IExp2Context) { s.e1 = v }

func (s *Exp2Context) SetE(v IExp2Context) { s.e = v }

func (s *Exp2Context) SetE2(v IExp2Context) { s.e2 = v }

func (s *Exp2Context) GetDir() string { return s.dir }

func (s *Exp2Context) GetNum() int { return s.num }

func (s *Exp2Context) SetDir(v string) { s.dir = v }

func (s *Exp2Context) SetNum(v int) { s.num = v }

func (s *Exp2Context) AllExp2() []IExp2Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExp2Context)(nil)).Elem())
	var tst = make([]IExp2Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExp2Context)
		}
	}

	return tst
}

func (s *Exp2Context) Exp2(i int) IExp2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp2Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExp2Context)
}

func (s *Exp2Context) NUM() antlr.TerminalNode {
	return s.GetToken(GramaticaParserNUM, 0)
}

func (s *Exp2Context) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *Exp2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterExp2(s)
	}
}

func (s *Exp2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitExp2(s)
	}
}

func (p *GramaticaParser) Exp2() (localctx IExp2Context) {
	return p.exp2(0)
}

func (p *GramaticaParser) exp2(_p int) (localctx IExp2Context) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExp2Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExp2Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, GramaticaParserRULE_exp2, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(64)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserT__4:
		{
			p.SetState(55)
			p.Match(GramaticaParserT__4)
		}
		{
			p.SetState(56)

			var _x = p.exp2(0)

			localctx.(*Exp2Context).e = _x
		}
		{
			p.SetState(57)
			p.Match(GramaticaParserT__5)
		}

		localctx.(*Exp2Context).num = localctx.(*Exp2Context).GetE().GetNum()
		localctx.(*Exp2Context).dir = localctx.(*Exp2Context).GetE().GetDir()

	case GramaticaParserNUM:
		{
			p.SetState(60)

			var _m = p.Match(GramaticaParserNUM)

			localctx.(*Exp2Context).n = _m
		}

		localctx.(*Exp2Context).num = 0
		localctx.(*Exp2Context).dir = (func() string {
			if localctx.(*Exp2Context).GetN() == nil {
				return ""
			} else {
				return localctx.(*Exp2Context).GetN().GetText()
			}
		}())

	case GramaticaParserID:
		{
			p.SetState(62)

			var _m = p.Match(GramaticaParserID)

			localctx.(*Exp2Context).id = _m
		}

		localctx.(*Exp2Context).num = 0
		localctx.(*Exp2Context).dir = (func() string {
			if localctx.(*Exp2Context).GetId() == nil {
				return ""
			} else {
				return localctx.(*Exp2Context).GetId().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(76)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExp2Context(p, _parentctx, _parentState)
				localctx.(*Exp2Context).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_exp2)
				p.SetState(66)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(67)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp2Context).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserT__0 || _la == GramaticaParserT__1) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp2Context).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(68)

					var _x = p.exp2(6)

					localctx.(*Exp2Context).e2 = _x
				}

				tmp = tmp - localctx.(*Exp2Context).GetE1().GetNum() - localctx.(*Exp2Context).GetE2().GetNum()
				localctx.(*Exp2Context).num = 1
				localctx.(*Exp2Context).dir = "t" + strconv.Itoa(tmp)
				gen.Gen(localctx.(*Exp2Context).dir + " = " + localctx.(*Exp2Context).GetE1().GetDir() + " " + (func() string {
					if localctx.(*Exp2Context).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Exp2Context).GetOp().GetText()
					}
				}()) + " " + localctx.(*Exp2Context).GetE2().GetDir())
				tmp = tmp + 1

			case 2:
				localctx = NewExp2Context(p, _parentctx, _parentState)
				localctx.(*Exp2Context).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_exp2)
				p.SetState(71)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(72)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Exp2Context).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserT__2 || _la == GramaticaParserT__3) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Exp2Context).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(73)

					var _x = p.exp2(5)

					localctx.(*Exp2Context).e2 = _x
				}

				tmp = tmp - localctx.(*Exp2Context).GetE1().GetNum() - localctx.(*Exp2Context).GetE2().GetNum()
				localctx.(*Exp2Context).num = 1
				localctx.(*Exp2Context).dir = "t" + strconv.Itoa(tmp)
				gen.Gen(localctx.(*Exp2Context).dir + " = " + localctx.(*Exp2Context).GetE1().GetDir() + " " + (func() string {
					if localctx.(*Exp2Context).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Exp2Context).GetOp().GetText()
					}
				}()) + " " + localctx.(*Exp2Context).GetE2().GetDir())
				tmp = tmp + 1

			}

		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// ICondContext is an interface to support dynamic dispatch.
type ICondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetC1 returns the c1 rule contexts.
	GetC1() ICondContext

	// GetC returns the c rule contexts.
	GetC() ICondContext

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpContext

	// GetOpr returns the opr rule contexts.
	GetOpr() IOprelContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpContext

	// GetC2 returns the c2 rule contexts.
	GetC2() ICondContext

	// SetC1 sets the c1 rule contexts.
	SetC1(ICondContext)

	// SetC sets the c rule contexts.
	SetC(ICondContext)

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpContext)

	// SetOpr sets the opr rule contexts.
	SetOpr(IOprelContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpContext)

	// SetC2 sets the c2 rule contexts.
	SetC2(ICondContext)

	// GetLv returns the lv attribute.
	GetLv() []string

	// GetLf returns the lf attribute.
	GetLf() []string

	// GetCad returns the cad attribute.
	GetCad() string

	// SetLv sets the lv attribute.
	SetLv([]string)

	// SetLf sets the lf attribute.
	SetLf([]string)

	// SetCad sets the cad attribute.
	SetCad(string)

	// IsCondContext differentiates from other interfaces.
	IsCondContext()
}

type CondContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lv     []string
	lf     []string
	cad    string
	c1     ICondContext
	op     antlr.Token
	c      ICondContext
	e1     IExpContext
	opr    IOprelContext
	e2     IExpContext
	c2     ICondContext
}

func NewEmptyCondContext() *CondContext {
	var p = new(CondContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_cond
	return p
}

func (*CondContext) IsCondContext() {}

func NewCondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondContext {
	var p = new(CondContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_cond

	return p
}

func (s *CondContext) GetParser() antlr.Parser { return s.parser }

func (s *CondContext) GetOp() antlr.Token { return s.op }

func (s *CondContext) SetOp(v antlr.Token) { s.op = v }

func (s *CondContext) GetC1() ICondContext { return s.c1 }

func (s *CondContext) GetC() ICondContext { return s.c }

func (s *CondContext) GetE1() IExpContext { return s.e1 }

func (s *CondContext) GetOpr() IOprelContext { return s.opr }

func (s *CondContext) GetE2() IExpContext { return s.e2 }

func (s *CondContext) GetC2() ICondContext { return s.c2 }

func (s *CondContext) SetC1(v ICondContext) { s.c1 = v }

func (s *CondContext) SetC(v ICondContext) { s.c = v }

func (s *CondContext) SetE1(v IExpContext) { s.e1 = v }

func (s *CondContext) SetOpr(v IOprelContext) { s.opr = v }

func (s *CondContext) SetE2(v IExpContext) { s.e2 = v }

func (s *CondContext) SetC2(v ICondContext) { s.c2 = v }

func (s *CondContext) GetLv() []string { return s.lv }

func (s *CondContext) GetLf() []string { return s.lf }

func (s *CondContext) GetCad() string { return s.cad }

func (s *CondContext) SetLv(v []string) { s.lv = v }

func (s *CondContext) SetLf(v []string) { s.lf = v }

func (s *CondContext) SetCad(v string) { s.cad = v }

func (s *CondContext) AllCond() []ICondContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICondContext)(nil)).Elem())
	var tst = make([]ICondContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICondContext)
		}
	}

	return tst
}

func (s *CondContext) Cond(i int) ICondContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICondContext)
}

func (s *CondContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *CondContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *CondContext) Oprel() IOprelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOprelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOprelContext)
}

func (s *CondContext) Marcador() IMarcadorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMarcadorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMarcadorContext)
}

func (s *CondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterCond(s)
	}
}

func (s *CondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitCond(s)
	}
}

func (p *GramaticaParser) Cond() (localctx ICondContext) {
	return p.cond(0)
}

func (p *GramaticaParser) cond(_p int) (localctx ICondContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewCondContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICondContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, GramaticaParserRULE_cond, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(82)

			var _m = p.Match(GramaticaParserT__6)

			localctx.(*CondContext).op = _m
		}
		{
			p.SetState(83)

			var _x = p.cond(8)

			localctx.(*CondContext).c = _x
		}

		localctx.(*CondContext).lv = localctx.(*CondContext).GetC().GetLf()
		localctx.(*CondContext).lf = localctx.(*CondContext).GetC().GetLv()

	case 2:
		{
			p.SetState(86)

			var _x = p.exp(0)

			localctx.(*CondContext).e1 = _x
		}
		{
			p.SetState(87)

			var _x = p.Oprel()

			localctx.(*CondContext).opr = _x
		}
		{
			p.SetState(88)

			var _x = p.exp(0)

			localctx.(*CondContext).e2 = _x
		}

		localctx.(*CondContext).lv = append(localctx.(*CondContext).lv, gen.NewEti())
		localctx.(*CondContext).lf = append(localctx.(*CondContext).lf, gen.NewEti())
		localctx.(*CondContext).cad = localctx.(*CondContext).GetE1().GetDir() + " " + localctx.(*CondContext).GetOpr().GetOp() + " " + localctx.(*CondContext).GetE2().GetDir()
		gen.Gen("if " + localctx.(*CondContext).cad + " then goto " + localctx.(*CondContext).lv[0])
		gen.Gen("goto " + localctx.(*CondContext).lf[0])

	case 3:
		{
			p.SetState(91)
			p.Match(GramaticaParserT__4)
		}
		{
			p.SetState(92)

			var _x = p.cond(0)

			localctx.(*CondContext).c = _x
		}
		{
			p.SetState(93)
			p.Match(GramaticaParserT__5)
		}

		localctx.(*CondContext).lv = localctx.(*CondContext).GetC().GetLv()
		localctx.(*CondContext).lf = localctx.(*CondContext).GetC().GetLf()

	case 4:
		{
			p.SetState(96)
			p.Match(GramaticaParserT__10)
		}

		localctx.(*CondContext).lv = append(localctx.(*CondContext).lv, gen.NewEti())
		localctx.(*CondContext).lf = append(localctx.(*CondContext).lf, gen.NewEti())
		gen.Gen("goto " + localctx.(*CondContext).lv[0])

	case 5:
		{
			p.SetState(98)
			p.Match(GramaticaParserT__11)
		}

		localctx.(*CondContext).lv = append(localctx.(*CondContext).lv, gen.NewEti())
		localctx.(*CondContext).lf = append(localctx.(*CondContext).lf, gen.NewEti())
		gen.Gen("goto " + localctx.(*CondContext).lf[0])

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewCondContext(p, _parentctx, _parentState)
				localctx.(*CondContext).c1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_cond)
				p.SetState(102)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(103)

					var _m = p.Match(GramaticaParserT__7)

					localctx.(*CondContext).op = _m
				}
				{
					p.SetState(104)
					p.Marcador(localctx.(*CondContext).GetC1().GetLv())
				}
				{
					p.SetState(105)

					var _x = p.cond(7)

					localctx.(*CondContext).c2 = _x
				}

				localctx.(*CondContext).lv = localctx.(*CondContext).GetC2().GetLv()
				localctx.(*CondContext).lf = gen.Unir(localctx.(*CondContext).GetC1().GetLf(), localctx.(*CondContext).GetC2().GetLf())

			case 2:
				localctx = NewCondContext(p, _parentctx, _parentState)
				localctx.(*CondContext).c1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_cond)
				p.SetState(108)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(109)

					var _m = p.Match(GramaticaParserT__8)

					localctx.(*CondContext).op = _m
				}
				{
					p.SetState(110)
					p.Marcador(localctx.(*CondContext).GetC1().GetLf())
				}
				{
					p.SetState(111)

					var _x = p.cond(6)

					localctx.(*CondContext).c2 = _x
				}

				gen.Soltar(localctx.(*CondContext).GetC1().GetLv())
				gen.Gen("if " + localctx.(*CondContext).GetC2().GetCad() + " then goto " + localctx.(*CondContext).GetC2().GetLf()[0])
				gen.Gen("goto " + localctx.(*CondContext).GetC2().GetLv()[0])
				localctx.(*CondContext).lv = localctx.(*CondContext).GetC2().GetLv()
				localctx.(*CondContext).lf = localctx.(*CondContext).GetC2().GetLf()

			case 3:
				localctx = NewCondContext(p, _parentctx, _parentState)
				localctx.(*CondContext).c1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_cond)
				p.SetState(114)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(115)

					var _m = p.Match(GramaticaParserT__9)

					localctx.(*CondContext).op = _m
				}
				{
					p.SetState(116)
					p.Marcador(localctx.(*CondContext).GetC1().GetLf())
				}
				{
					p.SetState(117)

					var _x = p.cond(5)

					localctx.(*CondContext).c2 = _x
				}

				localctx.(*CondContext).lf = localctx.(*CondContext).GetC2().GetLf()
				localctx.(*CondContext).lv = gen.Unir(localctx.(*CondContext).GetC1().GetLv(), localctx.(*CondContext).GetC2().GetLv())

			}

		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IMarcadorContext is an interface to support dynamic dispatch.
type IMarcadorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLista returns the lista attribute.
	GetLista() []string

	// SetLista sets the lista attribute.
	SetLista([]string)

	// IsMarcadorContext differentiates from other interfaces.
	IsMarcadorContext()
}

type MarcadorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lista  []string
}

func NewEmptyMarcadorContext() *MarcadorContext {
	var p = new(MarcadorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_marcador
	return p
}

func (*MarcadorContext) IsMarcadorContext() {}

func NewMarcadorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, lista []string) *MarcadorContext {
	var p = new(MarcadorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_marcador

	p.lista = lista

	return p
}

func (s *MarcadorContext) GetParser() antlr.Parser { return s.parser }

func (s *MarcadorContext) GetLista() []string { return s.lista }

func (s *MarcadorContext) SetLista(v []string) { s.lista = v }

func (s *MarcadorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MarcadorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MarcadorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterMarcador(s)
	}
}

func (s *MarcadorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitMarcador(s)
	}
}

func (p *GramaticaParser) Marcador(lista []string) (localctx IMarcadorContext) {
	this := p
	_ = this

	localctx = NewMarcadorContext(p, p.GetParserRuleContext(), p.GetState(), lista)
	p.EnterRule(localctx, 8, GramaticaParserRULE_marcador)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	gen.Soltar(lista)

	return localctx
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// GetN returns the n token.
	GetN() antlr.Token

	// GetId returns the id token.
	GetId() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// SetN sets the n token.
	SetN(antlr.Token)

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpresionContext

	// GetE returns the e rule contexts.
	GetE() IExpresionContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpresionContext

	// GetOpr returns the opr rule contexts.
	GetOpr() IOprelContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// SetE sets the e rule contexts.
	SetE(IExpresionContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpresionContext)

	// SetOpr sets the opr rule contexts.
	SetOpr(IOprelContext)

	// GetDir returns the dir attribute.
	GetDir() string

	// GetLv returns the lv attribute.
	GetLv() []string

	// GetLf returns the lf attribute.
	GetLf() []string

	// GetCad returns the cad attribute.
	GetCad() string

	// SetDir sets the dir attribute.
	SetDir(string)

	// SetLv sets the lv attribute.
	SetLv([]string)

	// SetLf sets the lf attribute.
	SetLf([]string)

	// SetCad sets the cad attribute.
	SetCad(string)

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	dir    string
	lv     []string
	lf     []string
	cad    string
	e1     IExpresionContext
	op     antlr.Token
	e      IExpresionContext
	n      antlr.Token
	id     antlr.Token
	e2     IExpresionContext
	opr    IOprelContext
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_expresion
	return p
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) GetOp() antlr.Token { return s.op }

func (s *ExpresionContext) GetN() antlr.Token { return s.n }

func (s *ExpresionContext) GetId() antlr.Token { return s.id }

func (s *ExpresionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpresionContext) SetN(v antlr.Token) { s.n = v }

func (s *ExpresionContext) SetId(v antlr.Token) { s.id = v }

func (s *ExpresionContext) GetE1() IExpresionContext { return s.e1 }

func (s *ExpresionContext) GetE() IExpresionContext { return s.e }

func (s *ExpresionContext) GetE2() IExpresionContext { return s.e2 }

func (s *ExpresionContext) GetOpr() IOprelContext { return s.opr }

func (s *ExpresionContext) SetE1(v IExpresionContext) { s.e1 = v }

func (s *ExpresionContext) SetE(v IExpresionContext) { s.e = v }

func (s *ExpresionContext) SetE2(v IExpresionContext) { s.e2 = v }

func (s *ExpresionContext) SetOpr(v IOprelContext) { s.opr = v }

func (s *ExpresionContext) GetDir() string { return s.dir }

func (s *ExpresionContext) GetLv() []string { return s.lv }

func (s *ExpresionContext) GetLf() []string { return s.lf }

func (s *ExpresionContext) GetCad() string { return s.cad }

func (s *ExpresionContext) SetDir(v string) { s.dir = v }

func (s *ExpresionContext) SetLv(v []string) { s.lv = v }

func (s *ExpresionContext) SetLf(v []string) { s.lf = v }

func (s *ExpresionContext) SetCad(v string) { s.cad = v }

func (s *ExpresionContext) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *ExpresionContext) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpresionContext) NUM() antlr.TerminalNode {
	return s.GetToken(GramaticaParserNUM, 0)
}

func (s *ExpresionContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *ExpresionContext) Oprel() IOprelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOprelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOprelContext)
}

func (s *ExpresionContext) Marcador() IMarcadorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMarcadorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMarcadorContext)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterExpresion(s)
	}
}

func (s *ExpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitExpresion(s)
	}
}

func (p *GramaticaParser) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *GramaticaParser) expresion(_p int) (localctx IExpresionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, GramaticaParserRULE_expresion, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(153)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserT__2:
		{
			p.SetState(128)

			var _m = p.Match(GramaticaParserT__2)

			localctx.(*ExpresionContext).op = _m
		}
		{
			p.SetState(129)

			var _x = p.expresion(14)

			localctx.(*ExpresionContext).e = _x
		}

		localctx.(*ExpresionContext).dir = gen.NewTemp()
		gen.Gen(localctx.(*ExpresionContext).dir + " = + " + localctx.(*ExpresionContext).GetE().GetDir())

	case GramaticaParserT__3:
		{
			p.SetState(132)

			var _m = p.Match(GramaticaParserT__3)

			localctx.(*ExpresionContext).op = _m
		}
		{
			p.SetState(133)

			var _x = p.expresion(13)

			localctx.(*ExpresionContext).e = _x
		}

		localctx.(*ExpresionContext).dir = gen.NewTemp()
		gen.Gen(localctx.(*ExpresionContext).dir + " = - " + localctx.(*ExpresionContext).GetE().GetDir())

	case GramaticaParserT__6:
		{
			p.SetState(136)

			var _m = p.Match(GramaticaParserT__6)

			localctx.(*ExpresionContext).op = _m
		}
		{
			p.SetState(137)

			var _x = p.expresion(12)

			localctx.(*ExpresionContext).e = _x
		}

		localctx.(*ExpresionContext).lv = localctx.(*ExpresionContext).GetE().GetLf()
		localctx.(*ExpresionContext).lf = localctx.(*ExpresionContext).GetE().GetLv()

	case GramaticaParserT__4:
		{
			p.SetState(140)
			p.Match(GramaticaParserT__4)
		}
		{
			p.SetState(141)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e = _x
		}
		{
			p.SetState(142)
			p.Match(GramaticaParserT__5)
		}

		localctx.(*ExpresionContext).dir = localctx.(*ExpresionContext).GetE().GetDir()
		localctx.(*ExpresionContext).lv = localctx.(*ExpresionContext).GetE().GetLv()
		localctx.(*ExpresionContext).lf = localctx.(*ExpresionContext).GetE().GetLf()
		localctx.(*ExpresionContext).cad = localctx.(*ExpresionContext).GetE().GetCad()

	case GramaticaParserNUM:
		{
			p.SetState(145)

			var _m = p.Match(GramaticaParserNUM)

			localctx.(*ExpresionContext).n = _m
		}

		localctx.(*ExpresionContext).dir = (func() string {
			if localctx.(*ExpresionContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).GetN().GetText()
			}
		}())
		localctx.(*ExpresionContext).cad = ""

	case GramaticaParserID:
		{
			p.SetState(147)

			var _m = p.Match(GramaticaParserID)

			localctx.(*ExpresionContext).id = _m
		}

		localctx.(*ExpresionContext).dir = (func() string {
			if localctx.(*ExpresionContext).GetId() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).GetId().GetText()
			}
		}())
		localctx.(*ExpresionContext).cad = ""

	case GramaticaParserT__10:
		{
			p.SetState(149)
			p.Match(GramaticaParserT__10)
		}

		localctx.(*ExpresionContext).dir = ""
		localctx.(*ExpresionContext).cad = ""
		localctx.(*ExpresionContext).lv = append(localctx.(*ExpresionContext).lv, gen.NewEti())
		localctx.(*ExpresionContext).lf = append(localctx.(*ExpresionContext).lf, gen.NewEti())
		gen.Gen("goto " + localctx.(*ExpresionContext).lv[0])

	case GramaticaParserT__11:
		{
			p.SetState(151)
			p.Match(GramaticaParserT__11)
		}

		localctx.(*ExpresionContext).dir = ""
		localctx.(*ExpresionContext).cad = ""
		localctx.(*ExpresionContext).lv = append(localctx.(*ExpresionContext).lv, gen.NewEti())
		localctx.(*ExpresionContext).lf = append(localctx.(*ExpresionContext).lf, gen.NewEti())
		gen.Gen("goto " + localctx.(*ExpresionContext).lf[0])

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(188)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(155)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(156)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GramaticaParserT__0)|(1<<GramaticaParserT__1)|(1<<GramaticaParserT__12))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(157)

					var _x = p.expresion(12)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).dir = gen.NewTemp()
				gen.Gen(localctx.(*ExpresionContext).dir + " = " + localctx.(*ExpresionContext).GetE1().GetDir() + " " + (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()) + " " + localctx.(*ExpresionContext).GetE2().GetDir())

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(160)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(161)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserT__2 || _la == GramaticaParserT__3) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(162)

					var _x = p.expresion(11)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).dir = gen.NewTemp()
				gen.Gen(localctx.(*ExpresionContext).dir + " = " + localctx.(*ExpresionContext).GetE1().GetDir() + " " + (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()) + " " + localctx.(*ExpresionContext).GetE2().GetDir())

			case 3:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(165)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(166)

					var _x = p.Oprel()

					localctx.(*ExpresionContext).opr = _x
				}
				{
					p.SetState(167)

					var _x = p.expresion(10)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).lv = append(localctx.(*ExpresionContext).lv, gen.NewEti())
				localctx.(*ExpresionContext).lf = append(localctx.(*ExpresionContext).lf, gen.NewEti())
				localctx.(*ExpresionContext).cad = localctx.(*ExpresionContext).GetE1().GetDir() + " " + localctx.(*ExpresionContext).GetOpr().GetOp() + " " + localctx.(*ExpresionContext).GetE2().GetDir()
				gen.Gen("if " + localctx.(*ExpresionContext).cad + " then goto " + localctx.(*ExpresionContext).lv[0])
				gen.Gen("goto " + localctx.(*ExpresionContext).lf[0])

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(170)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(171)

					var _m = p.Match(GramaticaParserT__7)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(172)
					p.Marcador(localctx.(*ExpresionContext).GetE1().GetLv())
				}
				{
					p.SetState(173)

					var _x = p.expresion(9)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).lv = localctx.(*ExpresionContext).GetE2().GetLv()
				localctx.(*ExpresionContext).lf = gen.Unir(localctx.(*ExpresionContext).GetE1().GetLf(), localctx.(*ExpresionContext).GetE2().GetLf())

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(176)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(177)

					var _m = p.Match(GramaticaParserT__8)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(178)
					p.Marcador(localctx.(*ExpresionContext).GetE1().GetLf())
				}
				{
					p.SetState(179)

					var _x = p.expresion(8)

					localctx.(*ExpresionContext).e2 = _x
				}

				gen.Soltar(localctx.(*ExpresionContext).GetE1().GetLv())
				gen.Gen("if " + localctx.(*ExpresionContext).GetE2().GetCad() + " then goto " + localctx.(*ExpresionContext).GetE2().GetLf()[0])
				gen.Gen("goto " + localctx.(*ExpresionContext).GetE2().GetLv()[0])
				localctx.(*ExpresionContext).lv = localctx.(*ExpresionContext).GetE2().GetLv()
				localctx.(*ExpresionContext).lf = localctx.(*ExpresionContext).GetE2().GetLf()

			case 6:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(182)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(183)

					var _m = p.Match(GramaticaParserT__9)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(184)
					p.Marcador(localctx.(*ExpresionContext).GetE1().GetLf())
				}
				{
					p.SetState(185)

					var _x = p.expresion(7)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).lf = localctx.(*ExpresionContext).GetE2().GetLf()
				localctx.(*ExpresionContext).lv = gen.Unir(localctx.(*ExpresionContext).GetE1().GetLv(), localctx.(*ExpresionContext).GetE2().GetLv())

			}

		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IOprelContext is an interface to support dynamic dispatch.
type IOprelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOpe returns the ope token.
	GetOpe() antlr.Token

	// SetOpe sets the ope token.
	SetOpe(antlr.Token)

	// GetOp returns the op attribute.
	GetOp() string

	// SetOp sets the op attribute.
	SetOp(string)

	// IsOprelContext differentiates from other interfaces.
	IsOprelContext()
}

type OprelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     string
	ope    antlr.Token
}

func NewEmptyOprelContext() *OprelContext {
	var p = new(OprelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_oprel
	return p
}

func (*OprelContext) IsOprelContext() {}

func NewOprelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OprelContext {
	var p = new(OprelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_oprel

	return p
}

func (s *OprelContext) GetParser() antlr.Parser { return s.parser }

func (s *OprelContext) GetOpe() antlr.Token { return s.ope }

func (s *OprelContext) SetOpe(v antlr.Token) { s.ope = v }

func (s *OprelContext) GetOp() string { return s.op }

func (s *OprelContext) SetOp(v string) { s.op = v }

func (s *OprelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OprelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OprelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterOprel(s)
	}
}

func (s *OprelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitOprel(s)
	}
}

func (p *GramaticaParser) Oprel() (localctx IOprelContext) {
	this := p
	_ = this

	localctx = NewOprelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GramaticaParserRULE_oprel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(205)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserT__13:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(193)

			var _m = p.Match(GramaticaParserT__13)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	case GramaticaParserT__14:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(195)

			var _m = p.Match(GramaticaParserT__14)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	case GramaticaParserT__15:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(197)

			var _m = p.Match(GramaticaParserT__15)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	case GramaticaParserT__16:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(199)

			var _m = p.Match(GramaticaParserT__16)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	case GramaticaParserT__17:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(201)

			var _m = p.Match(GramaticaParserT__17)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	case GramaticaParserT__18:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(203)

			var _m = p.Match(GramaticaParserT__18)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_instrucciones
	return p
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionContext)(nil)).Elem())
	var tst = make([]IInstruccionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionContext)
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), i)

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
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (p *GramaticaParser) Instrucciones() (localctx IInstruccionesContext) {
	this := p
	_ = this

	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GramaticaParserRULE_instrucciones)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GramaticaParserT__19 || _la == GramaticaParserID {
		{
			p.SetState(207)
			p.Instruccion()
		}

		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_instruccion
	return p
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Inst_if() IInst_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_ifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_ifContext)
}

func (s *InstruccionContext) Inst_asignacion() IInst_asignacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_asignacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_asignacionContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (p *GramaticaParser) Instruccion() (localctx IInstruccionContext) {
	this := p
	_ = this

	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GramaticaParserRULE_instruccion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(215)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserT__19:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(213)
			p.Inst_if()
		}

	case GramaticaParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(214)
			p.Inst_asignacion()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInst_ifContext is an interface to support dynamic dispatch.
type IInst_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpresionContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpresionContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpresionContext)

	// GetLsalida returns the lsalida attribute.
	GetLsalida() string

	// SetLsalida sets the lsalida attribute.
	SetLsalida(string)

	// IsInst_ifContext differentiates from other interfaces.
	IsInst_ifContext()
}

type Inst_ifContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	lsalida string
	e1      IExpresionContext
	e2      IExpresionContext
}

func NewEmptyInst_ifContext() *Inst_ifContext {
	var p = new(Inst_ifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_if
	return p
}

func (*Inst_ifContext) IsInst_ifContext() {}

func NewInst_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_ifContext {
	var p = new(Inst_ifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_if

	return p
}

func (s *Inst_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_ifContext) GetE1() IExpresionContext { return s.e1 }

func (s *Inst_ifContext) GetE2() IExpresionContext { return s.e2 }

func (s *Inst_ifContext) SetE1(v IExpresionContext) { s.e1 = v }

func (s *Inst_ifContext) SetE2(v IExpresionContext) { s.e2 = v }

func (s *Inst_ifContext) GetLsalida() string { return s.lsalida }

func (s *Inst_ifContext) SetLsalida(v string) { s.lsalida = v }

func (s *Inst_ifContext) AllBloque() []IBloqueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBloqueContext)(nil)).Elem())
	var tst = make([]IBloqueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBloqueContext)
		}
	}

	return tst
}

func (s *Inst_ifContext) Bloque(i int) IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Inst_ifContext) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *Inst_ifContext) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Inst_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterInst_if(s)
	}
}

func (s *Inst_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitInst_if(s)
	}
}

func (p *GramaticaParser) Inst_if() (localctx IInst_ifContext) {
	this := p
	_ = this

	localctx = NewInst_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GramaticaParserRULE_inst_if)

	localctx.(*Inst_ifContext).lsalida = gen.NewEti()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		p.Match(GramaticaParserT__19)
	}
	{
		p.SetState(218)

		var _x = p.expresion(0)

		localctx.(*Inst_ifContext).e1 = _x
	}

	gen.Soltar(localctx.(*Inst_ifContext).GetE1().GetLv())

	{
		p.SetState(220)
		p.Bloque()
	}

	gen.Gen("goto " + localctx.(*Inst_ifContext).lsalida)
	gen.Soltar(localctx.(*Inst_ifContext).GetE1().GetLf())

	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(222)
				p.Match(GramaticaParserT__20)
			}
			{
				p.SetState(223)
				p.Match(GramaticaParserT__19)
			}
			{
				p.SetState(224)

				var _x = p.expresion(0)

				localctx.(*Inst_ifContext).e2 = _x
			}

			gen.Soltar(localctx.(*Inst_ifContext).GetE2().GetLv())

			{
				p.SetState(226)
				p.Bloque()
			}

			gen.Gen("goto " + localctx.(*Inst_ifContext).lsalida)
			gen.Soltar(localctx.(*Inst_ifContext).GetE2().GetLf())

		}
		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GramaticaParserT__20 {
		{
			p.SetState(234)
			p.Match(GramaticaParserT__20)
		}
		{
			p.SetState(235)
			p.Bloque()
		}

	}
	gen.Gen(localctx.(*Inst_ifContext).lsalida + ":")

	return localctx
}

// IInst_asignacionContext is an interface to support dynamic dispatch.
type IInst_asignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetE returns the e rule contexts.
	GetE() IExpresionContext

	// SetE sets the e rule contexts.
	SetE(IExpresionContext)

	// IsInst_asignacionContext differentiates from other interfaces.
	IsInst_asignacionContext()
}

type Inst_asignacionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
	e      IExpresionContext
}

func NewEmptyInst_asignacionContext() *Inst_asignacionContext {
	var p = new(Inst_asignacionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_asignacion
	return p
}

func (*Inst_asignacionContext) IsInst_asignacionContext() {}

func NewInst_asignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_asignacionContext {
	var p = new(Inst_asignacionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_asignacion

	return p
}

func (s *Inst_asignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_asignacionContext) GetId() antlr.Token { return s.id }

func (s *Inst_asignacionContext) SetId(v antlr.Token) { s.id = v }

func (s *Inst_asignacionContext) GetE() IExpresionContext { return s.e }

func (s *Inst_asignacionContext) SetE(v IExpresionContext) { s.e = v }

func (s *Inst_asignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *Inst_asignacionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Inst_asignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_asignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inst_asignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterInst_asignacion(s)
	}
}

func (s *Inst_asignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitInst_asignacion(s)
	}
}

func (p *GramaticaParser) Inst_asignacion() (localctx IInst_asignacionContext) {
	this := p
	_ = this

	localctx = NewInst_asignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GramaticaParserRULE_inst_asignacion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)

		var _m = p.Match(GramaticaParserID)

		localctx.(*Inst_asignacionContext).id = _m
	}
	{
		p.SetState(241)
		p.Match(GramaticaParserT__21)
	}
	{
		p.SetState(242)

		var _x = p.expresion(0)

		localctx.(*Inst_asignacionContext).e = _x
	}
	gen.Gen((func() string {
		if localctx.(*Inst_asignacionContext).GetId() == nil {
			return ""
		} else {
			return localctx.(*Inst_asignacionContext).GetId().GetText()
		}
	}()) + " = " + localctx.(*Inst_asignacionContext).GetE().GetDir())

	return localctx
}

// IBloqueContext is an interface to support dynamic dispatch.
type IBloqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBloqueContext differentiates from other interfaces.
	IsBloqueContext()
}

type BloqueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloqueContext() *BloqueContext {
	var p = new(BloqueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_bloque
	return p
}

func (*BloqueContext) IsBloqueContext() {}

func NewBloqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueContext {
	var p = new(BloqueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_bloque

	return p
}

func (s *BloqueContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

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
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterBloque(s)
	}
}

func (s *BloqueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitBloque(s)
	}
}

func (p *GramaticaParser) Bloque() (localctx IBloqueContext) {
	this := p
	_ = this

	localctx = NewBloqueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GramaticaParserRULE_bloque)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Match(GramaticaParserT__22)
	}
	{
		p.SetState(246)
		p.Instrucciones()
	}
	{
		p.SetState(247)
		p.Match(GramaticaParserT__23)
	}

	return localctx
}

func (p *GramaticaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	case 2:
		var t *Exp2Context = nil
		if localctx != nil {
			t = localctx.(*Exp2Context)
		}
		return p.Exp2_Sempred(t, predIndex)

	case 3:
		var t *CondContext = nil
		if localctx != nil {
			t = localctx.(*CondContext)
		}
		return p.Cond_Sempred(t, predIndex)

	case 5:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GramaticaParser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GramaticaParser) Exp2_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GramaticaParser) Cond_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GramaticaParser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
