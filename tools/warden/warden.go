package warden

import (
	"fmt"
	"github.com/JEpifanio90/bulldog-cli/internal/aws"
	"github.com/JEpifanio90/bulldog-cli/internal/azure"
	"github.com/JEpifanio90/bulldog-cli/internal/gcp"
	"github.com/JEpifanio90/bulldog-cli/internal/models"
	"github.com/JEpifanio90/bulldog-cli/internal/travis_ci"
	"github.com/pterm/pterm"
	"os/exec"
)

var availableCmds []models.Command
var dumb = map[string][]string{
	"cloud":     {"aws", "gcp", "az"},
	"pipelines": {"travis"},
}

func FetchResources(filter *string) []models.Tenant {
	var tenants []models.Tenant
	setup(filter)

	for _, cmd := range availableCmds {
		executioner(cmd, &tenants)
	}

	return tenants
}

func setup(filter *string) {
	rawCmds := map[string][]string{
		"aws":    {"resourcegroupstaggingapi", "get-resources", "--no-paginate"},
		"gcp":    {"projects", "list", "--format", "json"},
		"az":     {"resource", "list", "--output", "json"},
		"travis": {"accounts"},
	}

	if *filter == "cloud" {
		delete(rawCmds, "travis")
	}

	if *filter == "pipelines" {
		delete(rawCmds, "aws")
		delete(rawCmds, "gcp")
		delete(rawCmds, "az")
	}

	for cmd, args := range rawCmds {
		if commandExists(cmd) {
			command := models.Command{Name: cmd, Args: args}
			availableCmds = append(availableCmds, command)
		} else {
			pterm.Warning.Println(fmt.Errorf("woof! It looks like you don't have the %v cli installed. Skipping it", cmd))
		}
	}
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func executioner(cmdMeta models.Command, tenants *[]models.Tenant) {
	rawOutput, err := exec.Command(cmdMeta.Name, cmdMeta.Args...).CombinedOutput()
	if err != nil {
		pterm.Error.Println(fmt.Errorf("warden: %v cli %v", cmdMeta.Name, err.Error()))
		return
	}

	switch cmdMeta.Name {
	case "aws":
		*tenants = append(*tenants, aws.ConvertToTenants(rawOutput)...)
	case "gcp":
		*tenants = append(*tenants, gcp.ConvertToTenants(rawOutput)...)
	case "az":
		*tenants = append(*tenants, azure.ConvertToTenants(rawOutput)...)
	case "travis":
		*tenants = append(*tenants, travis_ci.ConvertToTenants(rawOutput)...)
	}
}
