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
		Data: "# AWS Fargate\nAlerts.",
	}
}

func AwsLambdaeDescription() *resource.Resource {
	result := AwsLambdaDescriptionWithoutAvailableVersions()
	result.AvailableVersions = []string{"1.0.0"}

	return result
}

func AwsLambdaDescriptionWithoutAvailableVersions() *resource.Resource {
	return &resource.Resource{
		ID: resource.NewResourceID("AWS Lambda",
			"Description",
			[]string{"1.0.0", "1.0.1"}),
		Kind:       "Description",
		App:        "AWS Lambda",
		Version:    "1.0.0",
		AppVersion: []string{"1.0.0", "1.0.1"},
		Maintainers: []*resource.Maintainer{
			{
				Name: "sysdiglabs",
				Link: "github.com/sysdiglabs",
			},
		},
		Data: "# AWS Lambda\nDescription.",
	}
}
