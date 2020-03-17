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
	Dashboards        []*Dashboard
	Alerts            Alerts
}

type Maintainer struct {
	Name string
	Link string
}

type Alerts struct {
	PrometheusAlerts string
	SysdigAlerts     string
}

type Dashboard struct {
	Name        string
	Kind        string
	Image       string
	Description string
	Data        string
}
