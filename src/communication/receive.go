package communication

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"

	// dcs "github.com/shamuelmanrrique/petrynet/src/distconssim"
	u "github.com/shamuelmanrrique/petrynet/src/utils"
)

// Receive TODO
// func Receive( sim *dcs.SimulationEngineDist, connect *u.Connect ) error {
func Receive(chanInterf chan<- interface{}, connect *u.Connect) error {
	// time.Sleep(time.Duration(10) * time.Second)
	time.Sleep(5 * time.Second)
	var listener net.Listener
	var decoder *gob.Decoder
	var pack interface{}
	var red net.Conn
	var err error

	listener, err = net.Listen("tcp", connect.GetIDSubRed())
	u.Error(err, "Listen Error")
	defer listener.Close()

	for {

		red, err = listener.Accept()
		u.Error(err, "Server accept red error")
		// defer red.Close()

		decoder = gob.NewDecoder(red)
		err = decoder.Decode(&pack)
		u.Error(err, "Receive error  \n")
		fmt.Sprintln()
		// chanInterf <- pack

		// log.Println("[Receive] PACK", pack)
		// switch packNew := pack.(type) {
		// case u.Message:
		// 	chanMes <- packNew
		// 	// log.Println("[ReceiveM] ===> MESSAGE ", packNew, " DE ", packNew.GetFrom())
		// 	log.Println(" RECEIVE -->: from ", packNew.GetFrom(), " to ", packNew.GetTo(), "  || OBJ: ", packNew.GetTarg(),
		// 		"\n                     Vector: ", packNew.GetVector())
		// case u.Marker:
		// 	chanMar <- packNew
		// 	// log.Println("[ReceiveM] ----> Marker ", packNew, " DE ", packNew.GetCounter())
		// 	// log.Println(" RECEIVE -->: Init Marker:", packNew.Recoder, "  ||Counter:", packNew.GetCounter(),
		// 	// 	"\n                     Header: ", packNew.GetHeader(),
		// 	// 	"\n                     Channel: ", packNew.GetChannel())
		// 	log.Println(" RECEIVE -->: Init Marker:", packNew)
		// case string:
		// 	chanPoint <- packNew
		// 	// log.Println("[ReceiveM] ----> checkpoint ", packNew )
		// 	log.Println(" RECEIVE --> ACK from: ", packNew)

		// }

		// red.Close()
		// if connect.

	}

	return err
}
