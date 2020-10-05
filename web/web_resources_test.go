package web_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
)

var _ = Describe("HTTP API for resources", func() {
	Context("API V1 returns the same than legacy", func() {
		It("in /resources", func() {
			responseV1 := doGetRequest("/v1/resources")
			responseLegacy := doGetRequest("/resources")

			Expect(responseV1).To(Equal(responseLegacy))
		})

		It("in /resources/:kind/:app/:appVersion", func() {
			responseV1 := doGetRequest("/v1/resources/Description/aws-fargate/1.0.0")
			responseLegacy := doGetRequest("/resources/Description/aws-fargate/1.0.0")

			Expect(responseV1).To(Equal(responseLegacy))
		})

		It("in /resources/:kind/:app/:appVersion/:version", func() {
			responseV1 := doGetRequest("/v1/resources/Description/aws-fargate/1.0.0/2.0.0")
			responseLegacy := doGetRequest("/resources/Description/aws-fargate/1.0.0/2.0.0")

			Expect(responseV1).To(Equal(responseLegacy))
		})
	})

	Context("GET /v1/resources", func() {
		It("returns OK", func() {
			response := doGetRequest("/v1/resources")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an JSON response", func() {
			response := doGetRequest("/v1/resources")

			Expect(response.Header.Get("Content-Type"), "application/json")
		})
	})

	Context("GET /v1/resources/:kind/:app/:appVersion", func() {
		It("returns OK", func() {
			response := doGetRequest("/v1/resources/Description/aws-fargate/1.0.0")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an JSON response", func() {
			response := doGetRequest("/v1/resources/Description/aws-fargate/1.0.0")

			Expect(response.Header.Get("Content-Type"), "application/json")
		})

		Context("when app is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/v1/resources/Description/non-existent/1.0.0")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})

		Context("when app version is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/v1/resources/resources/Description/aws-fargate/non-existent")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})

		Context("when kind does not exist", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/v1/resources/non-existent/aws-fargate/1.0.0")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})

	Context("GET /v1/resources/:kind/:app/:appVersion/:version", func() {
		It("returns OK", func() {
			response := doGetRequest("/v1/resources/Description/aws-fargate/1.0.0/2.0.0")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an JSON response", func() {
			response := doGetRequest("/v1/resources/Description/aws-fargate/1.0.0/2.0.0")

			Expect(response.Header.Get("Content-Type"), "application/json")
		})

		Context("when app is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/v1/resources/Description/non-existent/1.0.0/1.0.0")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})

		Context("when app version is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/v1/resources/Description/aws-fargate/non-existent/1.0.0")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})

		Context("when kind does not exist", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/v1/resources/non-existent/aws-fargate/1.0.0/1.0.0")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})

		Context("when resource version does not exist", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/v1/resources/Description/aws-fargate/1.0.0/5.0.0")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})

})
