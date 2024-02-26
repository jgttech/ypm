package repo

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
		Name:    "repo",
		Suggest: true,
		Usage:   "Creates the initial repo setup for the YPM package manager",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			var (
				cwd     = utils.Cwd()
				project string
				pkgJson *conf.InitPackageJson
				ypmConf *conf.YpmConf
			)

			fmt.Println()

			status, _ := spinner.New{
				Msg: "Checking system requirements",
				Init: func() (int, string) {
					if etx.PkgExists && etx.YpmExists {
						return spinner.WARNING, "Repo config already exists, nothing to do."
					}

					if etx.PkgExists && !etx.YpmExists {
						pkgPath := path.Join(etx.RepoPath, etx.Env.Repofile.Name)
						json := utils.ReadJson[conf.PackageJson](pkgPath)

						project = json.Name + "@" + json.Version

						pkgJson = &conf.InitPackageJson{}
						pkgJson.Load(project)

						return spinner.SUCCESS, "YPM config created from existing 'package.json' config"
					}

					return spinner.SUCCESS, ""
				},
			}.Run()

			if status == spinner.SUCCESS {
				status, _ = spinner.New{
					Msg: "Configuring repository",
					Init: func() (int, string) {
						if project == "" {
							project = cmd.Args().First()
						}

						if project == "" {
							return spinner.FAILURE, "Missing required property: 'name'"
						}

						if !etx.PkgExists {
							pkgJson = &conf.InitPackageJson{}

							pkgJson.Load(project)
							pkgJson.Write(cwd)
						}

						if !etx.YpmExists {
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

						return spinner.SUCCESS, ""
					},
				}.Run()
			}

			notice.Success()
			return nil
		},
	}
}
