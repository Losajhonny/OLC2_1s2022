// Code generated from Gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 31, 161,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 6, 27, 130, 10, 27,
	13, 27, 14, 27, 131, 3, 28, 3, 28, 7, 28, 136, 10, 28, 12, 28, 14, 28,
	139, 11, 28, 3, 29, 3, 29, 3, 29, 3, 29, 7, 29, 145, 10, 29, 12, 29, 14,
	29, 148, 11, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 6, 30, 156,
	10, 30, 13, 30, 14, 30, 157, 3, 30, 3, 30, 3, 146, 2, 31, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45,
	24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 3, 2, 6, 3,
	2, 50, 59, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97,
	99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 164, 2, 3, 3, 2, 2, 2, 2, 5,
	3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2,
	2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3,
	2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59,
	3, 2, 2, 2, 3, 61, 3, 2, 2, 2, 5, 63, 3, 2, 2, 2, 7, 65, 3, 2, 2, 2, 9,
	67, 3, 2, 2, 2, 11, 69, 3, 2, 2, 2, 13, 71, 3, 2, 2, 2, 15, 73, 3, 2, 2,
	2, 17, 76, 3, 2, 2, 2, 19, 78, 3, 2, 2, 2, 21, 81, 3, 2, 2, 2, 23, 83,
	3, 2, 2, 2, 25, 85, 3, 2, 2, 2, 27, 90, 3, 2, 2, 2, 29, 96, 3, 2, 2, 2,
	31, 99, 3, 2, 2, 2, 33, 102, 3, 2, 2, 2, 35, 104, 3, 2, 2, 2, 37, 106,
	3, 2, 2, 2, 39, 109, 3, 2, 2, 2, 41, 112, 3, 2, 2, 2, 43, 114, 3, 2, 2,
	2, 45, 116, 3, 2, 2, 2, 47, 119, 3, 2, 2, 2, 49, 124, 3, 2, 2, 2, 51, 126,
	3, 2, 2, 2, 53, 129, 3, 2, 2, 2, 55, 133, 3, 2, 2, 2, 57, 140, 3, 2, 2,
	2, 59, 155, 3, 2, 2, 2, 61, 62, 7, 45, 2, 2, 62, 4, 3, 2, 2, 2, 63, 64,
	7, 47, 2, 2, 64, 6, 3, 2, 2, 2, 65, 66, 7, 35, 2, 2, 66, 8, 3, 2, 2, 2,
	67, 68, 7, 44, 2, 2, 68, 10, 3, 2, 2, 2, 69, 70, 7, 49, 2, 2, 70, 12, 3,
	2, 2, 2, 71, 72, 7, 39, 2, 2, 72, 14, 3, 2, 2, 2, 73, 74, 7, 40, 2, 2,
	74, 75, 7, 40, 2, 2, 75, 16, 3, 2, 2, 2, 76, 77, 7, 96, 2, 2, 77, 18, 3,
	2, 2, 2, 78, 79, 7, 126, 2, 2, 79, 80, 7, 126, 2, 2, 80, 20, 3, 2, 2, 2,
	81, 82, 7, 42, 2, 2, 82, 22, 3, 2, 2, 2, 83, 84, 7, 43, 2, 2, 84, 24, 3,
	2, 2, 2, 85, 86, 7, 118, 2, 2, 86, 87, 7, 116, 2, 2, 87, 88, 7, 119, 2,
	2, 88, 89, 7, 103, 2, 2, 89, 26, 3, 2, 2, 2, 90, 91, 7, 104, 2, 2, 91,
	92, 7, 99, 2, 2, 92, 93, 7, 110, 2, 2, 93, 94, 7, 117, 2, 2, 94, 95, 7,
	103, 2, 2, 95, 28, 3, 2, 2, 2, 96, 97, 7, 63, 2, 2, 97, 98, 7, 63, 2, 2,
	98, 30, 3, 2, 2, 2, 99, 100, 7, 35, 2, 2, 100, 101, 7, 63, 2, 2, 101, 32,
	3, 2, 2, 2, 102, 103, 7, 64, 2, 2, 103, 34, 3, 2, 2, 2, 104, 105, 7, 62,
	2, 2, 105, 36, 3, 2, 2, 2, 106, 107, 7, 64, 2, 2, 107, 108, 7, 63, 2, 2,
	108, 38, 3, 2, 2, 2, 109, 110, 7, 62, 2, 2, 110, 111, 7, 63, 2, 2, 111,
	40, 3, 2, 2, 2, 112, 113, 7, 63, 2, 2, 113, 42, 3, 2, 2, 2, 114, 115, 7,
	61, 2, 2, 115, 44, 3, 2, 2, 2, 116, 117, 7, 107, 2, 2, 117, 118, 7, 104,
	2, 2, 118, 46, 3, 2, 2, 2, 119, 120, 7, 103, 2, 2, 120, 121, 7, 110, 2,
	2, 121, 122, 7, 117, 2, 2, 122, 123, 7, 103, 2, 2, 123, 48, 3, 2, 2, 2,
	124, 125, 7, 125, 2, 2, 125, 50, 3, 2, 2, 2, 126, 127, 7, 127, 2, 2, 127,
	52, 3, 2, 2, 2, 128, 130, 9, 2, 2, 2, 129, 128, 3, 2, 2, 2, 130, 131, 3,
	2, 2, 2, 131, 129, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 54, 3, 2, 2,
	2, 133, 137, 9, 3, 2, 2, 134, 136, 9, 4, 2, 2, 135, 134, 3, 2, 2, 2, 136,
	139, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138, 56, 3,
	2, 2, 2, 139, 137, 3, 2, 2, 2, 140, 141, 7, 49, 2, 2, 141, 142, 7, 44,
	2, 2, 142, 146, 3, 2, 2, 2, 143, 145, 11, 2, 2, 2, 144, 143, 3, 2, 2, 2,
	145, 148, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 147,
	149, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 149, 150, 7, 44, 2, 2, 150, 151,
	7, 49, 2, 2, 151, 152, 3, 2, 2, 2, 152, 153, 8, 29, 2, 2, 153, 58, 3, 2,
	2, 2, 154, 156, 9, 5, 2, 2, 155, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2,
	157, 155, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159,
	160, 8, 30, 2, 2, 160, 60, 3, 2, 2, 2, 7, 2, 131, 137, 146, 157, 3, 8,
	2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'+'", "'-'", "'!'", "'*'", "'/'", "'%'", "'&&'", "'^'", "'||'", "'('",
	"')'", "'true'", "'false'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='",
	"'='", "';'", "'if'", "'else'", "'{'", "'}'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "NUM", "ID", "COMMENT", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"NUM", "ID", "COMMENT", "WHITESPACE",
}

type GramaticaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewGramaticaLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *GramaticaLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewGramaticaLexer(input antlr.CharStream) *GramaticaLexer {
	l := new(GramaticaLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Gramatica.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GramaticaLexer tokens.
const (
	GramaticaLexerT__0       = 1
	GramaticaLexerT__1       = 2
	GramaticaLexerT__2       = 3
	GramaticaLexerT__3       = 4
	GramaticaLexerT__4       = 5
	GramaticaLexerT__5       = 6
	GramaticaLexerT__6       = 7
	GramaticaLexerT__7       = 8
	GramaticaLexerT__8       = 9
	GramaticaLexerT__9       = 10
	GramaticaLexerT__10      = 11
	GramaticaLexerT__11      = 12
	GramaticaLexerT__12      = 13
	GramaticaLexerT__13      = 14
	GramaticaLexerT__14      = 15
	GramaticaLexerT__15      = 16
	GramaticaLexerT__16      = 17
	GramaticaLexerT__17      = 18
	GramaticaLexerT__18      = 19
	GramaticaLexerT__19      = 20
	GramaticaLexerT__20      = 21
	GramaticaLexerT__21      = 22
	GramaticaLexerT__22      = 23
	GramaticaLexerT__23      = 24
	GramaticaLexerT__24      = 25
	GramaticaLexerNUM        = 26
	GramaticaLexerID         = 27
	GramaticaLexerCOMMENT    = 28
	GramaticaLexerWHITESPACE = 29
)
