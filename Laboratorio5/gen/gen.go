package gen

import (
	"fmt"
	"strconv"
)

var tmp = 0
var eti = 0

func NewTemp() string {
	tmp = tmp + 1
	return "t" + strconv.Itoa(tmp)
}

func NewEti() string {
	eti = eti + 1
	return "L" + strconv.Itoa(eti)
}

func Soltar(eti []string) {
	for _, i := range eti {
		Gen(i + ":")
	}
}

func Unir(eti1 []string, eti2 []string) []string {
	lista := eti1
	for _, i := range eti2 {
		lista = append(lista, i)
	}
	return lista
}

// funcion para generar cadena
func Gen(out string) {
	fmt.Println(out)
}

func Mostrar(lv []string, lf []string) {
	fmt.Println("\nEtiquetas verdaderas")
	Soltar(lv)
	fmt.Println("\nEtiquetas falsas")
	Soltar(lf)
}
