package azure

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/pterm/pterm"

	"github.com/JEpifanio90/bulldog-cli/internal/models"
	"github.com/JEpifanio90/bulldog-cli/internal/savant"
)

func FetchResources(cmd models.Command) []models.Tenant {
	var tenants []models.Tenant

	if cmd.CLI {
		tenants = useAZCli()
	} else {
		tenants = useAPI()
	}

	return tenants
}

func useAPI() []models.Tenant {
	resp, err := http.Get(fmt.Sprintf(
		"https://management.azure.com/subscriptions/%s/resources?api-version=2021-04-01",
		os.Getenv("AZURE_SUBSCRIPTION")),
	)
	if err != nil {
		pterm.Error.Println(fmt.Errorf("az api: %w", err))
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		pterm.Error.Println(fmt.Errorf("az api: %w", err))
	}

	resources := parse(body, true)

	return convertToTenants(resources)
}

func useAZCli() []models.Tenant {
	rawOutput, err := exec.Command("az", []string{"resource", "list", "--output", "json"}...).CombinedOutput()
	if err != nil {
		pterm.Error.Println(fmt.Errorf("list command: az cli %w", err))

		return nil
	}

	azResources := parse(rawOutput, false)

	return convertToTenants(azResources)
}

func parse(rawOutput []byte, hasWrapper bool) []models.AZResource {
	if hasWrapper {
		var azWrapper models.AZResponse

		err := json.Unmarshal(rawOutput, &azWrapper)

		if err != nil {
			pterm.Error.Println(fmt.Errorf("az cli unmarshal: %w", err))

			return nil
		}

		return azWrapper.Value
	}

	var resources []models.AZResource
	err := json.Unmarshal(rawOutput, &resources)

	if err != nil {
		pterm.Error.Println(fmt.Errorf("az cli unmarshal: %w", err))

		return nil
	}

	return resources
}

func convertToTenants(azResources []models.AZResource) []models.Tenant {
	var tenants []models.Tenant

	for _, project := range azResources {
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
