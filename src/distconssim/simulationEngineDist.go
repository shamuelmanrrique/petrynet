package distconssim

import (
	"fmt"
	"time"

	cm "github.com/shamuelmanrrique/petrynet/src/communication"
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
	IlMisLefs    LefsDist              // Estructura de datos del simulador
	IlRelojLocal TypeClock             // Valor de mi reloj local
	IvResults    []ResultadoTransition // slice dinamico con los resultados
}

/*
-----------------------------------------------------------------
   METODO: NewMotorSimulation
   RECIBE: EStructura datos Lefs
   DEVUELVE: Nada
   PROPOSITO: Construir que recibe la estructura de datos con la que
	   Simulate, inicializa variables...
   HISTORIA DE CAMBIOS:
COMENTARIOS:
-----------------------------------------------------------------
*/
func MakeMotorSimulation(alLaLef LefsDist) SimulationEngineDist {
	m := SimulationEngineDist{}
	m.IlMisLefs = alLaLef
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

		// Si el valor de la transicion es negativo,indica que pertenece
		// a otra subred y el codigo global de la transicion es pasarlo
		// a positivo y restarle 1
		// ej: -3 -> transicion -(-3) -1 = 2
		if lEvent.ITransition < 0 {
			fmt.Println("Transicion Remota")
			lEvent.ITransition *= -1
			// lEvent.ITransition = Abs(lEvent.ITransition)

			addr := self.IlMisLefs.Post[lEvent.ITransition]
			fmt.Println(addr)
			// TODO aun no me queda claro lo que voy a enviar
			var message interface{}
			// var msg MsgI
			// msg = MsgEvent{le_evento}
			cm.Send(message, addr)

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
	fmt.Println("Aun sin agentes")

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
				ai_cicloinicial ciclos)
			Ciclo con el que terminamos
   DEVUELVE: Nada
   PROPOSITO: Simulate una RdP
   HISTORIA DE CAMBIOS:
COMENTARIOS:
-----------------------------------------------------------------
*/
func (self *SimulationEngineDist) Simulate(ai_cicloinicial, ai_nciclos TypeClock) {
	ld_ini := time.Now()

	// Inicializamos el reloj local
	// ------------------------------------------------------------------
	self.IlRelojLocal = ai_cicloinicial

	// Inicializamos las transiciones sensibilizadas, es decir, ver si con el
	// marcado inicial tenemos transiciones sensibilizadas
	// ------------------------------------------------------------------
	self.IlMisLefs.UpdateSensitive(self.IlRelojLocal)

	for self.IlRelojLocal <= ai_nciclos {
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
					self.IlRelojLocal = ai_nciclos + 1
				}
			}
		}
	}

	elapsedTime := time.Since(ld_ini)

	// Devolver los resultados de la simulacion
	self.RetornResults()
	result := "\n---------------------"
	result += "NUMERO DE TRANSICIONES DISPARADAS " +
		fmt.Sprintf("%d", len(self.IvResults)) + "\n"
	result += "TIEMPO SIMULADO en ciclos: " +
		fmt.Sprintf("%d", ai_nciclos-ai_cicloinicial) + "\n"
	result += "COSTE REAL SIMULACION: " +
		fmt.Sprintf("%v", elapsedTime.String()) + "\n"
	fmt.Println(result)
}
