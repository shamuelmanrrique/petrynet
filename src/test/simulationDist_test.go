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

var logMode string
var path string
var environment string
var subNetNames []string
var subNetIDS []string
var connect u.Connections

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
	subNetNames = strings.Split(cfg.Section(environment).Key("subNetName").String(), ",")
	subNetIDS = strings.Split(cfg.Section(environment).Key("subNetID").String(), ",")
	logMode = cfg.Section("general").Key("log").String()
	connect = u.NewConnec(subNetIDS)
}

func TestSubNet1(t *testing.T) {
	println(environment, path, subNetNames, subNetIDS, logMode)
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
	time.Sleep(160 * time.Second)
}

func TestSubNet2(t *testing.T) {
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
	time.Sleep(160 * time.Second)
}

func TestSubNet3(t *testing.T) {
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
	time.Sleep(160 * time.Second)
}
