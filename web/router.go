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

	router.GET("/resources", h.retrieveAllResourcesHandler)

	router.GET("/resources/:kind/:resource", h.retrieveOneResourcesHandler)
	router.GET("/resources/:kind/:resource/version/:version", h.retrieveOneResourceByVersionHandler)

	router.GET("/apps", h.retrieveAllAppsHandler)

	router.GET("/apps/:app", h.retrieveOneAppHandler)
	router.GET("/apps/:app/resources", h.retrieveAllResourcesFromAppHandler)

	router.GET("/health", h.healthCheckHandler)
	router.NotFound = h.notFound()
}
