package app_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sysdiglabs/promcat/pkg/app"
	"github.com/sysdiglabs/promcat/test/fixtures/apps"

	"database/sql"
	"os"
)

var _ = Describe("Postgres App Repository", func() {
	var repository app.Repository

	BeforeEach(func() {
		db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
		repository = app.NewPostgresRepository(db)

		db.Exec("TRUNCATE TABLE apps")
	})

	It("saves a new app", func() {
		repository.Save(apps.AwsFargate())

		retrieved, _ := repository.FindById("aws-fargate")

		Expect(retrieved).To(Equal(apps.AwsFargate()))
	})

	It("retrieves all existent apps", func() {
		repository.Save(apps.AwsFargate())

		retrieved, _ := repository.FindAll()

		Expect(retrieved).To(Equal([]*app.App{apps.AwsFargate()}))
	})

	Context("when querying by id", func() {
		Context("and app does not exist", func() {
			It("returns an error", func() {
				retrieved, err := repository.FindById("non existent id")

				Expect(retrieved).To(BeNil())
				Expect(err).To(MatchError(app.ErrAppNotFound))
			})
		})
	})
})
