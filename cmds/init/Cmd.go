package init

import (
	"context"
	"jgttech/ypm/conf"
	"jgttech/ypm/tui/spinner"
	"time"

	"github.com/urfave/cli/v3"
)

func Cmd(etx *conf.ExecutionContext) *cli.Command {
	return &cli.Command{
		Name:    "init",
		Usage:   "Initializes a YPM repo.",
		Suggest: true,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			spinnerIndicator := spinner.New{
				Msg: "Initializing workspace, please wait...",
				Init: func() (int, string) {
					time.Sleep(2 * time.Second)
					return 1, ""
				},
			}

			spinnerIndicator.Run()

			return nil
		},
	}
}
