package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	dcs "sd_petry_nets/src/distconssim"
	u "sd_petry_nets/src/utils"

	"gopkg.in/ini.v1"
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

	// Loading configuration file
	cfg, err := ini.Load("../config/go.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// Getting configuration values from .ini
	environment = cfg.Section("general").Key("environment").String()
	path = cfg.Section(environment).Key("mainPath").String()
	subNetNames = strings.Split(cfg.Section(environment).Key("subNetName5").String(), ",")
	subNetIDS = strings.Split(cfg.Section(environment).Key("subNetID5").String(), ",")
	logMode = cfg.Section("general").Key("log").String()
	connect = u.NewConnec(subNetIDS)
	// var conects = u.NewConnec(u.LocalIP5s)
}

func TestSSHDist5(t *testing.T) {
	for name, ip := range u.RemoteIP5 {
		addr := strings.Split(ip, ":")
		connection := u.InitSSH(addr[0])
		log.Println(connection, name, ip)
		go u.ExcecuteSSH(u.GoTest+name, connection)
	}
	time.Sleep(70 * time.Second)
}

func TestSubNet51(t *testing.T) {
	IDSubNet := connect.GetConnection(0)
	lfs := dcs.LefsDist{
		SubNet: dcs.TransitionList{
			// T0
			dcs.TransitionDist{
				IDGlobal:       0,
				IDLocal:        0,
				IiValorLef:     0,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{0, 1},
					dcs.TransitionConstant{-1, -1},
					dcs.TransitionConstant{-3, -1},
					dcs.TransitionConstant{-5, -1},
					dcs.TransitionConstant{-7, -1},
				},
			},
			// T9
			dcs.TransitionDist{
				IDGlobal:       9,
				IDLocal:        1,
				IiValorLef:     4,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{1, 4},
					dcs.TransitionConstant{0, -1},
				},
			},
		},
		Pre: dcs.Incidence{
			2: connect.GetConnection(1),
			4: connect.GetConnection(2),
			6: connect.GetConnection(3),
			8: connect.GetConnection(4),
		},
		Post: dcs.Incidence{
			1: connect.GetConnection(1),
			3: connect.GetConnection(2),
			5: connect.GetConnection(3),
			7: connect.GetConnection(4),
		},
	}
	// log.Println(IDSubNet)
	ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
	go dcs.Receive(ms, IDSubNet)
	time.Sleep(2 * time.Second)
	init := dcs.TypeClock(u.InitTransition)
	end := dcs.TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}

func TestSubNet52(t *testing.T) {
	IDSubNet := connect.GetConnection(1)
	lfs := dcs.LefsDist{
		SubNet: dcs.TransitionList{
			// T1
			dcs.TransitionDist{
				IDGlobal:       1,
				IDLocal:        0,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{0, 1},
					dcs.TransitionConstant{1, -1},
				},
			},
			// T2
			dcs.TransitionDist{
				IDGlobal:       2,
				IDLocal:        1,
				IiValorLef:     1,
				IiShotDuration: 2,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{1, 1},
					dcs.TransitionConstant{-9, -1},
				},
			},
		},
		Pre: dcs.Incidence{
			0: connect.GetConnection(0),
		},
		Post: dcs.Incidence{
			9: connect.GetConnection(0),
		},
	}
	ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
	go dcs.Receive(ms, IDSubNet)
	time.Sleep(2 * time.Second)
	init := dcs.TypeClock(u.InitTransition)
	end := dcs.TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}

func TestSubNet53(t *testing.T) {
	IDSubNet := connect.GetConnection(2)
	lfs := dcs.LefsDist{
		SubNet: dcs.TransitionList{
			// T3
			dcs.TransitionDist{
				IDGlobal:       3,
				IDLocal:        0,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{0, 1},
					dcs.TransitionConstant{1, -1},
				},
			},
			// T4
			dcs.TransitionDist{
				IDGlobal:       4,
				IDLocal:        1,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{1, 1},
					dcs.TransitionConstant{-9, -1},
				},
			},
		},
		Pre: dcs.Incidence{
			0: connect.GetConnection(0),
		},
		Post: dcs.Incidence{
			9: connect.GetConnection(0),
		},
	}
	ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
	go dcs.Receive(ms, IDSubNet)
	time.Sleep(2 * time.Second)
	init := dcs.TypeClock(u.InitTransition)
	end := dcs.TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}

func TestSubNet54(t *testing.T) {
	IDSubNet := connect.GetConnection(3)
	lfs := dcs.LefsDist{
		SubNet: dcs.TransitionList{
			// T5
			dcs.TransitionDist{
				IDGlobal:       5,
				IDLocal:        0,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{0, 1},
					dcs.TransitionConstant{1, -1},
				},
			},
			// T6
			dcs.TransitionDist{
				IDGlobal:       6,
				IDLocal:        1,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{1, 1},
					dcs.TransitionConstant{-9, -1},
				},
			},
		},
		Pre: dcs.Incidence{
			0: connect.GetConnection(0),
		},
		Post: dcs.Incidence{
			9: connect.GetConnection(0),
		},
	}
	ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
	go dcs.Receive(ms, IDSubNet)
	time.Sleep(1 * time.Second)
	init := dcs.TypeClock(u.InitTransition)
	end := dcs.TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}

func TestSubNet55(t *testing.T) {
	IDSubNet := connect.GetConnection(4)
	lfs := dcs.LefsDist{
		SubNet: dcs.TransitionList{
			// T7
			dcs.TransitionDist{
				IDGlobal:       7,
				IDLocal:        0,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{0, 1},
					dcs.TransitionConstant{1, -1},
				},
			},
			// T8
			dcs.TransitionDist{
				IDGlobal:       8,
				IDLocal:        1,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{1, 1},
					dcs.TransitionConstant{-9, -1},
				},
			},
		},
		Pre: dcs.Incidence{
			0: connect.GetConnection(0),
		},
		Post: dcs.Incidence{
			9: connect.GetConnection(0),
		},
	}
	ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
	go dcs.Receive(ms, IDSubNet)
	time.Sleep(1 * time.Second)
	init := dcs.TypeClock(u.InitTransition)
	end := dcs.TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}
