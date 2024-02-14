package repo

import (
	"context"

	create "jgttech/ypm/cmds/init"
	"jgttech/ypm/utils"

	"github.com/urfave/cli/v3"
)

func Cmd() *cli.Command {
	return &cli.Command{
		Name:    "repo",
		Suggest: true,
		Usage:   "Creates the initial repo setup for the YPM package manager",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			confPath := utils.Join("package.json")
			confExists := utils.PathExists(confPath)

			if !confExists {
				create.CreatePackageJson(ctx, cmd)
			}

			ypmPath := utils.Join(".ypm")
			ypmExists := utils.PathExists(ypmPath)

			if !ypmExists {
				CreateYpmConfig(ctx, cmd)
			}

			return nil
		},
	}
}
