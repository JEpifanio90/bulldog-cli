package list

import (
	"log"
	"os/exec"
	"sync"

	"github.com/JEpifanio90/bulldog-cli/internal/travis_ci"

	"github.com/JEpifanio90/bulldog-cli/internal/azure"

	"github.com/JEpifanio90/bulldog-cli/internal/gcp"

	"github.com/JEpifanio90/bulldog-cli/internal/aws"

	"github.com/JEpifanio90/bulldog-cli/internal/models"
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
		commands := fetchCommands()
		var tenants []models.Tenant
		var waitGroup sync.WaitGroup

		for _, cmd := range commands {
			waitGroup.Add(1)
			go executioner(cmd, &tenants, &waitGroup)
		}

		waitGroup.Wait()
		_ = pterm.DefaultSpinner.Stop()
		printer(&tenants)
		return nil
	},
}

func fetchCommands() []models.Command {
	commands := []models.Command{
		{Name: "aws", CLI: cliExists("aws")},
		{Name: "gcloud", CLI: cliExists("gcloud")},
		{Name: "az", CLI: cliExists("az")},
		{Name: "travis", CLI: cliExists("travis")},
	}

	if filter == "cloud" {
		commands = commands[:len(commands)-1]
	}

	if filter == "pipelines" {
		commands = commands[len(commands):]
	}

	return commands
}

func cliExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func executioner(cmd models.Command, tenants *[]models.Tenant, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	switch cmd.Name {
	case "aws":
		*tenants = append(*tenants, aws.FetchResources(cmd)...)
	case "gcloud":
		*tenants = append(*tenants, gcp.FetchResources(cmd)...)
	case "az":
		*tenants = append(*tenants, azure.FetchResources(cmd)...)
	case "travis":
		*tenants = append(*tenants, travis_ci.FetchResources(cmd)...)
	}
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
