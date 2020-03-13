package resource

import (
	"github.com/gosimple/slug"
)

type ResourceID struct {
	kind       string
	app        string
	appID      string
	appVersion string
}

func NewResourceID(app, kind string, appVersion []string) ResourceID {
	resourceID := ResourceID{
		kind:       kind,
		app:        app,
		appID:      "",
		appVersion: appVersion[0],
	}
	resourceID.appID = resourceID.generateAppID()
	return resourceID
}

func (r *ResourceID) generateAppID() string {
	return slug.Make(r.app)
}
