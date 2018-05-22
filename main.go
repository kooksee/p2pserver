package main

import (
	"time"
	"github.com/kooksee/p2pserver/config"
	"os"
	"github.com/urfave/cli"
	"sort"
	"github.com/kooksee/p2pserver/cmd"
)

const Version = "1.0"

func main() {
	cfg := config.GetCfg()
	cfg.InitLog()
	cmd.SetCfg(cfg)

	app := cli.NewApp()
	app.Compiled = time.Now()
	app.Version = Version
	app.Authors = []cli.Author{{
		Name:  "pike white",
		Email: "human@example.com",
	}}
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "home", Value: cfg.Home, Destination: &cfg.Home, Usage: "app config home"},
	}
	app.Before = func(c *cli.Context) error {
		cfg.LoadConfigFile()
		return nil
	}
	app.Commands = []cli.Command{
		cmd.ServerCmd(),
		cmd.ConfigCmd(),
		cmd.AccountCmd(),
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	if err := app.Run(os.Args); err != nil {
		panic(err.Error())
	}
}
