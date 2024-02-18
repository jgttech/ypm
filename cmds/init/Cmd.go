package init

import (
	"context"
	"jgttech/ypm/conf"
	tuiSpinner "jgttech/ypm/tui/tui-spinner"
	"time"

	"github.com/urfave/cli/v3"
)

func Cmd(etx *conf.ExecutionContext) *cli.Command {
	return &cli.Command{
		Name:    "init",
		Usage:   "Initializes a YPM repo.",
		Suggest: true,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			spinner := tuiSpinner.New()

			spinner.Run(tuiSpinner.Runner{
				Msg: "Initializing workspace, please wait...",
				Init: func() {
					time.Sleep(3 * time.Second)
				},
			})

			spinner.Run(tuiSpinner.Runner{
				Msg: "Updating repo state...",
				Init: func() {
					time.Sleep(3 * time.Second)
				},
			})

			return nil
		},
	}
}
