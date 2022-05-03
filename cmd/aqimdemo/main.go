package main

import (
	"log"
	"time"

	svr "github.com/nats-io/nats-server/v2/server"
	"github.com/siuyin/dflt"
	"github.com/siuyin/personal-assistant/command"
	"github.com/siuyin/personal-assistant/device"
	"github.com/siuyin/personal-assistant/internal/evt"
	"github.com/siuyin/personal-assistant/learn"
	"github.com/siuyin/personal-assistant/outp"
)

func main() {
	embedNATS()
	initSys()
	customHardwarePowerOn()

	select {} // wait forever
}

func embedNATS() {
	cfg := dflt.EnvString("LEAF_CONFIG", "./nats-leaf.conf")
	opts, err := svr.ProcessConfigFile(cfg)
	if err != nil {
		log.Fatal(err)
	}

	s, err := svr.NewServer(opts)
	if err != nil {
		log.Fatal(err)
	}
	s.Start()
	waitSvrRdy(s)
}

func waitSvrRdy(s *svr.Server) {
	timeout := time.Second
	for rdy := s.ReadyForConnections(timeout); !rdy; rdy = s.ReadyForConnections(timeout) {
	}
}

func initSys() {
	evt.Init()
	command.Init()
	learn.Init()
	outp.Init()
}

func customHardwarePowerOn() {
	device.PoweredOn()
}
