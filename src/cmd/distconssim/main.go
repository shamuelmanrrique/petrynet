package main

import (
	// "simbot/centralsim"
	cs "github.com/shamuelmanrrique/petrynet/src/centralsim"
	dcs "github.com/shamuelmanrrique/petrynet/src/distconssim"
)

func main() {
	lfs := dcs.LefsDist{ //Ejemplo PN documento adjunto
		SubNet: dcs.TransitionList{
			dcs.TransitionDist{
				IdLocal:        0,
				IiValorLef:     0,
				IiShotDuration: 1,
				IiListactes: []cs.TransitionConstant{
					cs.TransitionConstant{0, 1},
					cs.TransitionConstant{1, -1},
					cs.TransitionConstant{2, -1},
				},
			},
			dcs.TransitionDist{
				IdLocal:        1,
				IiValorLef:     1,
				IiShotDuration: 2,
				IiListactes: []cs.TransitionConstant{
					cs.TransitionConstant{1, 1},
					cs.TransitionConstant{2, -1},
				},
			},
			dcs.TransitionDist{
				IdLocal:        2,
				IiValorLef:     2,
				IiShotDuration: 1,
				IiListactes: []cs.TransitionConstant{
					cs.TransitionConstant{2, 2},
					cs.TransitionConstant{0, -1},
				},
			},
		},
	}
	ms := dcs.MakeMotorSimulation(lfs)
	ms.Simulate(0, 3) // ciclo 0 hasta ciclo 3
}
