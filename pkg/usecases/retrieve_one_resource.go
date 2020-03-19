package usecases

import "github.com/sysdiglabs/promcat/pkg/resource"

type RetrieveOneResource struct {
	ResourceRepository resource.Repository
}

func (r *RetrieveOneResource) Execute(app, kind, appVersion string) (res *resource.Resource, err error) {
	return r.ResourceRepository.FindById(resource.NewResourceID(app, kind, []string{appVersion}))
}
