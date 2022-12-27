package list

import (
	"log"

	"github.com/JEpifanio90/bulldog-cli/internal/models"
	"github.com/JEpifanio90/bulldog-cli/tools/warden"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

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
		tenants := warden.FetchResources(&filter)
		_ = pterm.DefaultSpinner.Stop()
		printer(&tenants)
		return nil
	},
}

func printer(tenants *[]models.Tenant) {
	tableData := pterm.TableData{{"Account ID", "Platform", "Name", "Type", "Region", "Tags"}}

	for _, tenant := range *tenants {
		tableData = append(tableData, []string{tenant.AccountID, tenant.Platform, tenant.Name, tenant.Type, tenant.Region})
	}

	err := pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()

	if err != nil {
		log.Fatalln(err)
	}
}
