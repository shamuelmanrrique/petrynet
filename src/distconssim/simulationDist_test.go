package distconssim

import (
	"testing"

	cm "github.com/shamuelmanrrique/petrynet/src/communication"
	u "github.com/shamuelmanrrique/petrynet/src/utils"
)

func TestSubNet0(t *testing.T) {
	conects := u.NewConnec(u.LocalIP3s)
	IDSubNet := conects[0]
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
		Pre: Incidence{
			1: conects.GetConnection(1),
			3: conects.GetConnection(2),
		},
		Post: Incidence{
			1: conects.GetConnection(1),
			3: conects.GetConnection(2),
		},
	}
	ms := MakeMotorSimulation(lfs, IDSubNet)
	go cm.Receive(make(chan<- interface{}), IDSubNet)
	ms.Simulate(0, 3) // ciclo 0 hasta ciclo 3
}

// func TestSubNet1(t *testing.T) {
// 	conects := u.NewConnec(u.LocalIP3s)
// 	fmt.Println(*conects[1])
// 	lfs := LefsDist{
// 		SubNet: TransitionList{
// 			// T1
// 			TransitionDist{
// 				IDGlobal:       1,
// 				IDLocal:        0,
// 				IiValorLef:     1,
// 				IiShotDuration: 1,
// 				IiListactes: []TransitionConstant{
// 					TransitionConstant{0, 1},
// 					TransitionConstant{1, -1},
// 				},
// 			},
// 			// T2
// 			TransitionDist{
// 				IDGlobal:       2,
// 				IDLocal:        1,
// 				IiValorLef:     1,
// 				IiShotDuration: 2,
// 				IiListactes: []TransitionConstant{
// 					TransitionConstant{1, 1},
// 					TransitionConstant{-5, -1},
// 				},
// 			},
// 		},
// 		Pre: Incidence{
// 			0: conects.GetConnection(0),
// 		},
// 		Post: Incidence{
// 			5: conects.GetConnection(0),
// 		},
// 	}
// 	ms := MakeMotorSimulation(lfs, conects.GetConnection(1))
// 	ms.Simulate(0, 3) // ciclo 0 hasta ciclo 3
// }

// func TestSubNet2(t *testing.T) {
// 	conects := u.NewConnec(u.LocalIP3s)
// 	lfs := LefsDist{
// 		SubNet: TransitionList{
// 			// T3
// 			TransitionDist{
// 				IDGlobal:       3,
// 				IDLocal:        0,
// 				IiValorLef:     1,
// 				IiShotDuration: 1,
// 				IiListactes: []TransitionConstant{
// 					TransitionConstant{0, 1},
// 					TransitionConstant{1, -1},
// 				},
// 			},
// 			// T4
// 			TransitionDist{
// 				IDGlobal:       4,
// 				IDLocal:        1,
// 				IiValorLef:     1,
// 				IiShotDuration: 1,
// 				IiListactes: []TransitionConstant{
// 					TransitionConstant{1, 1},
// 					TransitionConstant{-5, -1},
// 				},
// 			},
// 		},
// 		Pre: Incidence{
// 			0: conects.GetConnection(0),
// 		},
// 		Post: Incidence{
// 			5: conects.GetConnection(0),
// 		},
// 	}
// 	var connect *u.Connect
// 	ms := MakeMotorSimulation(lfs, connect)
// 	ms.Simulate(0, 3) // ciclo 0 hasta ciclo 3
// }

/*
func TestLefs(t *testing.T) {

}
*/
