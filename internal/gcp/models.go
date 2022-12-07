package gcp

type Project struct {
	CreateTime     string
	LifecycleState string
	Name           string
	ProjectId      string
	ProjectNumber  string
	Labels         map[string]string
}
