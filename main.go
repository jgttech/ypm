package main

import (
	"context"
	"jgttech/ypm/cmds/add"
	create "jgttech/ypm/cmds/init"
	"jgttech/ypm/cmds/remove"
	"jgttech/ypm/cmds/repo"
	"jgttech/ypm/cmds/sync"
	"jgttech/ypm/conf"
	"jgttech/ypm/utils"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	defer utils.Timer("Execution")()
	etx := conf.NewExecutionContext()

	cmd := &cli.Command{
		Name:    "ypm",
		Usage:   "Yet Another NodeJS Package Manager (ypm)",
		Suggest: true,
		Commands: []*cli.Command{
			create.Cmd(etx),
			repo.Cmd(etx),
			sync.Cmd(etx),
			add.Cmd(etx),
			remove.Cmd(etx),
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		utils.Check(err)
	}

	etx.Sync()
}
