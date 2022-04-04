package gen

import (
	"fmt"
)

func Gen(out string) {
	fmt.Print(out)
}

func Genln(out string) {
	fmt.Println(out)
}

func AddExpresionUnaria(dir string, op string, der string) {
	Genln(dir + " = " + op + " " + der)
}

func AddExpresion(dir string, izq string, op string, der string) {
	Genln(dir + " = " + izq + " " + op + " " + der)
}

func AddIf(izq string, op string, der string, label string) string {
	cad := izq + " " + op + " " + der
	Genln("if " + cad + " then goto " + label)
	return cad
}

func AddIfCad(cad string, label string) {
	Genln("if " + cad + " then goto " + label)
}

func AddGoto(label string) {
	Genln("goto " + label)
}

func AddLabel(label string) {
	Genln(label + ":")
}

func AddAsign(dir string, exp string) {
	Genln(dir + " = " + exp)
}

func AddSetArray(dir string, pos string, exp string) {
	Genln(dir + "[" + pos + "] = " + exp)
}

func AddGetArray(dir string, id string, pos string) {
	Genln(dir + " = " + id + "[" + pos + "]")
}
