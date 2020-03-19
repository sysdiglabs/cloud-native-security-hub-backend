package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sysdiglabs/promcat/test/fixtures/resources"

	"github.com/sysdiglabs/promcat/pkg/resource"
	"github.com/sysdiglabs/promcat/pkg/usecases"
)

var _ = Describe("RetrieveAllResources use case", func() {
	It("returns all the available resources", func() {
		existingResources := []*resource.Resource{resources.AwsFargateDescription(), resources.AwsFargateAlerts()}
		useCase := usecases.RetrieveAllResources{ResourceRepository: NewResourceRepository()}

		retrieved, _ := useCase.Execute()

		Expect(retrieved).To(Equal(existingResources))
	})
})
