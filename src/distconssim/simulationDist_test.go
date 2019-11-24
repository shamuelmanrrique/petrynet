package distconssim

import (
	u "github.com/shamuelmanrrique/petrynet/src/utils"
	"testing"
)

func TestSimulationEngine(t *testing.T) {
	//t.Skip("skipping test simulation.")
	lfs := LefsDist{ //Ejemplo PN documento adjunto
		SubNet: TransitionList{
			TransitionDist{
				IDGlobal:       0,
				IiValorLef:     0,
				IiShotDuration: 1,
				IiListactes: []TransitionConstant{
					TransitionConstant{0, 1},
					TransitionConstant{1, -1},
					TransitionConstant{2, -1},
				},
			},
			TransitionDist{
				IDGlobal:       1,
				IiValorLef:     1,
				IiShotDuration: 2,
				IiListactes: []TransitionConstant{
					TransitionConstant{1, 1},
					TransitionConstant{2, -1},
				},
			},
			TransitionDist{
				IDGlobal:       2,
				IiValorLef:     2,
				IiShotDuration: 1,
				IiListactes: []TransitionConstant{
					TransitionConstant{2, 2},
					TransitionConstant{0, -1},
				},
			},
		},
	}
	var connect *u.Connect
	ms := MakeMotorSimulation(lfs, connect)
	ms.Simulate(0, 3) // ciclo 0 hasta ciclo 3
}

/*
func TestTransition(t *testing.T) {

}

func TestLefs(t *testing.T) {

}
*/
