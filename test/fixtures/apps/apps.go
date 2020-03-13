package apps

import "github.com/sysdiglabs/prometheus-hub/pkg/app"

func AwsFargate() *app.App {
	return &app.App{
		ID:                "aws-fargate",
		Kind:              "App",
		Name:              "AWS Fargate",
		Keywords:          []string{"aws", "serverless", "containers"},
		AvailableVersions: []string{"1.0.0", "1.0.1"},
		Description:       "# AWS Fargate\n",
		ShortDescription:  "AWS serverless containers",
		Icon:              "https://upload.wikimedia.org/wikipedia/commons/thumb/9/93/Amazon_Web_Services_Logo.svg/500px-Amazon_Web_Services_Logo.svg.png",
		Website:           "https://aws.amazon.com/fargate/",
		Available:         true,
	}
}

func AwsLambda() *app.App {
	return &app.App{
		ID:                "aws-lambda",
		Kind:              "App",
		Name:              "AWS Lambda",
		Keywords:          []string{"aws", "serverless", "containers"},
		AvailableVersions: []string{"1.0.0", "1.0.1"},
		Description:       "# AWS Lambda\n",
		ShortDescription:  "AWS serverless functions",
		Icon:              "https://upload.wikimedia.org/wikipedia/commons/thumb/9/93/Amazon_Web_Services_Logo.svg/500px-Amazon_Web_Services_Logo.svg.png",
		Website:           "https://aws.amazon.com/lambda/",
		Available:         true,
	}
}
