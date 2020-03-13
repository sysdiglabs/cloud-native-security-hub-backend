package infrastructure

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/sysdiglabs/prometheus-hub/pkg/app"
)

func GetAppsFromPath(path string) ([]*app.App, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	var apps []*app.App

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".yaml" {
			app, err := getAppFromFile(path)
			if err != nil {
				return err
			}
			apps = append(apps, &app)
		}
		return nil
	})

	return apps, nil
}

func getAppFromFile(path string) (app app.App, err error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	defer file.Close()
	if err != nil {
		return
	}

	err = yaml.NewDecoder(file).Decode(&app)
	if err != nil {
		return
	}

	return
}
