package app

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gosimple/slug"
	"gopkg.in/yaml.v2"
)

//Kind defines the type of file containing the data
type Kind string

//APP is the kind of the files of the package app
const (
	APP Kind = "App"
)

//App defines the data structure of an application
type App struct {
	ID                string   `json:"id,omitempty" yaml:"-"`
	Kind              Kind     `json:"kind" yaml:"kind"`
	Name              string   `json:"name" yaml:"name"`
	Keywords          []string `json:"keywords" yaml:"keywords"`
	VersionsAvailable []string `json:"versionsAvailable" yaml:"versionsAvailable"`
	Description       string   `json:"description" yaml:"description"`
	ShortDescription  string   `json:"shortDescription" yaml:"shortDescription"`
	Icon              string   `json:"icon" yaml:"icon"`
	Website           string   `json:"website" yaml:"website"`
	Available         bool     `json:"available" yaml:"available"`
}

type appAlias App // Avoid stack overflow while marshalling / unmarshalling

func (r *App) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	res := appAlias{}
	err = unmarshal(&res)
	if err != nil {
		return
	}
	*r = App(res)
	r.ID = r.generateID()
	return
}

func (r *App) MarshalYAML() (interface{}, error) {
	x := appAlias(*r)
	r.ID = r.generateID()
	return yaml.Marshal(x)
}

func (r *App) UnmarshalJSON(data []byte) (err error) {
	res := appAlias{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	*r = App(res)
	r.ID = r.generateID()
	return
}

func (r *App) MarshalJSON() ([]byte, error) {
	x := appAlias(*r)
	r.ID = r.generateID()
	return json.Marshal(x)
}

func (r *App) Validate() error {
	var errors []string

	if r.Kind == "" {
		errors = append(errors, "the app must have a defined Kind")
	}

	if r.Icon == "" {
		errors = append(errors, "the app must have a valid icon")
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, ","))
	}

	return nil
}

func (r *App) generateID() string {
	return slug.Make(r.Name)
}
