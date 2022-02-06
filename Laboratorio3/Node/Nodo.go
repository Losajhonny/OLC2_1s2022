package node

type Nodo struct {
	Cad string
	Aux string
}

func NewNodo(cad string, aux string) Nodo {
	return Nodo{Cad: cad, Aux: aux}
}
