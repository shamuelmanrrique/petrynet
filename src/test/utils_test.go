package main

import (
	"time"
	"encoding/gob"
	"log"
	"os"
	"testing"

	dcs "sd_petry_nets/src/distconssim"
	u "sd_petry_nets/src/utils"
)

func init() {
	testing.Init()
	gob.Register(&u.Message{})
	gob.Register(&dcs.EventDist{})
	gob.Register(dcs.IndGlobalTrans(0))
	gob.Register(dcs.TypeClock(0))
	gob.Register(&dcs.LefsDist{})
	gob.Register(&dcs.TransitionConstant{})
	gob.Register(&dcs.TransitionList{})
}

// TestConnections create connections
func TestConnect(t *testing.T) {
	cons := u.NewConnec(u.RemoteIP3s)
	log.Println(cons)
}

func TestLog(t *testing.T) {
	file, err := os.OpenFile("../logs/[test-127.0.0.1]-subNetName.log",
		os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	log.SetOutput(file)

	cons := u.NewConnec(u.LocalIP3s)
	log.Println(cons)
}

func TestMinTime(t *testing.T) {
	var value = map[string]dcs.TypeClock{"127.0.1.1:5000": 3, "127.0.1.1:5001": 1, "127.0.1.1:5002": 2}
	// d = {320:1, 321:0, 322:3}
	// m := min(value, key=lambda value: value[k])
	// u.NewConnec(LocalIPs)
	log.Println(value)

}

func TestSendReceive(t *testing.T) {
	even := &dcs.EventDist{ITime: 4}
	tim := dcs.TypeClock(1)
	timp := &tim
	id := dcs.IndGlobalTrans(9)
	idp := &id
	addr := "127.0.1.1:5002"
	con := u.Connect{IDSubRed: "127.0.1.1:5002"}
	sim := new(dcs.SimulationEngineDist)
	go dcs.Receive(sim, con)
	time.Sleep(4 * time.Second)
	log.Println("EventDistr:", *even, "typeClock:", *timp, "idGLobal:", *idp)
	message := &u.Message{
		To:   addr,
		From: addr,
		Pack: even,
	}
	dcs.Send(message, addr)
	time.Sleep(2 * time.Second)

}
