package aws

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
		tenants = useAWSCli()
	} else {
		tenants = useCdk()
	}

	return tenants
}

func useAWSCli() []models.Tenant {
	rawOutput, err := exec.Command("aws", []string{"resourcegroupstaggingapi", "get-resources", "--no-paginate"}...).CombinedOutput()
	if err != nil {
		pterm.Error.Println(fmt.Errorf("list command: aws cli %v", err.Error()))
		return nil
	}
	var tenants []models.Tenant

	for _, value := range parse(rawOutput) {
		arn, _ := savant.ParseARN(value.ResourceARN)
		tags, _ := json.Marshal(value.Tags)

		tenants = append(
			tenants,
			models.Tenant{
				AccountID: arn.AccountID,
				Platform:  arn.Partition,
				Name:      arn.Resource,
				Type:      arn.Service,
				Region:    arn.Region,
				Tags:      string(tags),
			},
		)
	}

	return tenants
}

func useCdk() []models.Tenant {
	//// Load the Shared AWS Configuration (~/.aws/config)
	//ctx := context.TODO()
	//cfg, err := config.LoadDefaultConfig(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//resources, err := resourcegroupstaggingapi.Client.GetResources(ctx)
	//if err != nil {
	//	return
	//}
	//
	//pterm.Info.Println(resources)
	return make([]models.Tenant, 0)
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
