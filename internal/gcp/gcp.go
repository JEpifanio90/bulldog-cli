package gcp

import (
	"encoding/json"
	"fmt"

	"github.com/JEpifanio90/bulldog-cli/internal/models"
	"github.com/pterm/pterm"
)

func ConvertToTenants(rawOutput []byte) []models.Tenant {
	var tenants []models.Tenant

	for _, project := range parse(rawOutput) {
		tags, _ := json.Marshal(project.Labels)

		tenants = append(
			tenants,
			models.Tenant{
				AccountID: project.ProjectId,
				Platform:  "gcp",
				Name:      project.Name,
				Type:      "-",
				Region:    "-",
				Tags:      string(tags),
			},
		)
	}

	return tenants
}

func parse(rawOutput []byte) []models.GCPProject {
	var projects []models.GCPProject
	err := json.Unmarshal(rawOutput, &projects)

	if err != nil {
		pterm.Error.Println(fmt.Errorf("gcp cli unmarshal: %v", err))
		return nil
	}

	return projects
}
