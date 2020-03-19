package usecases

import (
	"github.com/sysdiglabs/promcat/pkg/resource"
)

type RetrieveAllResources struct {
	ResourceRepository resource.Repository
}

func (useCase *RetrieveAllResources) Execute() ([]*resource.Resource, error) {
	return useCase.ResourceRepository.FindAll()
}
