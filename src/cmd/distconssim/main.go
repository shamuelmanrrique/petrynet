package main

import (
	dcs "github.com/shamuelmanrrique/petrynet/src/distconssim"
	// cs "github.com/shamuelmanrrique/petrynet/src/centralsim"
)

func main() {
	lfs := dcs.LefsDist{ //Ejemplo PN documento adjunto
		SubNet: dcs.TransitionList{
			dcs.TransitionDist{
				IdLocal:        0,
				IiValorLef:     0,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{0, 1},
					dcs.TransitionConstant{1, -1},
					dcs.TransitionConstant{2, -1},
				},
			},
			dcs.TransitionDist{
				IdLocal:        1,
				IiValorLef:     1,
				IiShotDuration: 2,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{1, 1},
					dcs.TransitionConstant{2, -1},
				},
			},
			dcs.TransitionDist{
				IdLocal:        2,
				IiValorLef:     2,
				IiShotDuration: 1,
				IiListactes: []dcs.TransitionConstant{
					dcs.TransitionConstant{2, 2},
					dcs.TransitionConstant{0, -1},
				},
			},
		},
	}
	ms := dcs.MakeMotorSimulation(lfs)
	ms.Simulate(0, 3) // ciclo 0 hasta ciclo 3
}
