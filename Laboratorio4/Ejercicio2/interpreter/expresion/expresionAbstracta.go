package expresion

import "Ejercicio2/interpreter/tipos"

type ExpresionAbstracta interface {
	GetTipo() tipos.TipoAbstracto
	GetValor() string
}
