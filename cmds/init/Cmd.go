package init

import (
	"context"
	"encoding/json"
	"fmt"
	"jgttech/ypm/utils"
	"os"
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
			pkgPath := path.Join(utils.Cwd(), "package.json")
			pkgName := cmd.Args().First()

			if pkgName == "" {
				fmt.Println("Missing 'name' property.")
			}

			pkg := strings.Split(pkgName, "@")
			name := pkg[0]
			version := "0.0.1"

			if len(pkg) > 1 {
				version = pkg[1]
			}

			project := struct {
				Name    string `json:"name"`
				Version string `json:"version"`
			}{
				Name:    name,
				Version: version,
			}

			file, err := json.MarshalIndent(project, "", "  ")
			utils.Error(err)

			os.WriteFile(pkgPath, file, os.FileMode(0644))

			return nil
		},
	}
}
