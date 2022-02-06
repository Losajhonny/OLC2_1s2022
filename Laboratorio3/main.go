package main

import (
	"Laboratorio3/parser"
	"bufio"
	"log"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// referencia para la interfaz BaseListener
type TreeSyntax struct {
	*parser.BaseEjercicio1Listener
}

func NewTreeSyntax() *TreeSyntax {
	return new(TreeSyntax)
}

func interpretar(in string) {
	is := antlr.NewInputStream(in)

	// crear analizador lexico
	lexer := parser.NewEjercicio1Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// crear analizador sintactico
	p := parser.NewEjercicio1Parser(stream)

	antlr.ParseTreeWalkerDefault.Walk(NewTreeSyntax(), p.Start())
}

func main() {
	file, err := os.Open("entrada.txt")

	if err != nil {
		log.Fatal(err)
	}
	// se invoca hasta el final
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	text := ""

	for fileScanner.Scan() {
		text += fileScanner.Text()
	}

	// accion de interpretar
	interpretar(text)
}
