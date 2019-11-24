/*
PROPOSITO: Guardar la informacion de los eventos necesarios para
   la simulacion
-----------------------------------------------------------------
*/
package centralsim

import "fmt"

type Event struct {
	// Tiempo para el que debemos considerar el evento
	ITime TypeClock
	// A que transicion (indice transicion en subred)
	ITransition IndLocalTrans
	// Constante que mandamos
	IConst TypeConst
}

/*
-----------------------------------------------------------------
   METODO: NewEvent
   RECIBE: Tiempo, transicion y cte del evento a crear
   DEVUELVE: Event
   PROPOSITO: Crear evento con todos los datos del nuevo evento creados
-----------------------------------------------------------------
*/
func NewEvent(ai_tiempo TypeClock, ai_transicion IndLocalTrans, ai_cte TypeConst) *Event {
	e := new(Event)
	e.SetTime(ai_tiempo)
	e.SetTransition(ai_transicion)
	e.SetConst(ai_cte)
	return e
}

/*
-----------------------------------------------------------------
   METODO: SetTime
   RECIBE: Tiempo
   DEVUELVE: Nada
   PROPOSITO: Modificar el tiempo del evento
-----------------------------------------------------------------
*/
func (self *Event) SetTime(ai_tiempo TypeClock) {
	self.ITime = ai_tiempo
}

/*
-----------------------------------------------------------------
   METODO: SetTransition
   RECIBE: Identificador de la transicion (indice en array
	   de transiciones de esa subred)
   DEVUELVE: Nada
   PROPOSITO: Modificar la transicion del evento
-----------------------------------------------------------------
*/
func (self *Event) SetTransition(ai_transicion IndLocalTrans) {
	self.ITransition = ai_transicion
}

/*
-----------------------------------------------------------------
   METODO: SetConst
   RECIBE: Cte a transmitir
   DEVUELVE: Nada
   PROPOSITO: Modificar la cte del evento
-----------------------------------------------------------------
*/
func (self *Event) SetConst(ai_cte TypeConst) {
	self.IConst = ai_cte
}

/*
-----------------------------------------------------------------
   METODO: GetTime
   RECIBE: Nada
   DEVUELVE: El atributo ITime
   PROPOSITO: Recoger el tiempo del evento
-----------------------------------------------------------------
*/
func (self Event) GetTime() TypeClock {
	return self.ITime
}

/*
-----------------------------------------------------------------
   METODO: GetTransition
   RECIBE: Nada
   DEVUELVE: El atributo ITransition
   PROPOSITO: Recoger la transicion del evento
-----------------------------------------------------------------
*/
func (self Event) GetTransition() IndLocalTrans {
	return self.ITransition
}

/*
-----------------------------------------------------------------
   METODO: getConst
   RECIBE: Nada
   DEVUELVE: El atributo IConst
   PROPOSITO: Recoger la cte del evento
-----------------------------------------------------------------
*/
func (self Event) getConst() TypeConst {
	return self.IConst
}

/*
-----------------------------------------------------------------
   METODO: PrintEvent
   RECIBE: Nada
   DEVUELVE: Nada
   PROPOSITO: Visualizar los atributos de un evento para depurar errores
-----------------------------------------------------------------
*/
func (self Event) PrintEvent(i int) {
	fmt.Println("  EVENTO -> ", i)
	fmt.Println("    Tiempo: ", self.ITime)
	fmt.Println("    Transicion: ", self.ITransition)
	fmt.Println("    Constante: ", self.IConst)
}

//----------------------------------------------------------------------------

// EventList es el tipo que almacena la lista de eventos necesaria
// para los motores de	simulacion.
type EventList []Event

/*
-----------------------------------------------------------------
   METODO: make
   RECIBE: Longitud de la lista de eventos
DEVUELVE: Nada
   PROPOSITO: crear la lista de tamaño aiLongitud
-----------------------------------------------------------------
*/
func MakeEventList(aiLongitud int) EventList {
	return make(EventList, aiLongitud)
}

/*
-----------------------------------------------------------------
   METODO: Insertta
RECIBE: Evento a Inserttar
   DEVUELVE: Nada
PROPOSITO: Inserttar el evento en la lista de eventos, de forma que
   	la Insertcion sea ordenada por tiempo.
-----------------------------------------------------------------
*/
func (self *EventList) Insert(ae_evento Event) {
	var i int // INITIALIZED to 0 !!!

	//fmt.Println("Inserttar evento en lista : ", ae_evento, *self)

	// Obtengo la posicion ordenada del evento en slice con i
	for _, e := range *self {
		if e.GetTime() >= ae_evento.GetTime() {
			break
		}
		i++
	}

	//fmt.Println("POSICION a InsertTAR en lista de evnetos : ", i)
	*self = append((*self)[:i], append([]Event{ae_evento}, (*self)[i:]...)...)

	//fmt.Println("DESPUES de Inserttar : ", *self)
}

/*
 -----------------------------------------------------------------
    METODO: recoge_primer_evento
	RECIBE: Nada
    DEVUELVE: El primer evento de la lista
	PROPOSITO: Recoger el primer evento encolado
 -----------------------------------------------------------------
*/
func (self EventList) GetFirstEvent() Event {
	if len(self) > 0 {
		return self[0]
	} else {
		return Event{} //sino devuelve el tipo Event, zeroed
	}
}

/*
-----------------------------------------------------------------
   METODO: elimina_primer_evento
RECIBE: Nada
   DEVUELVE: Nada
PROPOSITO: Eliminar el primer evento encolado
-----------------------------------------------------------------
*/
func (self *EventList) DeleteFirstEvent() {
	if len(*self) > 0 {
		//suprimir con posibilidad de liberacion de memoria
		copy(*self, (*self)[1:])
		(*self)[len(*self)-1] = Event{} //pongo a zero el previo último Event
		(*self) = (*self)[:len(*self)-1]
	}
}

/*
-----------------------------------------------------------------
   METODO: longitud
RECIBE: Nada
   DEVUELVE: Numero de eventos encolados
PROPOSITO: Conocer el numero de elementos de la lista de eventos
-----------------------------------------------------------------
*/
func (self EventList) Length() int {
	return len(self)
}

/*
-----------------------------------------------------------------
   METODO: PrintEvent
RECIBE: Nada
   DEVUELVE: Nada
PROPOSITO: Imprimir la lista de eventos
-----------------------------------------------------------------
*/
func (self EventList) PrintEvent() {
	fmt.Println("Estructura EventList")
	for i, e := range self {
		e.PrintEvent(i)
	}
}
