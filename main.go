package main

import (
	"context"
	"jgttech/ypm/cmds/repo"
	"jgttech/ypm/sys"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	etx := sys.NewExecutionContext()

	cmd := &cli.Command{
		Name:    "ypm",
		Usage:   "Yet Another NodeJS Package Manager (ypm)",
		Suggest: true,
		Commands: []*cli.Command{
			repo.Main(etx),
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
