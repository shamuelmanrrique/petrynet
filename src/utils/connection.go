package utils

type Connection interface {
	GetId() string
	GetIp() string
	GetPort() string
	GetIds() []string
	GetAccept() bool
}

type Connect struct {
	Id, Ip, Port, Host string
	Ids                []string
	Delays             []int
	Accept             bool
}

var RemoteIPs = []string{"155.210.154.199", "155.210.154.197", "155.210.154.196"}
var RemoteFlags = []string{"-r=\"proof\" -i=\"155.210.154.199\" -t=\"155.210.154.197\" -d=\"5s\" -n=3 -m=true -p=\":1400\"", "-i=\"155.210.154.197\" -r=\"proof\" -n=3 -p=\":1400\"", "-i=\"155.210.154.196\" -r=\"proof\" -n=3 -p=\":1400\""}
var LocalFlags = []string{"-r=\"local\" -t=\"127.0.1.1:5002\" -d=\"5s\" -n=3 -m=true -p=\":5001\"", "-r=\"local\" -n=3 -p=\":5002\"", " -r=\"local\" -n=3 -p=\":5003\""}
var Command = make(map[string]string)

func (c Connect) GetId() string {
	return c.Id
}

func (c Connect) GetIp() string {
	return c.Ip
}

func (c Connect) GetPort() string {
	return c.Port
}

func (c Connect) GetAccept() bool {
	return c.Accept
}

func (c Connect) GetIds() []string {
	return c.Ids
}
