package web_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"

	"github.com/sysdiglabs/prometheus-hub/web"
)

func TestWeb(t *testing.T) {
	loadDatabaseFixture()
	RegisterFailHandler(Fail)
	RunSpecs(t, "Web Suite")
}

func loadDatabaseFixture() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`TRUNCATE TABLE latest_security_resources;
					  TRUNCATE TABLE security_resources;
					  TRUNCATE TABLE vendors;
					  TRUNCATE TABLE schema_migrations`)
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.Join("testdata", "fixture.sql") // relative path
	data, err := ioutil.ReadFile(path)
	_, err = db.Exec(string(data))
	if err != nil {
		log.Fatal(err)
	}
}

func doGetRequest(path string) *http.Response {
	request, _ := http.NewRequest("GET", path, nil)

	recorder := httptest.NewRecorder()
	router := web.NewRouter()
	router.ServeHTTP(recorder, request)

	return recorder.Result()
}
