package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sysdiglabs/promcat/pkg/resource"
	"github.com/sysdiglabs/promcat/pkg/usecases"

	"github.com/sysdiglabs/promcat/test/fixtures/resources"
)

var _ = Describe("RetrieveAllResourcesFromApp use case", func() {
	var useCase usecases.RetrieveAllResourcesFromApp

	BeforeEach(func() {
		useCase = usecases.RetrieveAllResourcesFromApp{
			ResourceRepository: newResourceRepositoryWithoutLambda(),
			AppRepository:      NewAppRepository(),
		}
	})

	It("returns all the avaliable resources for an App", func() {
		retrieved, _ := useCase.Execute("aws-fargate", "1.0.0")

		Expect(retrieved).To(Equal([]*resource.Resource{resources.AwsFargateDescription()}))
	})

	Context("when App does not exist", func() {
		It("returns App not found error", func() {
			retrieved, err := useCase.Execute("not-found", "1.0.0")

			Expect(retrieved).To(BeEmpty())
			Expect(err).To(HaveOccurred())
		})
	})

	Context("when vendor doesn't have resources", func() {
		It("returns an empty resource collection", func() {
			retrieved, err := useCase.Execute("lambda", "1.0.0")

			Expect(retrieved).To(BeEmpty())
			Expect(err).To(HaveOccurred())
		})
	})
})

func newResourceRepositoryWithoutLambda() resource.Repository {
	return resource.NewMemoryRepository(
		[]*resource.Resource{resources.AwsFargateDescription()},
	)
}
