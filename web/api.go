package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/sysdiglabs/prometheus-hub/pkg/resource"
	"github.com/sysdiglabs/prometheus-hub/pkg/usecases"
)

type HandlerRepository interface {
	notFound() http.HandlerFunc
	healthCheckHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	retrieveAllResourcesHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params)

	retrieveOneResourcesHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	retrieveOneResourceByVersionHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	retrieveAllAppsHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params)

	retrieveOneAppHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	retrieveAllResourcesFromAppHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type handlerRepository struct {
	factory usecases.Factory
	logger  *log.Logger
}

func NewHandlerRepository(logger *log.Logger) HandlerRepository {
	return &handlerRepository{
		factory: usecases.NewFactory(),
		logger:  logger,
	}
}

func (h *handlerRepository) notFound() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		h.logRequest(request, 404)
		http.NotFound(writer, request)
	}

}

func (h *handlerRepository) logRequest(request *http.Request, statusCode int) {
	if h.logger == nil {
		return
	}

	line := fmt.Sprintf("%d [%s] %s %s", statusCode, request.RemoteAddr, request.Method, request.URL)
	h.logger.Println(line)
}

func (h *handlerRepository) retrieveAllResourcesHandler(
	writer http.ResponseWriter,
	request *http.Request,
	_ httprouter.Params) {
	useCase := h.factory.NewRetrieveAllResourcesUseCase()
	resources, err := useCase.Execute()
	if err != nil {
		h.logRequest(request, 500)
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")

	h.logRequest(request, 200)
	json.NewEncoder(writer).Encode(collectionToDTO(resources))
}

func collectionToDTO(resources []*resource.Resource) []*resource.ResourceDTO {
	var result []*resource.ResourceDTO

	for _, current := range resources {
		result = append(result, resource.NewResourceDTO(current))
	}

	return result
}

func (h *handlerRepository) retrieveOneResourcesHandler(
	writer http.ResponseWriter,
	request *http.Request,
	params httprouter.Params) {

	useCase := h.factory.NewRetrieveOneResourceUseCase()
	resources, err := useCase.Execute(
		params.ByName("app"),
		params.ByName("kind"),
		params.ByName("appVersion"))
	if err != nil {
		writer.WriteHeader(404)
		h.logRequest(request, 404)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	h.logRequest(request, 200)
	json.NewEncoder(writer).Encode(resource.NewResourceDTO(resources))
}

func (h *handlerRepository) retrieveOneResourceByVersionHandler(
	writer http.ResponseWriter,
	request *http.Request,
	params httprouter.Params) {
	useCase := h.factory.NewRetrieveOneResourceByVersionUseCase()
	resources, err := useCase.Execute(
		params.ByName("app"),
		params.ByName("kind"),
		params.ByName("appVersion"),
		params.ByName("version"))
	if err != nil {
		writer.WriteHeader(404)
		h.logRequest(request, 404)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	h.logRequest(request, 200)
	json.NewEncoder(writer).Encode(resource.NewResourceDTO(resources))
}

func (h *handlerRepository) retrieveAllAppsHandler(
	writer http.ResponseWriter,
	request *http.Request,
	_ httprouter.Params) {
	useCase := h.factory.NewRetrieveAllAppsUseCase()
	apps, err := useCase.Execute()
	if err != nil {
		writer.WriteHeader(404)
		h.logRequest(request, 404)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	h.logRequest(request, 200)
	json.NewEncoder(writer).Encode(apps)
}

func (h *handlerRepository) retrieveOneAppHandler(
	writer http.ResponseWriter,
	request *http.Request,
	params httprouter.Params) {
	useCase := h.factory.NewRetrieveOneAppUseCase()
	vendor, err := useCase.Execute(params.ByName("app"))
	if err != nil {
		writer.WriteHeader(404)
		h.logRequest(request, 404)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	h.logRequest(request, 200)
	json.NewEncoder(writer).Encode(vendor)
}

func (h *handlerRepository) retrieveAllResourcesFromAppHandler(
	writer http.ResponseWriter,
	request *http.Request,
	params httprouter.Params) {
	useCase := h.factory.NewRetrieveAllResourcesFromAppUseCase()
	resources, err := useCase.Execute(params.ByName("app"))
	if err != nil {
		writer.WriteHeader(404)
		h.logRequest(request, 404)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	h.logRequest(request, 200)
	json.NewEncoder(writer).Encode(collectionToDTO(resources))
}

func (h *handlerRepository) healthCheckHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	h.logRequest(request, 200)
	writer.Header().Set("Content-Type", "text/plain")
	writer.Write([]byte("OK"))
}
