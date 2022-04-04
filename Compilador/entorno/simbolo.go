package entorno

type Simbolo struct {
	Id          string
	Tipo        string
	Dir         int
	NoDimension int
	Dimensiones []*Dimension
}

type Dimension struct {
	Linferior int
	Lsuperior int
}

func NewVar(id string, tipo string, dir int) *Simbolo {
	sim := new(Simbolo)
	sim.Id = id
	sim.Tipo = tipo
	sim.Dir = dir
	return sim
}

func NewVarArray(id string, tipo string, dir int, noDim int, dims []*Dimension) *Simbolo {
	sim := NewVar(id, tipo, dir)
	sim.NoDimension = noDim
	sim.Dimensiones = dims
	return sim
}

func NewDimension(inf int, sup int) *Dimension {
	dim := new(Dimension)
	dim.Linferior = inf
	dim.Lsuperior = sup
	return dim
}
