package aws

import (
	"encoding/json"
	"fmt"
	"github.com/JEpifanio90/bulldog-cli/internal/models"
	"github.com/JEpifanio90/bulldog-cli/tools/savant"
	"github.com/pterm/pterm"
)

func ConvertToTenants(rawOutput []byte) []models.Tenant {
	var tenants []models.Tenant

	for _, value := range parse(rawOutput) {
		arn, _ := savant.ParseARN(value.ResourceARN)
		tenants = append(
			tenants,
			models.Tenant{
				AccountID: arn.AccountID,
				Platform:  arn.Partition,
				Name:      arn.Resource,
				Type:      arn.Service,
				Region:    arn.Region,
				//Tags:      value.Tags,
			},
		)
	}

	return tenants
}

func parse(rawOutput []byte) []models.AWSResource {
	wrapper := models.AWSResourceWrapper{}
	err := json.Unmarshal(rawOutput, &wrapper)

	if err != nil {
		pterm.Error.Println(fmt.Errorf("aws cli unmarshal: %v", err))
		return nil
	}

	return wrapper.ResourceTagMappingList
}
