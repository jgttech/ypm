package sync

import (
	"context"
	"fmt"
	"jgttech/ypm/conf"

	"github.com/urfave/cli/v3"
)

func Cmd(etx *conf.ExecutionContext) *cli.Command {
	return &cli.Command{
		Name:    "sync",
		Suggest: true,
		Usage:   "Detect the current system configuration and ensure it is in sync with YPM.",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println("SYNC")
			return nil
		},
	}
}
