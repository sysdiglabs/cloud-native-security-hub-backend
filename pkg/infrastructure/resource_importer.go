package infrastructure

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/sysdiglabs/promcat/pkg/resource"
)

func GetResourcesFromPath(path string) ([]*resource.Resource, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	var resources []*resource.Resource

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".yaml" {
			if !strings.Contains(path, "/include/") {
				resource, err := getResourceFromFile(path)
				if err != nil {
					return err
				}
				resources = append(resources, resource)
			}
		}
		return nil
	})

	return resources, nil
}

func getResourceFromFile(path string) (*resource.Resource, error) {
	var dto resource.ResourceDTO
	var resourceEntity *resource.Resource

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	err = yaml.NewDecoder(file).Decode(&dto)
	if err != nil {
		return nil, err
	}

	resourceEntity = dto.ToEntity()

	// If in the configurations there are 'file' fields, fill the 'data' field with its content
	if resourceEntity.Configurations != nil {
		for _, configuration := range resourceEntity.Configurations {
			if configuration.File != "" {
				fileToIncludePath := fmt.Sprintf("%s/%s", filepath.Dir(path), configuration.File)
				bytes, err := ioutil.ReadFile(fileToIncludePath)
				if err != nil {
					fmt.Print(err)
				} else {
					configuration.Data = string(bytes)
				}
			}
		}
	}

	return resourceEntity, nil
}
