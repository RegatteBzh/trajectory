package wind

import (
	"image"

	"github.com/regattebzh/trajectory/wind"
)

// Speed is the speed of the wind (m/sec)
type Speed struct {
	SpeedU float32
	SpeedV float32
}

// SetWind set a wind speed
func (buffer Map) SetWind(loc image.Point, speed wind.Speed) {
	buffer.Data[loc.Y*buffer.Width+loc.X] = speed
}

// GetWind get a wind speed
func (buffer Map) GetWind(loc image.Point) wind.Speed {
	return buffer.Data[loc.Y*buffer.Width+loc.X]
}

// Compare compares values
func Compare(a Element, b Element) int {
	aLength := a.SpeedU*a.SpeedU + a.SpeedV*a.SpeedV
	bLength := b.SpeedU*b.SpeedU + b.SpeedV*b.SpeedV

	return (aLength - bLength)
}
