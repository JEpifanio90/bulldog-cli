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

type ARN struct {
	ARN               string
	Partition         string
	Service           string
	Region            string
	AccountID         string
	ResourceType      string
	Resource          string
	ResourceDelimiter string
}

type Resource struct {
	ResourceARN string
	Tags        []map[string]string
}

type ResourceWrapper struct {
	PaginationToken        string
	ResourceTagMappingList []Resource
}
