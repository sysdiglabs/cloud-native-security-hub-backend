package resource_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sysdiglabs/prometheus-hub/test/fixtures/resources"

	"database/sql"
	"os"

	"github.com/sysdiglabs/prometheus-hub/pkg/resource"
)

var _ = Describe("Postgres Resource Repository", func() {
	var repository resource.Repository

	BeforeEach(func() {
		db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
		repository = resource.NewPostgresRepository(db)

		db.Exec("TRUNCATE TABLE resources")
		db.Exec("TRUNCATE TABLE latest_resources")
	})

	It("saves a new resource", func() {
		repository.Save(resources.AwsFargateDescription())

		retrieved, _ := repository.FindById(resource.NewResourceID("aws-fargate",
			"Description",
			[]string{"1.0.0", "1.0.1"}))
		Expect(retrieved).To(Equal(resources.AwsFargateDescription()))
	})

	Context("when saving a resource", func() {
		Context("and there is already a resource with for this AppVersion", func() {
			It("returns an error", func() {
				repository.Save(resources.AwsFargateDescription())
				err := repository.Save(resources.AwsFargateDescription())
				Expect(err).To(HaveOccurred())
			})
		})
	})

	It("retrieves all existent resources", func() {
		repository.Save(resources.AwsFargateDescription())
		repository.Save(resources.AwsFargateAlerts())

		retrieved, _ := repository.FindAll()

		Expect(retrieved).To(Equal([]*resource.Resource{
			resources.AwsFargateDescription(),
			resources.AwsFargateAlerts()}))
	})

	Context("when querying by id", func() {
		Context("and resource is not found", func() {
			It("returns an error", func() {
				retrieved, err := repository.FindById(resource.NewResourceID("non-existent-app", "non-existent-kind", []string{"0.0.0"}))

				Expect(retrieved).To(BeNil())
				Expect(err).To(HaveOccurred())
			})
		})

		It("returns latest version of the resource", func() {
			fargateDescription := resources.AwsFargateDescription()
			repository.Save(fargateDescription)

			fargateDescription.Version = "2.0.0"
			repository.Save(fargateDescription)

			retrieved, _ := repository.FindById(resource.NewResourceID("aws-fargate", "Description", []string{"1.0.0"}))

			expected := resources.AwsFargateDescription()
			expected.Version = "2.0.0"
			expected.AvailableVersions = []string{"2.0.0", "1.0.0"}
			Expect(retrieved).To(Equal(expected))
		})

		Context("and version is specified as well", func() {
			It("returns the resource with the specified version", func() {
				fargateDescription := resources.AwsFargateDescription()
				repository.Save(fargateDescription)

				fargateDescription.Version = "2.0.0"
				repository.Save(fargateDescription)

				retrieved, _ := repository.FindByVersion(resource.NewResourceID("aws-fargate", "Description", []string{"1.0.0"}), "1.0.0")

				expected := resources.AwsFargateDescription()
				expected.AvailableVersions = []string{"2.0.0", "1.0.0"}
				Expect(retrieved).To(Equal(expected))
			})
		})
	})

	Context("when saving several versions for a resource", func() {
		It("returns all available versions, newer first", func() {
			fargateDescription := resources.AwsFargateDescription()
			repository.Save(fargateDescription)

			fargateDescription.Version = "2.0.0"
			repository.Save(fargateDescription)

			retrieved, _ := repository.FindById(resource.NewResourceID("aws-fargate", "Description", []string{"1.0.0"}))

			Expect(retrieved.AvailableVersions).To(Equal([]string{"2.0.0", "1.0.0"}))
		})
	})
})
