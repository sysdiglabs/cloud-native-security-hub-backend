package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sysdiglabs/promcat/test/fixtures/resources"

	"github.com/sysdiglabs/promcat/pkg/resource"
	"github.com/sysdiglabs/promcat/pkg/usecases"
)

var _ = Describe("RetrieveOneResource use case", func() {
	var useCase usecases.RetrieveOneResource

	BeforeEach(func() {
		useCase = usecases.RetrieveOneResource{
			ResourceRepository: NewResourceRepository(),
		}
	})

	It("returns one resource", func() {
		result, _ := useCase.Execute("AWS Fargate", "Description", "1.0.0")

		Expect(result).To(Equal(resources.AwsFargateDescription()))
	})

	Context("when resource does not exist", func() {
		It("returns resource not found error", func() {
			retrieved, err := useCase.Execute("notFound", "Description", "1.0.0")

			Expect(retrieved).To(BeNil())
			Expect(err).To(MatchError(resource.ErrResourceNotFound))
		})
	})
})
