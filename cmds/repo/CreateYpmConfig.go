package repo

import (
	"context"
	"jgttech/ypm/conf"
	"jgttech/ypm/utils"
	"os"

	"github.com/urfave/cli/v3"
)

const CONF_DIR string = ".ypm"
const REPO_CONFIG string = CONF_DIR + "/repo.yml"

func CreateYpmConfig(ctx context.Context, cmd *cli.Command) {
	// No need to read the args passed in because, by the time
	// we reach this point, the "package.json" exists and we
	// can use that information to build the initial config for
	// YPM, itself.

	pkgPath := utils.Join("package.json")
	pkgConf := utils.ReadJson[conf.InitPackageJson](pkgPath)
	ypmPath := utils.Join(CONF_DIR)

	os.MkdirAll(ypmPath, os.ModePerm)

	ypmConf := &conf.YpmConfigYaml{
		Workspaces: []conf.YpmConfigWorkspace{
			{
				Name: pkgConf.Name,
				Path: ".",
			},
		},
	}

	ypmConf.Write(REPO_CONFIG)
}
