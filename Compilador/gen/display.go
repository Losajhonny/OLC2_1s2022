package gen

type DNodo struct {
	Salida string
	Inicio string
	ContB  int
}

var Ptr int = 0

// pos 0 no tienen nada
var Display [10000]DNodo

func NextPtr() {
	Ptr = Ptr + 1
}

func PrevPtr() {
	Ptr = Ptr - 1
}
