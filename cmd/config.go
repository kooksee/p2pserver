package cmd

import (
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"os"
	"io/ioutil"
	"path"
	"fmt"
	"github.com/kooksee/crypt"
	"encoding/hex"
)

func genCFg(_ *cli.Context) error {
	EnsureDir(cfg.Home, os.FileMode(0755))
	EnsureDir(cfg.DbPath, os.FileMode(0755))
	EnsureDir(cfg.LogPath, os.FileMode(0755))

	cfg.GetExtIp()
	if cfg.AdvertiseUdpAddr == "" {
		cfg.AdvertiseUdpAddr = fmt.Sprintf("%s:%d", cfg.ExtIP, cfg.UdpPort)
	}

	if cfg.AdvertiseHttpAddr == "" {
		cfg.AdvertiseHttpAddr = fmt.Sprintf("%s:%d", cfg.ExtIP, cfg.HttpPort)
	}

	priv, err := crypto.GenerateKey()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	k := hex.EncodeToString(crypto.FromECDSA(priv))
	if err := ioutil.WriteFile(cfg.PriV, []byte(k), os.FileMode(0755)); err != nil {
		return err
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join(cfg.Home, "config.yaml"), data, os.FileMode(0755)); err != nil {
		return err
	}
	return nil
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
