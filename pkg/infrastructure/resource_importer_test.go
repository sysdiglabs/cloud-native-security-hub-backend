package infrastructure_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sysdiglabs/promcat/pkg/infrastructure"
	"github.com/sysdiglabs/promcat/pkg/resource"
	"github.com/sysdiglabs/promcat/test/fixtures/resources"
)

var _ = Describe("Resource importation from YAML files", func() {
	It("walks a directory and extract resources", func() {
		path := "../../test/fixtures/resources"
		parsed, _ := infrastructure.GetResourcesFromPath(path)

		Expect(parsed).To(Equal([]*resource.Resource{
			resources.AwsFargateAlertsWithoutAvailableVersions(),
			resources.AwsFargateDashboardsWithoutAvailableVersions(),
			resources.AwsFargateDescriptionWithoutAvailableVersions(),
			resources.AwsLambdaDashboardsWithoutAvailableVersions(),
		}))
	})

	Context("when path doesn't exist", func() {
		It("returns an error", func() {
			nonExistentPath := "../foo"

			parsed, err := infrastructure.GetResourcesFromPath(nonExistentPath)

			Expect(parsed).To(BeEmpty())
			Expect(err).To(HaveOccurred())
		})
	})
})
