package resources

import (
	"github.com/sysdiglabs/prometheus-hub/pkg/resource"
)

func AwsFargateDescription() *resource.Resource {
	result := AwsFargateDescriptionWithoutAvailableVersions()
	result.AvailableVersions = []string{"1.0.0"}

	return result
}

func AwsFargateDescriptionWithoutAvailableVersions() *resource.Resource {
	return &resource.Resource{
		ID: resource.NewResourceID("AWS Fargate",
			"Description",
			[]string{"1.0.0", "1.0.1"}),
		Kind:       "Description",
		App:        "AWS Fargate",
		Version:    "1.0.0",
		AppVersion: []string{"1.0.0", "1.0.1"},
		Maintainers: []*resource.Maintainer{
			{
				Name: "sysdiglabs",
				Link: "github.com/sysdiglabs",
			},
		},
		Data: "# AWS Fargate\nDescription.",
	}
}

func AwsFargateAlerts() *resource.Resource {
	result := AwsFargateAlertsWithoutAvailableVersions()
	result.AvailableVersions = []string{"1.0.0"}

	return result
}

func AwsFargateAlertsWithoutAvailableVersions() *resource.Resource {
	return &resource.Resource{
		ID: resource.NewResourceID("AWS Fargate",
			"Alerts",
			[]string{"1.0.0"}),
		Kind:       "Alerts",
		App:        "AWS Fargate",
		Version:    "1.0.0",
		AppVersion: []string{"1.0.0", "1.0.1"},
		Maintainers: []*resource.Maintainer{
			{
				Name: "sysdiglabs",
				Link: "github.com/sysdiglabs",
			},
		},
		Description: "Description of the alerts",
		Alerts: &resource.Alerts{
			PrometheusAlerts: "Prometheus Alert",
			SysdigAlerts:     "Sysdig Alert",
		},
	}
}

func AwsFargateDashboards() *resource.Resource {
	result := AwsFargateDashboardsWithoutAvailableVersions()
	result.AvailableVersions = []string{"1.0.0"}

	return result
}

func AwsFargateDashboardsWithoutAvailableVersions() *resource.Resource {
	return &resource.Resource{
		ID: resource.NewResourceID("AWS Fargate",
			"Dashboards",
			[]string{"1.0.0"}),
		Kind:       "Dashboards",
		App:        "AWS Fargate",
		Version:    "1.0.0",
		AppVersion: []string{"1.0.0", "1.0.1"},
		Maintainers: []*resource.Maintainer{
			{
				Name: "sysdiglabs",
				Link: "github.com/sysdiglabs",
			},
		},
		Description: "Description of the alerts",
		Dashboards: []*resource.Dashboard{
			{
				Name:        "Grafana Dashboard",
				Kind:        "Grafana",
				Image:       "url-of-grafana-image.png",
				Description: "Description of the Grafana dashboard",
				Data:        "{}",
			},
			{
				Name:        "Sysdig Dashboard",
				Kind:        "Sysdig",
				Image:       "url-of-sysdig-image.png",
				Description: "Description of the Sysdig dashboard",
				Data:        "{}",
			},
		},
	}
}
