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

var subNetNamesD5 []string

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
	pathTest = cfg.Section(environment).Key("mainPath").String()
	subNetNamesD5 = strings.Split(cfg.Section(environment).Key("subNetNameD5").String(), ",")
	subNetIDS = strings.Split(cfg.Section(environment).Key("subNetID5").String(), ",")
	// logMode = cfg.Section("general").Key("log").String()
	connect5 = u.NewConnec(subNetIDS)
}

func TestTime5Dist(t *testing.T) {
	println("------------------------------- ESTOY TestSSHDistTime5 ---------------------------------------")
	for i, ip := range subNetIDS5 {
		addr := strings.Split(ip, ":")
		connection := u.InitSSH(addr[0])

		println(pathTest+subNetNamesD5[i], ip, addr)

		go u.ExcecuteSSH(pathTest+subNetNamesD5[i], connection)
	}

	time.Sleep(80 * time.Second)
}

func TestSubNetD51(t *testing.T) {

	IDSubNet := connect5.GetConnection(0)
	if logMode {
		file, err := os.OpenFile("../logs/["+IDSubNet.GetIp()+"]-TestSubNetD51.log",
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		log.SetOutput(file)
	}

	println("------------>", IDSubNet.GetIDSubRed())
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
					dcs.TransitionConstant{-5, -1},
					dcs.TransitionConstant{-9, -1},
					dcs.TransitionConstant{-13, -1},
				},
			},
			// T17
			dcs.TransitionDist{
				IDGlobal:       17,
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
			4:  connect5.GetConnection(1),
			8:  connect5.GetConnection(2),
			12: connect5.GetConnection(3),
			16: connect5.GetConnection(4),
		},
		Post: dcs.Incidence{
			1:  connect5.GetConnection(1),
			5:  connect5.GetConnection(2),
			9:  connect5.GetConnection(3),
			13: connect5.GetConnection(4),
		},
	}
	// log.Println(IDSubNet)
	ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
	go dcs.Receive(ms, IDSubNet)
	time.Sleep(5 * time.Second)
	init := dcs.TypeClock(u.InitTransition)
	end := dcs.TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}

func TestSubNetD52(t *testing.T) {

	IDSubNet := connect5.GetConnection(1)
	if logMode {
		file, err := os.OpenFile("../logs/["+IDSubNet.GetIp()+"]-TestSubNetD52.log",
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		log.SetOutput(file)
	}
	println("------------>", IDSubNet.GetIDSubRed())
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
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{1, 1},
					dcs.TransitionConstant{2, -1},
				},
			},
			// T3
			dcs.TransitionDist{
				IDGlobal:       3,
				IDLocal:        2,
				IiValorLef:     1,
				IiShotDuration: 3,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{2, 1},
					dcs.TransitionConstant{3, -1},
				},
			},
			// T4
			dcs.TransitionDist{
				IDGlobal:       4,
				IDLocal:        3,
				IiValorLef:     1,
				IiShotDuration: 2,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{3, 1},
					dcs.TransitionConstant{-17, -1},
				},
			},
		},
		Pre: dcs.Incidence{
			0: connect5.GetConnection(0),
		},
		Post: dcs.Incidence{
			17: connect5.GetConnection(0),
		},
	}
	ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
	go dcs.Receive(ms, IDSubNet)
	time.Sleep(5 * time.Second)
	init := dcs.TypeClock(u.InitTransition)
	end := dcs.TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}

func TestSubNetD53(t *testing.T) {

	IDSubNet := connect5.GetConnection(2)
	if logMode {
		file, err := os.OpenFile("../logs/["+IDSubNet.GetIp()+"]-TestSubNetD53.log",
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		log.SetOutput(file)
	}
	println("------------>", IDSubNet.GetIDSubRed())
	lfs := dcs.LefsDist{
		SubNet: dcs.TransitionList{
			// T5
			dcs.TransitionDist{
				IDGlobal:       5,
				IDLocal:        0,
				IiValorLef:     1,
				IiShotDuration: 2,
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
					dcs.TransitionConstant{2, -1},
				},
			},
			// T7
			dcs.TransitionDist{
				IDGlobal:       7,
				IDLocal:        2,
				IiValorLef:     1,
				IiShotDuration: 3,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{2, 1},
					dcs.TransitionConstant{3, -1},
				},
			},
			// T8
			dcs.TransitionDist{
				IDGlobal:       8,
				IDLocal:        3,
				IiValorLef:     1,
				IiShotDuration: 2,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{3, 1},
					dcs.TransitionConstant{-17, -1},
				},
			},
		},
		Pre: dcs.Incidence{
			0: connect5.GetConnection(0),
		},
		Post: dcs.Incidence{
			17: connect5.GetConnection(0),
		},
	}
	ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
	go dcs.Receive(ms, IDSubNet)
	time.Sleep(5 * time.Second)
	init := dcs.TypeClock(u.InitTransition)
	end := dcs.TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}

func TestSubNetD54(t *testing.T) {

	IDSubNet := connect5.GetConnection(3)
	if logMode {
		file, err := os.OpenFile("../logs/["+IDSubNet.GetIp()+"]-TestSubNetD54.log",
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		log.SetOutput(file)
	}
	println("------------>", IDSubNet.GetIDSubRed())
	lfs := dcs.LefsDist{
		SubNet: dcs.TransitionList{
			// T9
			dcs.TransitionDist{
				IDGlobal:       9,
				IDLocal:        0,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{0, 1},
					dcs.TransitionConstant{1, -1},
				},
			},
			// T10
			dcs.TransitionDist{
				IDGlobal:       10,
				IDLocal:        1,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{1, 1},
					dcs.TransitionConstant{2, -1},
				},
			},
			// T11
			dcs.TransitionDist{
				IDGlobal:       11,
				IDLocal:        2,
				IiValorLef:     2,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{2, 1},
					dcs.TransitionConstant{3, -1},
				},
			},
			// T12
			dcs.TransitionDist{
				IDGlobal:       12,
				IDLocal:        3,
				IiValorLef:     1,
				IiShotDuration: 2,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{3, 1},
					dcs.TransitionConstant{-17, -1},
				},
			},
		},
		Pre: dcs.Incidence{
			0: connect5.GetConnection(0),
		},
		Post: dcs.Incidence{
			17: connect5.GetConnection(0),
		},
	}
	ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
	go dcs.Receive(ms, IDSubNet)
	time.Sleep(5 * time.Second)
	init := dcs.TypeClock(u.InitTransition)
	end := dcs.TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}

func TestSubNetD55(t *testing.T) {
	IDSubNet := connect5.GetConnection(4)
	if logMode {
		file, err := os.OpenFile("../logs/["+IDSubNet.GetIp()+"]-TestSubNetD55.log",
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		log.SetOutput(file)
	}
	println("------------>", IDSubNet.GetIDSubRed())
	lfs := dcs.LefsDist{
		SubNet: dcs.TransitionList{
			// T13
			dcs.TransitionDist{
				IDGlobal:       13,
				IDLocal:        0,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{0, 1},
					dcs.TransitionConstant{1, -1},
				},
			},
			// T14
			dcs.TransitionDist{
				IDGlobal:       14,
				IDLocal:        1,
				IiValorLef:     1,
				IiShotDuration: 2,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{1, 1},
					dcs.TransitionConstant{2, -1},
				},
			},
			// T15
			dcs.TransitionDist{
				IDGlobal:       15,
				IDLocal:        2,
				IiValorLef:     1,
				IiShotDuration: 3,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{2, 1},
					dcs.TransitionConstant{3, -1},
				},
			},
			// T16
			dcs.TransitionDist{
				IDGlobal:       16,
				IDLocal:        3,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{3, 1},
					dcs.TransitionConstant{-17, -1},
				},
			},
		},
		Pre: dcs.Incidence{
			0: connect5.GetConnection(0),
		},
		Post: dcs.Incidence{
			17: connect5.GetConnection(0),
		},
	}
	ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
	go dcs.Receive(ms, IDSubNet)
	time.Sleep(5 * time.Second)
	init := dcs.TypeClock(u.InitTransition)
	end := dcs.TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}
