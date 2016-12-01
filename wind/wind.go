package wind

import (
	"image"

	"github.com/regattebzh/trajectory/mapper"
)

// Speed is the speed of the wind (m/sec)
type Speed struct {
	SpeedU float32
	SpeedV float32
}

// SetWind set a wind speed
func SetWind(buffer mapper.Map, loc image.Point, speed Speed) {
	buffer.Data[loc.Y*buffer.Rect.Dx()+loc.X] = speed
}

// GetWind get a wind speed
func GetWind(buffer mapper.Map, loc image.Point) (Speed, bool) {
	speed, ok := buffer.Data[loc.Y*buffer.Rect.Dx()+loc.X].(Speed)
	return speed, ok
}

// Compare compares values
func (a Speed) Compare(b mapper.Element) (int, bool) {
	aLength := a.SpeedU*a.SpeedU + a.SpeedV*a.SpeedV
	bSpeed, ok := b.(Speed)
	if !ok {
		return 0, false
	}
	bLength := bSpeed.SpeedU*bSpeed.SpeedU + bSpeed.SpeedV*bSpeed.SpeedV

	return int(aLength - bLength), true
}
