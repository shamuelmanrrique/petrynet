package distconssim

import (
	"encoding/gob"
	"fmt"
	"testing"
	"time"

	u "github.com/shamuelmanrrique/petrynet/src/utils"
)

func init() {
	gob.Register(&u.Message{})
	gob.Register(&EventDist{})
	gob.Register(IndGlobalTrans(0))
	gob.Register(TypeClock(0))
	// gob.Register(&LefsDist{})
	// // gob.Register(TypeClock{})
	// // gob.Register(IndGlobalTrans{})
	// gob.Register(&TransitionConstant{})
	// gob.Register(&TransitionList{})
}

// TestConnections create connections
func TestConnect(t *testing.T) {
	var LocalIPs = []string{"127.0.1.1:5000", "127.0.1.1:5001", "127.0.1.1:5002", "127.0.1.1:5003",
		"127.0.1.1:5004", "127.0.1.1:5005", "127.0.1.1:5006"}
	u.NewConnec(LocalIPs)

}

func TestSendReceive(t *testing.T) {
	even := &EventDist{ITime: 4}
	tim := TypeClock(1)
	timp := &tim
	id := IndGlobalTrans(9)
	idp := &id
	addr := "127.0.1.1:5002"
	con := u.Connect{IDSubRed: "127.0.1.1:5002"}
	sim := new(SimulationEngineDist)
	go Receive(sim, con)
	time.Sleep(4 * time.Second)
	fmt.Println("EventDistr:", *even, "typeClock:", *timp, "idGLobal:", *idp)
	message := &u.Message{
		To:   addr,
		From: addr,
		Pack: tim,
	}
	Send(message, addr)
	time.Sleep(2 * time.Second)

}
