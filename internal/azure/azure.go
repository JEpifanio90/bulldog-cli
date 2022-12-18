package azure

import (
	"encoding/json"
	"fmt"
	"github.com/JEpifanio90/bulldog-cli/internal/models"
	"github.com/JEpifanio90/bulldog-cli/tools/savant"
	"github.com/pterm/pterm"
)

func ConvertToTenants(rawOutput []byte) []models.Tenant {
	var tenants []models.Tenant

	for _, project := range parse(rawOutput) {
		id, _ := savant.ParseAZ(project.ID)

		tenants = append(
			tenants,
			models.Tenant{
				AccountID: id,
				Platform:  "az",
				Name:      project.Name,
				Type:      project.Type,
				Region:    project.Location,
				//Tags:      project.Tags,
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
