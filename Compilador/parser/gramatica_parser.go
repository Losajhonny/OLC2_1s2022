// Code generated from Gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Gramatica

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "Compilador/gen"
import "Compilador/entorno"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 31, 149,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 5, 4, 52, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 7, 4, 87, 10, 4, 12, 4, 14, 4, 90, 11, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 104, 10, 5, 3,
	6, 7, 6, 107, 10, 6, 12, 6, 14, 6, 110, 11, 6, 3, 7, 3, 7, 5, 7, 114, 10,
	7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 134, 10, 9, 12, 9, 14, 9,
	137, 11, 9, 3, 9, 3, 9, 5, 9, 141, 10, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 2, 3, 6, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 4, 3,
	2, 6, 8, 3, 2, 3, 4, 2, 161, 2, 20, 3, 2, 2, 2, 4, 23, 3, 2, 2, 2, 6, 51,
	3, 2, 2, 2, 8, 103, 3, 2, 2, 2, 10, 108, 3, 2, 2, 2, 12, 113, 3, 2, 2,
	2, 14, 115, 3, 2, 2, 2, 16, 121, 3, 2, 2, 2, 18, 144, 3, 2, 2, 2, 20, 21,
	5, 10, 6, 2, 21, 22, 7, 2, 2, 3, 22, 3, 3, 2, 2, 2, 23, 24, 8, 3, 1, 2,
	24, 5, 3, 2, 2, 2, 25, 26, 8, 4, 1, 2, 26, 27, 7, 3, 2, 2, 27, 28, 5, 6,
	4, 16, 28, 29, 8, 4, 1, 2, 29, 52, 3, 2, 2, 2, 30, 31, 7, 4, 2, 2, 31,
	32, 5, 6, 4, 15, 32, 33, 8, 4, 1, 2, 33, 52, 3, 2, 2, 2, 34, 35, 7, 5,
	2, 2, 35, 36, 5, 6, 4, 14, 36, 37, 8, 4, 1, 2, 37, 52, 3, 2, 2, 2, 38,
	39, 7, 12, 2, 2, 39, 40, 5, 6, 4, 2, 40, 41, 7, 13, 2, 2, 41, 42, 8, 4,
	1, 2, 42, 52, 3, 2, 2, 2, 43, 44, 7, 28, 2, 2, 44, 52, 8, 4, 1, 2, 45,
	46, 7, 29, 2, 2, 46, 52, 8, 4, 1, 2, 47, 48, 7, 14, 2, 2, 48, 52, 8, 4,
	1, 2, 49, 50, 7, 15, 2, 2, 50, 52, 8, 4, 1, 2, 51, 25, 3, 2, 2, 2, 51,
	30, 3, 2, 2, 2, 51, 34, 3, 2, 2, 2, 51, 38, 3, 2, 2, 2, 51, 43, 3, 2, 2,
	2, 51, 45, 3, 2, 2, 2, 51, 47, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 52, 88,
	3, 2, 2, 2, 53, 54, 12, 13, 2, 2, 54, 55, 9, 2, 2, 2, 55, 56, 5, 6, 4,
	14, 56, 57, 8, 4, 1, 2, 57, 87, 3, 2, 2, 2, 58, 59, 12, 12, 2, 2, 59, 60,
	9, 3, 2, 2, 60, 61, 5, 6, 4, 13, 61, 62, 8, 4, 1, 2, 62, 87, 3, 2, 2, 2,
	63, 64, 12, 11, 2, 2, 64, 65, 5, 8, 5, 2, 65, 66, 5, 6, 4, 12, 66, 67,
	8, 4, 1, 2, 67, 87, 3, 2, 2, 2, 68, 69, 12, 10, 2, 2, 69, 70, 7, 9, 2,
	2, 70, 71, 5, 4, 3, 2, 71, 72, 5, 6, 4, 11, 72, 73, 8, 4, 1, 2, 73, 87,
	3, 2, 2, 2, 74, 75, 12, 9, 2, 2, 75, 76, 7, 10, 2, 2, 76, 77, 5, 4, 3,
	2, 77, 78, 5, 6, 4, 10, 78, 79, 8, 4, 1, 2, 79, 87, 3, 2, 2, 2, 80, 81,
	12, 8, 2, 2, 81, 82, 7, 11, 2, 2, 82, 83, 5, 4, 3, 2, 83, 84, 5, 6, 4,
	9, 84, 85, 8, 4, 1, 2, 85, 87, 3, 2, 2, 2, 86, 53, 3, 2, 2, 2, 86, 58,
	3, 2, 2, 2, 86, 63, 3, 2, 2, 2, 86, 68, 3, 2, 2, 2, 86, 74, 3, 2, 2, 2,
	86, 80, 3, 2, 2, 2, 87, 90, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3,
	2, 2, 2, 89, 7, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 91, 92, 7, 16, 2, 2, 92,
	104, 8, 5, 1, 2, 93, 94, 7, 17, 2, 2, 94, 104, 8, 5, 1, 2, 95, 96, 7, 18,
	2, 2, 96, 104, 8, 5, 1, 2, 97, 98, 7, 19, 2, 2, 98, 104, 8, 5, 1, 2, 99,
	100, 7, 20, 2, 2, 100, 104, 8, 5, 1, 2, 101, 102, 7, 21, 2, 2, 102, 104,
	8, 5, 1, 2, 103, 91, 3, 2, 2, 2, 103, 93, 3, 2, 2, 2, 103, 95, 3, 2, 2,
	2, 103, 97, 3, 2, 2, 2, 103, 99, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 104,
	9, 3, 2, 2, 2, 105, 107, 5, 12, 7, 2, 106, 105, 3, 2, 2, 2, 107, 110, 3,
	2, 2, 2, 108, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 11, 3, 2, 2,
	2, 110, 108, 3, 2, 2, 2, 111, 114, 5, 14, 8, 2, 112, 114, 5, 16, 9, 2,
	113, 111, 3, 2, 2, 2, 113, 112, 3, 2, 2, 2, 114, 13, 3, 2, 2, 2, 115, 116,
	7, 29, 2, 2, 116, 117, 7, 22, 2, 2, 117, 118, 5, 6, 4, 2, 118, 119, 7,
	23, 2, 2, 119, 120, 8, 8, 1, 2, 120, 15, 3, 2, 2, 2, 121, 122, 7, 24, 2,
	2, 122, 123, 5, 6, 4, 2, 123, 124, 8, 9, 1, 2, 124, 125, 5, 18, 10, 2,
	125, 135, 8, 9, 1, 2, 126, 127, 7, 25, 2, 2, 127, 128, 7, 24, 2, 2, 128,
	129, 5, 6, 4, 2, 129, 130, 8, 9, 1, 2, 130, 131, 5, 18, 10, 2, 131, 132,
	8, 9, 1, 2, 132, 134, 3, 2, 2, 2, 133, 126, 3, 2, 2, 2, 134, 137, 3, 2,
	2, 2, 135, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 140, 3, 2, 2, 2,
	137, 135, 3, 2, 2, 2, 138, 139, 7, 25, 2, 2, 139, 141, 5, 18, 10, 2, 140,
	138, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 143,
	8, 9, 1, 2, 143, 17, 3, 2, 2, 2, 144, 145, 7, 26, 2, 2, 145, 146, 5, 10,
	6, 2, 146, 147, 7, 27, 2, 2, 147, 19, 3, 2, 2, 2, 10, 51, 86, 88, 103,
	108, 113, 135, 140,
}
var literalNames = []string{
	"", "'+'", "'-'", "'!'", "'*'", "'/'", "'%'", "'&&'", "'^'", "'||'", "'('",
	"')'", "'true'", "'false'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='",
	"'='", "';'", "'if'", "'else'", "'{'", "'}'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "NUM", "ID", "COMMENT", "WHITESPACE",
}

var ruleNames = []string{
	"start", "marcador", "expresion", "oprel", "instrucciones", "instruccion",
	"inst_asignacion", "inst_if", "bloque",
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
	GramaticaParserT__24      = 25
	GramaticaParserNUM        = 26
	GramaticaParserID         = 27
	GramaticaParserCOMMENT    = 28
	GramaticaParserWHITESPACE = 29
)

// GramaticaParser rules.
const (
	GramaticaParserRULE_start           = 0
	GramaticaParserRULE_marcador        = 1
	GramaticaParserRULE_expresion       = 2
	GramaticaParserRULE_oprel           = 3
	GramaticaParserRULE_instrucciones   = 4
	GramaticaParserRULE_instruccion     = 5
	GramaticaParserRULE_inst_asignacion = 6
	GramaticaParserRULE_inst_if         = 7
	GramaticaParserRULE_bloque          = 8
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
		p.SetState(18)
		p.Instrucciones()
	}
	{
		p.SetState(19)
		p.Match(GramaticaParserEOF)
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
	p.EnterRule(localctx, 2, GramaticaParserRULE_marcador)

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
	_startState := 4
	p.EnterRecursionRule(localctx, 4, GramaticaParserRULE_expresion, _p)
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
	p.SetState(49)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserT__0:
		{
			p.SetState(24)

			var _m = p.Match(GramaticaParserT__0)

			localctx.(*ExpresionContext).op = _m
		}
		{
			p.SetState(25)

			var _x = p.expresion(14)

			localctx.(*ExpresionContext).e = _x
		}

		localctx.(*ExpresionContext).dir = gen.NewTemp()
		gen.GenOperacionUnaria(localctx.(*ExpresionContext).dir, (func() string {
			if localctx.(*ExpresionContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).GetOp().GetText()
			}
		}()), localctx.(*ExpresionContext).GetE().GetDir())

	case GramaticaParserT__1:
		{
			p.SetState(28)

			var _m = p.Match(GramaticaParserT__1)

			localctx.(*ExpresionContext).op = _m
		}
		{
			p.SetState(29)

			var _x = p.expresion(13)

			localctx.(*ExpresionContext).e = _x
		}

		localctx.(*ExpresionContext).dir = gen.NewTemp()
		gen.GenOperacionUnaria(localctx.(*ExpresionContext).dir, (func() string {
			if localctx.(*ExpresionContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).GetOp().GetText()
			}
		}()), localctx.(*ExpresionContext).GetE().GetDir())

	case GramaticaParserT__2:
		{
			p.SetState(32)

			var _m = p.Match(GramaticaParserT__2)

			localctx.(*ExpresionContext).op = _m
		}
		{
			p.SetState(33)

			var _x = p.expresion(12)

			localctx.(*ExpresionContext).e = _x
		}

		localctx.(*ExpresionContext).lv = localctx.(*ExpresionContext).GetE().GetLf()
		localctx.(*ExpresionContext).lf = localctx.(*ExpresionContext).GetE().GetLv()

	case GramaticaParserT__9:
		{
			p.SetState(36)
			p.Match(GramaticaParserT__9)
		}
		{
			p.SetState(37)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e = _x
		}
		{
			p.SetState(38)
			p.Match(GramaticaParserT__10)
		}

		localctx.(*ExpresionContext).dir = localctx.(*ExpresionContext).GetE().GetDir()
		localctx.(*ExpresionContext).lv = localctx.(*ExpresionContext).GetE().GetLv()
		localctx.(*ExpresionContext).lf = localctx.(*ExpresionContext).GetE().GetLf()
		localctx.(*ExpresionContext).cad = localctx.(*ExpresionContext).GetE().GetCad()

	case GramaticaParserNUM:
		{
			p.SetState(41)

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
			p.SetState(43)

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

	case GramaticaParserT__11:
		{
			p.SetState(45)
			p.Match(GramaticaParserT__11)
		}

		localctx.(*ExpresionContext).dir = ""
		localctx.(*ExpresionContext).cad = ""
		localctx.(*ExpresionContext).lv = append(localctx.(*ExpresionContext).lv, gen.NewEti())
		localctx.(*ExpresionContext).lf = append(localctx.(*ExpresionContext).lf, gen.NewEti())
		gen.GenGoto(localctx.(*ExpresionContext).lv[0])

	case GramaticaParserT__12:
		{
			p.SetState(47)
			p.Match(GramaticaParserT__12)
		}

		localctx.(*ExpresionContext).dir = ""
		localctx.(*ExpresionContext).cad = ""
		localctx.(*ExpresionContext).lv = append(localctx.(*ExpresionContext).lv, gen.NewEti())
		localctx.(*ExpresionContext).lf = append(localctx.(*ExpresionContext).lf, gen.NewEti())
		gen.GenGoto(localctx.(*ExpresionContext).lf[0])

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(84)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(51)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(52)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GramaticaParserT__3)|(1<<GramaticaParserT__4)|(1<<GramaticaParserT__5))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(53)

					var _x = p.expresion(12)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).dir = gen.NewTemp()
				gen.GenOperacion(localctx.(*ExpresionContext).dir, localctx.(*ExpresionContext).GetE1().GetDir(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetE2().GetDir())

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(56)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(57)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserT__0 || _la == GramaticaParserT__1) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(58)

					var _x = p.expresion(11)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).dir = gen.NewTemp()
				gen.GenOperacion(localctx.(*ExpresionContext).dir, localctx.(*ExpresionContext).GetE1().GetDir(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetE2().GetDir())

			case 3:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(61)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(62)

					var _x = p.Oprel()

					localctx.(*ExpresionContext).opr = _x
				}
				{
					p.SetState(63)

					var _x = p.expresion(10)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).lv = append(localctx.(*ExpresionContext).lv, gen.NewEti())
				localctx.(*ExpresionContext).lf = append(localctx.(*ExpresionContext).lf, gen.NewEti())
				localctx.(*ExpresionContext).cad = gen.GenIf(localctx.(*ExpresionContext).GetE1().GetDir(), localctx.(*ExpresionContext).GetOpr().GetOp(), localctx.(*ExpresionContext).GetE2().GetDir(), localctx.(*ExpresionContext).lv[0])
				gen.GenGoto(localctx.(*ExpresionContext).lf[0])

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(66)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(67)

					var _m = p.Match(GramaticaParserT__6)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(68)
					p.Marcador(localctx.(*ExpresionContext).GetE1().GetLv())
				}
				{
					p.SetState(69)

					var _x = p.expresion(9)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).lv = localctx.(*ExpresionContext).GetE2().GetLv()
				localctx.(*ExpresionContext).lf = gen.Unir(localctx.(*ExpresionContext).GetE1().GetLf(), localctx.(*ExpresionContext).GetE2().GetLf())

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(72)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(73)

					var _m = p.Match(GramaticaParserT__7)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(74)
					p.Marcador(localctx.(*ExpresionContext).GetE1().GetLf())
				}
				{
					p.SetState(75)

					var _x = p.expresion(8)

					localctx.(*ExpresionContext).e2 = _x
				}

				gen.Soltar(localctx.(*ExpresionContext).GetE1().GetLv())
				gen.GenIfCad(localctx.(*ExpresionContext).GetE2().GetCad(), localctx.(*ExpresionContext).GetE2().GetLf()[0])
				gen.GenGoto("goto " + localctx.(*ExpresionContext).GetE2().GetLv()[0])
				localctx.(*ExpresionContext).lv = localctx.(*ExpresionContext).GetE2().GetLv()
				localctx.(*ExpresionContext).lf = localctx.(*ExpresionContext).GetE2().GetLf()

			case 6:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(78)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(79)

					var _m = p.Match(GramaticaParserT__8)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(80)
					p.Marcador(localctx.(*ExpresionContext).GetE1().GetLf())
				}
				{
					p.SetState(81)

					var _x = p.expresion(7)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).lf = localctx.(*ExpresionContext).GetE2().GetLf()
				localctx.(*ExpresionContext).lv = gen.Unir(localctx.(*ExpresionContext).GetE1().GetLv(), localctx.(*ExpresionContext).GetE2().GetLv())

			}

		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 6, GramaticaParserRULE_oprel)

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

	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserT__13:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(89)

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
			p.SetState(91)

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
			p.SetState(93)

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
			p.SetState(95)

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
			p.SetState(97)

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
			p.SetState(99)

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
	p.EnterRule(localctx, 8, GramaticaParserRULE_instrucciones)
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
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GramaticaParserT__21 || _la == GramaticaParserID {
		{
			p.SetState(103)
			p.Instruccion()
		}

		p.SetState(108)
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

func (s *InstruccionContext) Inst_asignacion() IInst_asignacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_asignacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_asignacionContext)
}

func (s *InstruccionContext) Inst_if() IInst_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_ifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_ifContext)
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
	p.EnterRule(localctx, 10, GramaticaParserRULE_instruccion)

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

	p.SetState(111)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.Inst_asignacion()
		}

	case GramaticaParserT__21:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.Inst_if()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

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
	p.EnterRule(localctx, 12, GramaticaParserRULE_inst_asignacion)

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
		p.SetState(113)

		var _m = p.Match(GramaticaParserID)

		localctx.(*Inst_asignacionContext).id = _m
	}
	{
		p.SetState(114)
		p.Match(GramaticaParserT__19)
	}
	{
		p.SetState(115)

		var _x = p.expresion(0)

		localctx.(*Inst_asignacionContext).e = _x
	}
	{
		p.SetState(116)
		p.Match(GramaticaParserT__20)
	}
	gen.GenAsignacion((func() string {
		if localctx.(*Inst_asignacionContext).GetId() == nil {
			return ""
		} else {
			return localctx.(*Inst_asignacionContext).GetId().GetText()
		}
	}()), localctx.(*Inst_asignacionContext).GetE().GetDir())

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
	p.EnterRule(localctx, 14, GramaticaParserRULE_inst_if)
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
		p.SetState(119)
		p.Match(GramaticaParserT__21)
	}
	{
		p.SetState(120)

		var _x = p.expresion(0)

		localctx.(*Inst_ifContext).e1 = _x
	}

	localctx.(*Inst_ifContext).lsalida = gen.NewEti()
	gen.Soltar(localctx.(*Inst_ifContext).GetE1().GetLv())

	{
		p.SetState(122)
		p.Bloque()
	}

	gen.GenGoto(localctx.(*Inst_ifContext).lsalida)
	gen.Soltar(localctx.(*Inst_ifContext).GetE1().GetLf())

	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(124)
				p.Match(GramaticaParserT__22)
			}
			{
				p.SetState(125)
				p.Match(GramaticaParserT__21)
			}
			{
				p.SetState(126)

				var _x = p.expresion(0)

				localctx.(*Inst_ifContext).e2 = _x
			}

			gen.Soltar(localctx.(*Inst_ifContext).GetE2().GetLv())

			{
				p.SetState(128)
				p.Bloque()
			}

			gen.GenGoto(localctx.(*Inst_ifContext).lsalida)
			gen.Soltar(localctx.(*Inst_ifContext).GetE2().GetLf())

		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GramaticaParserT__22 {
		{
			p.SetState(136)
			p.Match(GramaticaParserT__22)
		}
		{
			p.SetState(137)
			p.Bloque()
		}

	}
	gen.GenDestino(localctx.(*Inst_ifContext).lsalida)

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
	p.EnterRule(localctx, 16, GramaticaParserRULE_bloque)

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
		p.SetState(142)
		p.Match(GramaticaParserT__23)
	}
	{
		p.SetState(143)
		p.Instrucciones()
	}
	{
		p.SetState(144)
		p.Match(GramaticaParserT__24)
	}

	return localctx
}

func (p *GramaticaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GramaticaParser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
