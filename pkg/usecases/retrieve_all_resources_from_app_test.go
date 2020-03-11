package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/resources"
)

var _ = Describe("RetrieveAllResourcesFromApp use case", func() {
	var useCase usecases.RetrieveAllResourcesFromApp

	BeforeEach(func() {
		useCase = usecases.RetrieveAllResourcesFromApp{
			ResourceRepository: newResourceRepositoryWithoutMongoDB(),
			VendorRepository:   NewVendorRepository(),
		}
	})

	It("returns all the avaliable resources for a vendor", func() {
		retrieved, _ := useCase.Execute("apache")

		Expect(retrieved).To(Equal([]*resource.Resource{resources.Apache()}))
	})

	Context("when vendor does not exist", func() {
		It("returns vendor not found error", func() {
			retrieved, err := useCase.Execute("not-found")

			Expect(retrieved).To(BeEmpty())
			Expect(err).To(HaveOccurred())
		})
	})

	PContext("when vendor doesn't have resources", func() {
		It("returns an empty resource collection", func() {
			retrieved, err := useCase.Execute("mongo")

			Expect(retrieved).To(BeEmpty())
			Expect(err).To(Succeed())
		})
	})
})

func newResourceRepositoryWithoutMongoDB() resource.Repository {
	return resource.NewMemoryRepository(
		[]*resource.Resource{resources.Apache()},
	)
}
