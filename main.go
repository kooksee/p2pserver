package main

import (
	"os"
	"github.com/kooksee/p2pserver/cmd"
	"github.com/kooksee/p2pserver/config"
)

const Version = "1.0"

func main() {
	home := "kdata"
	if len(os.Args) > 2 && os.Args[len(os.Args)-2] == "--home" {
		home = os.Args[len(os.Args)-1]
		os.Args = os.Args[:len(os.Args)-2]
	}
	cfg := config.GetCfg(home)
	cfg.Version = Version
	cmd.SetCfg(cfg)
	cmd.RunCmd(os.Args)
}
