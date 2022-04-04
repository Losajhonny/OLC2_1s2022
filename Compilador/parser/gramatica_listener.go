// Code generated from Gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GramaticaListener is a complete listener for a parse tree produced by GramaticaParser.
type GramaticaListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterMarcador is called when entering the marcador production.
	EnterMarcador(c *MarcadorContext)

	// EnterExpresion is called when entering the expresion production.
	EnterExpresion(c *ExpresionContext)

	// EnterOprel is called when entering the oprel production.
	EnterOprel(c *OprelContext)

	// EnterLref is called when entering the lref production.
	EnterLref(c *LrefContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterInst_declaracion is called when entering the inst_declaracion production.
	EnterInst_declaracion(c *Inst_declaracionContext)

	// EnterLdims is called when entering the ldims production.
	EnterLdims(c *LdimsContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// EnterInst_asignacion is called when entering the inst_asignacion production.
	EnterInst_asignacion(c *Inst_asignacionContext)

	// EnterInst_if is called when entering the inst_if production.
	EnterInst_if(c *Inst_ifContext)

	// EnterInst_switch_propuesta1 is called when entering the inst_switch_propuesta1 production.
	EnterInst_switch_propuesta1(c *Inst_switch_propuesta1Context)

	// EnterInst_switch is called when entering the inst_switch production.
	EnterInst_switch(c *Inst_switchContext)

	// EnterInst_while is called when entering the inst_while production.
	EnterInst_while(c *Inst_whileContext)

	// EnterInst_doWhile is called when entering the inst_doWhile production.
	EnterInst_doWhile(c *Inst_doWhileContext)

	// EnterInst_loop is called when entering the inst_loop production.
	EnterInst_loop(c *Inst_loopContext)

	// EnterInst_for is called when entering the inst_for production.
	EnterInst_for(c *Inst_forContext)

	// EnterInst_break is called when entering the inst_break production.
	EnterInst_break(c *Inst_breakContext)

	// EnterInst_continue is called when entering the inst_continue production.
	EnterInst_continue(c *Inst_continueContext)

	// EnterBloque is called when entering the bloque production.
	EnterBloque(c *BloqueContext)

	// EnterBloqueSinLlaves is called when entering the bloqueSinLlaves production.
	EnterBloqueSinLlaves(c *BloqueSinLlavesContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitMarcador is called when exiting the marcador production.
	ExitMarcador(c *MarcadorContext)

	// ExitExpresion is called when exiting the expresion production.
	ExitExpresion(c *ExpresionContext)

	// ExitOprel is called when exiting the oprel production.
	ExitOprel(c *OprelContext)

	// ExitLref is called when exiting the lref production.
	ExitLref(c *LrefContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitInst_declaracion is called when exiting the inst_declaracion production.
	ExitInst_declaracion(c *Inst_declaracionContext)

	// ExitLdims is called when exiting the ldims production.
	ExitLdims(c *LdimsContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)

	// ExitInst_asignacion is called when exiting the inst_asignacion production.
	ExitInst_asignacion(c *Inst_asignacionContext)

	// ExitInst_if is called when exiting the inst_if production.
	ExitInst_if(c *Inst_ifContext)

	// ExitInst_switch_propuesta1 is called when exiting the inst_switch_propuesta1 production.
	ExitInst_switch_propuesta1(c *Inst_switch_propuesta1Context)

	// ExitInst_switch is called when exiting the inst_switch production.
	ExitInst_switch(c *Inst_switchContext)

	// ExitInst_while is called when exiting the inst_while production.
	ExitInst_while(c *Inst_whileContext)

	// ExitInst_doWhile is called when exiting the inst_doWhile production.
	ExitInst_doWhile(c *Inst_doWhileContext)

	// ExitInst_loop is called when exiting the inst_loop production.
	ExitInst_loop(c *Inst_loopContext)

	// ExitInst_for is called when exiting the inst_for production.
	ExitInst_for(c *Inst_forContext)

	// ExitInst_break is called when exiting the inst_break production.
	ExitInst_break(c *Inst_breakContext)

	// ExitInst_continue is called when exiting the inst_continue production.
	ExitInst_continue(c *Inst_continueContext)

	// ExitBloque is called when exiting the bloque production.
	ExitBloque(c *BloqueContext)

	// ExitBloqueSinLlaves is called when exiting the bloqueSinLlaves production.
	ExitBloqueSinLlaves(c *BloqueSinLlavesContext)
}
