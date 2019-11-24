package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Retardos []time.Duration

type Targets []string

type CoordinatesInt interface {
	GetProcess()
	GetMaster()
	GetChandy()
	GetTimeDelay()
	GetTarget()
	GetIPsRem()
	GetRun()
	GetPort()
	GetExec()
	GetSshExc()
	GetIPuse()
}

type Coordinates struct {
	Process   int
	Master    bool
	TimeDelay Retardos
	Target    Targets
	Run       string
	IPsRem    string
	IPuse     string
	Port      string
	Exec      string
	Chandy    bool
	SshExc    bool
}

func (i *Retardos) String() string {
	return fmt.Sprint(*i)
}

func (i *Targets) String() string {
	return fmt.Sprint(*i)
}

func (i *Retardos) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("Delays flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

func (i *Targets) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("Delays flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		*i = append(*i, dt)
	}
	return nil
}

func (c Coordinates) GetProcess() int {
	return c.Process
}
func (c Coordinates) GetMaster() bool {
	return c.Master
}
func (c Coordinates) GetChandy() bool {
	return c.Chandy
}

func (c Coordinates) GetSshExc() bool {
	return c.SshExc
}
func (c Coordinates) GetTimeDelay() Retardos {
	return c.TimeDelay
}
func (c Coordinates) GetTarget() []string {
	return c.Target
}
func (c Coordinates) GetIPsRem() []string {
	aux := strings.Split(c.IPsRem, ",")
	return aux
}
func (c Coordinates) GetRun() string {
	return c.Run
}
func (c Coordinates) GetPort() string {
	return c.Port
}
func (c Coordinates) GetIPuse() string {
	return c.IPuse
}
func (c Coordinates) GetExec() string {
	return c.Exec
}
