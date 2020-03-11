package resource

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type ResourceDTO struct {
	Kind              string           `json:"kind" yaml:"kind"`
	App               string           `json:"app" yaml:"app"`
	Version           string           `json:"version" yaml:"version"`
	AvailableVersions []string         `json:"availableVersions" yaml:"-"`
	AppVersion        []string         `json:"appVersion" yaml:"appVersion"`
	Maintainers       []*MaintainerDTO `json:"maintainers" yaml:"maintainers"`
	Data              string           `json:"data" yaml:"data"`
}

type MaintainerDTO struct {
	Name string `json:"name" yaml:"name"`
	Link string `json:"link" yaml:"link"`
}

func NewResourceDTO(entity *Resource) *ResourceDTO {
	return &ResourceDTO{
		Kind:              entity.Kind,
		App:               entity.App,
		Version:           entity.Version,
		AvailableVersions: entity.AvailableVersions,
		AppVersion:        entity.AppVersion,
		Maintainers:       parseMaintainers(entity.Maintainers),
		Data:              entity.Data,
	}
}

func parseMaintainers(maintainers []*Maintainer) []*MaintainerDTO {
	var result []*MaintainerDTO

	for _, maintainer := range maintainers {
		result = append(result, &MaintainerDTO{
			Name: maintainer.Name,
			Link: maintainer.Link,
		})
	}

	return result
}

func (r *ResourceDTO) ToEntity() *Resource {
	return &Resource{
		ID: NewResourceID(r.App,
			r.Kind,
			r.AppVersion),
		Kind:              r.Kind,
		App:               r.App,
		Version:           r.Version,
		AvailableVersions: r.AvailableVersions,
		AppVersion:        r.AppVersion,
		Maintainers:       toEntityMaintainers(r.Maintainers),
		Data:              r.Data,
	}
}

func toEntityMaintainers(maintainers []*MaintainerDTO) []*Maintainer {
	var result []*Maintainer

	for _, maintainer := range maintainers {
		result = append(result, &Maintainer{
			Name: maintainer.Name,
			Link: maintainer.Link,
		})
	}

	return result
}

func (r ResourceDTO) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *ResourceDTO) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, &r)
}
