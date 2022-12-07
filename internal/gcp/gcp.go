package gcp

import (
	"encoding/json"
	"fmt"
	"github.com/pterm/pterm"
	"os/exec"
)

func FetchResources() []Project {
	var projects []Project

	// gcloud projects describe project ID
	gcpCli := exec.Command("gcloud", "projects", "list", "--format", "json")

	output, err := gcpCli.Output()

	if err != nil {
		pterm.Error.Println(fmt.Errorf("gcp cli %v", err.Error()))
		return nil
	}

	err = json.Unmarshal(output, &projects)

	if err != nil {
		pterm.Error.Println(fmt.Errorf("gcp cli unmarshal: %v", err))
		return nil
	}

	return projects
}
