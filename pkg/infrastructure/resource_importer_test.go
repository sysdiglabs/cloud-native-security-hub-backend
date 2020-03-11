package infrastructure_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sysdiglabs/prometheus-hub/pkg/infrastructure"
	"github.com/sysdiglabs/prometheus-hub/pkg/resource"
	"github.com/sysdiglabs/prometheus-hub/test/fixtures/resources"
)

var _ = Describe("Resource importation from YAML files", func() {
	It("walks a directory and extract resources", func() {
		path := "../../test/fixtures/resources"
		parsed, _ := infrastructure.GetResourcesFromPath(path)

		Expect(parsed).To(Equal([]*resource.Resource{
			resources.AwsFargateAlertsWithoutAvailableVersions(),
			resources.AwsFargateDescriptionWithoutAvailableVersions(),
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
