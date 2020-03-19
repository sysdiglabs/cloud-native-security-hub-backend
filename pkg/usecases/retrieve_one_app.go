package usecases

import (
	"github.com/sysdiglabs/promcat/pkg/app"
)

type RetrieveOneApp struct {
	AppRepository app.Repository
}

func (r *RetrieveOneApp) Execute(app string) (*app.App, error) {
	return r.AppRepository.FindById(app)
}
