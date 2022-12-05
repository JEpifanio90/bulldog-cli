// Package commands
/*
Copyright © 2022 Jose Epifanio jose.epifanio90@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package commands

import (
	"log"

	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"

	"github.com/JEpifanio90/bulldog-cli/internal/aws"
	"github.com/JEpifanio90/bulldog-cli/internal/gcp"
	"github.com/JEpifanio90/bulldog-cli/internal/savant"
	"github.com/JEpifanio90/bulldog-cli/models"
)

var tenants []models.Tenant
var List = cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "",
	Action: func(context *cli.Context) error {
		pterm.DefaultSpinner.Start("Fetching all of your resources from all the platforms...")
		awsResources, err := aws.FetchResources()
		if err != nil {
			pterm.Warning.Println("It looks like you don't have the AWS CLI installed! Skipping it...")
		}
		gcpProjects, err := gcp.FetchResources()
		if err != nil {
			pterm.Warning.Println("It looks like you don't have the GCP CLI installed! Skipping it...")
		}
		cloudOutput(&awsResources, &gcpProjects)
		pterm.DefaultSpinner.Stop()
		printer()
		return nil
	},
}

func cloudOutput(awsResources *[]aws.Resource, gcpProjects *[]gcp.Project) {
	if awsResources != nil && len(*awsResources) > 0 {
		for _, value := range *awsResources {
			arn, _ := savant.Parse(value.ResourceARN)
			tenants = append(
				tenants,
				models.Tenant{
					AccountID: arn.AccountID,
					Platform:  arn.Partition,
					Name:      arn.Resource,
					Type:      arn.Service,
					Region:    arn.Region,
					Tags:      value.Tags,
				},
			)
		}
	}

	if gcpProjects != nil && len(*gcpProjects) > 0 {
		for _, project := range *gcpProjects {
			tenants = append(
				tenants,
				models.Tenant{
					AccountID: project.ProjectId,
					Platform:  "gcp",
					Name:      project.Name,
					Type:      "-",
					Region:    "-",
					Tags:      nil,
				},
			)
		}
	}
}

func printer() {
	tableData := pterm.TableData{{"Account ID", "Platform", "Name", "Type", "Region", "Tags"}}

	for _, tenant := range tenants {
		tableData = append(tableData, []string{tenant.AccountID, tenant.Platform, tenant.Name, tenant.Type, tenant.Region})
	}

	err := pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()

	if err != nil {
		log.Fatalln(err)
	}
}
