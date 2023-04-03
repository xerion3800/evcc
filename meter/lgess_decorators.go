package meter

// Code generated by github.com/xerion3800/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/xerion3800/evcc/api"
)

func decorateLgEss(base *LgEss, meterEnergy func() (float64, error), battery func() (float64, error), batteryCapacity func() float64) api.Meter {
	switch {
	case battery == nil && batteryCapacity == nil && meterEnergy == nil:
		return base

	case battery == nil && batteryCapacity == nil && meterEnergy != nil:
		return &struct {
			*LgEss
			api.MeterEnergy
		}{
			LgEss: base,
			MeterEnergy: &decorateLgEssMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery != nil && batteryCapacity == nil && meterEnergy == nil:
		return &struct {
			*LgEss
			api.Battery
		}{
			LgEss: base,
			Battery: &decorateLgEssBatteryImpl{
				battery: battery,
			},
		}

	case battery != nil && batteryCapacity == nil && meterEnergy != nil:
		return &struct {
			*LgEss
			api.Battery
			api.MeterEnergy
		}{
			LgEss: base,
			Battery: &decorateLgEssBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decorateLgEssMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery == nil && batteryCapacity != nil && meterEnergy == nil:
		return &struct {
			*LgEss
			api.BatteryCapacity
		}{
			LgEss: base,
			BatteryCapacity: &decorateLgEssBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
		}

	case battery == nil && batteryCapacity != nil && meterEnergy != nil:
		return &struct {
			*LgEss
			api.BatteryCapacity
			api.MeterEnergy
		}{
			LgEss: base,
			BatteryCapacity: &decorateLgEssBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateLgEssMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy == nil:
		return &struct {
			*LgEss
			api.Battery
			api.BatteryCapacity
		}{
			LgEss: base,
			Battery: &decorateLgEssBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateLgEssBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy != nil:
		return &struct {
			*LgEss
			api.Battery
			api.BatteryCapacity
			api.MeterEnergy
		}{
			LgEss: base,
			Battery: &decorateLgEssBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateLgEssBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateLgEssMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}
	}

	return nil
}

type decorateLgEssBatteryImpl struct {
	battery func() (float64, error)
}

func (impl *decorateLgEssBatteryImpl) Soc() (float64, error) {
	return impl.battery()
}

type decorateLgEssBatteryCapacityImpl struct {
	batteryCapacity func() float64
}

func (impl *decorateLgEssBatteryCapacityImpl) Capacity() float64 {
	return impl.batteryCapacity()
}

type decorateLgEssMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decorateLgEssMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}
