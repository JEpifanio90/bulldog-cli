// Package gcp
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

package gcp

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func FetchResources() ([]Project, error) {
	var projects []Project

	// gcloud projects describe project ID
	gcpCli := exec.Command("gcloud", "projects", "list", "--format", "json")

	output, err := gcpCli.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("gcp cli cmd: %v", string(exitErr.Stderr))
		}
	}

	err = json.Unmarshal(output, &projects)
	if err != nil {
		return nil, fmt.Errorf("aws cli unmarshal: %v", err)
	}

	return projects, nil
}
