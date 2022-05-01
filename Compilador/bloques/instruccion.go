package bloques

import "strings"

const (
	COMMENT     string = "COMMENT"
	DECLARACION        = "DECLARACION"
	INCLUDE            = "INCLUDE"
	ASIGNACION         = "ASIGNACION"
	ETIQUETA           = "ETIQUETA"
	IF                 = "IF"
	GOTO               = "GOTO"
	PRINT              = "PRINT"
	INIMETODO          = "INIMETODO"
	FINMETODO          = "FINMETODO"
	RETURN             = "RETURN"
	LLAMADA            = "LLAMADA"
)

/*
**************************************************
	 Instruccion
		Almacena la informacion de la instruccion
**************************************************
*/
type Inst struct {
	Tipo      string
	Lider     bool
	Eliminado bool
	Codigo    string
	Linea     int
}

type IInst interface {
	EsEliminado() bool
	GetTipo() string
	GetLabel() string
	SetLabel(label string)
	GetDotInst() string
	EsSubExpComun(next *Inst) bool
	EsPropDeCopia(next *Inst) bool
	AplicarProp(actual *Inst)
	EstaCambiandoVariables(next *Inst) bool
	CambiarVariable(cambio *Inst) bool
}

/**
Compruba si la instruccion ha sido eliminada
*/
func (inst *Inst) EsEliminado() bool {
	return inst.Eliminado
}

func (inst *Inst) GetTipo() string {
	return inst.Tipo
}

/**
Se obtiene la etiqueta sin ; : y espacios
de una etiqueta destino, if o goto
*/
func (inst *Inst) GetLabel() string {
	if inst.GetTipo() == ETIQUETA {
		return strings.ReplaceAll(inst.Codigo, ":", "")
	} else if inst.GetTipo() == IF || inst.GetTipo() == GOTO {
		split := strings.Split(inst.Codigo, "goto")
		lbl := split[1]
		lbl = strings.ReplaceAll(lbl, " ", "")
		lbl = strings.ReplaceAll(lbl, ";", "")
		return lbl
	} else {
		return ""
	}
}

/**
Se modifica la etiqueta de if o goto para saltos de bloques
*/
func (inst *Inst) SetLabel(old string, new string) {
	inst.Codigo = strings.ReplaceAll(inst.Codigo, old, new)
}

// representacion en dot grapvhiz
func (inst *Inst) GetDotInst() string {
	return strings.ReplaceAll(inst.Codigo, "\"", "\\\"")
}

/**
Solo detecta si es una subexpresion comun
inst = inst actual
next = inst para comparar
*/
func (inst *Inst) EsSubExpComun(next *Inst) bool {
	/*
		a = b + c	= exp
		...
		x = b + c	= sub
	*/
	if inst.EsEliminado() || next.EsEliminado() {
		return false
	}
	inst_split := strings.Split(inst.Codigo, "=")
	exp := strings.ReplaceAll(inst_split[1], " ", "")
	exp = strings.ReplaceAll(exp, ";", "")
	next_split := strings.Split(next.Codigo, "=")
	sub := strings.ReplaceAll(next_split[1], " ", "")
	sub = strings.ReplaceAll(sub, ";", "")
	return exp == sub
}

/**
Solo detecta si es propagacion de copias
inst = inst actual
next = inst para comparar
*/
func (inst *Inst) EsPropDeCopia(next *Inst) bool {
	/*
		a = b + c	= tmp
		...
		x = a		= exp

		a = b[a]
		...
		x[c] = a
	*/
	if inst.EsEliminado() || next.EsEliminado() {
		return false
	}
	inst_split := strings.Split(inst.Codigo, "=")
	tmp := strings.ReplaceAll(inst_split[0], " ", "")
	next_split := strings.Split(next.Codigo, "=")
	exp := strings.ReplaceAll(next_split[1], " ", "")
	exp = strings.ReplaceAll(exp, ";", "")

	// comparar si inst actual no tiene asignado un acceso a rreglo
	// o si la instruccion a comparar no se este cambiando el valor de un acceso arreglo
	return tmp == exp && !(strings.Contains(inst_split[1], "[") && strings.Contains(next_split[0], "["))
}

/**
aplica la propagacion de copia
inst = instruccion que fue comparado
actual = instruccion que donde esta la expresion
*/
func (inst *Inst) AplicarProp(actual *Inst) {
	/*
		a = b + c	= actual
		...
		x = a		= inst

		x = b + c
	*/
	if inst.EsEliminado() || actual.EsEliminado() {
		return
	}
	inst_split := strings.Split(inst.Codigo, "=")
	actual_split := strings.Split(actual.Codigo, "=")
	newInst := inst_split[0] + "=" + actual_split[1]
	inst.Codigo = newInst
}

/**
Comprueba si las variables no han sido cambiados
inst = instruccion actual
next = instruccion para detectar cambio
*/
func (inst *Inst) EstaCambiandoVariables(next *Inst) bool {
	/*
		h = h + 1	es mejor no cambiar
		...
		h = h + 1


		izq = der
		a = b + c
		...
		x = b + c
		d = a + c
	*/
	if inst.EsEliminado() || next.EsEliminado() {
		return false
	}
	// instruccion actual
	inst_split := strings.Split(inst.Codigo, "=")
	inst_der := strings.ReplaceAll(inst_split[1], " ", "")
	// instruccion siguiente
	next_split := strings.Split(next.Codigo, "=")
	next_izq := strings.ReplaceAll(next_split[0], " ", "")
	// comprobacion
	if strings.Contains(next_izq, "[") {
		/*
			t1 = a[t1]		comprobrar que a no cambie
			...
			a[t2] = 10
		*/
		// separar a[t2]
		// izq   der
		// a  -- t2]
		return strings.Contains(inst_der, strings.Split(next_izq, "[")[0])
	} else {
		/*
			t1 = a[t1]		comprobrar que t1 no cambie
			...
			t1 = 10
			---------------------------
			a = b + c		comprobrar que b o c no cambie
			b = 10
		*/
		return strings.Contains(inst_der, next_izq)
	}
}

/**
Sucede cuando ocurre una subexpresion comun
a = b + c
...
x = b + c
...
d = x + d

ahora sería
a = b + c
...
...
d = a + d

la x se debe cambiar en a

inst: d = x + d
cambio: a = b + c
delete: x = b + c
*/
func (inst *Inst) CambiarVariable(cambio *Inst, delete *Inst) bool {
	cambio_split := strings.Split(cambio.Codigo, "=")
	cambio_izq := strings.ReplaceAll(cambio_split[0], " ", "")

	delete_split := strings.Split(delete.Codigo, "=")
	delete_izq := strings.ReplaceAll(delete_split[0], " ", "")
	if strings.Contains(inst.Codigo, delete_izq) {
		inst.Codigo = strings.ReplaceAll(inst.Codigo, delete_izq, cambio_izq)
		return true
	}
	return false
}

func NewInst(linea int) *Inst {
	inst := new(Inst)
	inst.Linea = linea
	return inst
}
