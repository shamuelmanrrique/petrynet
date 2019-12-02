package distconssim

import (
	"encoding/gob"
	"fmt"
	"testing"
	"time"

	u "github.com/shamuelmanrrique/petrynet/src/utils"
)

func init() {
	gob.Register(&u.Message{})
	gob.Register(&EventDist{})
	gob.Register(IndGlobalTrans(0))
	gob.Register(TypeClock(0))
	gob.Register(&LefsDist{})
	gob.Register(&TransitionConstant{})
	gob.Register(&TransitionList{})
}

func TestSubNet0(t *testing.T) {
	conects := u.NewConnec(u.RemoteIP3s)
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
	time.Sleep(2 * time.Second)
	init := TypeClock(u.InitTransition)
	end := TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	fmt.Println("SDT Termino en 10s")
	time.Sleep(160 * time.Second)
}

func TestSubNet1(t *testing.T) {
	conects := u.NewConnec(u.RemoteIP3s)
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
			0: conects.GetConnection(0),
		},
		Post: Incidence{
			5: conects.GetConnection(0),
		},
	}
	ms := MakeMotorSimulation(lfs, IDSubNet)
	go Receive(ms, IDSubNet)
	time.Sleep(2 * time.Second)
	init := TypeClock(u.InitTransition)
	end := TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	fmt.Println("SDT Termino en 10s")
	time.Sleep(160 * time.Second)
}

func TestSubNet2(t *testing.T) {
	conects := u.NewConnec(u.RemoteIP3s)
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
			0: conects.GetConnection(0),
		},
		Post: Incidence{
			5: conects.GetConnection(0),
		},
	}
	ms := MakeMotorSimulation(lfs, IDSubNet)
	go Receive(ms, IDSubNet)
	time.Sleep(2 * time.Second)
	init := TypeClock(u.InitTransition)
	end := TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	fmt.Println("SDT Termino en 10s")
	time.Sleep(160 * time.Second)
}
