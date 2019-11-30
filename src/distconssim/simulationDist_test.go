package distconssim

import (
	"fmt"
	"testing"
	"time"

	u "github.com/shamuelmanrrique/petrynet/src/utils"
)

func TestSubNet0(t *testing.T) {
	time.Sleep(4 * time.Second)
	conects := u.NewConnec(u.LocalIP3s)
	IDSubNet := conects.GetConnection(0)
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
					TransitionConstant{-3, -1},
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
			2: conects.GetConnection(1),
			4: conects.GetConnection(2),
		},
		Post: Incidence{
			1: conects.GetConnection(1),
			3: conects.GetConnection(2),
		},
	}
	// fmt.Println(IDSubNet)
	ms := MakeMotorSimulation(lfs, IDSubNet)
	go Receive(ms, IDSubNet)
	ms.Simulate(0, 3) // ciclo 0 hasta ciclo 3
	time.Sleep(10 * time.Second)
}

func TestSubNet1(t *testing.T) {
	// time.Sleep(3 * time.Second)
	conects := u.NewConnec(u.LocalIP3s)
	IDSubNet := conects.GetConnection(1)
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
		Pre: Incidence{
			5: conects.GetConnection(0),
		},
		Post: Incidence{
			0: conects.GetConnection(0),
		},
	}
	// ms := MakeMotorSimulation(lfs, conects.GetConnection(1))
	// ms.Simulate(0, 3) // ciclo 0 hasta ciclo 3
	ms := MakeMotorSimulation(lfs, IDSubNet)
	fmt.Println("^^^^^^^^^^^^ESCUCHANDO", IDSubNet)
	// Receive(ms, IDSubNet)
	go Receive(ms, IDSubNet)
	ms.Simulate(0, 3) // ciclo 0 hasta ciclo 3
	time.Sleep(10 * time.Second)
}

func TestSubNet2(t *testing.T) {
	// time.Sleep(2 * time.Second)
	conects := u.NewConnec(u.LocalIP3s)
	IDSubNet := conects.GetConnection(2)
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
		Pre: Incidence{
			5: conects.GetConnection(0),
		},
		Post: Incidence{
			0: conects.GetConnection(0),
		},
	}
	ms := MakeMotorSimulation(lfs, IDSubNet)
	go Receive(ms, IDSubNet)
	ms.Simulate(0, 3) // ciclo 0 hasta ciclo 3
	time.Sleep(10 * time.Second)
}

/*
func TestLefs(t *testing.T) {

}
*/
