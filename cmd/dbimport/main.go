package main

import (
	"log"
	"os"

	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/sysdiglabs/prometheus-hub/pkg/app"
	"github.com/sysdiglabs/prometheus-hub/pkg/infrastructure"
	"github.com/sysdiglabs/prometheus-hub/pkg/resource"
)

func main() {
	log.Println("Starting database importing job")

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	migrateDatabase(db)
	importResources(db)
	importApps(db)
}

func migrateDatabase(db *sql.DB) {
	log.Println("Applying migrations")

	config := &postgres.Config{}
	driver, err := postgres.WithInstance(db, config)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: Change this for production
	migrator, err := migrate.NewWithDatabaseInstance("file://db/migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	err = migrator.Up()
	if err != nil {
		if err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

}

func importResources(db *sql.DB) {
	log.Println("Importing resources")

	resources, err := infrastructure.GetResourcesFromPath(os.Getenv("RESOURCES_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	repository := resource.NewPostgresRepository(db)

	for _, resource := range resources {
		err = repository.Save(resource)
		if err != nil {
			log.Println(err)
		}
	}
}

func importApps(db *sql.DB) {
	log.Println("Importing apps")

	apps, err := infrastructure.GetAppsFromPath(os.Getenv("APPS_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	repository := app.NewPostgresRepository(db)
	for _, app := range apps {
		err = repository.Save(app)
		if err != nil {
			log.Println(err)
		}
	}
}
