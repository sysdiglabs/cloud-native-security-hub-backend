package resource

type ResourceID struct {
	kind       string
	app        string
	appVersion string
}

func NewResourceID(app, kind string, appVersion []string) ResourceID {
	return ResourceID{
		kind:       kind,
		app:        app,
		appVersion: appVersion[0],
	}
}
