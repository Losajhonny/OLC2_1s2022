// Code generated from Gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Gramatica

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 8, 45, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 26,
	10, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 36, 10, 6,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 43, 10, 7, 3, 7, 2, 2, 8, 2, 4, 6,
	8, 10, 12, 2, 2, 2, 41, 2, 14, 3, 2, 2, 2, 4, 17, 3, 2, 2, 2, 6, 25, 3,
	2, 2, 2, 8, 27, 3, 2, 2, 2, 10, 35, 3, 2, 2, 2, 12, 42, 3, 2, 2, 2, 14,
	15, 5, 4, 3, 2, 15, 16, 7, 2, 2, 3, 16, 3, 3, 2, 2, 2, 17, 18, 5, 8, 5,
	2, 18, 19, 5, 6, 4, 2, 19, 5, 3, 2, 2, 2, 20, 21, 7, 3, 2, 2, 21, 22, 5,
	8, 5, 2, 22, 23, 5, 6, 4, 2, 23, 26, 3, 2, 2, 2, 24, 26, 3, 2, 2, 2, 25,
	20, 3, 2, 2, 2, 25, 24, 3, 2, 2, 2, 26, 7, 3, 2, 2, 2, 27, 28, 5, 12, 7,
	2, 28, 29, 5, 10, 6, 2, 29, 9, 3, 2, 2, 2, 30, 31, 7, 4, 2, 2, 31, 32,
	5, 12, 7, 2, 32, 33, 5, 10, 6, 2, 33, 36, 3, 2, 2, 2, 34, 36, 3, 2, 2,
	2, 35, 30, 3, 2, 2, 2, 35, 34, 3, 2, 2, 2, 36, 11, 3, 2, 2, 2, 37, 38,
	7, 5, 2, 2, 38, 39, 5, 4, 3, 2, 39, 40, 7, 6, 2, 2, 40, 43, 3, 2, 2, 2,
	41, 43, 7, 7, 2, 2, 42, 37, 3, 2, 2, 2, 42, 41, 3, 2, 2, 2, 43, 13, 3,
	2, 2, 2, 5, 25, 35, 42,
}
var literalNames = []string{
	"", "'+'", "'*'", "'('", "')'",
}
var symbolicNames = []string{
	"", "MAS", "POR", "PARI", "PARD", "NUM", "WHITESPACE",
}

var ruleNames = []string{
	"start", "e", "ep", "t", "tp", "f",
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

// GramaticaParser tokens.
const (
	GramaticaParserEOF        = antlr.TokenEOF
	GramaticaParserMAS        = 1
	GramaticaParserPOR        = 2
	GramaticaParserPARI       = 3
	GramaticaParserPARD       = 4
	GramaticaParserNUM        = 5
	GramaticaParserWHITESPACE = 6
)

// GramaticaParser rules.
const (
	GramaticaParserRULE_start = 0
	GramaticaParserRULE_e     = 1
	GramaticaParserRULE_ep    = 2
	GramaticaParserRULE_t     = 3
	GramaticaParserRULE_tp    = 4
	GramaticaParserRULE_f     = 5
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

func (s *StartContext) CopyFrom(ctx *StartContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StatContext struct {
	*StartContext
}

func NewStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatContext {
	var p = new(StatContext)

	p.StartContext = NewEmptyStartContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StartContext))

	return p
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) E() IEContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *StatContext) EOF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserEOF, 0)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitStat(s)
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

	localctx = NewStatContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)
		p.E()
	}
	{
		p.SetState(13)
		p.Match(GramaticaParserEOF)
	}

	return localctx
}

// IEContext is an interface to support dynamic dispatch.
type IEContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEContext differentiates from other interfaces.
	IsEContext()
}

type EContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEContext() *EContext {
	var p = new(EContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_e
	return p
}

func (*EContext) IsEContext() {}

func NewEContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EContext {
	var p = new(EContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_e

	return p
}

func (s *EContext) GetParser() antlr.Parser { return s.parser }

func (s *EContext) CopyFrom(ctx *EContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *EContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpContext struct {
	*EContext
}

func NewExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpContext {
	var p = new(ExpContext)

	p.EContext = NewEmptyEContext()
	p.parser = parser
	p.CopyFrom(ctx.(*EContext))

	return p
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) T() ITContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (s *ExpContext) Ep() IEpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEpContext)
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

func (p *GramaticaParser) E() (localctx IEContext) {
	this := p
	_ = this

	localctx = NewEContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GramaticaParserRULE_e)

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

	localctx = NewExpContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(15)
		p.T()
	}
	{
		p.SetState(16)
		p.Ep()
	}

	return localctx
}

// IEpContext is an interface to support dynamic dispatch.
type IEpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEpContext differentiates from other interfaces.
	IsEpContext()
}

type EpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEpContext() *EpContext {
	var p = new(EpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_ep
	return p
}

func (*EpContext) IsEpContext() {}

func NewEpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EpContext {
	var p = new(EpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_ep

	return p
}

func (s *EpContext) GetParser() antlr.Parser { return s.parser }

func (s *EpContext) CopyFrom(ctx *EpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *EpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EpExpContext struct {
	*EpContext
}

func NewEpExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EpExpContext {
	var p = new(EpExpContext)

	p.EpContext = NewEmptyEpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*EpContext))

	return p
}

func (s *EpExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterEpExp(s)
	}
}

func (s *EpExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitEpExp(s)
	}
}

type SumContext struct {
	*EpContext
}

func NewSumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SumContext {
	var p = new(SumContext)

	p.EpContext = NewEmptyEpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*EpContext))

	return p
}

func (s *SumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumContext) MAS() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMAS, 0)
}

func (s *SumContext) T() ITContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (s *SumContext) Ep() IEpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEpContext)
}

func (s *SumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterSum(s)
	}
}

func (s *SumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitSum(s)
	}
}

func (p *GramaticaParser) Ep() (localctx IEpContext) {
	this := p
	_ = this

	localctx = NewEpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GramaticaParserRULE_ep)

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

	p.SetState(23)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserMAS:
		localctx = NewSumContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(18)
			p.Match(GramaticaParserMAS)
		}
		{
			p.SetState(19)
			p.T()
		}
		{
			p.SetState(20)
			p.Ep()
		}

	case GramaticaParserEOF, GramaticaParserPARD:
		localctx = NewEpExpContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITContext is an interface to support dynamic dispatch.
type ITContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTContext differentiates from other interfaces.
	IsTContext()
}

type TContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTContext() *TContext {
	var p = new(TContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_t
	return p
}

func (*TContext) IsTContext() {}

func NewTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TContext {
	var p = new(TContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_t

	return p
}

func (s *TContext) GetParser() antlr.Parser { return s.parser }

func (s *TContext) CopyFrom(ctx *TContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TermContext struct {
	*TContext
}

func NewTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TermContext {
	var p = new(TermContext)

	p.TContext = NewEmptyTContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TContext))

	return p
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) F() IFContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFContext)
}

func (s *TermContext) Tp() ITpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITpContext)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *GramaticaParser) T() (localctx ITContext) {
	this := p
	_ = this

	localctx = NewTContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GramaticaParserRULE_t)

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

	localctx = NewTermContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.F()
	}
	{
		p.SetState(26)
		p.Tp()
	}

	return localctx
}

// ITpContext is an interface to support dynamic dispatch.
type ITpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTpContext differentiates from other interfaces.
	IsTpContext()
}

type TpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTpContext() *TpContext {
	var p = new(TpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_tp
	return p
}

func (*TpContext) IsTpContext() {}

func NewTpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TpContext {
	var p = new(TpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_tp

	return p
}

func (s *TpContext) GetParser() antlr.Parser { return s.parser }

func (s *TpContext) CopyFrom(ctx *TpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MulContext struct {
	*TpContext
}

func NewMulContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulContext {
	var p = new(MulContext)

	p.TpContext = NewEmptyTpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TpContext))

	return p
}

func (s *MulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulContext) POR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPOR, 0)
}

func (s *MulContext) F() IFContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFContext)
}

func (s *MulContext) Tp() ITpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITpContext)
}

func (s *MulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterMul(s)
	}
}

func (s *MulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitMul(s)
	}
}

type EpTermContext struct {
	*TpContext
}

func NewEpTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EpTermContext {
	var p = new(EpTermContext)

	p.TpContext = NewEmptyTpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TpContext))

	return p
}

func (s *EpTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterEpTerm(s)
	}
}

func (s *EpTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitEpTerm(s)
	}
}

func (p *GramaticaParser) Tp() (localctx ITpContext) {
	this := p
	_ = this

	localctx = NewTpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GramaticaParserRULE_tp)

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

	p.SetState(33)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserPOR:
		localctx = NewMulContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.Match(GramaticaParserPOR)
		}
		{
			p.SetState(29)
			p.F()
		}
		{
			p.SetState(30)
			p.Tp()
		}

	case GramaticaParserEOF, GramaticaParserMAS, GramaticaParserPARD:
		localctx = NewEpTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFContext is an interface to support dynamic dispatch.
type IFContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFContext differentiates from other interfaces.
	IsFContext()
}

type FContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFContext() *FContext {
	var p = new(FContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_f
	return p
}

func (*FContext) IsFContext() {}

func NewFContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FContext {
	var p = new(FContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_f

	return p
}

func (s *FContext) GetParser() antlr.Parser { return s.parser }

func (s *FContext) CopyFrom(ctx *FContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParentesisContext struct {
	*FContext
}

func NewParentesisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentesisContext {
	var p = new(ParentesisContext)

	p.FContext = NewEmptyFContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FContext))

	return p
}

func (s *ParentesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentesisContext) PARI() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPARI, 0)
}

func (s *ParentesisContext) E() IEContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *ParentesisContext) PARD() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPARD, 0)
}

func (s *ParentesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterParentesis(s)
	}
}

func (s *ParentesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitParentesis(s)
	}
}

type NumeroContext struct {
	*FContext
}

func NewNumeroContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumeroContext {
	var p = new(NumeroContext)

	p.FContext = NewEmptyFContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FContext))

	return p
}

func (s *NumeroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumeroContext) NUM() antlr.TerminalNode {
	return s.GetToken(GramaticaParserNUM, 0)
}

func (s *NumeroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.EnterNumero(s)
	}
}

func (s *NumeroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GramaticaListener); ok {
		listenerT.ExitNumero(s)
	}
}

func (p *GramaticaParser) F() (localctx IFContext) {
	this := p
	_ = this

	localctx = NewFContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GramaticaParserRULE_f)

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

	p.SetState(40)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserPARI:
		localctx = NewParentesisContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(35)
			p.Match(GramaticaParserPARI)
		}
		{
			p.SetState(36)
			p.E()
		}
		{
			p.SetState(37)
			p.Match(GramaticaParserPARD)
		}

	case GramaticaParserNUM:
		localctx = NewNumeroContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.Match(GramaticaParserNUM)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
