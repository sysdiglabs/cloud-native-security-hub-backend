package usecases

import (
	"github.com/sysdiglabs/prometheus-hub/pkg/resource"
)

type RetrieveAllResources struct {
	ResourceRepository resource.Repository
}

func (useCase *RetrieveAllResources) Execute() ([]*resource.Resource, error) {
	return useCase.ResourceRepository.FindAll()
}
