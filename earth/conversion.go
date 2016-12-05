package earth

/**
 *  ra = equatorial ray
 *  rb = polar ray
 *  Lat = Latitude angle [-90; 90]
 *  Lon = Longitude angle [0; 360]
 */

import (
	"math"
)

// RA is the ray at equator
const RA = 6378137

// RB is the ray at the pole
const RB = 6356752

// VectorL is a 2D vector that representes lengths
type VectorL struct {
	U float64 // Longitude length
	V float64 // Latitude length
}

// VectorA is a 2D vector that representes angles
type VectorA struct {
	Lon float64
	Lat float64
}

func degreeToRadian(angle float64) float64 {
	return angle * math.Pi / 180
}

// LocalRayLatitude compute local ray on latitude
func LocalRayLatitude(Lon float64) float64 {
	angle := degreeToRadian(Lon)
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return math.Sqrt(RA*RA*cos*cos + RB*RB*sin*sin)
}

// LocalRayLongitude compute local ray on longitude
func LocalRayLongitude(Lon float64) float64 {
	angle := degreeToRadian(Lon)
	return LocalRayLatitude(Lon) * math.Cos(angle)
}

// Angle2length convert angle (degree) in length (meter)
func Angle2length(angle VectorA, Lon float64) (length VectorL) {
	length = VectorL{
		U: math.Pi * angle.Lon * LocalRayLongitude(Lon) / 180,
		V: math.Pi * angle.Lat * LocalRayLatitude(Lon) / 180,
	}
	return
}

// Length2Angle convert length (meter) in angle (degree)
func Length2Angle(length VectorL, Lon float64) (angle VectorA) {
	angle = VectorA{
		Lon: 180 * length.U / (math.Pi * LocalRayLongitude(Lon)),
		Lat: 180 * length.V / (math.Pi * LocalRayLatitude(Lon)),
	}
	return
}
