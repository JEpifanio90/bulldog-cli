package savant

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/JEpifanio90/bulldog-cli/internal/models"
)

func TestArnParse(t *testing.T) {
	arn, _ := ParseARN("arn:partition:service:region:account-id:resource-id")
	output := &models.AWSArn{
		ARN:               "arn",
		Partition:         "partition",
		Service:           "service",
		Region:            "region",
		AccountID:         "account-id",
		Resource:          "resource-id",
		ResourceType:      "",
		ResourceDelimiter: "",
	}

	fmt.Println(arn, output)
	if !reflect.DeepEqual(arn, output) {
		t.Error("ARN should be equal")
	}
}
