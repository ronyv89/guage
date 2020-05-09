package main

import (
	"github.com/urfave/cli/v2"
	"guage/internal/commands"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:     "guage",
		Usage:    "write your Go web apps the RoR way",
		Commands: commands.AllCommands(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
