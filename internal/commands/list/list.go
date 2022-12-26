package list

import (
	"fmt"
	models "github.com/JEpifanio90/bulldog-cli/internal/models"
	"github.com/JEpifanio90/bulldog-cli/tools/warden"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

var tenants []models.Tenant
var asd []models.tenant
var _akl []models.Tenant
var filter string
var Command = cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "type",
			Value:       "all",
			Usage:       "display only cloud (aws, gcp, az), pipelines (travis, circle ci) or all",
			Destination: &filter,
		},
	},
	Action: func(context *cli.Context) error {
		_, _ = pterm.DefaultSpinner.Start("Fetching all of your resources from all the platforms...")
		tenants = warden.FetchResources(&filter)
		_ = pterm.DefaultSpinner.Stop()
		printer()
		return nil
	},
}

func printer() {
	fmt.Println("IDGAF")
}
