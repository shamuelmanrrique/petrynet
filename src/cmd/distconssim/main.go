package main

import (
	"encoding/gob"
	dcs "github.com/shamuelmanrrique/petrynet/src/distconssim"

	u "github.com/shamuelmanrrique/petrynet/src/utils"
)

func init() {
	gob.Register(u.Message{})
	gob.Register(dcs.EventDist{})
	gob.Register(dcs.LefsDist{})
	gob.Register(dcs.TransitionConstant{})
	gob.Register(dcs.TransitionList{})
}

func main() {
	value := []string{"155.210.154.199"}
	connection := u.InitSSH(value[0])

	go u.RunCommand("ls|grep ;cat 155.210.154.208:1400.log.txt", connection)

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
}
