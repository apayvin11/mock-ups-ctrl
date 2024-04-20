package util

import "math"

func Float32To2Uint16(val float32) []uint16 {
	u := math.Float32bits(val)
	return []uint16{uint16(u >> 16), uint16(u)}
}
