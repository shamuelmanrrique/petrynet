/*
PROPOSITO:
		- Estructura de datos para gestionar las Lefs de las transiciones. Sera
		rellenada a partir de los datos obtenidos de la entrada de
		datos del usuario y es diferente para cada una de las subredes,
		es decir cada simulador local tendra una instancia distinta de
		esta clase.
-----------------------------------------------------------------
*/
package centralsim

import "fmt"

type TypeIndexSubnet int32

//----------------------------------------------------------------------------

// Lefs es el tipo de datos principal que gestiona el disparo de transiciones.
type Lefs struct {
	Subnet TransitionList // Slice de transiciones de esta subred
	// Identificadores de las transiciones sensibilizadas para
	// T = Reloj local actual. Slice que funciona como Stack
	IsTransSensib StackTransitions
	IlEvents      EventList //Lista de eventos a procesar
}

/*
-----------------------------------------------------------------
   METODO: NewLefs
   RECIBE: Lista de transiciones
   DEVUELVE: Nada
   PROPOSITO: crear nueva estructura LEF
-----------------------------------------------------------------
*/
func NewLefs(listaTransiciones TransitionList) Lefs {
	l := Lefs{}
	l.Subnet = listaTransiciones
	l.IsTransSensib = nil
	l.IlEvents = nil

	return l
}

/*
-----------------------------------------------------------------
   METODO: AddEvents
   RECIBE: Evento a a�adir
   DEVUELVE: OK si todo va bien o ERROR en caso contrario
   PROPOSITO: A�ade a la lista de eventos
-----------------------------------------------------------------
*/
func (self *Lefs) AddEvents(ae_evento Event) bool {
	self.IlEvents.Insert(ae_evento)
	return true
}

/*
-----------------------------------------------------------------
   METODO: AddSensitive
   RECIBE: Transicion sensibilizada a a�adir
   DEVUELVE: OK si todo va bien o ERROR en caso contrario
   PROPOSITO: A�ade a la lista de transiciones sensibilizadas
-----------------------------------------------------------------
*/
func (self *Lefs) AddSensitive(ai_transicion IndLocalTrans) bool {
	self.IsTransSensib.push(ai_transicion)
	return true // OK
}

/*
-----------------------------------------------------------------
   METODO: TimeFirstEvent
   RECIBE: Nada
   DEVUELVE: El valor del tiempo del primer evento de la lista de eventos.
	  -1 si ocurrio un error o no hay eventos.
   PROPOSITO: Visualizar el valor temporal del primer evento para conocer
	   posteriormente si debemos avanzar el reloj local
-----------------------------------------------------------------
*/
func (self Lefs) TimeFirstEvent() TypeClock {
	if self.IlEvents.length() > 0 {
		le_evento := self.IlEvents.GetFirstEvent()
		return le_evento.ITime
	} else {
		return -1
	}
}

/*
-----------------------------------------------------------------
   METODO: ThereEvent
   RECIBE: Tiempo del reloj local
   DEVUELVE: true si quedan eventos para ese tiempo o false en caso contrario
   PROPOSITO: Conocer si restan eventos disponibles para el tiempo dado
-----------------------------------------------------------------
*/
func (self Lefs) ThereEvent(ai_tiempo TypeClock) bool {
	if self.TimeFirstEvent() == ai_tiempo {
		return true
	} else {
		return false
	}
}

/*
-----------------------------------------------------------------
   METODO: ThereSensitive
   RECIBE: Nada
   DEVUELVE: true si las hay o false en caso contrario
   PROPOSITO: Conocer si tenemos funciones sensibilizadas
	COMENTARIOS: Se supone que previamente a la invocacion a esta funcion
	   se ha tenido que llamar a UpdateSensitive (reloj_local)
-----------------------------------------------------------------
*/
func (self Lefs) ThereSensitive() bool {
	return !self.IsTransSensib.isEmpty()
}

/*
-----------------------------------------------------------------
   METODO: GetSensitive
   RECIBE: Nada
   DEVUELVE: El identificador de la primera transicion sensibilizada
	 o -1 en caso contrario
   PROPOSITO: Coger el primer identificador de la lista de transiciones
	 sensibilizadas
-----------------------------------------------------------------
*/
func (self *Lefs) GetSensitive() IndLocalTrans {
	if (*self).IsTransSensib.isEmpty() {
		return -1
	} else {
		return (*self).IsTransSensib.pop()
	}
}

/*
-----------------------------------------------------------------
   METODO: GetFirstEvent
   RECIBE: Nada
   DEVUELVE: El primer evento de la lista de eventos
   PROPOSITO: Coger el primer evento de la lista de eventos
-----------------------------------------------------------------
*/
func (self *Lefs) GetFirstEvent() Event {
	/* fmt.Println("Lista antes de eliminar primer evento :")
	(*self).IlEvents.PrintEvent()
	*/
	le_evento := (*self).IlEvents.GetFirstEvent()
	(*self).IlEvents.DeleteFirstEvent()
	/*fmt.Println("Lista DESPUES de eliminar primer evento :")
	(*self).IlEvents.PrintEvent()
	*/
	return le_evento
}

/*
-----------------------------------------------------------------
   METODO: UpdateSensitive
   RECIBE: Valor del reloj local actual para el que queremos saber las
	  transiciones sensibilizadas
   DEVUELVE: OK si todo fue bien o ERROR en caso contrario
   PROPOSITO: Que esta funcion sirva para recorrerse toda la lista de transiciones
	   e Inserttar aquellas en la pila de transiciones sensibilizadas.
COMENTARIOS: Me recorro todo el array de transiciones, por lo que deberiamos
	   invocar a esta funcion cuando ya hayan sido a�adidas todas las transiciones.
-----------------------------------------------------------------
*/
func (self *Lefs) UpdateSensitive(ai_relojlocal TypeClock) bool {
	for li_i, t := range (*self).Subnet {
		if t.IiValorLef <= 0 && t.ITime == ai_relojlocal {
			(*self).IsTransSensib.push(IndLocalTrans(li_i))
		}
	}
	return true
}

/*
-----------------------------------------------------------------
   METODO: UpdateTime
   RECIBE: Codigo de la transicion y nuevo valor del tiempo
   DEVUELVE: true si todo fue bien o false en caso contrario
   PROPOSITO: Modificar el tiempo de la transicion dada
-----------------------------------------------------------------
*/
func (self *Lefs) UpdateTime(il_tr IndLocalTrans, ai_ti TypeClock) bool {
	// Algunas comprobaciones...
	if il_tr >= 0 && il_tr < self.Subnet.length() {
		// Modificacion del tiempo
		self.Subnet[il_tr].ITime = ai_ti
		return true
	} else { // index out of range
		return false
	}
}

/*
-----------------------------------------------------------------
   METODO: UpdateFuncValue
   RECIBE: Codigo de la transicion y valor con el que modificar
		OJO, no es el valor definitivo, sino la CTE a a�adir al valor que tenia
		antes la funcion
   DEVUELVE: OK si todo fue bien o ERROR en caso contrario
   PROPOSITO: Modificar valor de funcion de sensibilizacion de transicion dada
-----------------------------------------------------------------
*/
func (self *Lefs) UpdateFuncValue(ilTr IndLocalTrans, aiValLef TypeConst) bool {
	// Algunas comprobaciones...
	if ilTr >= 0 && ilTr < self.Subnet.length() {
		// Modificacion del valor de la funcion lef
		self.Subnet[ilTr].IiValorLef += aiValLef
		return true
	} else { // Out of range
		return false
	}
}

/*
-----------------------------------------------------------------
   METODO: Shoot
   RECIBE: Indice en el vector de la transicion a Shoot
   DEVUELVE: OK si todo fue bien o ERROR en caso contrario
   PROPOSITO: Shoot una transicion. Esto es, generar todos los eventos
	   ocurridos por el disparo de una transicion
-----------------------------------------------------------------
*/
func (self *Lefs) Shoot(ilTr IndLocalTrans) bool {
	// Algunas comprobaciones...
	if ilTr >= 0 && ilTr < self.Subnet.length() {
		// Prepare 3 local variables
		tiTrans := self.Subnet[ilTr].ITime             // time to spread to new events
		tiDur := self.Subnet[ilTr].Ii_duracion_disparo //time length
		listCtes := self.Subnet[ilTr].Ii_listactes     // list of TransCtes

		// La CTE de la primera trans., hace referencia a la cte a mandar a
		// TRANS. QUE SE HA DISPARADO, y va con tiempo igual al de la transicion
		// tiempo, cod_transicion, cte
		self.AddEvents(Event{tiTrans, listCtes[0].INextTrans, listCtes[0].Cnstnt})

		// Generamos eventos ocurridos por disparo de transicion ilTr
		for _, trCo := range listCtes[1:] {
			// tiempo = tiempo de la transicion + coste disparo
			self.AddEvents(Event{tiTrans + tiDur, trCo.INextTrans, trCo.Cnstnt})
		}

		return true
	} else {
		return false
	}
}

/*
-----------------------------------------------------------------
   METODO: PrintEventTransitions
   RECIBE: Nada
   DEVUELVE: Nada
   PROPOSITO: Imprimir las transiciones para depurar errores
-----------------------------------------------------------------
*/
func (self Lefs) PrintEventTransitions() {
	fmt.Println(" ")
	fmt.Println("------IMPRIMIMOS LA LISTA DE TRANSICIONES---------")
	for _, tr := range self.Subnet {
		tr.PrintEvent_valores()
	}
	fmt.Println("------FINAL DE LA LISTA DE TRANSICIONES---------")
	fmt.Println(" ")
}

/*
-----------------------------------------------------------------
   METODO: PrintEvent
   RECIBE: Nada
   DEVUELVE: Nada
   PROPOSITO: Imprimir los atributos de la clase para depurar errores
-----------------------------------------------------------------
*/
func (self Lefs) PrintEvent() {

	fmt.Println("STRUCT LEFS")
	//fmt.Println ("\tNº transiciones: ", self.ii_indice)
	fmt.Println("\tNº transiciones: ", self.Subnet.length())

	if self.IsTransSensib.isEmpty() {
		fmt.Println("\tLISTA TRANSICIONES SENSIBILIZADAS VACIA")
	} else {
		fmt.Println("\tLista transciones sensibilizadas :")
		for _, iTr := range self.IsTransSensib {
			fmt.Print(iTr, " ")
			fmt.Println(" ")
		}
	}
	fmt.Println("------Lista transiciones---------")
	for _, tr := range self.Subnet {
		tr.PrintEvent()
	}
	fmt.Println("------Final lista transiciones---------")

	fmt.Println("-----------Lista eventos---------")
	self.IlEvents.PrintEvent()
	fmt.Println("-----------Final lista eventos---------")
	fmt.Println("FINAL ESTRUCTURA LEFS")
}
