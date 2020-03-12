package usecases_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sysdiglabs/prometheus-hub/pkg/app"
	"github.com/sysdiglabs/prometheus-hub/pkg/resource"
	"github.com/sysdiglabs/prometheus-hub/test/fixtures/apps"
	"github.com/sysdiglabs/prometheus-hub/test/fixtures/resources"
)

func TestUsecases(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Usecases Suite")
}

func NewResourceRepository() resource.Repository {
	return resource.NewMemoryRepository(
		[]*resource.Resource{resources.AwsFargateDescription(), resources.AwsFargateAlerts()},
	)
}

func NewAppRepository() app.Repository {
	return app.NewMemoryRepository(
		[]*app.App{apps.AwsFargate()},
	)
}
