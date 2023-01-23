package gcp

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/pterm/pterm"
	"google.golang.org/api/cloudresourcemanager/v1"

	"github.com/JEpifanio90/bulldog-cli/internal/models"
)

func FetchResources(cmd models.Command) []models.Tenant {
	var tenants []models.Tenant

	if cmd.CLI {
		tenants = useCli()
	} else {
		tenants = useCdk()
	}

	return tenants
}

func useCli() []models.Tenant {
	rawOutput, err := exec.Command("gcloud", []string{"projects", "list", "--format", "json"}...).CombinedOutput()
	if err != nil {
		pterm.Error.Println(fmt.Errorf("list command: gcp cli %w", err))

		return nil
	}

	return parseAndConvert(rawOutput, true)
}

func useCdk() []models.Tenant {
	client, err := cloudresourcemanager.NewService(context.Background())
	if err != nil {
		pterm.Error.Println(fmt.Errorf("gcp sdk authentication: %w", err))
		return make([]models.Tenant, 0)
	}

	response, err := client.Projects.List().Do()
	if err != nil {
		pterm.Warning.Println(fmt.Errorf("gcp projects: %w", err))
		return make([]models.Tenant, 0)
	}

	var projects []cloudresourcemanager.Project

	for _, pointerProject := range response.Projects {
		projects = append(projects, *pointerProject)
	}

	return parseAndConvert(projects, false)
}

func parseAndConvert[T []byte | []cloudresourcemanager.Project](input T, unmarshal bool) []models.Tenant {
	var projects []cloudresourcemanager.Project
	if _, ok := any(input).([]byte); unmarshal && ok {
		err := json.Unmarshal(any(input).([]byte), &projects)

		if err != nil {
			pterm.Error.Println(fmt.Errorf("gcp unmarshal: %w", err))
			return nil
		}
	} else {
		projects = any(input).([]cloudresourcemanager.Project)
	}

	var tenants []models.Tenant

	for _, project := range projects {
		tags, _ := json.Marshal(project.Labels)
		tenants = append(
			tenants,
			models.Tenant{
				AccountID: project.ProjectId,
				Platform:  "gcp",
				Name:      project.Name,
				Type:      string(rune(project.ProjectNumber)),
				Region:    "-",
				Tags:      string(tags),
			},
		)
	}

	return tenants
}
