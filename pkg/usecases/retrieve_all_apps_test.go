package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sysdiglabs/prometheus-hub/pkg/app"
	"github.com/sysdiglabs/prometheus-hub/pkg/usecases"
	"github.com/sysdiglabs/prometheus-hub/test/fixtures/apps"
)

var _ = Describe("RetrieveAllApps use case", func() {
	It("returns all the available apps", func() {
		existingApps := []*app.App{apps.AwsFargate()}
		useCase := usecases.RetrieveAllApps{AppRepository: NewAppRepository()}

		retrieved, _ := useCase.Execute()

		Expect(retrieved).To(Equal(existingApps))
	})
})
