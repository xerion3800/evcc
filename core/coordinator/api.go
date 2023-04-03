package coordinator

import "github.com/xerion3800/evcc/api"

// API is the coordinator API
type API interface {
	GetVehicles() []api.Vehicle
	Acquire(api.Vehicle)
	Release(api.Vehicle)
	IdentifyVehicleByStatus() api.Vehicle
}
