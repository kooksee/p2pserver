package main

import (
	"github.com/kooksee/p2pserver/cmd"
	"github.com/kooksee/p2pserver/config"
)

func main() {
	cfg := config.NewCfg("kdata")
	cfg.InitConfigFile()
	cfg.InitLog()

	cmd.Init()
	cmd.RunCmd()
}
