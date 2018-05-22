package cmd

import (
	"github.com/json-iterator/go"
	"github.com/kooksee/log"
	"github.com/kooksee/p2pserver/config"
	"github.com/urfave/cli"
)

var (
	json   = jsoniter.ConfigCompatibleWithStandardLibrary
	logger log.Logger
	cfg    *config.Config
)

func Init() {
	cfg = config.GetCfg()
	logger = config.GetLog("package", "cmd")
}

func homeFlag() cli.StringFlag     { return cli.StringFlag{Name: "home", Value: cfg.Home, Destination: &cfg.Home, Usage: "app config home"} }
func httpPortFlag() cli.IntFlag    { return cli.IntFlag{Name: "hp", Value: cfg.HttpPort, Destination: &cfg.HttpPort, Usage: "http port"} }
func udpPortFlag() cli.IntFlag     { return cli.IntFlag{Name: "up", Value: cfg.UdpPort, Destination: &cfg.UdpPort, Usage: "udp port"} }
func logLevelFlag() cli.StringFlag { return cli.StringFlag{Name: "ll", Value: cfg.LogLevel, Destination: &cfg.LogLevel, Usage: "log level"} }
