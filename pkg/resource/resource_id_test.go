package resource_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sysdiglabs/prometheus-hub/pkg/resource"
)

var _ = Describe("ResourceID", func() {
	Context("when comparing ResourceID's", func() {
		It("considers the App of the resource", func() {
			one := resource.NewResourceID("aws-fargate", "Description", []string{"1.0.0"})
			other := resource.NewResourceID("aws-lambda", "Description", []string{"1.0.0"})

			Expect(one).NotTo(Equal(other))

		})

		It("considers the kind of the resource", func() {
			one := resource.NewResourceID("aws-fargate", "Description", []string{"1.0.0"})
			other := resource.NewResourceID("aws-fargate", "Alerts", []string{"1.0.0"})

			Expect(one).NotTo(Equal(other))
		})

		It("considers the version of the App", func() {
			one := resource.NewResourceID("aws-fargate", "Description", []string{"1.0.0"})
			other := resource.NewResourceID("aws-fargate", "Description", []string{"1.0.1"})

			Expect(one).NotTo(Equal(other))
		})
	})
})
