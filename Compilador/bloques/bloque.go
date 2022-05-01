package bloques

import (
	"github.com/colegno/arraylist"
)

var Labels map[string]*Bloque = make(map[string]*Bloque)

type Bloque struct {
	// Id nombre del bloque
	Id     string
	Labels *arraylist.List
	Insts  *arraylist.List
	Refs   *arraylist.List
}

/*
**************************************************
	Bloque basico
		Almacena las instrucciones aplicando
		el algoritmo 8.5
**************************************************
*/

type IBloque interface {
	GetId() string
	GetDotGrafo() string
	// Implementado en el archivo optimizar.go
	OptimizarSubExpresionesComunesLocal()
	OptimizarPropagacionDeCopiasLocal()
}

func (bloque *Bloque) GetId() string {
	return bloque.Id
}

// Generar grafo solo de un bloque
func (bloque *Bloque) GetDotGrafo() string {
	cadena := ""
	for _, inst_ := range bloque.Insts.ToArray() {
		inst := inst_.(*Inst)
		cadena += inst.GetDotInst() + "\\n"
	}
	ncad := bloque.GetId() + "[label=\"" + cadena + "\",tooltip=\"" + bloque.GetId() + "\"]\n"
	return ncad
}

func NewBloque(id string) *Bloque {
	bloque := new(Bloque)
	bloque.Id = id
	bloque.Labels = arraylist.New()
	bloque.Insts = arraylist.New()
	bloque.Refs = arraylist.New()
	return bloque
}
