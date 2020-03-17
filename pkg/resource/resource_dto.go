package resource

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type ResourceDTO struct {
	Kind              string           `json:"kind" yaml:"kind"`
	App               string           `json:"app" yaml:"app"`
	AppID             string           `json:"appID" yaml:"-"`
	Version           string           `json:"version" yaml:"version"`
	AvailableVersions []string         `json:"availableVersions" yaml:"-"`
	AppVersion        []string         `json:"appVersion" yaml:"appVersion"`
	Maintainers       []*MaintainerDTO `json:"maintainers" yaml:"maintainers"`
	Data              string           `json:"data" yaml:"data"`
	Dashboards        []*DashboardDTO  `json:"dashboards" yaml:"dashboards"`
	Alerts            AlertsDTO        `json:"alerts" yaml:"alerts"`
}

type MaintainerDTO struct {
	Name string `json:"name" yaml:"name"`
	Link string `json:"link" yaml:"link"`
}

type AlertsDTO struct {
	PrometheusAlerts string `json:"prometheusAlerts" yaml:"prometheusAlerts"`
	SysdigAlerts     string `json:"sysdigAlerts" yaml:"sysdigAlerts"`
}

type DashboardDTO struct {
	Name        string `json:"name" yaml:"name"`
	Kind        string `json:"kind" yaml:"kind"`
	Image       string `json:"image" yaml:"image"`
	Description string `json:"description" yaml:"description"`
	Data        string `json:"data" yaml:"data"`
}

func NewResourceDTO(entity *Resource) *ResourceDTO {
	return &ResourceDTO{
		Kind:              entity.Kind,
		App:               entity.App,
		AppID:             entity.ID.appID,
		Version:           entity.Version,
		AvailableVersions: entity.AvailableVersions,
		AppVersion:        entity.AppVersion,
		Maintainers:       parseMaintainers(entity.Maintainers),
		Data:              entity.Data,
		Dashboards:        parseDashboards(entity.Dashboards),
		Alerts:            parseAlerts(entity.Alerts),
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

func parseAlerts(alerts Alerts) AlertsDTO {
	return AlertsDTO{
		PrometheusAlerts: alerts.PrometheusAlerts,
		SysdigAlerts:     alerts.SysdigAlerts,
	}
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
		Dashboards:        toEntityDashboards(r.Dashboards),
		Alerts:            toEntityAlerts(r.Alerts),
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

func toEntityAlerts(alerts AlertsDTO) Alerts {
	return Alerts{
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
