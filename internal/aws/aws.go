package aws

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/aws/aws-sdk-go-v2/config"
	resourceGroups "github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi/types"
	"github.com/pterm/pterm"

	"github.com/JEpifanio90/bulldog-cli/internal/models"
	"github.com/JEpifanio90/bulldog-cli/internal/savant"
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
	rawOutput, err := exec.Command("aws", []string{"resourcegroupstaggingapi", "get-resources", "--no-paginate"}...).CombinedOutput()
	if err != nil {
		pterm.Error.Println(fmt.Errorf("list command: aws cli %v", err.Error()))
		return nil
	}

	return parseAndConvert(rawOutput, true)
}

func useCdk() []models.Tenant {
	// Load the Shared AWS Configuration (~/.aws/config)
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	svc := resourceGroups.NewFromConfig(cfg)
	rawResp, err := svc.GetResources(context.TODO(), nil)
	if err != nil {
		pterm.Warning.Println("woof! It looks like you don't have any aws accounts! Skipping it...")
		return nil
	}

	return parseAndConvert(rawResp.ResourceTagMappingList, false)
}

func parseAndConvert[T []byte | []types.ResourceTagMapping](input T, unmarshal bool) []models.Tenant {
	wrapper := resourceGroups.GetResourcesOutput{}

	// TODO: Check else condition
	if _, ok := any(input).([]byte); unmarshal && ok {
		err := json.Unmarshal(any(input).([]byte), &wrapper)
		if err != nil {
			pterm.Error.Println(fmt.Errorf("aws cli unmarshal: %w", err))
			return nil
		}
	}

	var tenants []models.Tenant

	for _, value := range wrapper.ResourceTagMappingList {
		arn, _ := savant.ParseARN(*value.ResourceARN)
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
