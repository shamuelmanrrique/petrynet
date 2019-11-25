package distconssim

import (
	u "github.com/shamuelmanrrique/petrynet/src/utils"
	"testing"
)

func TestSubRed0(t *testing.T) {
	lfs := LefsDist{
		SubNet: TransitionList{
			// T0
			TransitionDist{
				IDGlobal:       0,
				IDLocal:        0,
				IiValorLef:     0,
				IiShotDuration: 1,
				IiListactes: []TransitionConstant{
					TransitionConstant{0, 1},
					TransitionConstant{-1, -1},
					TransitionConstant{-2, -1},
				},
			},
			// T5
			TransitionDist{
				IDGlobal:       5,
				IDLocal:        1,
				IiValorLef:     2,
				IiShotDuration: 1,
				IiListactes: []TransitionConstant{
					TransitionConstant{1, 2},
					TransitionConstant{0, -1},
				},
			},
		},
	}
	var connect *u.Connect
	ms := MakeMotorSimulation(lfs, connect)
	ms.Simulate(0, 3) // ciclo 0 hasta ciclo 3
}

func TestSubRed1(t *testing.T) {
	lfs := LefsDist{
		SubNet: TransitionList{
			// T1
			TransitionDist{
				IDGlobal:       1,
				IDLocal:        0,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []TransitionConstant{
					TransitionConstant{0, 1},
					TransitionConstant{1, -1},
				},
			},
			// T2
			TransitionDist{
				IDGlobal:       2,
				IDLocal:        1,
				IiValorLef:     1,
				IiShotDuration: 2,
				IiListactes: []TransitionConstant{
					TransitionConstant{1, 1},
					TransitionConstant{-5, -1},
				},
			},
		},
	}
	var connect *u.Connect
	ms := MakeMotorSimulation(lfs, connect)
	ms.Simulate(0, 3) // ciclo 0 hasta ciclo 3
}

func TestSubRed2(t *testing.T) {
	lfs := LefsDist{
		SubNet: TransitionList{
			// T3
			TransitionDist{
				IDGlobal:       3,
				IDLocal:        0,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []TransitionConstant{
					TransitionConstant{0, 1},
					TransitionConstant{1, -1},
				},
			},
			// T4
			TransitionDist{
				IDGlobal:       4,
				IDLocal:        1,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []TransitionConstant{
					TransitionConstant{1, 1},
					TransitionConstant{-5, -1},
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
