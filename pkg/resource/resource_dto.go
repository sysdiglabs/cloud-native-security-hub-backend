package resource

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// ResourceDTO Allows to parse the resources from and to files
type ResourceDTO struct {
	Kind              string              `json:"kind" yaml:"kind"`
	App               string              `json:"app" yaml:"app"`
	AppID             string              `json:"appID" yaml:"-"`
	Version           string              `json:"version" yaml:"version"`
	AvailableVersions []string            `json:"availableVersions" yaml:"-"`
	AppVersion        []string            `json:"appVersion" yaml:"appVersion"`
	Maintainers       string              `json:"maintainers,omitempty" yaml:"maintainers"`
	Description       string              `json:"description,omitempty" yaml:"description"`
	Data              string              `json:"data,omitempty" yaml:"data"`
	Configurations    []*ConfigurationDTO `json:"configurations,omitempty" yaml:"configurations"`
}

// ConfigurationDTO Allows to parse configurations for dashboards and alerts
type ConfigurationDTO struct {
	Name        string `json:"name,omitempty" yaml:"name"`
	Kind        string `json:"kind" yaml:"kind"`
	Image       string `json:"image,omitempty" yaml:"image"`
	Description string `json:"description,omitempty" yaml:"description"`
	File        string `json:"file,omitempty" yaml:"file"`
	Data        string `json:"data,omitempty" yaml:"data"`
}

// NewResourceDTO Creates a new resourceDTO from an entity Resource
func NewResourceDTO(entity *Resource) *ResourceDTO {
	resource := ResourceDTO{
		Kind:              entity.Kind,
		App:               entity.App,
		AppID:             entity.ID.appID,
		Version:           entity.Version,
		AvailableVersions: entity.AvailableVersions,
		AppVersion:        entity.AppVersion,
		Maintainers:       entity.Maintainers,
		Description:       entity.Description,
		Data:              entity.Data,
		Configurations:    parseConfigurations(entity.Configurations),
	}

	return &resource
}

func parseConfigurations(configurations []*Configuration) []*ConfigurationDTO {
	var result []*ConfigurationDTO

	for _, configuration := range configurations {
		result = append(result, &ConfigurationDTO{
			Name:        configuration.Name,
			Kind:        configuration.Kind,
			Image:       configuration.Image,
			Description: configuration.Description,
			File:        configuration.File,
			Data:        configuration.Data,
		})
	}

	return result
}

// ToEntity Converts a ResourceDTO into an entity Resource
func (r *ResourceDTO) ToEntity() *Resource {
	resource := Resource{
		ID: NewResourceID(r.App,
			r.Kind,
			r.AppVersion),
		Kind:              r.Kind,
		App:               r.App,
		Version:           r.Version,
		AvailableVersions: r.AvailableVersions,
		AppVersion:        r.AppVersion,
		Maintainers:       r.Maintainers,
		Description:       r.Description,
		Data:              r.Data,
		Configurations:    toEntityConfigurations(r.Configurations),
	}

	return &resource

}

func toEntityConfigurations(configurations []*ConfigurationDTO) []*Configuration {
	var result []*Configuration

	for _, configuration := range configurations {
		result = append(result, &Configuration{
			Name:        configuration.Name,
			Kind:        configuration.Kind,
			Image:       configuration.Image,
			Description: configuration.Description,
			File:        configuration.File,
			Data:        configuration.Data,
		})
	}

	return result
}

// Value Returns the ResourceDTO parsed as JSON
func (r ResourceDTO) Value() (driver.Value, error) {
	return json.Marshal(r)
}

// Scan Parses an interface and returns a ResourceDTO with its content
func (r *ResourceDTO) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, &r)
}
