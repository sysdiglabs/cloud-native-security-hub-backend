package usecases

import "github.com/sysdiglabs/prometheus-hub/pkg/app"

type RetrieveOneApp struct {
	AppRepository app.Repository
}

func (r *RetrieveOneApp) Execute(appID string) (*app.App, error) {
	return r.AppRepository.FindById(appID)
}
