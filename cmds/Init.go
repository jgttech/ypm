package cmds

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func Init() *cli.Command {
	return &cli.Command{
		Name:  "init",
		Usage: "Initializes a YPM repo.",
		Action: func(ctx context.Context, c *cli.Command) error {
			fmt.Println("INITIALIZING")
			return nil
		},
	}
}
