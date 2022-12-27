package travis_ci

import (
	"encoding/json"
	"fmt"

	"github.com/JEpifanio90/bulldog-cli/internal/models"
	"github.com/pterm/pterm"
)

func ConvertToTenants(rawOutput []byte) []models.Tenant {
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
