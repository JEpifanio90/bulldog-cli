package azure

import (
	"encoding/json"
	"fmt"
	"github.com/pterm/pterm"
	"os/exec"
)

func FetchResources() []Resource {
	var resources []Resource

	azureCli := exec.Command("az", "resource", "list", "-o", "json")

	output, err := azureCli.Output()

	if err != nil {
		pterm.Error.Println(fmt.Errorf("az cli %v", err.Error()))
		return nil
	}

	err = json.Unmarshal(output, &resources)

	if err != nil {
		pterm.Error.Println(fmt.Errorf("az cli unmarshal: %v", err))
		return nil
	}

	return resources
}
