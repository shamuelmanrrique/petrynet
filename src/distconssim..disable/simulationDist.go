package distconssim

import (
	"fmt"
	"github.com/shamuelmanrrique/petrynet/src/centralsim"
	"github.com/shamuelmanrrique/petrynet/src/communication"
	"time"
)

// SimulationEngine is the basic data type for simulation execution
type SimulationEngineDist struct {
	*centralsim.SimulationEngine
	IAddress string // Direccion de escucha de mensajes
	Iport    string // Puerto de escucha de mensaje
	ILookout bool
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
func MakeMotorSimulationDist(alLaLef LefDists) SimulationEngine {
	m := SimulationEngineDist{}
	m.IlMisLefs= alLaLef
	return m
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
	var leEvent Event

	for self.IlMisLefs.ThereEvent(ai_tiempo) {
		leEvent = self.IlMisLefs.GetFirstEvent()

		// Si el valor de la transicion es negativo,indica que pertenece
		// a otra subred y el codigo global de la transicion es pasarlo
		// a positivo y restarle 1
		// ej: -3 -> transicion -(-3) -1 = 2
		if leEvent.ITransition >= 0 {
			// Establecer nuevo valor de la funcion
			self.IlMisLefs.UpdateFuncValue(leEvent.ITransition,
				leEvent.IConst)
			// Establecer nuevo valor del tiempo
			self.IlMisLefs.UpdateTime(leEvent.ITransition,
				leEvent.ITime)
		}
	// } else {
	// 		fmt.Println("Transicion Remota")
	// 		// Changing subred index and search address
	// 		leEvent.IlMisLefs *= -1
	// 		addr := self.IlMisLefs   il_mislefs.Il_pos[leEvent.Ii_transicion]
	// 		var msg MsgI
	// 		msg = MsgEvent{leEvent}
	// 		communication.Send(message, addr)
		}
	}
}

// //Enviar mensaje a traves de la red de forma codificada
// func (self *SimulationEngine) send_message(msg MsgI, addr string) {
// 	conn, err := net.Dial("tcp", addr)
// 	if err != nil {
// 		fmt.Println("Dial error, retrying..:", err.Error())
// 	} else {
// 		// Encode and send data
// 		encoder := gob.NewEncoder(conn)
// 		err = encoder.Encode(&msg)
// 		fmt.Println("Mensaje: ", msg, " enviado a: ", addr)
// 		// Close connection
// 		conn.Close()
// 	}
// }
