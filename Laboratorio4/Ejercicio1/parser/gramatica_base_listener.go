// Code generated from Gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGramaticaListener is a complete listener for a parse tree produced by GramaticaParser.
type BaseGramaticaListener struct{}

var _ GramaticaListener = &BaseGramaticaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGramaticaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGramaticaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGramaticaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGramaticaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseGramaticaListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseGramaticaListener) ExitStat(ctx *StatContext) {}

// EnterExp is called when production exp is entered.
func (s *BaseGramaticaListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseGramaticaListener) ExitExp(ctx *ExpContext) {}

// EnterSum is called when production Sum is entered.
func (s *BaseGramaticaListener) EnterSum(ctx *SumContext) {}

// ExitSum is called when production Sum is exited.
func (s *BaseGramaticaListener) ExitSum(ctx *SumContext) {}

// EnterEpExp is called when production EpExp is entered.
func (s *BaseGramaticaListener) EnterEpExp(ctx *EpExpContext) {}

// ExitEpExp is called when production EpExp is exited.
func (s *BaseGramaticaListener) ExitEpExp(ctx *EpExpContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseGramaticaListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseGramaticaListener) ExitTerm(ctx *TermContext) {}

// EnterMul is called when production Mul is entered.
func (s *BaseGramaticaListener) EnterMul(ctx *MulContext) {}

// ExitMul is called when production Mul is exited.
func (s *BaseGramaticaListener) ExitMul(ctx *MulContext) {}

// EnterEpTerm is called when production EpTerm is entered.
func (s *BaseGramaticaListener) EnterEpTerm(ctx *EpTermContext) {}

// ExitEpTerm is called when production EpTerm is exited.
func (s *BaseGramaticaListener) ExitEpTerm(ctx *EpTermContext) {}

// EnterParentesis is called when production parentesis is entered.
func (s *BaseGramaticaListener) EnterParentesis(ctx *ParentesisContext) {}

// ExitParentesis is called when production parentesis is exited.
func (s *BaseGramaticaListener) ExitParentesis(ctx *ParentesisContext) {}

// EnterNumero is called when production numero is entered.
func (s *BaseGramaticaListener) EnterNumero(ctx *NumeroContext) {}

// ExitNumero is called when production numero is exited.
func (s *BaseGramaticaListener) ExitNumero(ctx *NumeroContext) {}
