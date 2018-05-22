package cmd

import (
	"time"
	"sort"
	"github.com/urfave/cli"
)

func RunCmd(args []string) {
	app := cli.NewApp()
	app.Compiled = time.Now()
	app.Version = cfg.Version
	app.Authors = []cli.Author{{Name: "pike white", Email: "human@example.com"}}
	app.Commands = []cli.Command{ServerCmd(), ConfigCmd(), AccountCmd()}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	if err := app.Run(args); err != nil {
		cfg.Log().Error(err.Error())
	}
}
