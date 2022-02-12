/*
	go mod init Ejercicio1
	go get github.com/antlr/antlr4/runtime/Go/antlr
	antlr -o parser Gramatica.g4
	go run main.go
*/

package main

import (
	"Ejercicio1/listener"
	"Ejercicio1/parser"
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

	// activate listener
	listen := listener.Listener{}
	antlr.ParseTreeWalkerDefault.Walk(&listen, tree)

	// print value
	fmt.Println(input, "=", listen.Pop())
}

func main() {
	input := "2 + 5 * 3"
	analizar(input)
}
