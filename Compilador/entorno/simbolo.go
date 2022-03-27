package entorno

type Simbolo struct {
	Id   string
	Tipo string
	Dims []*Dim
	Dir  int
}

type Dim struct {
	Linf int
	Lsup int
}

func NewSimbolo(id string, tipo string, dir int) *Simbolo {
	return &Simbolo{Id: id, Tipo: tipo, Dir: dir}
}

func NewSimboloArr(id string, dims []*Dim, dir int) *Simbolo {
	return &Simbolo{Id: id, Dims: dims, Dir: dir}
}

func NewDim(inf int, sup int) *Dim {
	return &Dim{Linf: inf, Lsup: sup}
}
