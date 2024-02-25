package remove

import (
	"context"
	"fmt"
	"jgttech/ypm/conf"
	"jgttech/ypm/tui/spinner"

	"github.com/urfave/cli/v3"
)

const YPM_REPO_DESTROY = "ypmrepo"

func Cmd(etx *conf.ExecutionContext) *cli.Command {
	return &cli.Command{
		Name:    "remove",
		Suggest: true,
		Usage:   "Used to remove a package or the repo configuration itself.",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			pkgName := cmd.Args().First()

			spinner.New{
				Msg: "Verifying YPM config",
				Init: func() (int, string) {
					return spinner.SUCCESS, ""
				},
			}.Run()

			if pkgName == YPM_REPO_DESTROY {
				// spinner.New{
				// 	Msg: "Removing YPM repo config",
				// 	Init: func() (int, string) {
				// 		return spinner.SUCCESS, ""
				// 	},
				// }.Run()
			} else {
				fmt.Println("TBD: Need to remove ", pkgName)
			}

			return nil
		},
	}
}
