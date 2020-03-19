package usecases_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sysdiglabs/promcat/pkg/app"
	"github.com/sysdiglabs/promcat/pkg/resource"
	"github.com/sysdiglabs/promcat/test/fixtures/apps"
	"github.com/sysdiglabs/promcat/test/fixtures/resources"
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
