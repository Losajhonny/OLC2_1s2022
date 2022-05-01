package bloques

import (
	"fmt"
)

/**
Funciones de caja
*/
func (caja *Caja) OptimizarLocal() {
	for _, bloque_ := range caja.Bloques.ToArray() {
		bloque := bloque_.(*Bloque)
		bloque.OptimizarSubExpresionesComunesLocal()
		bloque.OptimizarPropagacionDeCopiasLocal()
	}
}

func (caja *Caja) OptimizarSubExpresionesComunesLocal() {
	for _, bloque_ := range caja.Bloques.ToArray() {
		bloque := bloque_.(*Bloque)
		bloque.OptimizarSubExpresionesComunesLocal()
	}
}

func (caja *Caja) OptimizarPropagacionDeCopiasLocal() {
	for _, bloque_ := range caja.Bloques.ToArray() {
		bloque := bloque_.(*Bloque)
		bloque.OptimizarPropagacionDeCopiasLocal()
	}
}

/**
Funciones de bloque
*/
func (bloque *Bloque) OptimizarSubExpresionesComunesLocal() {
	for i := 0; i < bloque.Insts.Len(); i++ {
		actual := bloque.Insts.GetValue(i).(*Inst)
		cambio := false
		if actual.GetTipo() != ASIGNACION {
			continue
		}
		if actual.EstaCambiandoVariables(actual) {
			cambio = true
		}
		for j := i + 1; j < bloque.Insts.Len(); j++ {
			next := bloque.Insts.GetValue(j).(*Inst)
			if next.GetTipo() != ASIGNACION {
				continue
			}
			// se cambiaron las variables?
			if actual.EstaCambiandoVariables(next) {
				cambio = true
			}
			// detectar subexpresion comun
			if actual.EsSubExpComun(next) {
				if !cambio {
					fmt.Println("****************************************")
					fmt.Println("Subexpresion comun detectado")
					fmt.Println("\tExpresion: " + actual.Codigo)
					fmt.Println("\tSubexpres: " + next.Codigo)
					fmt.Println("Expresión eliminada")
					fmt.Println("\tSubexpres: " + next.Codigo)
					fmt.Println("Expresión actualizada:")
					// reemplazar variables
					for k := j + 1; k < bloque.Insts.Len(); k++ {
						reemplazar := bloque.Insts.GetValue(k).(*Inst)
						expAnt := reemplazar.Codigo
						if reemplazar.CambiarVariable(actual, next) {
							fmt.Println("\t\tanterior: " + expAnt)
							fmt.Println("\t\tactual  : " + reemplazar.Codigo)
							fmt.Println("\t\t----------")
						}
					}
					next.Codigo = ""
					next.Eliminado = true
					fmt.Println("****************************************")
					fmt.Println()
					break
				}
			}
		}
	}
}

func (bloque *Bloque) OptimizarPropagacionDeCopiasLocal() {
	for i := 0; i < bloque.Insts.Len(); i++ {
		actual := bloque.Insts.GetValue(i).(*Inst)
		cambio := false
		if actual.GetTipo() != ASIGNACION {
			continue
		}
		for j := i + 1; j < bloque.Insts.Len(); j++ {
			next := bloque.Insts.GetValue(j).(*Inst)
			if next.GetTipo() != ASIGNACION {
				continue
			}
			// se cambiaron las variables?
			if actual.EstaCambiandoVariables(next) {
				cambio = true
			}
			// detectar propagacion de copia
			if actual.EsPropDeCopia(next) {
				if !cambio {
					fmt.Println("****************************************")
					fmt.Println("Propagación de copia detectado")
					fmt.Println("\tActual: " + actual.Codigo)
					fmt.Println("\tCopia: " + next.Codigo)
					fmt.Println("Expresión actualizada")
					fmt.Println("\tAntes: " + next.Codigo)
					next.AplicarProp(actual)
					fmt.Println("\tAntes: " + next.Codigo)
					fmt.Println("****************************************")
					fmt.Println()
					break
				}
			}
		}
	}
}
