package init

import (
	"context"

	"github.com/urfave/cli/v3"
)

func Cmd() *cli.Command {
	return &cli.Command{
		Name:    "init",
		Usage:   "Initializes a YPM repo.",
		Suggest: true,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			CreatePackageJson(ctx, cmd)
			CreateYpmConfig(ctx, cmd)

			return nil
		},
	}
}
