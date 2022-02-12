// Code generated from Gramatica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GramaticaListener is a complete listener for a parse tree produced by GramaticaParser.
type GramaticaListener interface {
	antlr.ParseTreeListener

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterSum is called when entering the Sum production.
	EnterSum(c *SumContext)

	// EnterEpExp is called when entering the EpExp production.
	EnterEpExp(c *EpExpContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterMul is called when entering the Mul production.
	EnterMul(c *MulContext)

	// EnterEpTerm is called when entering the EpTerm production.
	EnterEpTerm(c *EpTermContext)

	// EnterParentesis is called when entering the parentesis production.
	EnterParentesis(c *ParentesisContext)

	// EnterNumero is called when entering the numero production.
	EnterNumero(c *NumeroContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitSum is called when exiting the Sum production.
	ExitSum(c *SumContext)

	// ExitEpExp is called when exiting the EpExp production.
	ExitEpExp(c *EpExpContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitMul is called when exiting the Mul production.
	ExitMul(c *MulContext)

	// ExitEpTerm is called when exiting the EpTerm production.
	ExitEpTerm(c *EpTermContext)

	// ExitParentesis is called when exiting the parentesis production.
	ExitParentesis(c *ParentesisContext)

	// ExitNumero is called when exiting the numero production.
	ExitNumero(c *NumeroContext)
}
