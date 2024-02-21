package geo

import "github.com/torabian/fireback/modules/workspaces"

type GeoCityFieldMap struct {
	Name workspaces.TranslatedString `yaml:"name"`

	Province workspaces.TranslatedString `yaml:"province"`

	State workspaces.TranslatedString `yaml:"state"`

	Country workspaces.TranslatedString `yaml:"country"`
}
