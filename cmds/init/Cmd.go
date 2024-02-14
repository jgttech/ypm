package init

import (
	"context"
	"fmt"
	"jgttech/ypm/utils"

	"github.com/urfave/cli/v3"
)

func Cmd() *cli.Command {
	return &cli.Command{
		Name:    "init",
		Usage:   "Initializes a YPM repo.",
		Suggest: true,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			confPath := utils.Join("package.json")
			fmt.Println(confPath)
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
