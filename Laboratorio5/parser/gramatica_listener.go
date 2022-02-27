// Code generated from Gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GramaticaListener is a complete listener for a parse tree produced by GramaticaParser.
type GramaticaListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterExp2 is called when entering the exp2 production.
	EnterExp2(c *Exp2Context)

	// EnterCond is called when entering the cond production.
	EnterCond(c *CondContext)

	// EnterMarcador is called when entering the marcador production.
	EnterMarcador(c *MarcadorContext)

	// EnterExpresion is called when entering the expresion production.
	EnterExpresion(c *ExpresionContext)

	// EnterOprel is called when entering the oprel production.
	EnterOprel(c *OprelContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterInst_if is called when entering the inst_if production.
	EnterInst_if(c *Inst_ifContext)

	// EnterInst_asignacion is called when entering the inst_asignacion production.
	EnterInst_asignacion(c *Inst_asignacionContext)

	// EnterBloque is called when entering the bloque production.
	EnterBloque(c *BloqueContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitExp2 is called when exiting the exp2 production.
	ExitExp2(c *Exp2Context)

	// ExitCond is called when exiting the cond production.
	ExitCond(c *CondContext)

	// ExitMarcador is called when exiting the marcador production.
	ExitMarcador(c *MarcadorContext)

	// ExitExpresion is called when exiting the expresion production.
	ExitExpresion(c *ExpresionContext)

	// ExitOprel is called when exiting the oprel production.
	ExitOprel(c *OprelContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitInst_if is called when exiting the inst_if production.
	ExitInst_if(c *Inst_ifContext)

	// ExitInst_asignacion is called when exiting the inst_asignacion production.
	ExitInst_asignacion(c *Inst_asignacionContext)

	// ExitBloque is called when exiting the bloque production.
	ExitBloque(c *BloqueContext)
}
