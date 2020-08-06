package distconssim

import (
	"encoding/gob"
	"log"
	"strings"
	"testing"
	"time"

	u "sd_petry_nets/src/utils"
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

var conects = u.NewConnec(u.LocalIP5s)

func TestSubNetR50(t *testing.T) {
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
					TransitionConstant{-5, -1},
					TransitionConstant{-7, -1},
				},
			},
			// T9
			TransitionDist{
				IDGlobal:       9,
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
			2: conects.GetConnection(1),
			4: conects.GetConnection(2),
			6: conects.GetConnection(3),
			8: conects.GetConnection(4),
		},
		Post: Incidence{
			1: conects.GetConnection(1),
			3: conects.GetConnection(2),
			5: conects.GetConnection(3),
			7: conects.GetConnection(4),
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

func TestSubNetR51(t *testing.T) {
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
					TransitionConstant{-9, -1},
				},
			},
		},
		Pre: Incidence{
			0: conects.GetConnection(0),
		},
		Post: Incidence{
			9: conects.GetConnection(0),
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

func TestSubNetR52(t *testing.T) {
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
					TransitionConstant{-9, -1},
				},
			},
		},
		Pre: Incidence{
			0: conects.GetConnection(0),
		},
		Post: Incidence{
			9: conects.GetConnection(0),
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

func TestSubNetR53(t *testing.T) {
	IDSubNet := conects.GetConnection(3)
	lfs := LefsDist{
		SubNet: TransitionList{
			// T5
			TransitionDist{
				IDGlobal:       5,
				IDLocal:        0,
				IiValorLef:     1,
				IiShotDuration: 1,
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
					TransitionConstant{-9, -1},
				},
			},
		},
		Pre: Incidence{
			0: conects.GetConnection(0),
		},
		Post: Incidence{
			9: conects.GetConnection(0),
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

func TestSubNetR54(t *testing.T) {
	IDSubNet := conects.GetConnection(4)
	lfs := LefsDist{
		SubNet: TransitionList{
			// T7
			TransitionDist{
				IDGlobal:       7,
				IDLocal:        0,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []TransitionConstant{
					TransitionConstant{0, 1},
					TransitionConstant{1, -1},
				},
			},
			// T8
			TransitionDist{
				IDGlobal:       8,
				IDLocal:        1,
				IiValorLef:     1,
				IiShotDuration: 1,
				IiListactes: []TransitionConstant{
					TransitionConstant{1, 1},
					TransitionConstant{-9, -1},
				},
			},
		},
		Pre: Incidence{
			0: conects.GetConnection(0),
		},
		Post: Incidence{
			9: conects.GetConnection(0),
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

func TestSSHDist5(t *testing.T) {
	for name, ip := range u.RemoteIP5 {
		addr := strings.Split(ip, ":")
		connection := u.InitSSH(addr[0])
		log.Println(connection, name, ip)
		go u.ExcecuteSSH(u.GoTest+name, connection)
	}
	time.Sleep(70 * time.Second)
}
