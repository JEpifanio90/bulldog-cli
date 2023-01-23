package travis_ci

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/pterm/pterm"

	"github.com/JEpifanio90/bulldog-cli/internal/models"
)

func FetchResources(cmd models.Command) []models.Tenant {
	var tenants []models.Tenant

	if cmd.CLI {
		tenants = useCli()
	}

	return tenants
}

func useCli() []models.Tenant {
	rawOutput, err := exec.Command("travis", []string{"accounts"}...).CombinedOutput()
	if err != nil {
		pterm.Error.Println(fmt.Errorf("list command: az cli %v", err.Error()))
		return nil
	}
	var tenants []models.Tenant

	for _, travisPipe := range parse(rawOutput) {
		tenants = append(
			tenants,
			models.Tenant{
				AccountID: "-",
				Platform:  "Travis",
				Name:      travisPipe,
				Type:      "-",
				Region:    "-",
				Tags:      "-",
			},
		)
	}

	return tenants
}

func parse(rawOutput []byte) []string {
	var accounts []string
	err := json.Unmarshal(rawOutput, &accounts)

	if err != nil {
		pterm.Error.Println(fmt.Errorf("travis cli unmarshal: %v", err))
		return nil
	}

	return accounts
}
