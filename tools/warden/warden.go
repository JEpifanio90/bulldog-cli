package warden

import (
	"fmt"
	"github.com/JEpifanio90/bulldog-cli/internal/aws"
	"github.com/JEpifanio90/bulldog-cli/internal/azure"
	"github.com/JEpifanio90/bulldog-cli/internal/gcp"
	"github.com/JEpifanio90/bulldog-cli/internal/models"
	"github.com/pterm/pterm"
	"os/exec"
)

var availableCmds []models.Command

func init() {
	rawCmds := map[string][]string{
		"aws": {"resourcegroupstaggingapi", "get-resources", "--no-paginate"},
		"gcp": {"projects", "list", "--format", "json"},
		"az":  {"resource", "list", "--output", "json"},
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

func FetchResources() []models.Tenant {
	var tenants []models.Tenant

	for _, cmd := range availableCmds {
		executioner(cmd, &tenants)
	}

	return tenants
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func executioner(cmdMeta models.Command, tenants *[]models.Tenant) {
	cmd := exec.Command(cmdMeta.Name, cmdMeta.Args...)
	rawOutput, err := cmd.Output()

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
	}
}
