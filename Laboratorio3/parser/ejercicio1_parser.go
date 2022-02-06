// Code generated from Ejercicio1.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Ejercicio1

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "Laboratorio3/Node"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 7, 20, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 6, 3, 16, 10, 3, 13, 3, 14, 3, 17, 3, 3, 2, 2, 4, 2, 4, 2, 2, 2, 18,
	2, 6, 3, 2, 2, 2, 4, 15, 3, 2, 2, 2, 6, 7, 7, 3, 2, 2, 7, 8, 5, 4, 3, 2,
	8, 9, 7, 2, 2, 3, 9, 10, 8, 2, 1, 2, 10, 3, 3, 2, 2, 2, 11, 12, 7, 4, 2,
	2, 12, 13, 7, 6, 2, 2, 13, 14, 7, 5, 2, 2, 14, 16, 8, 3, 1, 2, 15, 11,
	3, 2, 2, 2, 16, 17, 3, 2, 2, 2, 17, 15, 3, 2, 2, 2, 17, 18, 3, 2, 2, 2,
	18, 5, 3, 2, 2, 2, 3, 17,
}
var literalNames = []string{
	"", "'int'", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "CORI", "CORD", "NUMBER", "WHITESPACE",
}

var ruleNames = []string{
	"start", "lista",
}

type Ejercicio1Parser struct {
	*antlr.BaseParser
}

// NewEjercicio1Parser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *Ejercicio1Parser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewEjercicio1Parser(input antlr.TokenStream) *Ejercicio1Parser {
	this := new(Ejercicio1Parser)
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
	this.GrammarFileName = "Ejercicio1.g4"

	return this
}

func mostrar(out string) {
	fmt.Println(out)
}

// Ejercicio1Parser tokens.
const (
	Ejercicio1ParserEOF        = antlr.TokenEOF
	Ejercicio1ParserT__0       = 1
	Ejercicio1ParserCORI       = 2
	Ejercicio1ParserCORD       = 3
	Ejercicio1ParserNUMBER     = 4
	Ejercicio1ParserWHITESPACE = 5
)

// Ejercicio1Parser rules.
const (
	Ejercicio1ParserRULE_start = 0
	Ejercicio1ParserRULE_lista = 1
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t token.
	GetT() antlr.Token

	// SetT sets the t token.
	SetT(antlr.Token)

	// GetNa returns the na rule contexts.
	GetNa() IListaContext

	// SetNa sets the na rule contexts.
	SetNa(IListaContext)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	t      antlr.Token
	na     IListaContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Ejercicio1ParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Ejercicio1ParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) GetT() antlr.Token { return s.t }

func (s *StartContext) SetT(v antlr.Token) { s.t = v }

func (s *StartContext) GetNa() IListaContext { return s.na }

func (s *StartContext) SetNa(v IListaContext) { s.na = v }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(Ejercicio1ParserEOF, 0)
}

func (s *StartContext) Lista() IListaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Ejercicio1Listener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Ejercicio1Listener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *Ejercicio1Parser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Ejercicio1ParserRULE_start)

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
		p.SetState(4)

		var _m = p.Match(Ejercicio1ParserT__0)

		localctx.(*StartContext).t = _m
	}
	{
		p.SetState(5)

		var _x = p.Lista()

		localctx.(*StartContext).na = _x
	}
	{
		p.SetState(6)
		p.Match(Ejercicio1ParserEOF)
	}

	mostrar(localctx.(*StartContext).GetNa().GetN().Cad + "integer" + localctx.(*StartContext).GetNa().GetN().Aux)

	return localctx
}

// IListaContext is an interface to support dynamic dispatch.
type IListaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num token.
	GetNum() antlr.Token

	// SetNum sets the num token.
	SetNum(antlr.Token)

	// GetN returns the n attribute.
	GetN() node.Nodo

	// SetN sets the n attribute.
	SetN(node.Nodo)

	// IsListaContext differentiates from other interfaces.
	IsListaContext()
}

type ListaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	n      node.Nodo
	num    antlr.Token
}

func NewEmptyListaContext() *ListaContext {
	var p = new(ListaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Ejercicio1ParserRULE_lista
	return p
}

func (*ListaContext) IsListaContext() {}

func NewListaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaContext {
	var p = new(ListaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Ejercicio1ParserRULE_lista

	return p
}

func (s *ListaContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaContext) GetNum() antlr.Token { return s.num }

func (s *ListaContext) SetNum(v antlr.Token) { s.num = v }

func (s *ListaContext) GetN() node.Nodo { return s.n }

func (s *ListaContext) SetN(v node.Nodo) { s.n = v }

func (s *ListaContext) AllCORI() []antlr.TerminalNode {
	return s.GetTokens(Ejercicio1ParserCORI)
}

func (s *ListaContext) CORI(i int) antlr.TerminalNode {
	return s.GetToken(Ejercicio1ParserCORI, i)
}

func (s *ListaContext) AllCORD() []antlr.TerminalNode {
	return s.GetTokens(Ejercicio1ParserCORD)
}

func (s *ListaContext) CORD(i int) antlr.TerminalNode {
	return s.GetToken(Ejercicio1ParserCORD, i)
}

func (s *ListaContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(Ejercicio1ParserNUMBER)
}

func (s *ListaContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(Ejercicio1ParserNUMBER, i)
}

func (s *ListaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Ejercicio1Listener); ok {
		listenerT.EnterLista(s)
	}
}

func (s *ListaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Ejercicio1Listener); ok {
		listenerT.ExitLista(s)
	}
}

func (p *Ejercicio1Parser) Lista() (localctx IListaContext) {
	this := p
	_ = this

	localctx = NewListaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Ejercicio1ParserRULE_lista)
	localctx.(*ListaContext).SetN(node.NewNodo("", ""))
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
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == Ejercicio1ParserCORI {
		{
			p.SetState(9)
			p.Match(Ejercicio1ParserCORI)
		}
		{
			p.SetState(10)

			var _m = p.Match(Ejercicio1ParserNUMBER)

			localctx.(*ListaContext).num = _m
		}
		{
			p.SetState(11)
			p.Match(Ejercicio1ParserCORD)
		}

		localctx.(*ListaContext).n.Cad += "arreglo(" + (func() string {
			if localctx.(*ListaContext).GetNum() == nil {
				return ""
			} else {
				return localctx.(*ListaContext).GetNum().GetText()
			}
		}()) + ","
		localctx.(*ListaContext).n.Aux += ")"

		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
