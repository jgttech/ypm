package repo

import (
	"context"
	"jgttech/ypm/conf"
	"jgttech/ypm/exceptions"
	"log"

	"jgttech/ypm/utils"

	"github.com/urfave/cli/v3"
)

func Cmd(etx *conf.ExecutionContext) *cli.Command {
	return &cli.Command{
		Name:    "repo",
		Suggest: true,
		Usage:   "Creates the initial repo setup for the YPM package manager",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			pkgName := etx.Env.Repofile.Name
			pkgPath := utils.Join(pkgName)
			pkgExists := utils.PathExists(pkgPath)

			ypmName := etx.Env.Lockfile.Name
			ypmPath := utils.Join(ypmName)
			ypmExists := utils.PathExists(ypmPath)

			if pkgExists && ypmExists {
				log.Fatal("Repo configuration exists, can't initialize the repo.")
			}

			if pkgExists && !ypmExists {
				log.Fatal("A '" + pkgName + "' file exists, can't initialize the repo.")
			}

			if !pkgExists && ypmExists {
				log.Fatal("A '" + ypmName + "' file exists, can't initialize the repo.")
			}

			cwd := utils.Cwd()
			project := cmd.Args().First()

			if project == "" {
				exceptions.RequiredPropertyNotFound("name")
			}

			var pkgJson *conf.InitPackageJson

			if !pkgExists {
				pkgJson = &conf.InitPackageJson{}

				pkgJson.Load(project)
				pkgJson.Write(cwd)
			}

			var ypmConf *conf.YpmConf

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
