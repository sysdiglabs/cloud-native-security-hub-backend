package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sysdiglabs/promcat/test/fixtures/apps"

	"github.com/sysdiglabs/promcat/pkg/app"
	"github.com/sysdiglabs/promcat/pkg/usecases"
)

var _ = Describe("RetrieveOneApp use case", func() {
	var useCase usecases.RetrieveOneApp

	BeforeEach(func() {
		useCase = usecases.RetrieveOneApp{
			AppRepository: NewAppRepository(),
		}
	})

	It("returns one app", func() {
		result, _ := useCase.Execute("aws-fargate")

		Expect(result).To(Equal(apps.AwsFargate()))
	})

	Context("when app does not exist", func() {
		It("returns app not found error", func() {
			retrieved, err := useCase.Execute("nonExistent")

			Expect(retrieved).To(BeNil())
			Expect(err).To(MatchError(app.ErrAppNotFound))
		})
	})
})
