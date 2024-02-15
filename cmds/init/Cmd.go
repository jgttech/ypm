package init

import (
	"context"
	"fmt"
	"jgttech/ypm/conf"

	"github.com/urfave/cli/v3"
)

func Cmd(etx *conf.ExecutionContext) *cli.Command {
	return &cli.Command{
		Name:    "init",
		Usage:   "Initializes a YPM repo.",
		Suggest: true,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println("INIT")
			return nil
		},
	}
}
