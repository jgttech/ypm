package repo

import (
	"context"
	"fmt"
	"jgttech/ypm/sys"

	"github.com/urfave/cli/v3"
)

func Main(etx *sys.ExecutionContext) *cli.Command {
	var directory bool

	return &cli.Command{
		Name:    "repo",
		Usage:   "Create a YPM root repo configuration",
		Suggest: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "directory",
				Aliases:     []string{"D"},
				Destination: &directory,
				Usage:       "Indicates that the repo command should create a directory",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			repo := NewRepoConf(cmd.Args().First())

			fmt.Printf("%#v", repo)

			return nil
		},
	}
}
