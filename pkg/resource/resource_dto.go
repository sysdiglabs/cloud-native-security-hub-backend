package resource

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"log"
)

type ResourceDTO struct {
	Kind              string           `json:"kind" yaml:"kind"`
	App               string           `json:"app" yaml:"app"`
	AppID             string           `json:"appID" yaml:"-"`
	Version           string           `json:"version" yaml:"version"`
	AvailableVersions []string         `json:"availableVersions" yaml:"-"`
	AppVersion        []string         `json:"appVersion" yaml:"appVersion"`
	Maintainers       []*MaintainerDTO `json:"maintainers" yaml:"maintainers"`
	Description       string           `json:"description,omitempty" yaml:"description"`
	Data              string           `json:"data,omitempty" yaml:"data"`
	Dashboards        []*DashboardDTO  `json:"dashboards,omitempty" yaml:"dashboards"`
	Alerts            *AlertsDTO       `json:"alerts,omitempty" yaml:"alerts"`
}

type MaintainerDTO struct {
	Name string `json:"name" yaml:"name"`
	Link string `json:"link" yaml:"link"`
}

type AlertsDTO struct {
	PrometheusAlerts string `json:"prometheusAlerts,omitempty" yaml:"prometheusAlerts"`
	SysdigAlerts     string `json:"sysdigAlerts,omitempty" yaml:"sysdigAlerts"`
}

type DashboardDTO struct {
	Name        string `json:"name" yaml:"name"`
	Kind        string `json:"kind" yaml:"kind"`
	Image       string `json:"image" yaml:"image"`
	Description string `json:"description" yaml:"description"`
	Data        string `json:"data" yaml:"data"`
}

func NewResourceDTO(entity *Resource) *ResourceDTO {
	resource := ResourceDTO{
		Kind:              entity.Kind,
		App:               entity.App,
		AppID:             entity.ID.appID,
		Version:           entity.Version,
		AvailableVersions: entity.AvailableVersions,
		AppVersion:        entity.AppVersion,
		Maintainers:       parseMaintainers(entity.Maintainers),
		Description:       entity.Description,
		Data:              entity.Data,
		Dashboards:        parseDashboards(entity.Dashboards),
	}

	if entity.Kind == "Alerts" {
		if entity.Alerts == nil {
			log.Printf("Warning: Resource of kind 'Alerts' without field Alerts. App: %s", string(entity.App))
			resource.Alerts = &AlertsDTO{
				PrometheusAlerts: "",
				SysdigAlerts:     "",
			}
		} else {
			resource.Alerts = parseAlerts(*entity.Alerts)
		}
	}

	return &resource
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

func parseDashboards(dashboards []*Dashboard) []*DashboardDTO {
	var result []*DashboardDTO

	for _, dashboard := range dashboards {
		result = append(result, &DashboardDTO{
			Name:        dashboard.Name,
			Kind:        dashboard.Kind,
			Image:       dashboard.Image,
			Description: dashboard.Description,
			Data:        dashboard.Data,
		})
	}

	return result
}

func parseAlerts(alerts Alerts) *AlertsDTO {
	return &AlertsDTO{
		PrometheusAlerts: alerts.PrometheusAlerts,
		SysdigAlerts:     alerts.SysdigAlerts,
	}
}

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
		Maintainers:       toEntityMaintainers(r.Maintainers),
		Description:       r.Description,
		Data:              r.Data,
		Dashboards:        toEntityDashboards(r.Dashboards),
	}

	if r.Kind == "Alerts" {
		if r.Alerts == nil {
			log.Printf("Resource of kind 'Alerts' without field Alerts. AppID: %s", string(r.AppID))
			resource.Alerts = &Alerts{
				PrometheusAlerts: "",
				SysdigAlerts:     "",
			}
		} else {
			resource.Alerts = toEntityAlerts(*r.Alerts)
		}
	} else {
		resource.Alerts = nil
	}

	return &resource

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

func toEntityDashboards(dashboards []*DashboardDTO) []*Dashboard {
	var result []*Dashboard

	for _, dashboard := range dashboards {
		result = append(result, &Dashboard{
			Name:        dashboard.Name,
			Kind:        dashboard.Kind,
			Image:       dashboard.Image,
			Description: dashboard.Description,
			Data:        dashboard.Data,
		})
	}

	return result
}

func toEntityAlerts(alerts AlertsDTO) *Alerts {
	return &Alerts{
		PrometheusAlerts: alerts.PrometheusAlerts,
		SysdigAlerts:     alerts.SysdigAlerts,
	}
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
