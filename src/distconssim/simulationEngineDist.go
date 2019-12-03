package distconssim

import (
	"log"
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
   METODO: FireEnabledTransitions
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
func (self *SimulationEngineDist) FireEnabledTransitions(clock TypeClock) {
	for self.IlMisLefs.ThereSensitive() { //while
		liCodTrans := self.IlMisLefs.GetSensitive()
		self.IlMisLefs.Shoot(liCodTrans)

		// Anotar el Resultado que disparo la liCodTrans en tiempoaiLocalClock
		self.IvResults = append(self.IvResults,
			ResultadoTransition{liCodTrans, clock})
	}
}

/*
-----------------------------------------------------------------
   METODO: TreatMenssage
   RECIBE: *Message
   DEVUELVE: Nothing
   PROPOSITO: Esta funcion maneja una parte muy importante del
	la comunicacion en la simulacion distribuida, recibe un mensaje
	extrae el paquete de ese mensaje y realiza la accion correspondiente
	dependiendo del typo del paquete
   HISTORIA DE CAMBIOS:
   COMENTARIOS:
-----------------------------------------------------------------
*/
func (self *SimulationEngineDist) TreatMenssage(msm *u.Message) {
	switch pack := msm.GetPack().(type) {
	// Si es un evento lo add a mi lista de eventos
	case *EventDist:
		log.Println("SED[Receive] -EventDist ==> :", *pack, "MSM", *msm)
		IDTrans := pack.GetTransition()
		pack.SetTransition(self.GetIDTransition(IDTrans))
		self.IlMisLefs.AddEvents(*pack)
	// Si es lookahead lo guardo y checkeo que me llegen todos
	case TypeClock:
		log.Println("SED[Receive] -TypeClock ==>  ", pack, "MSM", *msm)
		self.IlMisLefs.Lookout[msm.GetFrom()] = pack
		self.IlMisLefs.CheckLookout()
		// Si recibo msm null envio lookahead
	case IndGlobalTrans:
		log.Println("SED[Receive] -IndGlobalTrans ==> (NULL) ", pack, "MSM", *msm)
		timeD := self.IlRelojLocal + self.IlMisLefs.TimeDuration(pack)
		message := &u.Message{
			To:   msm.GetFrom(),
			From: self.connect.GetIDSubRed(),
			Pack: timeD,
		}
		Send(message, message.GetTo())
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
			// u.DistMsm("Send Event to Remote Transition")
			IDtrans *= -1
			lEvent.SetTransition(IDtrans)
			// create msm with event
			message := &u.Message{
				To:   self.IlMisLefs.Post[IDtrans].GetIDSubRed(),
				From: self.connect.GetIDSubRed(),
				Pack: &lEvent,
			}
			Send(message, message.GetTo())
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
	// u.DistUnic("Wait agent")
	subNets := self.IlMisLefs.Pre
	for idTrans, conn := range subNets {
		message := &u.Message{
			To:   conn.GetIDSubRed(),
			From: self.connect.GetIDSubRed(),
			Pack: idTrans,
		}
		// Add element to lookout map
		self.IlMisLefs.SetLookout(conn.GetIDSubRed(), TypeClock(-1))

		// Send Lookahead
		Send(message, message.GetTo())
	}

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
	eTime := self.IlMisLefs.TimeFirstEvent()
	newTime := self.IlMisLefs.MinTime()

	self.IlMisLefs.Lookout = make(map[string]TypeClock)

	if eTime <= newTime && eTime != -1 {
		log.Println("NEXT CLOCK...... : eTime", eTime)
		return eTime
	}

	log.Println("NEXT CLOCK...... : newTime", newTime)
	return newTime
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
func (self SimulationEngineDist) RetornResults() {
	u.DistWall()
	u.DistWall()
	log.Println("RESULTADOS DEL SIMULADOR LOCAL " + u.NetName)
	u.DistWall()

	if len(self.IvResults) == 0 {
		log.Println("No esperes ningun resultado...")
	} else {
		for _, li_result := range self.IvResults {
			log.Println("TIEMPO: ", li_result.ValorRelojDisparo, " -> TRANSICION: ",
				self.GetIDLocalTrans(li_result.CodTransition))
			//  ID LOCAL Transition li_result.CodTransition,
		}
	}
}

/*
-----------------------------------------------------------------
   METODO: GetIDTransition
   RECIBE: Valor del reloj local actual para el que queremos saber las
	  transiciones sensibilizadas
   DEVUELVE: OK si todo fue bien o ERROR en caso contrario
   PROPOSITO: Que esta funcion sirva para recorrerse toda la lista de transiciones
	   e Inserttar aquellas en la pila de transiciones sensibilizadas.
-----------------------------------------------------------------
*/
func (self *SimulationEngineDist) GetIDTransition(id IndGlobalTrans) IndGlobalTrans {

	for _, transition := range self.IlMisLefs.SubNet {
		if id == transition.IDGlobal {
			return IndGlobalTrans(transition.IDLocal)
		}
	}

	return IndGlobalTrans(id)
}

func (self *SimulationEngineDist) GetIDLocalTrans(id IndGlobalTrans) IndGlobalTrans {
	for _, transition := range self.IlMisLefs.SubNet {
		if id == IndGlobalTrans(transition.IDLocal) {
			return transition.IDGlobal
		}
	}

	return IndGlobalTrans(id)
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
	ldInit := time.Now()

	// Inicializamos el reloj local
	// ------------------------------------------------------------------
	self.IlRelojLocal = initCycle

	// Inicializamos las transiciones sensibilizadas, es decir, ver si con el
	// marcado inicial tenemos transiciones sensibilizadas
	// ------------------------------------------------------------------
	self.IlMisLefs.UpdateSensitive(self.IlRelojLocal)

	for self.IlRelojLocal <= endCycle {
		// self.IlMisLefs.PrintEvent() //DEPURACION
		log.Println("RELOJ LOCAL !!!  = ", self.IlRelojLocal)
		u.DistWall()

		if initCycle == 0 {
			time.Sleep(2 * time.Second)
		}

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
			// Envio msm null a los demas para que me envien sus lookaheads
			// ------------------------------------------------------------------
			self.WaitAgents()

			// Chequeo que me llegaron todos los lookahead en caso contrario espero
			// ------------------------------------------------------------------
			self.IlMisLefs.CheckLookout()
			if !Active {
				time.Sleep(4 * time.Second)
			}

			if !self.IlMisLefs.ThereEvent(self.IlRelojLocal) {
				self.IlRelojLocal = self.AdvanceTime()

				if self.IlRelojLocal == -1 {
					self.IlRelojLocal = endCycle + 1
				}
			}
		}
	}

	// Close Receive
	self.connect.SetAccept(true)

	elapsedTime := time.Since(ldInit)

	// Print los resultados de la simulacion
	self.RetornResults()
	u.DistL()
	log.Println("NUMERO DE TRANSICIONES DISPARADAS", len(self.IvResults))
	log.Println("TIEMPO SIMULADO en ciclos: ", endCycle-initCycle)
	log.Println("COSTE REAL SIMULACION: ", elapsedTime.String())
	u.DistWall()
	u.DistWall()
}
