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

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterInst_asignacion is called when entering the inst_asignacion production.
	EnterInst_asignacion(c *Inst_asignacionContext)

	// EnterLref is called when entering the lref production.
	EnterLref(c *LrefContext)

	// EnterInst_declaracion is called when entering the inst_declaracion production.
	EnterInst_declaracion(c *Inst_declaracionContext)

	// EnterLdims is called when entering the ldims production.
	EnterLdims(c *LdimsContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// EnterInst_if is called when entering the inst_if production.
	EnterInst_if(c *Inst_ifContext)

	// EnterInst_switch_propuesta1 is called when entering the inst_switch_propuesta1 production.
	EnterInst_switch_propuesta1(c *Inst_switch_propuesta1Context)

	// EnterInst_switch_propuesta2 is called when entering the inst_switch_propuesta2 production.
	EnterInst_switch_propuesta2(c *Inst_switch_propuesta2Context)

	// EnterInst_while is called when entering the inst_while production.
	EnterInst_while(c *Inst_whileContext)

	// EnterInst_doWhile is called when entering the inst_doWhile production.
	EnterInst_doWhile(c *Inst_doWhileContext)

	// EnterInst_loop is called when entering the inst_loop production.
	EnterInst_loop(c *Inst_loopContext)

	// EnterBloque is called when entering the bloque production.
	EnterBloque(c *BloqueContext)

	// EnterBloqueSinLLaves is called when entering the bloqueSinLLaves production.
	EnterBloqueSinLLaves(c *BloqueSinLLavesContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

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

	// ExitInst_asignacion is called when exiting the inst_asignacion production.
	ExitInst_asignacion(c *Inst_asignacionContext)

	// ExitLref is called when exiting the lref production.
	ExitLref(c *LrefContext)

	// ExitInst_declaracion is called when exiting the inst_declaracion production.
	ExitInst_declaracion(c *Inst_declaracionContext)

	// ExitLdims is called when exiting the ldims production.
	ExitLdims(c *LdimsContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)

	// ExitInst_if is called when exiting the inst_if production.
	ExitInst_if(c *Inst_ifContext)

	// ExitInst_switch_propuesta1 is called when exiting the inst_switch_propuesta1 production.
	ExitInst_switch_propuesta1(c *Inst_switch_propuesta1Context)

	// ExitInst_switch_propuesta2 is called when exiting the inst_switch_propuesta2 production.
	ExitInst_switch_propuesta2(c *Inst_switch_propuesta2Context)

	// ExitInst_while is called when exiting the inst_while production.
	ExitInst_while(c *Inst_whileContext)

	// ExitInst_doWhile is called when exiting the inst_doWhile production.
	ExitInst_doWhile(c *Inst_doWhileContext)

	// ExitInst_loop is called when exiting the inst_loop production.
	ExitInst_loop(c *Inst_loopContext)

	// ExitBloque is called when exiting the bloque production.
	ExitBloque(c *BloqueContext)

	// ExitBloqueSinLLaves is called when exiting the bloqueSinLLaves production.
	ExitBloqueSinLLaves(c *BloqueSinLLavesContext)
}
