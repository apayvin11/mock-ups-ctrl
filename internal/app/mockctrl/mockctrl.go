package mockctrl

import (
	"github.com/alex11prog/mock-ups-ctrl/internal/app/util"
	"github.com/tbrandon/mbserver"
)

type battery struct {
	voltage float32
	temp    float32
	resist  float32
}

type batteryGroup struct {
	voltage   float32
	current   float32
	batteries []battery
}

var ctrlMockParams = []batteryGroup{
	{
		voltage: 48,
		current: 10,
		batteries: []battery{
			{
				voltage: 12,
				temp:    25,
				resist:  5,
			},
			{
				voltage: 13,
				temp:    24,
				resist:  4,
			},
			{
				voltage: 11,
				temp:    23,
				resist:  6,
			},
			{
				voltage: 12,
				temp:    22,
				resist:  7,
			},
		},
	},
	{
		voltage: 50,
		current: 12.5,
		batteries: []battery{
			{
				voltage: 14,
				temp:    24,
				resist:  7,
			},
			{
				voltage: 11,
				temp:    25,
				resist:  6,
			},
			{
				voltage: 13.5,
				temp:    23,
				resist:  5,
			},
			{
				voltage: 11.5,
				temp:    21,
				resist:  4,
			},
		},
	},
}

// InitMockModbusParams initializes registers with mock parameters
func InitMockModbusParams(s *mbserver.Server) {
	for groupIndex, group := range ctrlMockParams {
		regGroup := groupIndex << 8
		// filling InputRegisters according to the modbus register map
		copy(s.InputRegisters[regGroup:], util.Float32To2Uint16(group.voltage))
		copy(s.InputRegisters[regGroup|0x02:], util.Float32To2Uint16(group.current))
		for batteryIndex, battery := range group.batteries {
			regBat := (batteryIndex + 1) << 4
			copy(s.InputRegisters[regGroup|regBat:], util.Float32To2Uint16(battery.voltage))
			copy(s.InputRegisters[regGroup|regBat|0x02:], util.Float32To2Uint16(battery.temp))
			copy(s.InputRegisters[regGroup|regBat|0x04:], util.Float32To2Uint16(battery.resist))
		}
	}
}