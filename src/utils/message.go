package utils

import "log"

// Msm int
type Msm interface {
	GetTo() string
	GetFrom() string
	GetPack() interface{}
}

// Mensssage Struct
type Message struct {
	To, From string
	Pack     interface{} // Can be Type Event, LookA, IndGlobalTrans
}

func (m Message) GetTo() string {
	return m.To
}

func (m Message) GetFrom() string {
	return m.From
}

func (m Message) GetPack() interface{} {
	return m.Pack
}

func DistMsm(s string) {
	log.Printf("###################### MAIN  %s ########################### \n", s)
}

func DistWall() {
	log.Println("############################################################################")
}

func DistL() {
	log.Println("----------------------------------------------------------------------------")
}

func DistUnic(s string) {
	log.Printf("#########################  %s ################################# \n", s)
}
