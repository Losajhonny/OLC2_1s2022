package bloques

import "github.com/colegno/arraylist"

/*
	Algoritmo 8.5: Particionar instrucciones de tres direcciones en bloques básicos.

	ENTRADA: Una secuencia de instrucciones de tres direcciones.

	SALIDA: Una lista de los bloques básicos para la secuencia en la que cada instrucción se asigna exactamente a un bloque básico.

	MÉTODO:
		En primer lugar, determinamos las instrucciones en el código intermedio que son
		líderes; es decir, las primeras instrucciones en algún bloque básico. La instrucción que va justo
		después del programa intermedio no se incluye como líder. Las reglas para buscar líderes son:

		1. La primera instrucción de tres direcciones en el código intermedio es líder.
		2. Cualquier instrucción que sea el destino de un salto condicional o incondicional es líder.
		3. Cualquier instrucción que siga justo después de un salto condicional o incondicional es líder.
*/

// permite label lider
func DeterminarLider(insts *arraylist.List) {
	// primera instruccion lider
	lider := true
	for i := 0; i < insts.Len(); i++ {
		inst := insts.GetValue(i).(*Inst)
		inst.Lider = lider
		lider = false

		if inst.Tipo == ETIQUETA {
			// cualquier instruccion que sea destino de un salto condicional o incondicional es lider
			lider = true
		} else if inst.Tipo == IF {
			// cualquier instruccion que siga justo después de un salto condicional es lider
			if i+1 < insts.Len() {
				next := insts.GetValue(i + 1).(*Inst)
				// si la siguiente instruccion es goto, entonces no es lider
				if next.Tipo != GOTO {
					lider = true
				}
			}
		} else if inst.Tipo == GOTO {
			// cualquier instruccion que siga justo después de un salto incondicional es lider
			lider = true
		} else {
			// cualquier label puede ser lider
			if i+1 < insts.Len() {
				next := insts.GetValue(i + 1).(*Inst)
				if next.Tipo == ETIQUETA {
					lider = true
				}
			}
		}
	}
}

// manejar el label lider
func GenerarBloques(insts *arraylist.List) *arraylist.List {
	var lbs *arraylist.List = arraylist.New()
	var bloques *arraylist.List = arraylist.New()
	var bloque *Bloque

	for i := 0; i < insts.Len(); i++ {
		inst := insts.GetValue(i).(*Inst)
		if inst.Lider {
			if inst.Tipo == ETIQUETA {
				lbs.Add(inst)
			} else {
				bloque = NewBloque("B" + GetId())
				bloque.Labels = lbs
				bloque.Insts.Add(inst)

				// dejar labels globales
				for _, lbl_ := range lbs.ToArray() {
					lbl := lbl_.(*Inst).GetLabel()
					Labels[lbl] = bloque
				}

				lbs = arraylist.New()
				bloques.Add(bloque)
			}
		} else {
			bloque.Insts.Add(inst)
		}
	}
	return bloques
}

/*
	Grafos de flujo

	Una vez que un programa de código intermedio se particiona en bloques básicos, representamos
	el flujo de control entre ellos mediante un grafo de flujo. Los nodos del grafo de flujo son los
	nodos básicos. Hay una flecha del bloque B al bloque C sí, y sólo si es posible que la primera
	instrucción en el bloque C vaya justo después de la última instrucción en el bloque B. Hay dos
	formas en las que podría justificarse dicha flecha:

	*	Hay un salto condicional o incondicional desde el final de B hasta el inicio de C.
	*	C sigue justo después de B en el orden original de las instrucciones de tres direcciones, y
		B no termina en un salto incondicional.
*/
func GenerarGrafo(bloques *arraylist.List) {
	for i := 0; i < bloques.Len(); i++ {
		bloque := bloques.GetValue(i).(*Bloque)

		for j := 0; j < bloque.Insts.Len(); j++ {
			inst := bloque.Insts.GetValue(j).(*Inst)

			// Si hay un salto condicional o incondicional desde el final de B hasta el inicio de C.
			if inst.Tipo == IF || inst.Tipo == GOTO {
				label := inst.GetLabel()
				// ref de bloque
				ref, _ := Labels[label]
				bloque.Refs.Add(ref)
				inst.SetLabel(label, ref.GetId())
			}

			// si no es última inst continuar
			if j != bloque.Insts.Len()-1 {
				continue
			}

			// C sigue justo después de B en el orden original de las instrucciones de tres direcciones, y
			// B no termina en un salto incondicional
			if inst.Tipo != GOTO {
				if i+1 < bloques.Len() {
					next := bloques.GetValue(i + 1).(*Bloque)
					bloque.Refs.Add(next)
				}
			}
		}
	}
}

// Generar grafo general
func GetDotGrafo(nodos *arraylist.List) string {
	cadena := "digraph {\n"
	cadena += "node[fontname=\"Helvetica,Arial,sans-serif\", shape=rectangle, color=lightblue2, style=filled]\n"
	for _, nodo_ := range nodos.ToArray() {
		nodo := nodo_.(*Caja)
		cadena += nodo.GetDotGrafo()
	}

	var nodoAux *Caja
	for _, nodo_ := range nodos.ToArray() {
		nodo := nodo_.(*Caja)
		if nodoAux != nil {
			cadena += nodoAux.Bloques.GetValue(nodoAux.Bloques.Len()-1).(*Bloque).GetId() + " -> " + nodo.Bloques.GetValue(0).(*Bloque).GetId() + "\n"
		}
		nodoAux = nodo
	}
	// nodos de entrada y salida
	nodoIni := nodos.GetValue(0).(*Caja)
	nodoFin := nodos.GetValue(nodos.Len() - 1).(*Caja)

	bIni := nodoIni.Bloques.GetValue(0).(*Bloque)
	bFin := nodoFin.Bloques.GetValue(nodoFin.Bloques.Len() - 1).(*Bloque)

	nEntrada := "N" + GetId()
	nSalida := "N" + GetId()

	cadena += nEntrada + "[label=\"Entrada\",tooltip=\"Entrada\"]\n"
	cadena += nSalida + "[label=\"Salida\",tooltip=\"Salida\"]\n"
	cadena += nEntrada + " -> " + bIni.GetId() + "\n"
	cadena += bFin.GetId() + " -> " + nSalida + "\n"
	cadena += "}\n"
	return cadena
}
