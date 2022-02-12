package main

import (
	"Ejercicio2/parser"
	"Ejercicio2/visitor"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func analizar(input string) {
	// create input stream
	stream := antlr.NewInputStream(input)
	// create lexer
	lexer := parser.NewGramaticaLexer(stream)
	// create tokenStream
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// create parser
	parser := parser.NewGramaticaParser(tokenStream)
	// get tree
	tree := parser.Start()

	// patron visitor
	visit := visitor.Visitor{}

	// print value
	exp := visit.Visit(tree)
	fmt.Println("tipo", exp.GetTipo().ToString())
	fmt.Println(input, "=", exp.GetValor())
}

func main() {
	input := "2 + 2 * 5 + 3"
	analizar(input)
}
