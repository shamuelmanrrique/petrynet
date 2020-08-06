package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"gopkg.in/ini.v1"

	dcs "sd_petry_nets/src/distconssim"
	u "sd_petry_nets/src/utils"
)

func init() {
	testing.Init()
	gob.Register(&u.Message{})
	gob.Register(&dcs.EventDist{})
	gob.Register(dcs.IndGlobalTrans(0))
	gob.Register(dcs.TypeClock(0))
	gob.Register(&dcs.LefsDist{})
	gob.Register(&dcs.TransitionConstant{})
	gob.Register(&dcs.TransitionList{})
}

func TestSimpleSSH(t *testing.T) {
	connection := u.InitSSH("127.0.0.1")
	println(connection, "localhost", "ip")
	go u.ExcecuteSSH("ls", connection)
	time.Sleep(5 * time.Second)
}

func TestSSH(t *testing.T) {
	var logMode string
	var path string
	var environment string
	var subNetNames []string
	var subNetIDS []string

	// Loading configuration file
	cfg, err := ini.Load("../config/go.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// Getting configuration values from .ini
	environment = cfg.Section("general").Key("environment").String()
	path = cfg.Section(environment).Key("mainPath").String()
	subNetNames = strings.Split(cfg.Section(environment).Key("subNetName").String(), ",")
	subNetIDS = strings.Split(cfg.Section(environment).Key("subNetID").String(), ",")
	logMode = cfg.Section("general").Key("log").String()

	for i, ip := range subNetIDS {
		addr := strings.Split(ip, ":")
		connection := u.InitSSH(addr[0])
		
		println(path+" -name="+subNetNames[i]+" -log="+logMode, ip)

		go u.ExcecuteSSH(path+" -name="+subNetNames[i]+" -log="+logMode, connection)
	}

	time.Sleep(50 * time.Second)
}
