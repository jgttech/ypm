package init

import (
	"context"
	"fmt"
	"jgttech/ypm/utils"
	"path"

	"github.com/urfave/cli/v3"
)

func Cmd() *cli.Command {
	return &cli.Command{
		Name:    "init",
		Usage:   "Initializes a YPM repo.",
		Suggest: true,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			cwd := path.Join(utils.Cwd())
			confPath := path.Join(cwd, "package.json")
			confExists := utils.PathExists(confPath)

			if !confExists {
				CreatePackageJson(ctx, cmd)
			} else {
				fmt.Println("Package already has a \"package.json\". No \"package.json\" created.")
			}

			return nil
		},
	}
}
