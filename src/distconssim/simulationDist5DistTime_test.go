package distconssim

import (
	"encoding/gob"
	"log"
	"strings"
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

var connect = u.NewConnec(u.LocalIP5s)

func TestSubNetRD50(t *testing.T) {
	IDSubNet := connect.GetConnection(0)
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
					TransitionConstant{-5, -1},
					TransitionConstant{-9, -1},
					TransitionConstant{-13, -1},
				},
			},
			// T17
			TransitionDist{
				IDGlobal:       17,
				IDLocal:        1,
				IiValorLef:     4,
				IiShotDuration: 1,
				IiListactes: []TransitionConstant{
					TransitionConstant{1, 4},
					TransitionConstant{0, -1},
				},
			},
		},
		Pre: Incidence{
			4:  connect.GetConnection(1),
			8:  connect.GetConnection(2),
			12: connect.GetConnection(3),
			16: connect.GetConnection(4),
		},
		Post: Incidence{
			1:  connect.GetConnection(1),
			5:  connect.GetConnection(2),
			9:  connect.GetConnection(3),
			13: connect.GetConnection(4),
		},
	}
	// log.Println(IDSubNet)
	ms := MakeMotorSimulation(lfs, IDSubNet)
	go Receive(ms, IDSubNet)
	time.Sleep(2 * time.Second)
	init := TypeClock(u.InitTransition)
	end := TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}

func TestSubNetRD51(t *testing.T) {
	IDSubNet := connect.GetConnection(1)
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
				IiShotDuration: 1,
				IiListactes: []TransitionConstant{
					TransitionConstant{1, 1},
					TransitionConstant{2, -1},
				},
			},
			// T3
			TransitionDist{
				IDGlobal:       3,
				IDLocal:        2,
				IiValorLef:     1,
				IiShotDuration: 3,
				IiListactes: []TransitionConstant{
					TransitionConstant{2, 1},
					TransitionConstant{3, -1},
				},
			},
			// T4
			TransitionDist{
				IDGlobal:       4,
				IDLocal:        3,
				IiValorLef:     1,
				IiShotDuration: 2,
				IiListactes: []TransitionConstant{
					TransitionConstant{3, 1},
					TransitionConstant{-17, -1},
				},
			},
		},
		Pre: Incidence{
			0: connect.GetConnection(0),
		},
		Post: Incidence{
			17: connect.GetConnection(0),
		},
	}
	ms := MakeMotorSimulation(lfs, IDSubNet)
	go Receive(ms, IDSubNet)
	time.Sleep(2 * time.Second)
	init := TypeClock(u.InitTransition)
	end := TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}

func TestSubNetRD52(t *testing.T) {
	IDSubNet := connect.GetConnection(2)
	lfs := LefsDist{
		SubNet: TransitionList{
			// T5
			TransitionDist{
				IDGlobal:       5,
				IDLocal:        0,
				IiValorLef:     1,
				IiShotDuration: 2,
				IiListactes: []TransitionConstant{
					TransitionConstant{0, 1},
					TransitionConstant{1, -1},
				},
			},
			// T6
			TransitionDist{
				IDGlobal:       6,
				IDLocal:        1,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []TransitionConstant{
					TransitionConstant{1, 1},
					TransitionConstant{2, -1},
				},
			},
			// T7
			TransitionDist{
				IDGlobal:       7,
				IDLocal:        2,
				IiValorLef:     1,
				IiShotDuration: 3,
				IiListactes: []TransitionConstant{
					TransitionConstant{2, 1},
					TransitionConstant{3, -1},
				},
			},
			// T8
			TransitionDist{
				IDGlobal:       8,
				IDLocal:        3,
				IiValorLef:     1,
				IiShotDuration: 2,
				IiListactes: []TransitionConstant{
					TransitionConstant{3, 1},
					TransitionConstant{-17, -1},
				},
			},
		},
		Pre: Incidence{
			0: connect.GetConnection(0),
		},
		Post: Incidence{
			17: connect.GetConnection(0),
		},
	}
	ms := MakeMotorSimulation(lfs, IDSubNet)
	go Receive(ms, IDSubNet)
	time.Sleep(2 * time.Second)
	init := TypeClock(u.InitTransition)
	end := TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}

func TestSubNetRD53(t *testing.T) {
	IDSubNet := connect.GetConnection(3)
	lfs := LefsDist{
		SubNet: TransitionList{
			// T9
			TransitionDist{
				IDGlobal:       9,
				IDLocal:        0,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []TransitionConstant{
					TransitionConstant{0, 1},
					TransitionConstant{1, -1},
				},
			},
			// T10
			TransitionDist{
				IDGlobal:       10,
				IDLocal:        1,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []TransitionConstant{
					TransitionConstant{1, 1},
					TransitionConstant{2, -1},
				},
			},
			// T11
			TransitionDist{
				IDGlobal:       11,
				IDLocal:        2,
				IiValorLef:     2,
				IiShotDuration: 1,
				IiListactes: []TransitionConstant{
					TransitionConstant{2, 1},
					TransitionConstant{3, -1},
				},
			},
			// T12
			TransitionDist{
				IDGlobal:       12,
				IDLocal:        3,
				IiValorLef:     1,
				IiShotDuration: 2,
				IiListactes: []TransitionConstant{
					TransitionConstant{3, 1},
					TransitionConstant{-17, -1},
				},
			},
		},
		Pre: Incidence{
			0: connect.GetConnection(0),
		},
		Post: Incidence{
			17: connect.GetConnection(0),
		},
	}
	ms := MakeMotorSimulation(lfs, IDSubNet)
	go Receive(ms, IDSubNet)
	time.Sleep(1 * time.Second)
	init := TypeClock(u.InitTransition)
	end := TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}

func TestSubNetRD54(t *testing.T) {
	IDSubNet := connect.GetConnection(4)
	lfs := LefsDist{
		SubNet: TransitionList{
			// T13
			TransitionDist{
				IDGlobal:       13,
				IDLocal:        0,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []TransitionConstant{
					TransitionConstant{0, 1},
					TransitionConstant{1, -1},
				},
			},
			// T14
			TransitionDist{
				IDGlobal:       14,
				IDLocal:        1,
				IiValorLef:     1,
				IiShotDuration: 2,
				IiListactes: []TransitionConstant{
					TransitionConstant{1, 1},
					TransitionConstant{2, -1},
				},
			},
			// T15
			TransitionDist{
				IDGlobal:       15,
				IDLocal:        2,
				IiValorLef:     1,
				IiShotDuration: 3,
				IiListactes: []TransitionConstant{
					TransitionConstant{2, 1},
					TransitionConstant{3, -1},
				},
			},
			// T16
			TransitionDist{
				IDGlobal:       16,
				IDLocal:        3,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []TransitionConstant{
					TransitionConstant{3, 1},
					TransitionConstant{-17, -1},
				},
			},
		},
		Pre: Incidence{
			0: connect.GetConnection(0),
		},
		Post: Incidence{
			17: connect.GetConnection(0),
		},
	}
	ms := MakeMotorSimulation(lfs, IDSubNet)
	go Receive(ms, IDSubNet)
	time.Sleep(1 * time.Second)
	init := TypeClock(u.InitTransition)
	end := TypeClock(u.EndTransition)
	ms.Simulate(init, end) // ciclo 0 hasta ciclo 3
	log.Println("SDT Termino en 10s")
	time.Sleep(100 * time.Second)
}

func TestSSHDistTime5(t *testing.T) {
	for name, ip := range u.RemoteIP5T {
		addr := strings.Split(ip, ":")
		connection := u.InitSSH(addr[0])
		log.Println(connection, name, ip)
		go u.ExcecuteSSH(u.GoTest+name, connection)
	}
	time.Sleep(70 * time.Second)
}
