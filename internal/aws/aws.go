// Package aws
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

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/pterm/pterm"
	"os/exec"
)

func FetchResources() []Resource {
	wrapper := ResourceWrapper{}

	awsCli := exec.Command("aws", "resourcegroupstaggingapi", "get-resources", "--no-paginate")

	output, err := awsCli.Output()

	if err != nil {
		pterm.Error.Println(fmt.Errorf("aws cli %v", err.Error()))
		return nil
	}

	err = json.Unmarshal(output, &wrapper)

	if err != nil {
		pterm.Error.Println(fmt.Errorf("aws cli unmarshal: %v", err))
		return nil
	}

	return wrapper.ResourceTagMappingList
}
