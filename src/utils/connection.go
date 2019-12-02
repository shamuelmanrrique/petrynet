package utils

import (
	"strings"
)

var InitTransition = 0
var EndTransition = 3
var Dir = "/home/a802400/"
var GoMainLog = "/usr/local/go/bin/go run /home/a802400/go/src/github.com/shamuelmanrrique/petrynet/src/cmd/distconssim/main.go"
var GoMain = "/usr/local/go/bin/go run /home/a802400/go/src/github.com/shamuelmanrrique/petrynet/src/cmd/distconssim/main.go -l=false"
var GoTest = "/usr/local/go/bin/go test /home/a802400/go/src/github.com/shamuelmanrrique/petrynet/src/distconssim -timeout 499s -v -run "

// var GoTest = "/usr/local/go/bin/go test /home/a802400/go/src/github.com/shamuelmanrrique/petrynet/src/distconssim -timeout 99s -v -run "
var SSHIPs = []string{"155.210.154.199,155.210.154.209,155.210.154.208"}
var LocalIP3s = []string{"127.0.1.1:5000", "127.0.1.1:5001", "127.0.1.1:5002"}

// Inf Remote IPs
var RemoteIP3 = map[string]string{"TestSubNet0": "155.210.154.199", "TestSubNet1": "155.210.154.200", "TestSubNet2": "155.210.154.204"}
var RemoteIP3s = []string{"155.210.154.199:1400", "155.210.154.200:1400", "155.210.154.204:1400"}

var LocalIPs = []string{"127.0.1.1:5000", "127.0.1.1:5001", "127.0.1.1:5002", "127.0.1.1:5003",
	"127.0.1.1:5004", "127.0.1.1:5005", "127.0.1.1:5006"}

// var RemoteIP2s = []string{"155.210.154.207", "155.210.154.208", "155.210.154.209"}
var RemoteIP1s = []string{"155.210.154.199:1400", "155.210.154.197:1400", "155.210.154.196:1400"}
var Command = make(map[string]string)

// Connections is an array of connection
type Connections []Connect

// Connection is a interface
type Connection interface {
	GetIDSubRed() string
	GetIp() string
	GetPort() string
	GetIDs() []string
	GetAccept() bool
	SetAccept(b bool)
}

// Connect is a struct that contains information about connection
type Connect struct {
	IDSubRed, IP, Port string
	IDs                []string
	Delays             []int
	Accept             bool
}

func (c Connect) GetIDSubRed() string {
	return c.IDSubRed
}

func (c Connect) GetIp() string {
	return c.IP
}

func (c Connect) GetPort() string {
	return c.Port
}

func (c Connect) GetAccept() bool {
	return c.Accept
}

func (c Connect) GetIds() []string {
	return c.IDs
}

func (c *Connect) SetAccept(b bool) {
	c.Accept = b
}

// NewConnec will create slice of Connect
func NewConnec(IPs []string) Connections {
	var connections Connections
	for _, val := range IPs {
		addr := strings.Split(val, ":")
		conn := Connect{}
		conn.IP = addr[0]
		conn.Port = addr[1]
		conn.Accept = false
		conn.IDSubRed = val
		connections = append(connections, conn)
	}
	return connections
}

// GetConnection return connection by Index in slices
func (c Connections) GetConnection(n int) Connect {
	for i, connect := range c {
		if i == n {
			return connect
		}
	}
	return Connect{}
}

// func (c *u.Connect) SetAddrs(inc Incidence) {
// 	for _, addr := range inc {
// 		*c.IDs = append(*c.IDs, addr)
// 	}
// }
