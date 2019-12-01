package distconssim

import (
	"encoding/gob"
	"log"
	"net"

	u "github.com/shamuelmanrrique/petrynet/src/utils"
)

// Send function
func Send(pack interface{}, addr string) error {
	var connection net.Conn
	var err error
	var encoder *gob.Encoder

	// log.Println(" ++> SEND+++++++++++++++:", addr)
	connection, err = net.Dial("tcp", addr)
	u.Error(err, "Error Sending message")
	defer connection.Close()

	log.Println(" ++> SEND Message:", addr, "MSM", pack)
	encoder = gob.NewEncoder(connection)
	// err = encoder.Encode(&pack)
	err = encoder.Encode(&pack)
	u.Error(err, "Error Encoding message")

	// switch packNew := pack.(type) {
	// case *f.Message:
	// 	log.Println(" ++> SEND MSM: from ", packNew.GetFrom(), " to ", packNew.GetTo(), "  || OBJ: ", packNew.GetTarg(),
	// 		"\n                     Vector: ", packNew.GetVector())
	// case *f.Marker:
	// 	log.Println(" ++> SEND Marker: Init Marker", packNew)
	// case *string:
	// 	log.Println(" ++> SEND Count: ", packNew)

	// }

	return err

}
