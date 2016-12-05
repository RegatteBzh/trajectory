package etopo

import (
	"image"

	"github.com/regattebzh/trajectory/mapper"
)

// Altitude is the speed of the wind (m/sec)
type Altitude int16

// SetAltitude set a wind speed
func SetAltitude(buffer mapper.Map, loc image.Point, altitude Altitude) {
	buffer.Set(loc, altitude)
}

// GetAltitude get a wind speed
func GetAltitude(buffer mapper.Map, loc image.Point) (Altitude, bool) {
	alt, ok := buffer.Get(loc).(Altitude)
	return alt, ok
}

// Compare compares values
func (a Altitude) Compare(b mapper.Element) (int, bool) {

	bAlt, ok := b.(Altitude)
	if !ok {
		return 0, false
	}

	return int(a - bAlt), true
}
