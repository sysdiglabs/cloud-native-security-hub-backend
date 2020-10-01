package web_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
)

var _ = Describe("HTTP API for apps", func() {
	Context("API V1 returns the same than legacy", func() {
		It("in /apps", func() {
			responseV1 := doGetRequest("/v1/apps")
			responseLegacy := doGetRequest("/apps")

			Expect(responseV1).To(Equal(responseLegacy))
		})

		It("in /apps/:name", func() {
			responseV1 := doGetRequest("/v1/apps/aws-fargate")
			responseLegacy := doGetRequest("/apps/aws-fargate")

			Expect(responseV1).To(Equal(responseLegacy))
		})

		It("in /apps/:name/:appVersion/resources", func() {
			responseV1 := doGetRequest("/v1/apps/aws-fargate/1.0.0/resources")
			responseLegacy := doGetRequest("/apps/aws-fargate/1.0.0/resources")

			Expect(responseV1).To(Equal(responseLegacy))
		})
	})

	Context("GET /v1/apps", func() {
		It("returns OK", func() {
			response := doGetRequest("/v1/apps")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an JSON response", func() {
			response := doGetRequest("/v1/apps")

			Expect(response.Header.Get("Content-Type"), "application/json")
		})
	})

	Context("GET /v1/apps/:name", func() {
		It("returns OK", func() {
			response := doGetRequest("/v1/apps/aws-fargate")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an JSON response", func() {
			response := doGetRequest("/v1/apps/aws-fargate")

			Expect(response.Header.Get("Content-Type"), "application/json")
		})

		Context("when name is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/v1/apps/non-existent")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})

	Context("GET /v1/apps/:name/:appVersion/resources", func() {
		It("returns OK", func() {
			response := doGetRequest("/v1/apps/aws-fargate/1.0.0/resources")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an JSON response", func() {
			response := doGetRequest("/v1/apps/aws-fargate/1.0.0/resources")

			Expect(response.Header.Get("Content-Type"), "application/json")
		})

		Context("when name is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/v1/apps/non-existent/resources")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})

		Context("when app version is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/v1/apps/aws-fargate/1.3/resources")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})
})
