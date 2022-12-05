package azure

type Resource struct {
	ID          string
	Kind        string
	Location    string
	Name        string
	CreatedTime string
	Type        string
	Tags        map[string]string
}
