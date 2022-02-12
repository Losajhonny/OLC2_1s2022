package listener

import (
	"Ejercicio1/parser"
	"strconv"
)

type Listener struct {
	*parser.BaseGramaticaListener
	stack []int
}

func (l *Listener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *Listener) Pop() int {
	if len(l.stack) < 1 {
		panic("empty stack")
	}
	result := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]
	return result
}

// ExitStat is called when exiting the stat production.
func (l *Listener) ExitStat(c *parser.StatContext) {

}

// ExitExp is called when exiting the exp production.
func (l *Listener) ExitExp(c *parser.ExpContext) {

}

// ExitSum is called when exiting the Sum production.
func (l *Listener) ExitSum(c *parser.SumContext) {
	right, left := l.Pop(), l.Pop()
	l.push(left + right)
}

// ExitEpExp is called when exiting the EpExp production.
func (l *Listener) ExitEpExp(c *parser.EpExpContext) {

}

// ExitTerm is called when exiting the term production.
func (l *Listener) ExitTerm(c *parser.TermContext) {

}

// ExitMul is called when exiting the Mul production.
func (l *Listener) ExitMul(c *parser.MulContext) {
	right, left := l.Pop(), l.Pop()
	l.push(left * right)
}

// ExitEpTerm is called when exiting the EpTerm production.
func (l *Listener) ExitEpTerm(c *parser.EpTermContext) {

}

// ExitParentesis is called when exiting the parentesis production.
func (l *Listener) ExitParentesis(c *parser.ParentesisContext) {

}

// ExitNumero is called when exiting the numero production.
func (l *Listener) ExitNumero(c *parser.NumeroContext) {
	num, err := strconv.Atoi(c.NUM().GetText())
	if err != nil {
		panic(err.Error())
	}
	l.push(num)
}
