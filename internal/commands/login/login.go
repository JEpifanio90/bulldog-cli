package login

import (
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

var Command = cli.Command{
	Name:    "login",
	Aliases: []string{},
	Usage:   "",
	Flags:   []cli.Flag{},
	Action: func(context *cli.Context) error {
		pterm.Info.Println("BOOOM! Login")

		return nil
	},
}
