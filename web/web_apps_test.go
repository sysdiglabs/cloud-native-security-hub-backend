package web_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
)

var _ = Describe("HTTP API for apps", func() {
	Context("GET /apps", func() {
		It("returns OK", func() {
			response := doGetRequest("/apps")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an JSON response", func() {
			response := doGetRequest("/apps")

			Expect(response.Header.Get("Content-Type"), "application/json")
		})
	})

	Context("GET /app/:name", func() {
		It("returns OK", func() {
			response := doGetRequest("/apps/aws-fargate")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an JSON response", func() {
			response := doGetRequest("/apps/aws-fargate")

			Expect(response.Header.Get("Content-Type"), "application/json")
		})

		Context("when name is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/apps/non-existent")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})

	Context("GET /apps/:name/:appVersion/resources", func() {
		It("returns OK", func() {
			response := doGetRequest("/apps/aws-fargate/1.0.0/resources")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an JSON response", func() {
			response := doGetRequest("/apps/aws-fargate/1.0.0/resources")

			Expect(response.Header.Get("Content-Type"), "application/json")
		})

		Context("when name is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/apps/non-existent/resources")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})
})
