/*
PROPOSITO:
		- Tipo abstracto para realizar la simulacion de una (sub)RdP.
HISTORIA DE CAMBIOS:
COMENTARIOS:
		- El resultado de una simulacion local sera un slice dinamico de
		componentes, de forma que cada una de ella sera una structura estatica de
		dos enteros, el primero de ellos sera el codigo de la transicion
		disparada y el segundo sera el valor del reloj local para el que se
		disparo.
-----------------------------------------------------------------
*/
package centralsim

import (
	"fmt"
	"time"
)

// TypeClock defines integer size for holding time.
type TypeClock int64

// ResultadoTransition holds fired transition id and time of firing
type ResultadoTransition struct {
	CodTransition     IndLocalTrans
	ValorRelojDisparo TypeClock
}

// SimulationEngine is the basic data type for simulation execution
type SimulationEngine struct {
	ilMisLefs    Lefs                  // Estructura de datos del simulador
	ilRelojLocal TypeClock             // Valor de mi reloj local
	ivResults    []ResultadoTransition // slice dinamico con los resultados
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
func MakeMotorSimulation(alLaLef Lefs) SimulationEngine {
	m := SimulationEngine{}
	m.ilMisLefs = alLaLef
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
func (self *SimulationEngine) fireEnabledTransitions(aiLocalClock TypeClock) {
	for self.ilMisLefs.ThereSensitive() { //while
		liCodTrans := self.ilMisLefs.GetSensitive()
		self.ilMisLefs.Shoot(liCodTrans)

		// Anotar el Resultado que disparo la liCodTrans en tiempoaiLocalClock
		self.ivResults = append(self.ivResults,
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
func (self *SimulationEngine) TreatEvent(ai_tiempo TypeClock) {
	var le_evento Event

	for self.ilMisLefs.ThereEvent(ai_tiempo) {
		le_evento = self.ilMisLefs.GetFirstEvent()

		// Si el valor de la transicion es negativo,indica que pertenece
		// a otra subred y el codigo global de la transicion es pasarlo
		// a positivo y restarle 1
		// ej: -3 -> transicion -(-3) -1 = 2
		if le_evento.ITransition >= 0 {
			// Establecer nuevo valor de la funcion
			self.ilMisLefs.UpdateFuncValue(le_evento.ITransition,
				le_evento.IConst)
			// Establecer nuevo valor del tiempo
			self.ilMisLefs.UpdateTime(le_evento.ITransition,
				le_evento.ITime)
		}
	}
}

/*
-----------------------------------------------------------------
   METODO: waitAgents
   RECIBE: Nada
   DEVUELVE: Nada
   PROPOSITO: Espera a que lleguen todos los agentes que hemos enviado
	   anteriormente, para recibir nuevos eventos o el mensaje "No voy
		a generar nada hasta T"
   HISTORIA DE CAMBIOS:
COMENTARIOS:
-----------------------------------------------------------------
*/
func (self *SimulationEngine) waitAgents() {
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
func (self *SimulationEngine) AdvanceTime() TypeClock {
	nextTime := self.ilMisLefs.TimeFirstEvent()
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
func (self SimulationEngine) RetornResults() string {
	resultados := "----------------------------------------\n"
	resultados += "Resultados del simulador local\n"
	resultados += "----------------------------------------\n"
	if len(self.ivResults) == 0 {
		resultados += "No esperes ningun resultado...\n"
	}

	for _, li_result := range self.ivResults {
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
func (self *SimulationEngine) Simulate(ai_cicloinicial, ai_nciclos TypeClock) {
	ld_ini := time.Now()

	// Inicializamos el reloj local
	// ------------------------------------------------------------------
	self.ilRelojLocal = ai_cicloinicial

	// Inicializamos las transiciones sensibilizadas, es decir, ver si con el
	// marcado inicial tenemos transiciones sensibilizadas
	// ------------------------------------------------------------------
	self.ilMisLefs.UpdateSensitive(self.ilRelojLocal)

	for self.ilRelojLocal <= ai_nciclos {
		self.ilMisLefs.PrintEvent() //DEPURACION
		fmt.Println("RELOJ LOCAL !!!  = ", self.ilRelojLocal)

		// Si existen transiciones sensibilizadas para reloj local las disparamos
		// ------------------------------------------------------------------
		if self.ilMisLefs.ThereSensitive() {
			self.fireEnabledTransitions(self.ilRelojLocal)
		}

		//self.ilMisLefs.IlEvents.PrintEvent()

		// Si existen eventos para el reloj local los tratamos
		// ------------------------------------------------------------------
		if self.ilMisLefs.ThereEvent(self.ilRelojLocal) {
			self.TreatEvent(self.ilRelojLocal)
		}

		// Los nuevos eventos han podido sensibilizar nuevas transiciones
		// ------------------------------------------------------------------
		self.ilMisLefs.UpdateSensitive(self.ilRelojLocal)

		// Tras tratar todos los eventos, si no nos quedan transiciones
		// sensibilizadas no podemos Simulate nada mas, luego esperamos a
		// los agentes y si no nos generan nuevos eventos procedemos a avanzar
		// el reloj local
		// ------------------------------------------------------------------
		if !self.ilMisLefs.ThereSensitive() {
			// self.waitAgents()
			if !self.ilMisLefs.ThereEvent(self.ilRelojLocal) {
				self.ilRelojLocal = self.AdvanceTime()

				if self.ilRelojLocal == -1 {
					self.ilRelojLocal = ai_nciclos + 1
				}
			}
		}
	}

	elapsedTime := time.Since(ld_ini)

	// Devolver los resultados de la simulacion
	self.RetornResults()
	result := "\n---------------------"
	result += "NUMERO DE TRANSICIONES DISPARADAS " +
		fmt.Sprintf("%d", len(self.ivResults)) + "\n"
	result += "TIEMPO SIMULADO en ciclos: " +
		fmt.Sprintf("%d", ai_nciclos-ai_cicloinicial) + "\n"
	result += "COSTE REAL SIMULACION: " +
		fmt.Sprintf("%v", elapsedTime.String()) + "\n"
	fmt.Println(result)
}
