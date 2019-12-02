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
	gob.Register(u.Message{})
	gob.Register(dcs.EventDist{})
	gob.Register(dcs.LefsDist{})
	gob.Register(dcs.TransitionConstant{})
	gob.Register(dcs.TransitionList{})
	flag.StringVar(&name, "n", "SubRed", "SubRed Name")
	flag.StringVar(&ip, "ip", "127.0.1.1:1400", "IP que se usara en el proceso")
	flag.BoolVar(&checklog, "l", true, "Send output to file true otherwise false")
}

func main() {
	flag.Parse()

	if checklog {
		file, err := os.OpenFile(name+"_"+ip+"_log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		log.SetOutput(file)
	}

	var LocalIPs = []string{"127.0.1.1:5000", "127.0.1.1:5001", "127.0.1.1:5002", "127.0.1.1:5003",
		"127.0.1.1:5004", "127.0.1.1:5005", "127.0.1.1:5006"}
	cons := u.NewConnec(LocalIPs)
	log.Println(cons)
	fmt.Println("puta")

	if false {
		//t.Skip("skipping test simulation.")
		lfs := dcs.LefsDist{ //Ejemplo PN documento adjunto
			SubNet: dcs.TransitionList{
				dcs.TransitionDist{
					IDGlobal:       0,
					IiValorLef:     0,
					IiShotDuration: 1,
					IiListactes: []dcs.TransitionConstant{
						dcs.TransitionConstant{0, 1},
						dcs.TransitionConstant{1, -1},
						dcs.TransitionConstant{2, -1},
					},
				},
				dcs.TransitionDist{
					IDGlobal:       1,
					IiValorLef:     1,
					IiShotDuration: 2,
					IiListactes: []dcs.TransitionConstant{
						dcs.TransitionConstant{1, 1},
						dcs.TransitionConstant{2, -1},
					},
				},
				dcs.TransitionDist{
					IDGlobal:       2,
					IiValorLef:     2,
					IiShotDuration: 1,
					IiListactes: []dcs.TransitionConstant{
						dcs.TransitionConstant{2, 2},
						dcs.TransitionConstant{0, -1},
					},
				},
			},
		}
		var connect u.Connect
		ms := dcs.MakeMotorSimulation(lfs, connect)
		ms.Simulate(0, 3) // ciclo 0 hasta ciclo 3
	}

	time.Sleep(5 * time.Second)
}
