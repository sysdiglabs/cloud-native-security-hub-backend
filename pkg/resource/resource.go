package resource

type Resource struct {
	ID                ResourceID
	Kind              string
	App               string
	Version           string
	AvailableVersions []string
	AppVersion        []string
	Maintainers       []*Maintainer
	Data              string
}

type Maintainer struct {
	Name string
	Link string
}
