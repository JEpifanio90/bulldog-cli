package main

import (
	"fmt"
	"os"

	"github.com/JEpifanio90/bulldog-cli/internal/commands/list"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "bulldog",
		Usage: "is a CLI that empowers developers by giving them full control over their cloud accounts and pipelines",
		Commands: []*cli.Command{
			{
				Name:        "tenant",
				Aliases:     []string{"tnt"},
				Usage:       "...",
				Subcommands: []*cli.Command{&list.Command},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		pterm.Error.Println(fmt.Errorf("bulldog main: %v", err))
	}
}
