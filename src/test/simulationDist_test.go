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

var logMode bool
var path string
var pathTest string
var environment string
var subNetNames []string
var subNetIDS []string
var connect u.Connections

// var err error

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
	pathTest = cfg.Section(environment).Key("testPath").String()
	subNetNames = strings.Split(cfg.Section(environment).Key("subNetName").String(), ",")
	subNetIDS = strings.Split(cfg.Section(environment).Key("subNetID").String(), ",")
	log, err := cfg.Section("general").Key("log").Bool()
	logMode = log
	connect = u.NewConnec(subNetIDS)
}

func TestDist(t *testing.T) {
	println("------------------------------- ESTOY TestSSHDist ---------------------------------------")
	println("TestSSHDist", strings.Join(subNetIDS, ", "))
	for i, ip := range subNetIDS {
		addr := strings.Split(ip, ":")
		connection := u.InitSSH(addr[0])

		println(pathTest+subNetNames[i], ip, addr)

		go u.ExcecuteSSH(pathTest+subNetNames[i], connection)
	}

	time.Sleep(80 * time.Second)
}

func TestSubNet1(t *testing.T) {
	println(environment, path, subNetNames, subNetIDS, logMode)

	IDSubNet := connect.GetConnection(0)
	if logMode {
		file, err := os.OpenFile("../logs/["+IDSubNet.GetIp()+"]-TestSubNet1.log",
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		log.SetOutput(file)
	}

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
				},
			},
			// T5
			dcs.TransitionDist{
				IDGlobal:       5,
				IDLocal:        1,
				IiValorLef:     2,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{1, 2},
					dcs.TransitionConstant{0, -1},
				},
			},
		},
		Pre: dcs.Incidence{
			2: connect.GetConnection(1),
			4: connect.GetConnection(2),
		},
		Post: dcs.Incidence{
			1: connect.GetConnection(1),
			3: connect.GetConnection(2),
		},
	}
	// log.Println(IDSubNet)
	ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
	go dcs.Receive(ms, IDSubNet)
	// time.Sleep(2 * time.Second)
	init := dcs.TypeClock(u.InitTransition)
	end := dcs.TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}

func TestSubNet2(t *testing.T) {
	println(environment, path, subNetNames, subNetIDS, logMode)

	IDSubNet := connect.GetConnection(1)
	if logMode {
		file, err := os.OpenFile("../logs/["+IDSubNet.GetIp()+"]-TestSubNet2.log",
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		log.SetOutput(file)
	}

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
					dcs.TransitionConstant{-5, -1},
				},
			},
		},
		Pre: dcs.Incidence{
			0: connect.GetConnection(0),
		},
		Post: dcs.Incidence{
			5: connect.GetConnection(0),
		},
	}
	ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
	go dcs.Receive(ms, IDSubNet)
	// time.Sleep(2 * time.Second)
	init := dcs.TypeClock(u.InitTransition)
	end := dcs.TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}

func TestSubNet3(t *testing.T) {
	println(environment, path, subNetNames, subNetIDS, logMode)

	IDSubNet := connect.GetConnection(2)
	if logMode {
		file, err := os.OpenFile("../logs/["+IDSubNet.GetIp()+"]-TestSubNet3.log",
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		log.SetOutput(file)
	}

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
					dcs.TransitionConstant{-5, -1},
				},
			},
		},
		Pre: dcs.Incidence{
			0: connect.GetConnection(0),
		},
		Post: dcs.Incidence{
			5: connect.GetConnection(0),
		},
	}
	ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
	go dcs.Receive(ms, IDSubNet)
	// time.Sleep(2 * time.Second)
	init := dcs.TypeClock(u.InitTransition)
	end := dcs.TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}
