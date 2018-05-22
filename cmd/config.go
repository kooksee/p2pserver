package cmd

import (
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"os"
	"io/ioutil"
	"path"
	"fmt"
)

func genCFg(_ *cli.Context) error {
	cfg.GetExtIp()
	cfg.InitDb()
	cfg.AdvertiseUdpAddr = fmt.Sprintf("%s:%d", cfg.ExtIP, cfg.UdpPort)
	cfg.AdvertiseHttpAddr = fmt.Sprintf("%s:%d", cfg.ExtIP, cfg.HttpPort)
	data, _ := yaml.Marshal(cfg)
	return ioutil.WriteFile(path.Join(cfg.Home, "config.yaml"), data, os.FileMode(0755))
}

func showCFg(_ *cli.Context) error {
	d, _ := json.MarshalToString(cfg)
	logger.Info(d)
	return nil
}

func ConfigCmd() cli.Command {
	return cli.Command{
		Name:    "config",
		Aliases: []string{"cfg"},
		Usage:   "config manager",
		Flags: []cli.Flag{
			httpPortFlag(),
			udpPortFlag(),
			logLevelFlag(),
		},
		Subcommands: cli.Commands{
			{
				Name:   "gen",
				Action: genCFg,
			},
			{
				Name:   "show",
				Action: showCFg,
			},
		},
		Action: func(c *cli.Context) error {
			return nil
		},
	}
}
