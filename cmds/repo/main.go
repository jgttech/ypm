package repo

import (
	"context"
	"fmt"
	"jgttech/ypm/notice"
	"jgttech/ypm/sys"
	"jgttech/ypm/tui/confirm"

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
			var input string

			// repo := NewRepoConf(cmd.Args().First())

			fmt.Println()

			if directory {
				confirm.New{
					Message: "Is this directory correct?",
					Validate: func(s string) error {
						return nil
					},
				}.Run()
			}

			fmt.Println(input)

			// status, _ := spinner.New{
			// 	Msg: "Verifying YPM config",
			// 	Init: func() (int, string) {
			// 		if etx.PkgExists && etx.RepoExists {
			// 			return spinner.WARNING, "Repo config already exists, nothing to do."
			// 		}
			//
			// 		return spinner.SUCCESS, ""
			// 	},
			// }.Run()

			// if status == spinner.SUCCESS {
			// 	status, _ = spinner.New{
			// 		Msg: "Configuring repository",
			// 		Init: func() (int, string) {
			// 			var pkgPath string
			// 			var pkgJson *conf.PackageJson
			// 			var repoConf *conf.YpmConf
			//
			// 			fmt.Printf("%#v\n", repo)
			//
			// 			if !etx.PkgExists && !etx.RepoExists {
			// 				fmt.Println("TESTING")
			// 			}
			//
			// 			return spinner.SUCCESS, ""
			// 		},
			// 	}.Run()
			// }

			notice.Done()
			return nil
		},
	}
}
