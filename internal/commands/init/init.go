package init

import "github.com/urfave/cli/v2"

func InitCommands() *cli.Command {
	return &cli.Command{
		Name:    "new",
		Aliases: []string{"n"},
		Usage:   "create a new Guage project",
		Action: func(context *cli.Context) error {

			return nil
		},
	}
}
