package infrastructure_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sysdiglabs/promcat/test/fixtures/apps"

	"github.com/sysdiglabs/promcat/pkg/app"
	"github.com/sysdiglabs/promcat/pkg/infrastructure"
)

var _ = Describe("App importation from YAML files", func() {
	It("walks a directory and extract resources", func() {
		path := "../../test/fixtures/apps"
		parsed, _ := infrastructure.GetAppsFromPath(path)

		Expect(parsed).To(Equal([]*app.App{
			apps.AwsFargate(),
			apps.AwsLambda(),
		}))
	})

	Context("when path doesn't exist", func() {
		It("returns an error", func() {
			nonExistentPath := "../foo"

			parsed, err := infrastructure.GetAppsFromPath(nonExistentPath)

			Expect(parsed).To(BeEmpty())
			Expect(err).To(HaveOccurred())
		})
	})
})
