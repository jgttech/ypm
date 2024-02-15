package main

import (
	"context"
	"os"

	"jgttech/ypm/cmds/add"
	create "jgttech/ypm/cmds/init"
	"jgttech/ypm/cmds/repo"
	"jgttech/ypm/cmds/sync"
	"jgttech/ypm/conf"
	"jgttech/ypm/utils"

	"github.com/urfave/cli/v3"
)

func main() {
	// Create an execution context that updates and loads
	// the state of the repo before running any commands.
	// This creates the capability to ensure that the latest
	// changes to the repo are recorded and updated within
	// the configuration.
	etx := &conf.ExecutionContext{}

	// Loads the config into memory.
	etx.Load()

	cmd := &cli.Command{
		Name:    "ypm",
		Usage:   "Yet Another NodeJS Package Manager (ypm)",
		Suggest: true,
		Commands: []*cli.Command{
			create.Cmd(etx),
			repo.Cmd(etx),
			sync.Cmd(etx),
			add.Cmd(etx),
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		utils.Check(err)
	}
}
