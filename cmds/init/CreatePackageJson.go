package init

import (
	"context"
	"jgttech/ypm/conf"
	"jgttech/ypm/errors"
	"jgttech/ypm/utils"
	"log"
	"path"
	"strings"

	"github.com/urfave/cli/v3"
	"golang.org/x/mod/semver"
)

func CreatePackageJson(ctx context.Context, cmd *cli.Command) {
	cwd := path.Join(utils.Cwd())
	projectName := cmd.Args().First()

	if projectName == "" {
		errors.RequiredPropertyNotFound("name")
	}

	idx := strings.LastIndex(projectName, "@")
	name := projectName
	version := "0.0.1"

	if idx != -1 && idx != 0 {
		name = projectName[0:idx]
		confVersion := projectName[idx+1:]
		isValid := semver.IsValid("v" + confVersion)

		if !isValid {
			log.Fatal("Invalid version. Must be a valid semver format. Did you have an extra \"@\" character somewhere?")
		}

		version = confVersion
	}

	pkg := conf.InitPackageJson{
		Name:    name,
		Version: version,
	}

	pkg.Write(cwd)
}
