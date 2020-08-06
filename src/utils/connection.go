package utils

import (
	"strings"
)

// InitTransition time to init transtition
var InitTransition = 0

// EndTransition time to end transtition
var EndTransition = 4

var Dir = "/home/a802400/"
var NetName = ""

// Inf Remote IPs
var RemoteIP3 = map[string]string{"TestSubNet0": "155.210.154.199:1400", "TestSubNet1": "155.210.154.200:1400", "TestSubNet2": "155.210.154.204:1400"}
var RemoteIP3s = []string{"155.210.154.199:1400", "155.210.154.200:1400", "155.210.154.204:1400"}
var RemoteIP5 = map[string]string{"TestSubNetR50": "155.210.154.199:1400", "TestSubNetR51": "155.210.154.200:1400", "TestSubNetR52": "155.210.154.204:1400", "TestSubNetR53": "155.210.154.209:1400", "TestSubNetR54": "155.210.154.210:1400"}
var RemoteIP5T = map[string]string{"TestSubNetRD50": "155.210.154.199:1400", "TestSubNetRD51": "155.210.154.200:1400", "TestSubNetRD52": "155.210.154.204:1400", "TestSubNetRD53": "155.210.154.209:1400", "TestSubNetRD54": "155.210.154.210:1400"}
var RemoteIP5s = []string{"155.210.154.199:1400", "155.210.154.200:1400", "155.210.154.204:1400", "155.210.154.204:1400", "155.210.154.210:1400"}
var GoMainLog = "/usr/local/go/bin/go run /home/a802400/go/src/sd_petry_nets/src/cmd/distconssim/main.go"
var GoMain = "/usr/local/go/bin/go run /home/a802400/go/src/sd_petry_nets/src/cmd/distconssim/main.go -l=false"
var GoTest = "/usr/local/go/bin/go test -timeout 499s -v /home/a802400/go/src/sd_petry_nets/src/distconssim -run "

// Inf Local IPs
var LocalIP3 = map[string]string{"TestSubNetL0": "127.0.1.1:5000", "TestSubNetL1": "127.0.1.1:5001", "TestSubNetL2": "127.0.1.1:5002"}
var LocalIP3s = []string{"127.0.1.1:5000", "127.0.1.1:5001", "127.0.1.1:5002"}
var LocalIP5 = map[string]string{"TestSubNetL0": "127.0.1.1:5000", "TestSubNetL1": "127.0.1.1:5001", "TestSubNetL2": "127.0.1.1:5002", "TestSubNetL3": "127.0.1.1:5003", "TestSubNetL4": "127.0.1.1:5004"}
var LocalIP5s = []string{"127.0.1.1:5000", "127.0.1.1:5001", "127.0.1.1:5002", "127.0.1.1:5003", "127.0.1.1:5004"}
var GoLocalMainLog = "go run ~/go/src/sd_petry_nets/src/cmd/distconssim/main.go"
var GoLocalTest = "go test ~/go/src/sd_petry_nets/src/distconssim -timeout 499s -v -run "

// var LocalIPs = []string{"127.0.1.1:5000", "127.0.1.1:5001", "127.0.1.1:5002", "127.0.1.1:5003",
// 	"127.0.1.1:5004"}

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
