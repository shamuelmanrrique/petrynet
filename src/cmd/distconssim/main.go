package main

import (
	dcs "github.com/shamuelmanrrique/petrynet/src/distconssim"

	u "github.com/shamuelmanrrique/petrynet/src/utils"
)

func main() {
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
	var connect *u.Connect
	ms := dcs.MakeMotorSimulation(lfs, connect)
	ms.Simulate(0, 3) // ciclo 0 hasta ciclo 3
}
