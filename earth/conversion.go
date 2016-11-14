package earth

/**
 *  ra = equatorial ray
 *  rb = polar ray
 *  phi = latitude angle [-90; 90]
 *  lambda = longitude angle [0; 360]
 */

import (
	"math"
)

const ra = 6378137
const rb = 6356752

// VectorL is a 2D vector that representes lengths
type VectorL struct {
	U float64
	V float64
}

// VectorA is a 2D vector that representes angles
type VectorA struct {
	Lambda float64
	Phi    float64
}

func localRay(lambda float64) (ray float64) {
	angle := lambda * math.Pi / 180
	ray = ra*math.Cos(angle) + rb*math.Sin(angle)
	return
}

// Angle2length convert angle (degree) in length (meter)
func Angle2length(angle VectorA, lambda float64) (length VectorL) {
	length = VectorL{
		math.Pi * angle.Lambda * localRay(lambda) / 180,
		math.Pi * angle.Phi * localRay(lambda) / 180,
	}
	return
}

// Length2Angle convert length (meter) in angle (degree)
func Length2Angle(length VectorL, lambda float64) (angle VectorA) {
	angle = VectorA{
		180 * length.U / (math.Pi * localRay(lambda)),
		180 * length.V / (math.Pi * localRay(lambda)),
	}
	return
}
