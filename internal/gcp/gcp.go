package gcp

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/JEpifanio90/bulldog-cli/internal/models"
	"github.com/pterm/pterm"
)

func FetchResources(cmd models.Command) []models.Tenant {
	var tenants []models.Tenant

	if cmd.CLI {
		tenants = useGCloud()
	}

	return tenants
}

func useGCloud() []models.Tenant {
	rawOutput, err := exec.Command("gcloud", []string{"projects", "list", "--format", "json"}...).CombinedOutput()
	if err != nil {
		pterm.Error.Println(fmt.Errorf("list command: gcp cli %v", err.Error()))
		return nil
	}
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
