package resource

// Resource Any resource for an App
type Resource struct {
	ID                ResourceID
	Kind              string
	App               string
	Version           string
	AvailableVersions []string
	AppVersion        []string
	Maintainers       []*Maintainer
	Description       string
	Data              string
	Configurations    []*Configuration
}

// Maintainer Name and link of the maintainer of the Resource
type Maintainer struct {
	Name string
	Link string
}

// Configuration Metadata and data of the configuration.
// This is used for Resources of kind Dashboards and Alerts
type Configuration struct {
	Name        string
	Kind        string
	Image       string
	Description string
	File        string
	Data        string
}
