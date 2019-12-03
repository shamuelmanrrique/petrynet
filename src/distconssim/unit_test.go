package distconssim

import (
	"encoding/gob"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	u "github.com/shamuelmanrrique/petrynet/src/utils"
)

func init() {
	testing.Init()
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
	cons := u.NewConnec(u.RemoteIP3s)
	log.Println(cons)
}

func TestLog(t *testing.T) {
	file, err := os.OpenFile("~"+u.Dir+"log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	log.SetOutput(file)

	cons := u.NewConnec(u.LocalIP3s)
	log.Println(cons)
}

func TestSSH(t *testing.T) {
	for name, ip := range u.RemoteIP3 {
		addr := strings.Split(ip, ":")
		connection := u.InitSSH(addr[0])
		log.Println(connection, name, ip)
		go u.ExcecuteSSH(u.GoMainLog+" -i="+ip+" -n="+name, connection)
		// go u.ExcecuteSSH(u.GoMainLog+" -ip="+ip+" -n="+name, connection)
		// go u.ExcecuteSSH(u.GoTest+name, connection)
		// Ready
		// go u.ExcecuteSSH(u.GoTest+"TestConnect", connection)
		// go u.ExcecuteSSH(u.GoTest+"TestLog", connection)
	}
	time.Sleep(50 * time.Second)
}

func TestSSHRemote(t *testing.T) {
	for _, ip := range u.RemoteIP3 {
		addr := strings.Split(ip, ":")
		log.Println(addr[0])
		connection := u.InitSSH(addr[0])
		go u.ExcecuteSSH(u.GoTest+"TestLog", connection)
	}
	time.Sleep(30 * time.Second)

}

func TestSSHLOCALPetry(t *testing.T) {
	for testS, ip := range u.LocalIP3 {
		addr := strings.Split(ip, ":")
		log.Println(addr[0])
		connection := u.InitSSH(addr[0])
		log.Println(" -ip="+ip+" -n="+testS, connection)
		// go u.ExcecuteSSH(u.GoMainLog+" -ip="+ip+" -n="+testS, connection)
		// go u.ExcecuteSSH(u.GoLocalTest+"TestLog", connection)
		break
	}

	time.Sleep(300 * time.Second)

}
func TestMinTime(t *testing.T) {
	var value = map[string]TypeClock{"127.0.1.1:5000": 3, "127.0.1.1:5001": 1, "127.0.1.1:5002": 2}
	// d = {320:1, 321:0, 322:3}
	// m := min(value, key=lambda value: value[k])
	// u.NewConnec(LocalIPs)
	log.Println(value)

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
	log.Println("EventDistr:", *even, "typeClock:", *timp, "idGLobal:", *idp)
	message := &u.Message{
		To:   addr,
		From: addr,
		Pack: even,
	}
	Send(message, addr)
	time.Sleep(2 * time.Second)

}
