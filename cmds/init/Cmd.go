package init

import (
	"context"
	"fmt"
	"jgttech/ypm/conf"
	"jgttech/ypm/notice"
	"jgttech/ypm/tui/spinner"
	"jgttech/ypm/utils"
	"path"

	"github.com/urfave/cli/v3"
)

func Cmd(etx *conf.ExecutionContext) *cli.Command {
	return &cli.Command{
		Name:    "init",
		Usage:   "Initializes a YPM repo.",
		Suggest: true,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			var pkg string

			fmt.Println()

			spinner.New{
				Msg: "Verifying YPM config",
				Init: func() (int, string) {
					if etx.RepoPath == etx.WorkingPath {
						return spinner.FAILURE, "Can't create a new package in repo root"
					}

					if !etx.PkgExists || !etx.YpmExists {
						return spinner.FAILURE, "YPM is not configured"
					}

					return spinner.SUCCESS, ""
				},
			}.Run()

			spinner.New{
				Msg: "Configuring package",
				Init: func() (int, string) {
					if pkg == "" {
						pkg = cmd.Args().First()

						if pkg == "" {
							return spinner.FAILURE, "Missing required repo name (i.e <name>[@version])"
						}
					}

					pkgPath := path.Join(etx.WorkingPath, etx.Env.Repofile.Name)
					pkgExists := utils.PathExists(pkgPath)

					if !pkgExists {
						pkgJson := &conf.InitPackageJson{}

						pkgJson.Load(pkg)

						if isValid := etx.IsValidPackageName(pkgJson.Name); !isValid {
							return spinner.FAILURE, "A package named '" + pkgJson.Name + "' already exists."
						}

						pkgJson.Write(etx.WorkingPath)
					}

					return spinner.SUCCESS, ""
				},
			}.Run()

			notice.Success()
			return nil
		},
	}
}
