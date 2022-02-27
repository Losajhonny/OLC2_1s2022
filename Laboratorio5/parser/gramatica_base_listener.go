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

// EnterStart is called when production start is entered.
func (s *BaseGramaticaListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseGramaticaListener) ExitStart(ctx *StartContext) {}

// EnterExp is called when production exp is entered.
func (s *BaseGramaticaListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseGramaticaListener) ExitExp(ctx *ExpContext) {}

// EnterExp2 is called when production exp2 is entered.
func (s *BaseGramaticaListener) EnterExp2(ctx *Exp2Context) {}

// ExitExp2 is called when production exp2 is exited.
func (s *BaseGramaticaListener) ExitExp2(ctx *Exp2Context) {}

// EnterCond is called when production cond is entered.
func (s *BaseGramaticaListener) EnterCond(ctx *CondContext) {}

// ExitCond is called when production cond is exited.
func (s *BaseGramaticaListener) ExitCond(ctx *CondContext) {}

// EnterMarcador is called when production marcador is entered.
func (s *BaseGramaticaListener) EnterMarcador(ctx *MarcadorContext) {}

// ExitMarcador is called when production marcador is exited.
func (s *BaseGramaticaListener) ExitMarcador(ctx *MarcadorContext) {}

// EnterExpresion is called when production expresion is entered.
func (s *BaseGramaticaListener) EnterExpresion(ctx *ExpresionContext) {}

// ExitExpresion is called when production expresion is exited.
func (s *BaseGramaticaListener) ExitExpresion(ctx *ExpresionContext) {}

// EnterOprel is called when production oprel is entered.
func (s *BaseGramaticaListener) EnterOprel(ctx *OprelContext) {}

// ExitOprel is called when production oprel is exited.
func (s *BaseGramaticaListener) ExitOprel(ctx *OprelContext) {}

// EnterInstrucciones is called when production instrucciones is entered.
func (s *BaseGramaticaListener) EnterInstrucciones(ctx *InstruccionesContext) {}

// ExitInstrucciones is called when production instrucciones is exited.
func (s *BaseGramaticaListener) ExitInstrucciones(ctx *InstruccionesContext) {}

// EnterInstruccion is called when production instruccion is entered.
func (s *BaseGramaticaListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BaseGramaticaListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterInst_if is called when production inst_if is entered.
func (s *BaseGramaticaListener) EnterInst_if(ctx *Inst_ifContext) {}

// ExitInst_if is called when production inst_if is exited.
func (s *BaseGramaticaListener) ExitInst_if(ctx *Inst_ifContext) {}

// EnterInst_asignacion is called when production inst_asignacion is entered.
func (s *BaseGramaticaListener) EnterInst_asignacion(ctx *Inst_asignacionContext) {}

// ExitInst_asignacion is called when production inst_asignacion is exited.
func (s *BaseGramaticaListener) ExitInst_asignacion(ctx *Inst_asignacionContext) {}

// EnterBloque is called when production bloque is entered.
func (s *BaseGramaticaListener) EnterBloque(ctx *BloqueContext) {}

// ExitBloque is called when production bloque is exited.
func (s *BaseGramaticaListener) ExitBloque(ctx *BloqueContext) {}
