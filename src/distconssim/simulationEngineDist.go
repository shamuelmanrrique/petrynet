package distconssim

import (
	"fmt"
	"time"

	u "github.com/shamuelmanrrique/petrynet/src/utils"
)

// TypeClock defines integer size for holding time.
type TypeClock int64

// ResultadoTransition holds fired transition id and time of firing
type ResultadoTransition struct {
	CodTransition     IndGlobalTrans
	ValorRelojDisparo TypeClock
}

// SimulationEngineDist is the basic data type for simulation execution
type SimulationEngineDist struct {
	connect      u.Connect             // Embided connect struct
	IlMisLefs    LefsDist              // Estructura de datos del simulador
	IlRelojLocal TypeClock             // Valor de mi reloj local
	IvResults    []ResultadoTransition // slice dinamico con los resultados
}

/*
-----------------------------------------------------------------
   METODO: NewMotorSimulation
   RECIBE: EStructura datos Lefs, Estructura datos Connect
   DEVUELVE: Nada
   PROPOSITO: Construir que recibe la estructura de datos con la que
	   Simulate, inicializa variables...
   HISTORIA DE CAMBIOS:
COMENTARIOS:
-----------------------------------------------------------------
*/
func MakeMotorSimulation(alLaLef LefsDist, connect u.Connect) *SimulationEngineDist {
	m := &SimulationEngineDist{}
	m.IlMisLefs = alLaLef
	m.connect = connect
	return m
}

/*
-----------------------------------------------------------------
   METODO: Shoot_transiciones_sensibilizadas
   RECIBE: Valor del reloj local
   DEVUELVE: Nada
   PROPOSITO: Accede a la lista de transiciones sensibilizadas y procede con su
	   disparo, lo que generara nuevos eventos y modificara el marcado de la
		transicion disparada. Igualmente anotara en los resultados el disparo de
		cada transicion para el reloj actual dado
   HISTORIA DE CAMBIOS:
COMENTARIOS:
-----------------------------------------------------------------
*/
func (self *SimulationEngineDist) FireEnabledTransitions(aiLocalClock TypeClock) {
	for self.IlMisLefs.ThereSensitive() { //while
		liCodTrans := self.IlMisLefs.GetSensitive()
		self.IlMisLefs.Shoot(liCodTrans)

		// Anotar el Resultado que disparo la liCodTrans en tiempoaiLocalClock
		self.IvResults = append(self.IvResults,
			ResultadoTransition{liCodTrans, aiLocalClock})
	}
}

/*
-----------------------------------------------------------------
   METODO: TreatEvent
   RECIBE: Tiempo para el que trataremos los eventos
   DEVUELVE: Nada
   PROPOSITO: Accede a la lista de eventos y trata todos aquellos con tiempo
	   igual al recibido. Al tratar los eventos se modificaran los valores de
		las funciones de sensibilizacion de algunas transiciones, por lo que puede
		que tengamos nuevas transiciones sensibilizadas.
   HISTORIA DE CAMBIOS:
   COMENTARIOS:
-----------------------------------------------------------------
*/
func (self *SimulationEngineDist) TreatEvent(ai_tiempo TypeClock) {
	var lEvent EventDist

	for self.IlMisLefs.ThereEvent(ai_tiempo) {
		lEvent = self.IlMisLefs.GetFirstEvent()

		IDtrans := lEvent.GetTransition()
		if IDtrans < 0 {
			fmt.Println("Transicion Remota")
			lEvent.SetTransition(IDtrans * (-1))

			// // m := self.IlMisLefs.Post[trans]
			// fmt.Println("connection to send", self.IlMisLefs.Post[IDtrans])
			// message := u.Message{
			// 	To: m.GetIDSubRed(),
			// 	// To: self.IlMisLefs.Post[lEvent.ITransition],
			// 	//TODO
			// 	// From: self.Connect.GetId(),
			// 	Pack: lEvent,
			// }
			// cm.Send(message, message.GetTo())
		} else {
			// Establecer nuevo valor de la funcion
			self.IlMisLefs.UpdateFuncValue(lEvent.ITransition,
				lEvent.IConst)
			// Establecer nuevo valor del tiempo
			self.IlMisLefs.UpdateTime(lEvent.ITransition,
				lEvent.ITime)
		}
	}
}

/*
-----------------------------------------------------------------
   METODO: WaitAgents
   RECIBE: Nada
   DEVUELVE: Nada
   PROPOSITO: Espera a que lleguen todos los agentes que hemos enviado
	   anteriormente, para recibir nuevos eventos o el mensaje "No voy
		a generar nada hasta T"
   HISTORIA DE CAMBIOS:
COMENTARIOS:
-----------------------------------------------------------------
*/
func (self *SimulationEngineDist) WaitAgents() {
	fmt.Println("TODO")
	// subNets := self.IlMisLefs.Pre
	// // TODOREVISAR POR STRUCT
	// self.Connect.Accept = false
	// for idTrans, addr := range subNets {
	// 	message := u.Message{
	// 		To: addr,
	// 		//TODO
	// 		From: self.Connect.GetId(),
	// 		Pack: idTrans,
	// 	}
	// 	self.IlMisLefs.Lookout[addr] = -1
	// 	cm.Send(message, message.GetTo())
	// }

	// for !self.Connect.GetAccept() {
	// }
	// return

}

/*
-----------------------------------------------------------------
   METODO: AdvanceTime
   RECIBE: Nada
   DEVUELVE: Nada
   PROPOSITO: Modifica el reloj local con el minimo tiempo de entre los
	   recibidos por los agentes o del primer evento encolado en la lista
		de eventos
   HISTORIA DE CAMBIOS:
COMENTARIOS:
-----------------------------------------------------------------
*/
func (self *SimulationEngineDist) AdvanceTime() TypeClock {
	nextTime := self.IlMisLefs.TimeFirstEvent()
	fmt.Println("NEXT CLOCK...... : ", nextTime)
	return nextTime
}

/*
-----------------------------------------------------------------
   METODO: RetornResults
   RECIBE: Nada
   DEVUELVE: Nada
   PROPOSITO: Mostrar los resultados de la simulacion
   HISTORIA DE CAMBIOS:

COMENTARIOS:
-----------------------------------------------------------------
*/
func (self SimulationEngineDist) RetornResults() string {
	resultados := "----------------------------------------\n"
	resultados += "Resultados del simulador local\n"
	resultados += "----------------------------------------\n"
	if len(self.IvResults) == 0 {
		resultados += "No esperes ningun resultado...\n"
	}

	for _, li_result := range self.IvResults {
		resultados +=
			"TIEMPO: " + fmt.Sprintf("%v", li_result.ValorRelojDisparo) +
				" -> TRANSICION: " + fmt.Sprintf("%v", li_result.CodTransition) + "\n"
	}

	fmt.Println(resultados)
	return resultados
}

/*
-----------------------------------------------------------------
   METODO: Simulate
   RECIBE: Ciclo con el que partimos (por si el marcado recibido no
				se corresponde al inicial sino a uno obtenido tras Simulate
				initCycle ciclos)
			Ciclo con el que terminamos
   DEVUELVE: Nada
   PROPOSITO: Simulate una RdP
   HISTORIA DE CAMBIOS:
COMENTARIOS:
-----------------------------------------------------------------
*/
func (self *SimulationEngineDist) Simulate(initCycle, endCycle TypeClock) {
	// u.DistMsm(self.connect.GetIDSubRed())
	ldInit := time.Now()

	// Inicializamos el reloj local
	// ------------------------------------------------------------------
	self.IlRelojLocal = initCycle

	// Inicializamos las transiciones sensibilizadas, es decir, ver si con el
	// marcado inicial tenemos transiciones sensibilizadas
	// ------------------------------------------------------------------
	self.IlMisLefs.UpdateSensitive(self.IlRelojLocal)

	for self.IlRelojLocal <= endCycle {
		self.IlMisLefs.PrintEvent() //DEPURACION
		fmt.Println("RELOJ LOCAL !!!  = ", self.IlRelojLocal)

		// Si existen transiciones sensibilizadas para reloj local las disparamos
		// ------------------------------------------------------------------
		if self.IlMisLefs.ThereSensitive() {
			self.FireEnabledTransitions(self.IlRelojLocal)
		}

		//self.IlMisLefs.IlEvents.PrintEvent()

		// Si existen eventos para el reloj local los tratamos
		// ------------------------------------------------------------------
		if self.IlMisLefs.ThereEvent(self.IlRelojLocal) {
			self.TreatEvent(self.IlRelojLocal)
		}

		// Los nuevos eventos han podido sensibilizar nuevas transiciones
		// ------------------------------------------------------------------
		self.IlMisLefs.UpdateSensitive(self.IlRelojLocal)

		// Tras tratar todos los eventos, si no nos quedan transiciones
		// sensibilizadas no podemos Simulate nada mas, luego esperamos a
		// los agentes y si no nos generan nuevos eventos procedemos a avanzar
		// el reloj local
		// ------------------------------------------------------------------
		if !self.IlMisLefs.ThereSensitive() {
			// self.WaitAgents()
			if !self.IlMisLefs.ThereEvent(self.IlRelojLocal) {
				self.IlRelojLocal = self.AdvanceTime()

				if self.IlRelojLocal == -1 {
					self.IlRelojLocal = endCycle + 1
				}
			}
		}
	}

	elapsedTime := time.Since(ldInit)

	// Devolver los resultados de la simulacion
	self.RetornResults()
	result := "\n---------------------"
	result += "NUMERO DE TRANSICIONES DISPARADAS " +
		fmt.Sprintf("%d", len(self.IvResults)) + "\n"
	result += "TIEMPO SIMULADO en ciclos: " +
		fmt.Sprintf("%d", endCycle-initCycle) + "\n"
	result += "COSTE REAL SIMULACION: " +
		fmt.Sprintf("%v", elapsedTime.String()) + "\n"
	fmt.Println(result)
}
