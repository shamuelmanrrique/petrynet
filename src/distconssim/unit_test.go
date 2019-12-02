package distconssim

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
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
	cons := u.NewConnec(LocalIPs)
	fmt.Println(cons)
}

func TestSSH(t *testing.T) {
	value := map[string]string{"TestSubNet0": "155.210.154.199", "TestSubNet1": "155.210.154.200", "TestSubNet2": "155.210.154.204"}
	for name, ip := range value {
		connection := u.InitSSH(ip)
		fmt.Println(connection)
		go u.RunCommand(u.GoMainLog+" -ip="+ip+" -n="+name, connection)
		// go u.RunCommand(u.GoTest+" TestLog  >> 1.txt", connection)
	}
	time.Sleep(300 * time.Second)
}

func TestLog(t *testing.T) {
	file, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	log.SetOutput(file)

	var LocalIPs = []string{"127.0.1.1:5000", "127.0.1.1:5001", "127.0.1.1:5002", "127.0.1.1:5003",
		"127.0.1.1:5004", "127.0.1.1:5005", "127.0.1.1:5006"}
	cons := u.NewConnec(LocalIPs)
	log.Println(cons)

}

func TestSSHPetry(t *testing.T) {
	value := map[string]string{"TestSubNet0": "155.210.154.199", "TestSubNet1": "155.210.154.200", "TestSubNet2": "155.210.154.204"}

	for testS, ip := range value {
		// for _, ip := range value {
		connection := u.InitSSH(ip)
		fmt.Println(connection)
		go u.RunCommand(u.GoMainLog+" -ip="+ip+" -n="+testS, connection)
		// go u.RunCommand(u.GoTest+" TestConnect", connection)

	}

	time.Sleep(300 * time.Second)

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
