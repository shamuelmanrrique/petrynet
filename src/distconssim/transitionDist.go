package distconssim

import (
	"fmt"
)

// TransitionList is a list of transitions themselves
type TransitionList []TransitionDist //Slice de transiciones como Lista

// Length return length of TransitionList with type adapted to IndLocalTrans
func (self TransitionList) Length() IndGlobalTrans {
	return IndGlobalTrans(len(self))
}

//--------------------------------------------------------------------------

// IndGlobalTrans is a index of a transition in the global list
type IndGlobalTrans int32

// IndLocalTrans is a index of a transition in the local lefs list
type IndLocalTrans int32

//TypeConst is the constant to propagate in lefs
type TypeConst int32

// TransitionConstant is
type TransitionConstant struct {
	INextTrans IndGlobalTrans
	// INextTransGl IndGlobalTrans
	Const TypeConst
}

//------------------------------------------------------------------------

// -----------------------------------------------------------------------
// Tipo abstracto Transition para guardar la informacion de una transicion
// -----------------------------------------------------------------------
type TransitionDist struct {
	IDGlobal IndGlobalTrans // indice en la tabla global de transiciones
	IDLocal  IndLocalTrans
	// iiValorLef es el valor que tiene la funcion de
	// sensibilizacion en el instante de tiempo que nos da
	// la variable ITime
	IiValorLef TypeConst
	ITime      TypeClock

	// tiempo que dura el disparo de la transicion
	IiShotDuration TypeClock

	// vector de transiciones a las que tengo que propagar cte
	// cuando se dispare esta transicion, junto con la cte que
	// tengo que propagar
	IiListactes []TransitionConstant
}

/*
	-----------------------------------------------------------------
	   METODO: PrintEvent
	   RECIBE: Nada
	   DEVUELVE: Nada
	   PROPOSITO: Imprimir los atributos de la clase para depurar errores
		-----------------------------------------------------------------
*/
func (self *TransitionDist) PrintEvent() {
	fmt.Println("Dato Transicion:")
	fmt.Println("IDGLOBAL: ", self.IDGlobal)
	fmt.Println(" VALOR LEF: ", self.IiValorLef)
	fmt.Println(" TIEMPO: ", self.ITime)
	fmt.Println(" DURACION DISPARO: ", self.IiShotDuration)
	fmt.Println(" LISTA DE CTES: ")
	for _, v := range self.IiListactes {
		fmt.Println("\tTRANSICION: ", v.INextTrans, "\t\tCTE: ", v.Const)
	}
}

/*
	-----------------------------------------------------------------
   METODO: PrintEventValues
   RECIBE: Nada
   DEVUELVE: Nada
   PROPOSITO: Imprimir simplemente el valor de la transicion
	COMENTARIO : es solo de lectura
	-----------------------------------------------------------------
*/
func (self TransitionDist) PrintEventValues() {
	fmt.Println("Transicion -> ")
	fmt.Println("\tIDGLOBAL: ", self.IDGlobal)
	fmt.Println("\t\tVALOR LEF: ", self.IiValorLef)
	fmt.Println("\t\tTIEMPO: ", self.ITime)
}

//----------------------------------------------------------------------

// Stack Transition is a Stack of transitions indices
type StackTransitions []IndGlobalTrans

//Push transition id to stack
func (self *StackTransitions) push(i_tr IndGlobalTrans) {
	*self = append(*self, i_tr)
}

//Pop transition id from stack
func (self *StackTransitions) pop() IndGlobalTrans {
	if (*self).isEmpty() {
		return -1
	} else {
		i_tr := (*self)[len(*self)-1]  // obtener dato de lo alto de la pila
		*self = (*self)[:len(*self)-1] //desempilar
		return i_tr
	}
}

func (self StackTransitions) isEmpty() bool {
	return len(self) == 0
}
