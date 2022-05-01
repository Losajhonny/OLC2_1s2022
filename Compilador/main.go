/*
	go mod init Compilador
	go get github.com/antlr/antlr4/runtime/Go/antlr
	antlr -o parser Gramatica.g4
	go run main.go
*/

package main

import (
	"Compilador/bloques"
	"Compilador/parser"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Listener struct {
	*parser.BaseGramaticaListener
}

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
	listen := Listener{}
	antlr.ParseTreeWalkerDefault.Walk(&listen, tree)
}

func optimizarCodigo() {
	cajas := bloques.LecturaC3D("salida.txt")

	for _, caja_ := range cajas.ToArray() {
		caja := caja_.(*bloques.Caja)
		bloques.DeterminarLider(caja.Insts)
		caja.Bloques = bloques.GenerarBloques(caja.Insts)
		bloques.GenerarGrafo(caja.Bloques)
		caja.GetDotGrafo()
		caja.OptimizarSubExpresionesComunesLocal()
		caja.OptimizarPropagacionDeCopiasLocal()
	}

	dot := bloques.GetDotGrafo(cajas)
	fmt.Println(dot)
}

func main() {
	optimizarCodigo()
}
