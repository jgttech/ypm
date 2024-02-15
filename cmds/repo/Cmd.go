package repo

import (
	"context"
	"jgttech/ypm/conf"

	"jgttech/ypm/errors"
	"jgttech/ypm/utils"

	"github.com/urfave/cli/v3"
)

func Cmd(etx *conf.ExecutionContext) *cli.Command {
	return &cli.Command{
		Name:    "repo",
		Suggest: true,
		Usage:   "Creates the initial repo setup for the YPM package manager",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			cwd := utils.Cwd()
			project := cmd.Args().First()

			if project == "" {
				errors.RequiredPropertyNotFound("name")
			}

			var pkgJson *conf.InitPackageJson
			pkgPath := utils.Join(etx.Env.Repofile.Name)
			pkgExists := utils.PathExists(pkgPath)

			if !pkgExists {
				pkgConf := utils.ParsePackageArgs(project)
				pkgJson = &conf.InitPackageJson{
					Name:    pkgConf.Name,
					Version: pkgConf.Version,
				}

				pkgJson.Write(cwd)
			}

			var ypmConf *conf.YpmConf
			ypmPath := utils.Join(etx.Env.Lockfile.Name)
			ypmExists := utils.PathExists(ypmPath)

			if !ypmExists {
				ypmConf = &conf.YpmConf{
					LockfileVersion: etx.Env.Lockfile.Version,
				}

				if pkgJson != nil {
					ypmConf.Packages = []conf.YpmConfPackage{
						{
							Name:    pkgJson.Name,
							Version: pkgJson.Version,
							Path:    ".",
						},
					}
				}

				ypmConf.Write(cwd)
			}

			return nil
		},
	}
}
