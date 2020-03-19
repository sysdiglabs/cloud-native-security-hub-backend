package usecases

import "github.com/sysdiglabs/promcat/pkg/resource"

type RetrieveOneResourceByVersion struct {
	ResourceRepository resource.Repository
}

func (r *RetrieveOneResourceByVersion) Execute(app, kind, appVersion, version string) (res *resource.Resource, err error) {
	return r.ResourceRepository.FindByVersion(
		resource.NewResourceID(app, kind, []string{appVersion}),
		version)
}
