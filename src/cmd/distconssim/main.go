package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	dcs "github.com/shamuelmanrrique/petrynet/src/distconssim"

	u "github.com/shamuelmanrrique/petrynet/src/utils"
)

var ip string
var name string
var checklog bool

func init() {
	gob.Register(&u.Message{})
	gob.Register(&dcs.EventDist{})
	gob.Register(dcs.IndGlobalTrans(0))
	gob.Register(dcs.TypeClock(0))
	gob.Register(&dcs.LefsDist{})
	gob.Register(&dcs.TransitionConstant{})
	gob.Register(&dcs.TransitionList{})
	flag.StringVar(&name, "n", "SubRed", "SubRed Name")
	flag.StringVar(&ip, "i", "127.0.1.1:1400", "IP que se usara en el proceso")
	flag.BoolVar(&checklog, "l", true, "Send output to file true otherwise false")
}

func main() {
	flag.Parse()

	fmt.Println(checklog)

	if checklog {
		file, err := os.OpenFile(name+"_"+ip+"_log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		log.SetOutput(file)
	}

	conects := u.NewConnec(u.LocalIP3s)
	fmt.Println(conects)
	if ip == conects.GetConnection(0).GetIDSubRed() {
		u.DistUnic(name)
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
		// fmt.Println(IDSubNet)
		ms := dcs.MakeMotorSimulation(lfs, IDSubNet)
		go dcs.Receive(ms, IDSubNet)
		time.Sleep(1 * time.Second)
		init := dcs.TypeClock(u.InitTransition)
		end := dcs.TypeClock(u.EndTransition)
		ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
		fmt.Println("SDT Termino en 10s")

	}

	if ip == conects.GetConnection(1).GetIDSubRed() {
		u.DistUnic(name)
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
		time.Sleep(1 * time.Second)
		init := dcs.TypeClock(u.InitTransition)
		end := dcs.TypeClock(u.EndTransition)
		ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
		fmt.Println("SDT Termino en 10s")

	}

	if ip == conects.GetConnection(2).GetIDSubRed() {
		u.DistUnic(name)
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
		time.Sleep(1 * time.Second)
		init := dcs.TypeClock(u.InitTransition)
		end := dcs.TypeClock(u.EndTransition)
		ms.Simulate(init, end) // ciclo 0 hasta ciclo 3

	}

	time.Sleep(50 * time.Second)
}
