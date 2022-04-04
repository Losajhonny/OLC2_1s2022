package entorno

type Simbolo struct {
	Id            string
	TipoDato      string
	TipoSimbolo   string
	Ambito        string
	Referencia    bool
	Dir           int
	NoDimensiones int
	Dimensiones   []*Dimension
}

type Dimension struct {
	Linferior int
	Lsuperior int
}

func NewVar(id string, tipoDato string /*tipoSimbolo string, referencia bool, ambito string,*/, dir int) *Simbolo {
	sim := new(Simbolo)
	sim.Id = id
	sim.TipoDato = tipoDato
	/*sim.TipoSimbolo = tipoSimbolo
	sim.Referencia = referencia
	sim.Ambito = ambito*/
	sim.Dir = dir
	return sim
}

func NewVarArray(id string, tipoDato string /*tipoSimbolo string, referencia bool, ambito string, */, dir int, noDims int, dims []*Dimension) *Simbolo {
	sim := NewVar(id, tipoDato /*tipoSimbolo, referencia, ambito,*/, dir)
	sim.NoDimensiones = noDims
	sim.Dimensiones = dims
	return sim
}

func NewDimension(inf int, sup int) *Dimension {
	dim := new(Dimension)
	dim.Linferior = inf
	dim.Lsuperior = sup
	return dim
}
