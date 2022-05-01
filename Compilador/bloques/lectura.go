package bloques

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/colegno/arraylist"
)

func LecturaC3D(path string) *arraylist.List {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fscan := bufio.NewScanner(file)

	var caja *Caja
	lista := arraylist.New()
	noLinea := 0

	for fscan.Scan() {
		noLinea = noLinea + 1
		text := fscan.Text()
		text = strings.ReplaceAll(text, "\t", "")
		text = strings.ReplaceAll(text, "\n", "")
		text = strings.ReplaceAll(text, "\r", "")
		fmt.Println(text)
		if text == "" {
			continue
		}
		inst := GetInstruccion(text, noLinea)

		if inst.Tipo == INCLUDE {
			caja = NewCaja("cluster_" + GetId())
			lista.Add(caja)
		} else if inst.Tipo == INIMETODO {
			caja = NewCaja("cluster_" + GetId())
			lista.Add(caja)
		}
		caja.Insts.Add(inst)
	}
	return lista
}

func GetInstruccion(txt string, noLinea int) *Inst {
	nodo := new(Inst)
	nodo.Codigo = txt
	nodo.Linea = noLinea
	if strings.HasPrefix(txt, "//") || strings.HasPrefix(txt, "/*") {
		nodo.Tipo = COMMENT
	} else if strings.HasPrefix(txt, "#") {
		nodo.Tipo = INCLUDE
	} else if strings.HasPrefix(txt, "float") {
		nodo.Tipo = DECLARACION
	} else if strings.HasPrefix(txt, "if") {
		nodo.Tipo = IF
	} else if strings.HasPrefix(txt, "goto") {
		nodo.Tipo = GOTO
	} else if strings.HasSuffix(txt, ":") {
		nodo.Tipo = ETIQUETA
	} else if strings.HasPrefix(txt, "printf") {
		nodo.Tipo = PRINT
	} else if strings.HasPrefix(txt, "return") {
		nodo.Tipo = RETURN
	} else if strings.HasSuffix(txt, "{") {
		nodo.Tipo = INIMETODO
	} else if strings.Contains(txt, "}") {
		nodo.Tipo = FINMETODO
	} else if strings.Contains(txt, "()") && strings.HasSuffix(txt, ";") {
		nodo.Tipo = LLAMADA
	} else {
		nodo.Tipo = ASIGNACION
	}
	return nodo
}
