package web

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
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

	router.GET("/v1/resources", h.retrieveAllResourcesHandler)
	router.GET("/resources", h.retrieveAllResourcesHandler)

	router.GET("/v1/resources/:kind/:app/:appVersion", h.retrieveOneResourcesHandler)
	router.GET("/resources/:kind/:app/:appVersion", h.retrieveOneResourcesHandler)

	router.GET("/v1/resources/:kind/:app/:appVersion/:version", h.retrieveOneResourceByVersionHandler)
	router.GET("/resources/:kind/:app/:appVersion/:version", h.retrieveOneResourceByVersionHandler)

	router.GET("/v1/apps", h.retrieveAllAppsHandler)
	router.GET("/apps", h.retrieveAllAppsHandler)

	router.GET("/v1/apps/:app", h.retrieveOneAppHandler)
	router.GET("/apps/:app", h.retrieveOneAppHandler)

	router.GET("/v1/apps/:app/:appVersion/resources", h.retrieveAllResourcesFromAppHandler)
	router.GET("/apps/:app/:appVersion/resources", h.retrieveAllResourcesFromAppHandler)

	router.GET("/health", h.healthCheckHandler)
	router.NotFound = h.notFound()
}
