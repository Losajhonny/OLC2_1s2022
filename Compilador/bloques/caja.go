package bloques

import "github.com/colegno/arraylist"

type Caja struct {
	// id para el grafo en graphviz
	Id      string
	Insts   *arraylist.List
	Bloques *arraylist.List
}

type ICaja interface {
	GetId() string
	GetDotGrafo() string
	// Implementado en el archivo optimizar.go
	OptimizarSubExpresionesComunesLocal()
	OptimizarPropagacionDeCopiasLocal()
	OptimizarLocal()
}

func (bloque *Caja) GetId() string {
	return bloque.Id
}

// Generar grafo solo de una caja
func (caja *Caja) GetDotGrafo() string {
	cadena := "subgraph " + caja.GetId() + " {\n"
	// nodos
	for _, bloque_ := range caja.Bloques.ToArray() {
		bloque := bloque_.(*Bloque)
		cadena += bloque.GetDotGrafo()
	}
	// flechas
	for _, bloque_ := range caja.Bloques.ToArray() {
		bloque := bloque_.(*Bloque)
		for _, ref_ := range bloque.Refs.ToArray() {
			ref := ref_.(*Bloque)
			cadena += bloque.GetId() + " -> " + ref.GetId() + "\n"
		}
	}
	cadena += "}\n"
	return cadena
}

func NewCaja(id string) *Caja {
	caja := new(Caja)
	caja.Id = id
	caja.Insts = arraylist.New()
	return caja
}
