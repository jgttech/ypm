package add

import (
	"context"
	"fmt"
	"jgttech/ypm/conf"

	"github.com/urfave/cli/v3"
)

func Cmd(etx *conf.ExecutionContext) *cli.Command {
	return &cli.Command{
		Name:  "add",
		Usage: "Fetches and adds a package",
		Action: func(ctx context.Context, c *cli.Command) error {
			fmt.Println("ADDING A PACKAGE!!")
			return nil
		},
	}
}
