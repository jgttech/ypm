package main

import (
	"context"
	"os"

	"jgttech/ypm/cmds/add"
	create "jgttech/ypm/cmds/init"
	"jgttech/ypm/utils"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:    "ypm",
		Usage:   "Yet Another NodeJS Package Manager (ypm)",
		Suggest: true,
		Commands: []*cli.Command{
			add.Cmd(),
			create.Cmd(),
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		utils.Check(err)
	}
}
