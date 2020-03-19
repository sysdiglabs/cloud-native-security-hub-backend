package resource

import (
	"github.com/gosimple/slug"
)

// ResourceID is defined as a combination of kind, appId and appVersion
type ResourceID struct {
	kind       string
	app        string
	appID      string
	appVersion string
}

// NewResourceID Generates a new ResourceID
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
