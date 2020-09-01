package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	dcs "sd_petry_nets/src/distconssim"

	"gopkg.in/ini.v1"

	u "sd_petry_nets/src/utils"
)

var ip string
var checklog bool
var environment string
var subNetName string
var subNetID []string
var subNetNames []string

func init() {
	// Register all interface to use
	gob.Register(&u.Message{})
	gob.Register(&dcs.EventDist{})
	gob.Register(dcs.IndGlobalTrans(0))
	gob.Register(dcs.TypeClock(0))
	gob.Register(&dcs.LefsDist{})
	gob.Register(&dcs.TransitionConstant{})
	gob.Register(&dcs.TransitionList{})

	flag.StringVar(&subNetName, "name", "subNet1", "Insert name like subNet# (# is a number 1-3) ")
	flag.BoolVar(&checklog, "log", false, "Send output to file true otherwise false")
}

func main() {
	// Parcing flags
	flag.Parse()

	// Loading configuration file
	cfg, err := ini.Load("../../config/go.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// Getting configuration values from .ini
	environment = cfg.Section("general").Key("environment").String()
	subNetID = strings.Split(cfg.Section(environment).Key("subNetID").String(), ",")
	subNetNames = strings.Split(cfg.Section(environment).Key("subNetName").String(), ",")

	for i, name := range subNetNames {
		if name == subNetName {
			ip = subNetID[i]
			break
		}
	}

	// Writting output in log if checklog is true
	if checklog {
		file, err := os.OpenFile("../../logs/["+ip+"]-"+subNetName+".log",
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		log.SetOutput(file)
	}

	println(subNetID, ip, subNetName)
	// Get Ip
	conects := u.NewConnec(subNetID)

	println(ip, conects.GetConnection(0).GetIDSubRed())

	if ip == conects.GetConnection(0).GetIDSubRed() {
		u.NetName = subNetName + " " + ip
		u.DistUnic(u.NetName)
		IDSubNet := conects.GetConnection(0)
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
				2: conects.GetConnection(1),
				4: conects.GetConnection(2),
			},
			Post: dcs.Incidence{
				1: conects.GetConnection(1),
				3: conects.GetConnection(2),
			},
		}
		log.Println(IDSubNet)
		ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
		go dcs.Receive(ms, IDSubNet)
		time.Sleep(4 * time.Second)
		init := dcs.TypeClock(u.InitTransition)
		end := dcs.TypeClock(u.EndTransition)
		ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
		log.Println("SDT Termino en 10s")

	}

	if ip == conects.GetConnection(1).GetIDSubRed() {
		u.NetName = subNetName + " " + ip
		u.DistUnic(u.NetName)
		IDSubNet := conects.GetConnection(1)
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
				0: conects.GetConnection(0),
			},
			Post: dcs.Incidence{
				5: conects.GetConnection(0),
			},
		}
		ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
		go dcs.Receive(ms, IDSubNet)
		time.Sleep(4 * time.Second)
		init := dcs.TypeClock(u.InitTransition)
		end := dcs.TypeClock(u.EndTransition)
		ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
		log.Println("SDT Termino en 10s")

	}

	if ip == conects.GetConnection(2).GetIDSubRed() {
		u.NetName = subNetName + " " + ip
		u.DistUnic(u.NetName)
		IDSubNet := conects.GetConnection(2)
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
				0: conects.GetConnection(0),
			},
			Post: dcs.Incidence{
				5: conects.GetConnection(0),
			},
		}
		ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
		go dcs.Receive(ms, IDSubNet)
		time.Sleep(4 * time.Second)
		init := dcs.TypeClock(u.InitTransition)
		end := dcs.TypeClock(u.EndTransition)
		ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
		log.Println("SDT Termino en 10s")

	}

	time.Sleep(50 * time.Second)
}
