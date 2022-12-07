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
