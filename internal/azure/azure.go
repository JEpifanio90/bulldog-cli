package azure

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/JEpifanio90/bulldog-cli/internal/models"
	"github.com/JEpifanio90/bulldog-cli/internal/savant"
	"github.com/pterm/pterm"
)

func FetchResources(cmd models.Command) []models.Tenant {
	var tenants []models.Tenant

	if cmd.CLI {
		tenants = useAZCli()
	}

	return tenants
}

func useAZCli() []models.Tenant {
	rawOutput, err := exec.Command("az", []string{"resource", "list", "--output", "json"}...).CombinedOutput()
	if err != nil {
		pterm.Error.Println(fmt.Errorf("list command: az cli %v", err.Error()))
		return nil
	}
	var tenants []models.Tenant

	for _, project := range parse(rawOutput) {
		id, _ := savant.ParseAZ(project.ID)
		tags, _ := json.Marshal(project.Tags)

		tenants = append(
			tenants,
			models.Tenant{
				AccountID: id,
				Platform:  "az",
				Name:      project.Name,
				Type:      project.Type,
				Region:    project.Location,
				Tags:      string(tags),
			},
		)
	}

	return tenants
}

func parse(rawOutput []byte) []models.AZResource {
	var resources []models.AZResource
	err := json.Unmarshal(rawOutput, &resources)

	if err != nil {
		pterm.Error.Println(fmt.Errorf("az cli unmarshal: %v", err))
		return nil
	}

	return resources
}
