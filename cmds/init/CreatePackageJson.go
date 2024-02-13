package init

import (
	"context"
	"jgttech/ypm/conf"
	"jgttech/ypm/errors"
	"jgttech/ypm/utils"
	"path"

	"github.com/urfave/cli/v3"
)

func CreatePackageJson(ctx context.Context, cmd *cli.Command) {
	cwd := path.Join(utils.Cwd())
	project := cmd.Args().First()

	if project == "" {
		errors.RequiredPropertyNotFound("name")
	}

	pkgConf := utils.ParsePackageArgs(project)
	pkg := conf.InitPackageJson{
		Name:    pkgConf.Name,
		Version: pkgConf.Version,
	}

	pkg.Write(cwd)
}
