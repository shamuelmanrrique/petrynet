package distconssim

import (
	"encoding/gob"
	"log"
	"net"

	u "github.com/shamuelmanrrique/petrynet/src/utils"
)

/*
-----------------------------------------------------------------
   METODO: Send
   RECIBE: interface(Message will send) and address
   DEVUELVE: Nothing
   PROPOSITO: Send message to other subnets
-----------------------------------------------------------------
*/
func Send(pack interface{}, addr string) error {
	var connection net.Conn
	var err error
	var encoder *gob.Encoder

	connection, err = net.Dial("tcp", addr)
	u.Error(err, "Error Sending message")
	defer connection.Close()

	log.Println(" ++> SEND Message:", addr, "MSM", pack)
	encoder = gob.NewEncoder(connection)
	err = encoder.Encode(&pack)
	u.Error(err, "Error Encoding message")

	return err

}
