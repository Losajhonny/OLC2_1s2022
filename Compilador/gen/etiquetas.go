package gen

import (
	"fmt"
	"strconv"
)

var label = 0

func NewLabel() string {
	label = label + 1
	return "L" + strconv.Itoa(label)
}

func Soltar(lb []string) {
	for _, label := range lb {
		AddLabel(label)
	}
}

func Unir(lb1 []string, lb2 []string) []string {
	lista := lb1
	for _, label := range lb2 {
		lista = append(lista, label)
	}
	return lista
}

func Mostrar(lv []string, lf []string) {
	fmt.Println("\nEtiquetas verdaderas")
	Soltar(lv)
	fmt.Println("\nEtiquetas falsas")
	Soltar(lf)
}
