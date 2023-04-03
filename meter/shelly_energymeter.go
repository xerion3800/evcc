package meter

import (
	"github.com/xerion3800/evcc/api"
	"github.com/xerion3800/evcc/meter/shelly"
	"github.com/xerion3800/evcc/util"
)

// Shelly meter implementation
func init() {
	registry.Add("shelly-energymeter", NewShellyEnergyMeterFromConfig)
}

// NewShellyFromConfig creates a Shelly charger from generic config
func NewShellyEnergyMeterFromConfig(other map[string]interface{}) (api.Meter, error) {
	var cc struct {
		URI      string
		User     string
		Password string
		Channel  int
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	conn, err := shelly.NewConnection(cc.URI, cc.User, cc.Password, cc.Channel)
	if err != nil {
		return nil, err
	}

	return shelly.NewEnergyMeter(conn), nil
}
