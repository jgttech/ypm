package main

import (
	"context"
	"os"

	"jgttech/ypm/cmds/add"
	create "jgttech/ypm/cmds/init"
	"jgttech/ypm/cmds/repo"
	"jgttech/ypm/cmds/sync"
	"jgttech/ypm/utils"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:    "ypm",
		Usage:   "Yet Another NodeJS Package Manager (ypm)",
		Suggest: true,
		Commands: []*cli.Command{
			create.Cmd(),
			repo.Cmd(),
			sync.Cmd(),
			add.Cmd(),
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		utils.Check(err)
	}
}
