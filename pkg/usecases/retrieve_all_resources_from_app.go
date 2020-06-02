package usecases

import (
	"fmt"

	"github.com/sysdiglabs/promcat/pkg/app"
	"github.com/sysdiglabs/promcat/pkg/resource"
)

type RetrieveAllResourcesFromApp struct {
	AppRepository      app.Repository
	ResourceRepository resource.Repository
}

func (r *RetrieveAllResourcesFromApp) Execute(appID string, appVersion string) (res []*resource.Resource, err error) {
	app, err := r.AppRepository.FindById(appID)
	if err != nil {
		return
	}

	resources, err := r.ResourceRepository.FindAll()
	if err != nil {
		return
	}

	for _, r := range resources {
		if app.Name == r.App {
			for _, resourceAppVersion := range r.AppVersion {
				if appVersion == resourceAppVersion {
					res = append(res, r)
				}
			}
		}
	}

	if len(res) == 0 {
		err = fmt.Errorf("no resources available for this app")
		return res, err
	}

	return res, err
}
