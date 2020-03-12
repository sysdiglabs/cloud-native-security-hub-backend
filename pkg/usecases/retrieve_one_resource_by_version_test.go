package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sysdiglabs/prometheus-hub/test/fixtures/resources"

	"github.com/sysdiglabs/prometheus-hub/pkg/resource"
	"github.com/sysdiglabs/prometheus-hub/pkg/usecases"
)

var _ = Describe("RetrieveOneResourceByVersion use case", func() {
	var useCase usecases.RetrieveOneResourceByVersion

	BeforeEach(func() {
		useCase = usecases.RetrieveOneResourceByVersion{
			ResourceRepository: newResourceRepositoryWithVersions(),
		}
	})

	It("returns one resource", func() {
		result, _ := useCase.Execute("AWS Fargate", "Description", "1.0.0", "1.0.1")

		apacheWithSpecificVersion := resources.AwsFargateDescription()
		apacheWithSpecificVersion.Version = "1.0.1"
		Expect(result).To(Equal(apacheWithSpecificVersion))
	})

	Context("when version does not exist", func() {
		It("returns an error", func() {
			result, err := useCase.Execute("aws-fargate", "Description", "1.0.0.", "2.0.0")

			Expect(result).To(BeNil())
			Expect(err).To(MatchError(resource.ErrResourceNotFound))
		})
	})
})

func newResourceRepositoryWithVersions() resource.Repository {
	fargate := resources.AwsFargateDescription()
	fargate.Version = "1.0.1"

	return resource.NewMemoryRepository(
		[]*resource.Resource{
			resources.AwsFargateDescription(),
			fargate,
		},
	)
}
