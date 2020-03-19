package usecases

import "github.com/sysdiglabs/promcat/pkg/app"

type RetrieveAllApps struct {
	AppRepository app.Repository
}

func (useCase *RetrieveAllApps) Execute() ([]*app.App, error) {
	return useCase.AppRepository.FindAll()
}
