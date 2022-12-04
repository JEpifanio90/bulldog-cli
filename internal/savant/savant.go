// Package savant
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

package savant

import (
	"errors"
	"github.com/JEpifanio90/bulldog-cli/internal/aws"
	"strings"
)

var (
	// ErrMalformed is returned when the ARN appears to be invalid.
	ErrMalformed = errors.New("malformed ARN")

	// ErrVariablesNotSupported is returned when the ARN contains policy
	// variables.
	ErrVariablesNotSupported = errors.New("policy variables are not supported")
)

// Parse accepts and ARN string and attempts to break it into component parts.
func Parse(arn string) (*aws.ARN, error) {
	pieces := strings.SplitN(arn, ":", 6)

	if err := validate(arn, pieces); err != nil {
		return nil, err
	}

	components := &aws.ARN{
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

func validate(arn string, pieces []string) error {
	if strings.Contains(arn, "${") {
		return ErrVariablesNotSupported
	}
	if len(pieces) < 6 {
		return ErrMalformed
	}

	return nil
}
