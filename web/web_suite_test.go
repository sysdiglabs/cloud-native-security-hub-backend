package web_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/sysdiglabs/prometheus-hub/pkg/app"
	"github.com/sysdiglabs/prometheus-hub/pkg/resource"
	"github.com/sysdiglabs/prometheus-hub/test/fixtures/apps"
	"github.com/sysdiglabs/prometheus-hub/test/fixtures/resources"
	"github.com/sysdiglabs/prometheus-hub/web"
)

func TestWeb(t *testing.T) {
	loadDatabaseData()
	RegisterFailHandler(Fail)
	RunSpecs(t, "Web Suite")
}

func loadDatabaseData() {
	var appRepository app.Repository
	var resourceRepository resource.Repository

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	db.Exec(`TRUNCATE TABLE latest_security_resources;
					  TRUNCATE TABLE security_resources;
					  TRUNCATE TABLE apps;`)

	appRepository = app.NewPostgresRepository(db)
	appRepository.Save(apps.AwsFargate())
	appRepository.Save(apps.AwsLambda())

	resourceRepository = resource.NewPostgresRepository(db)

	fargateDescription := resources.AwsFargateDescription()
	resourceRepository.Save(fargateDescription)
	fargateDescription.Version = "2.0.0"
	resourceRepository.Save(fargateDescription)
	resourceRepository.Save(resources.AwsFargateAlerts())

}

func doGetRequest(path string) *http.Response {
	request, _ := http.NewRequest("GET", path, nil)

	recorder := httptest.NewRecorder()
	router := web.NewRouter()
	router.ServeHTTP(recorder, request)

	return recorder.Result()
}
