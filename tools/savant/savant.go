package savant

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/JEpifanio90/bulldog-cli/internal/models"
)

var (
	ErrMalformed             = errors.New("malformed ARN")
	ErrVariablesNotSupported = errors.New("policy variables are not supported")
)

func ParseARN(arn string) (*models.AWSArn, error) {
	pieces := strings.SplitN(arn, ":", 6)

	if err := validate(arn, pieces); err != nil {
		return nil, err
	}

	components := &models.AWSArn{
		ARN:       pieces[0],
		Partition: pieces[1],
		Service:   pieces[2],
		Region:    pieces[3],
		AccountID: pieces[4],
	}

	if n := strings.Count(pieces[5], ":"); n > 0 {
		components.ResourceDelimiter = ":"
		resourceParts := strings.SplitN(pieces[5], ":", 2)
		components.ResourceType = resourceParts[0]
		components.Resource = resourceParts[1]
	} else {
		if m := strings.Count(pieces[5], "/"); m == 0 {
			components.Resource = pieces[5]
		} else {
			components.ResourceDelimiter = "/"
			resourceParts := strings.SplitN(pieces[5], "/", 2)
			components.ResourceType = resourceParts[0]
			components.Resource = resourceParts[1]
		}
	}
	return components, nil
}

func ParseAZ(resourceID string) (string, error) {
	const patternText = `(?i)subscriptions/(.+)/resourceGroups/(.+)/providers/(.+?)/(.+?)/(.+)`
	resourcePattern := regexp.MustCompile(patternText)
	match := resourcePattern.FindStringSubmatch(resourceID)

	if len(match) == 0 {
		return "", fmt.Errorf("savant: parsing failed for resource ID %s", resourceID)
	}

	v := strings.Split(match[5], "/")
	return v[len(v)-1], nil
}

func validate(arn string, pieces []string) error {
	if strings.Contains(arn, "${") {
		return ErrVariablesNotSupported
	}
	if len(pieces) < 6 {
		return ErrMalformed
	}

	return nil
}
