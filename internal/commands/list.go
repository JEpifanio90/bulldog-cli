package commands

import (
	"log"

	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"

	"github.com/JEpifanio90/bulldog-cli/internal/aws"
	"github.com/JEpifanio90/bulldog-cli/internal/gcp"
	"github.com/JEpifanio90/bulldog-cli/internal/savant"
	"github.com/JEpifanio90/bulldog-cli/models"
)

var tenants []models.Tenant
var List = cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "",
	Action: func(context *cli.Context) error {
		pterm.DefaultSpinner.Start("Getting your resources across all the platforms...")
		awsResources := aws.FetchResources()
		gcpProjects := gcp.FetchResources()
		cloudOutput(&awsResources, &gcpProjects)
		pterm.DefaultSpinner.Stop()
		printer()
		return nil
	},
}

func cloudOutput(awsResources *[]aws.Resource, gcpProjects *[]gcp.Project) {
	if awsResources != nil && len(*awsResources) > 0 {
		for _, value := range *awsResources {
			arn, _ := savant.Parse(value.ResourceARN)
			tenants = append(
				tenants,
				models.Tenant{
					AccountID: arn.AccountID,
					Platform:  arn.Partition,
					Name:      arn.Resource,
					Type:      arn.Service,
					Region:    arn.Region,
					Tags:      value.Tags,
				},
			)
		}
	}

	if gcpProjects != nil && len(*gcpProjects) > 0 {
		for _, project := range *gcpProjects {
			tenants = append(
				tenants,
				models.Tenant{
					AccountID: project.ProjectId,
					Platform:  "gcp",
					Name:      project.Name,
					Type:      "-",
					Region:    "-",
					Tags:      nil,
				},
			)
		}
	}
}

func printer() {
	tableData := pterm.TableData{{"Account ID", "Platform", "Name", "Type", "Region", "Tags"}}

	for _, tenant := range tenants {
		tableData = append(tableData, []string{tenant.AccountID, tenant.Platform, tenant.Name, tenant.Type, tenant.Region})
	}

	err := pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()

	if err != nil {
		log.Fatalln(err)
	}
}
