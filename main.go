package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
	"jgttech/ypm/cmds"
)

func main() {
	cmd := &cli.Command{
		Name:  "ypm",
		Usage: "Yet Another NodeJS Package Manager (ypm)",
		Commands: []*cli.Command{
			cmds.Init(),
			cmds.Add(),
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
