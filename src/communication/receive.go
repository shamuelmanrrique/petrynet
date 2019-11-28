package communication

import (
	"encoding/gob"
	"log"
	"net"

	// dcs "github.com/shamuelmanrrique/petrynet/src/distconssim"
	u "github.com/shamuelmanrrique/petrynet/src/utils"
)

// Receive a msm and check by type of packet received
func Receive(chanInterf chan<- interface{}, connect u.Connect) error {
	var listener net.Listener
	var decoder *gob.Decoder
	var pack interface{}
	var red net.Conn
	var err error

	listener, err = net.Listen("tcp", connect.GetIDSubRed())
	u.Error(err, "Listen Error")
	defer listener.Close()

receiveChannel:
	for {
		red, err = listener.Accept()
		u.Error(err, "Server accept red error")

		decoder = gob.NewDecoder(red)
		err = decoder.Decode(&pack)
		u.Error(err, "Receive error  \n")

		log.Println("[Receive] PACK", pack)
		switch packNew := pack.(type) {
		case u.Message:
			// chanMes <- packNew
			log.Println("[ReceiveM] ===> MESSAGE ", packNew, " DE ", packNew.GetFrom())

		default:
			u.Error(nil, "ERROR Receive type")
		}

		// TODO CHANGE BY GLOBAL SIMULATION NUMBER
		if connect.GetAccept() {
			red.Close()
			break receiveChannel
		}

	}

	return err
}
