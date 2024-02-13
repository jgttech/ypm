package repo

import (
	"context"
	"jgttech/ypm/utils"
	"os"
	"path"

	"github.com/urfave/cli/v3"
)

const CONF_DIR string = ".ypm"
const WORKSPACES string = CONF_DIR + "/workspaces"

func CreateYpmConfig(ctx context.Context, cmd *cli.Command) {
	// No need to read the args passed in because, by the time
	// we reach this point, the "package.json" exists and we
	// can use that information to build the initial config for
	// YPM, itself.

	// pkgPath := path.Join(utils.Cwd(), "package.json")
	// pkgConf := fsutils.ReadJson[conf.InitPackageJson](pkgPath)
	ypmPath := path.Join(utils.Cwd(), CONF_DIR)

	os.MkdirAll(ypmPath, os.ModePerm)
}
