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

type AZResource struct {
	ID          string
	Kind        string
	Location    string
	Name        string
	CreatedTime string
	Type        string
	Tags        map[string]string
}

type GCPProject struct {
	CreateTime     string
	LifecycleState string
	Name           string
	ProjectId      string
	ProjectNumber  string
	Labels         map[string]string
}
