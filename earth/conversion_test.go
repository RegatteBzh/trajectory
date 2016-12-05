package earth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalRayLatitudePole(t *testing.T) {
	northPolarRay := localRayLatitude(-90)
	assert.Equal(t, rb, int(northPolarRay+0.5), "should get latitude ray at the north pole")
	southPolarRay := localRayLatitude(90)
	assert.Equal(t, rb, int(southPolarRay+0.5), "should get latitude ray at the south pole")
}

func TestLocalRayLatitudeEquator(t *testing.T) {
	equatorRay := localRayLatitude(0)
	assert.Equal(t, ra, int(equatorRay+0.5), "should get latitude ray at the equator")
}

func TestLocalRayLongitudePole(t *testing.T) {
	northPolarRay := localRayLongitude(-90)
	assert.Equal(t, 0, int(northPolarRay+0.5), "should get longitude ray 0 at the north pole")
	southPolarRay := localRayLongitude(90)
	assert.Equal(t, 0, int(southPolarRay+0.5), "should get longitude ray 0 at the south pole")
}

func TestLocalRayLongitudeEquator(t *testing.T) {
	equatorRay := localRayLongitude(0)
	assert.Equal(t, ra, int(equatorRay+0.5), "should get longitude ray at the equator")
}

func TestAngle2lengthEquator1Minute(t *testing.T) {
	vectLat := Angle2length(
		VectorA{
			Lon: 0,
			Lat: float64(1) / 60,
		},
		0,
	)
	assert.Equal(t, 0, int(vectLat.U+0.5), "should get longitude length at the equator for 0 deg")
	assert.Equal(t, 1855, int(vectLat.V+0.5), "should get latitude length at the equator for 1 minute")
	vectLon := Angle2length(
		VectorA{
			Lon: float64(1) / 60,
			Lat: 0,
		},
		0,
	)
	assert.Equal(t, 1855, int(vectLon.U+0.5), "should get longitude length at the equator for 1 minute")
	assert.Equal(t, 0, int(vectLon.V+0.5), "should get latitude length at the equator for 0 deg")
}

func TestAngle2lengthSouthPole1Minute(t *testing.T) {
	vectLat := Angle2length(
		VectorA{
			Lon: 0,
			Lat: float64(1) / 60,
		},
		90,
	)
	assert.Equal(t, 0, int(vectLat.U+0.5), "should get longitude length at the south pole for 0 deg")
	assert.Equal(t, 1849, int(vectLat.V+0.5), "should get latitude length at the south pole for 1 minute")
	vectLon := Angle2length(
		VectorA{
			Lon: float64(1) / 60,
			Lat: 0,
		},
		90,
	)
	assert.Equal(t, 0, int(vectLon.U+0.5), "should get longitude length at the south pole for 1 minute")
	assert.Equal(t, 0, int(vectLon.V+0.5), "should get latitude length at the south pole for 0 deg")

}

func TestAngle2lengthNothPole1Minute(t *testing.T) {
	vectLat := Angle2length(
		VectorA{
			Lon: 0,
			Lat: float64(1) / 60,
		},
		-90,
	)
	assert.Equal(t, 0, int(vectLat.U+0.5), "should get longitude length at the north pole for 0 deg")
	assert.Equal(t, 1849, int(vectLat.V+0.5), "should get latitude length at the north pole for 1 minute")
	vectLon := Angle2length(
		VectorA{
			Lon: float64(1) / 60,
			Lat: 0,
		},
		-90,
	)
	assert.Equal(t, 0, int(vectLon.U+0.5), "should get longitude length at the north pole for 1 minute")
	assert.Equal(t, 0, int(vectLon.V+0.5), "should get latitude length at the north pole for 0 deg")

}
