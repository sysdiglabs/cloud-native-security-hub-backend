package web_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/sysdiglabs/prometheus-hub/web"
)

var _ = Describe("HTTP Server Logging", func() {
	It("is logging requests", func() {
		request, _ := http.NewRequest("GET", "/resources/Description/AWS Fargate/1.0.0/2.0.0", nil)
		recorder := httptest.NewRecorder()

		buff := &bytes.Buffer{}
		router := web.NewRouterWithLogger(log.New(buff, "", 0))
		router.ServeHTTP(recorder, request)

		Expect("200 [] GET /resources/Description/AWS%20Fargate/1.0.0/2.0.0\n").To(Equal(buff.String()))
	})
})
