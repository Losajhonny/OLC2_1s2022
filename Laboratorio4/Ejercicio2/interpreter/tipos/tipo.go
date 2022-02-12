package tipos

type Tipo struct {
	Nombre string
}

func NewTipo(nombre string) Tipo {
	return Tipo{Nombre: nombre}
}

func (t *Tipo) ToString() string {
	return t.Nombre
}

// type TipoArreglo struct {
// 	NoDim int
// 	De    TipoAbstracto
// }

// func NewTipoArreglo(noDim int, de TipoAbstracto) TipoArreglo {
// 	return TipoArreglo{NoDim: noDim, De: de}
// }

// func (t *TipoArreglo) ToString() string {
// 	nombre := ""
// 	for i := 0; i < t.NoDim; i++ {
// 		nombre += "[" + t.ToString() + "]"
// 	}
// 	return nombre
// }
