package init

import (
	"context"
	"fmt"
	"jgttech/ypm/conf"
	"jgttech/ypm/utils"
	"path"
	"strings"

	"github.com/urfave/cli/v3"
)

func Cmd() *cli.Command {
	return &cli.Command{
		Name:    "init",
		Usage:   "Initializes a YPM repo.",
		Suggest: true,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			cwd := path.Join(utils.Cwd())
			projectName := cmd.Args().First()

			if projectName == "" {
				fmt.Println("Missing 'name' property.")
			}

			project := strings.Split(projectName, "@")
			name := project[0]
			version := "0.0.1"

			if len(project) > 1 {
				version = project[1]
			}

			pkg := conf.InitPackageJson{
				Name:    name,
				Version: version,
			}

			pkg.Write(cwd)

			return nil
		},
	}
}
