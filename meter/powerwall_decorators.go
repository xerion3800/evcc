package meter

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decoratePowerWall(base *PowerWall, meterEnergy func() (float64, error), battery func() (float64, error), batteryCapacity func() float64) api.Meter {
	switch {
	case battery == nil && batteryCapacity == nil && meterEnergy == nil:
		return base

	case battery == nil && batteryCapacity == nil && meterEnergy != nil:
		return &struct {
			*PowerWall
			api.MeterEnergy
		}{
			PowerWall: base,
			MeterEnergy: &decoratePowerWallMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery != nil && batteryCapacity == nil && meterEnergy == nil:
		return &struct {
			*PowerWall
			api.Battery
		}{
			PowerWall: base,
			Battery: &decoratePowerWallBatteryImpl{
				battery: battery,
			},
		}

	case battery != nil && batteryCapacity == nil && meterEnergy != nil:
		return &struct {
			*PowerWall
			api.Battery
			api.MeterEnergy
		}{
			PowerWall: base,
			Battery: &decoratePowerWallBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decoratePowerWallMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery == nil && batteryCapacity != nil && meterEnergy == nil:
		return &struct {
			*PowerWall
			api.BatteryCapacity
		}{
			PowerWall: base,
			BatteryCapacity: &decoratePowerWallBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
		}

	case battery == nil && batteryCapacity != nil && meterEnergy != nil:
		return &struct {
			*PowerWall
			api.BatteryCapacity
			api.MeterEnergy
		}{
			PowerWall: base,
			BatteryCapacity: &decoratePowerWallBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decoratePowerWallMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy == nil:
		return &struct {
			*PowerWall
			api.Battery
			api.BatteryCapacity
		}{
			PowerWall: base,
			Battery: &decoratePowerWallBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decoratePowerWallBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy != nil:
		return &struct {
			*PowerWall
			api.Battery
			api.BatteryCapacity
			api.MeterEnergy
		}{
			PowerWall: base,
			Battery: &decoratePowerWallBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decoratePowerWallBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decoratePowerWallMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}
	}

	return nil
}

type decoratePowerWallBatteryImpl struct {
	battery func() (float64, error)
}

func (impl *decoratePowerWallBatteryImpl) Soc() (float64, error) {
	return impl.battery()
}

type decoratePowerWallBatteryCapacityImpl struct {
	batteryCapacity func() float64
}

func (impl *decoratePowerWallBatteryCapacityImpl) Capacity() float64 {
	return impl.batteryCapacity()
}

type decoratePowerWallMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decoratePowerWallMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}
