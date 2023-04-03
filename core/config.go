package core

import "github.com/xerion3800/evcc/api"

// configProvider gives access to configuration repository
type configProvider interface {
	Meter(string) (api.Meter, error)
	Charger(string) (api.Charger, error)
	Vehicle(string) (api.Vehicle, error)
}
