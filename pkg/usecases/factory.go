package usecases

import (
	"database/sql"
	"log"
	"os"

	"github.com/sysdiglabs/prometheus-hub/pkg/app"
	"github.com/sysdiglabs/prometheus-hub/pkg/resource"
)

type Factory interface {
	NewRetrieveAllResourcesUseCase() *RetrieveAllResources
	NewRetrieveOneResourceUseCase() *RetrieveOneResource
	NewRetrieveOneResourceByVersionUseCase() *RetrieveOneResourceByVersion

	NewRetrieveAllAppsUseCase() *RetrieveAllApps

	NewRetrieveOneAppUseCase() *RetrieveOneApp
	NewRetrieveAllResourcesFromAppUseCase() *RetrieveAllResourcesFromApp

	NewResourcesRepository() resource.Repository
	NewAppRepository() app.Repository
}

func NewFactory() Factory {
	factory := &factory{}
	factory.db = factory.newDB()
	factory.resourceRepository = factory.NewResourcesRepository()
	factory.appRepository = factory.NewAppRepository()
	return factory
}

// TODO: Instantiate useCases only once
type factory struct {
	db                 *sql.DB
	appRepository      app.Repository
	resourceRepository resource.Repository
}

func (f *factory) NewRetrieveAllResourcesUseCase() *RetrieveAllResources {
	return &RetrieveAllResources{ResourceRepository: f.resourceRepository}
}

func (f *factory) NewRetrieveOneResourceUseCase() *RetrieveOneResource {
	return &RetrieveOneResource{ResourceRepository: f.resourceRepository}
}

func (f *factory) NewRetrieveOneResourceByVersionUseCase() *RetrieveOneResourceByVersion {
	return &RetrieveOneResourceByVersion{ResourceRepository: f.resourceRepository}
}

func (f *factory) NewRetrieveAllAppsUseCase() *RetrieveAllApps {
	return &RetrieveAllApps{
		AppRepository: f.appRepository,
	}
}

func (f *factory) NewRetrieveOneAppUseCase() *RetrieveOneApp {
	return &RetrieveOneApp{AppRepository: f.appRepository}
}

func (f *factory) NewRetrieveAllResourcesFromAppUseCase() *RetrieveAllResourcesFromApp {
	return &RetrieveAllResourcesFromApp{
		AppRepository:      f.appRepository,
		ResourceRepository: f.resourceRepository,
	}
}

func (f *factory) NewResourcesRepository() resource.Repository {
	return resource.NewPostgresRepository(f.db)
}

func (f *factory) NewAppRepository() app.Repository {
	return app.NewPostgresRepository(f.db)
}

func (f *factory) newDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return db
}
