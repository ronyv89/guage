package commands

import (
	"github.com/urfave/cli/v2"
	"guage/internal/commands/init"
)

func AllCommands() []*cli.Command {
	return []*cli.Command{
		init.InitCommands(),
	}
}
