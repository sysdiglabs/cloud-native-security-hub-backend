package infrastructure

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/sysdiglabs/prometheus-hub/pkg/resource"
)

func GetResourcesFromPath(path string) ([]*resource.Resource, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	var resources []*resource.Resource

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".yaml" {
			resource, err := getResourceFromFile(path)
			if err != nil {
				return err
			}
			resources = append(resources, resource)
		}
		return nil
	})

	return resources, nil
}

func getResourceFromFile(path string) (*resource.Resource, error) {
	var dto resource.ResourceDTO

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	err = yaml.NewDecoder(file).Decode(&dto)
	if err != nil {
		return nil, err
	}

	return dto.ToEntity(), nil
}
