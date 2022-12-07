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
