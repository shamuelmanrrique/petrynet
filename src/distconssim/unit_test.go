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
	gob.Register(&LefsDist{})
	gob.Register(&TransitionConstant{})
	gob.Register(&TransitionList{})
}

// TestConnections create connections
func TestConnect(t *testing.T) {
	var LocalIPs = []string{"127.0.1.1:5000", "127.0.1.1:5001", "127.0.1.1:5002", "127.0.1.1:5003",
		"127.0.1.1:5004", "127.0.1.1:5005", "127.0.1.1:5006"}
	u.NewConnec(LocalIPs)

}

func TestMinTime(t *testing.T) {
	var value = map[string]TypeClock{"127.0.1.1:5000": 3, "127.0.1.1:5001": 1, "127.0.1.1:5002": 2}
	// d = {320:1, 321:0, 322:3}
	// m := min(value, key=lambda value: value[k])
	// u.NewConnec(LocalIPs)
	fmt.Println(value)

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
		Pack: even,
	}
	Send(message, addr)
	time.Sleep(2 * time.Second)

}

func TestSSH(t *testing.T) {
	value := []string{"155.210.154.199"}
	connection := u.InitSSH(value[0])

	// "/usr/local/go/bin/go run /home/a802400/go/src/practice1/app/main.go
	// fmt.Println("ssh to:", defaultAddresses[i], len(defaultAddresses), i)
	// go RunCommand("cd "+dir+" && go test -run "+subnets[i], conn)
	go u.RunCommand("ls", connection)

}
