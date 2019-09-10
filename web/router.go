package web

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func NewRouter() http.Handler {
	router := httprouter.New()
	registerOn(router, nil)

	return cors.Default().Handler(router)
}

func NewRouterWithLogger(logger *log.Logger) http.Handler {
	router := httprouter.New()
	registerOn(router, logger)

	return cors.Default().Handler(router)
}

func registerOn(router *httprouter.Router, logger *log.Logger) {
	h := NewHandlerRepository(logger)
	router.GET("/resources", h.retrieveAllResourcesHandler)
	router.GET("/resources/:hash", h.retrieveOneResourcesHandler)
	router.GET("/resources/:hash/raw.yaml", h.retrieveOneResourcesRawHandler)
	router.GET("/vendors", h.retrieveAllVendorsHandler)
	router.GET("/vendors/:vendor", h.retrieveOneVendorsHandler)
	router.GET("/vendors/:vendor/resources", h.retrieveAllResourcesFromVendorHandler)
	router.NotFound = h.notFound()
}
