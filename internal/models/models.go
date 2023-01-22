package models

type Tenant struct {
	AccountID string
	Platform  string
	Name      string
	Type      string
	Region    string
	Tags      string
}

type Command struct {
	Name string
	CLI  bool
}

type AWSArn struct {
	ARN               string
	Partition         string
	Service           string
	Region            string
	AccountID         string
	ResourceType      string
	Resource          string
	ResourceDelimiter string
}

type AZResponse struct {
	NextLink string
	Value    []AZResource
}

type AZResource struct {
	ID          string
	Kind        string
	Location    string
	Name        string
	CreatedTime string
	Type        string
	Tags        map[string]string
}
