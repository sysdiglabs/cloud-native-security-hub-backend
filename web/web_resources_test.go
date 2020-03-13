package web_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
)

var _ = Describe("HTTP API for resources", func() {
	Context("GET /resources", func() {
		It("returns OK", func() {
			response := doGetRequest("/resources")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an JSON response", func() {
			response := doGetRequest("/resources")

			Expect(response.Header.Get("Content-Type"), "application/json")
		})
	})

	Context("GET /resources/:kind/:app/:appVersion", func() {
		It("returns OK", func() {
			response := doGetRequest("/resources/Description/AWS Fargate/1.0.0")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an JSON response", func() {
			response := doGetRequest("/resources/Description/AWS Fargate/1.0.0")

			Expect(response.Header.Get("Content-Type"), "application/json")
		})

		Context("when app is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/resources/Description/non-existent/1.0.0")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})

		Context("when app version is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/resources/resources/Description/AWS Fargate/non-existent")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})

		Context("when kind does not exist", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/resources/non-existent/AWS Fargate/1.0.0")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})

	Context("GET /resources/:kind/:app/:appVersion/:version", func() {
		It("returns OK", func() {
			response := doGetRequest("/resources/Description/AWS Fargate/1.0.0/2.0.0")

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns an JSON response", func() {
			response := doGetRequest("/resources/Description/AWS Fargate/1.0.0/2.0.0")

			Expect(response.Header.Get("Content-Type"), "application/json")
		})

		Context("when app is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/resources/Description/non-existent/1.0.0/1.0.0")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})

		Context("when app version is not found", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/resources/Description/AWS Fargate/non-existent/1.0.0")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})

		Context("when kind does not exist", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/resources/non-existent/AWS Fargate/1.0.0/1.0.0")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})

		Context("when resource version does not exist", func() {
			It("returns a NOTFOUND", func() {
				response := doGetRequest("/resources/Description/AWS Fargate/1.0.0/5.0.0")

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})

})
