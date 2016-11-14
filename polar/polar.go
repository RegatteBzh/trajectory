package polar

// Polar is a polar for a given sail and a given angle
type Polar struct {
	Angle float64
	speed []float64
}

// SailCharacteristic is the characteristic of a sail
type SailCharacteristic struct {
	Winds  []float64
	Polars []Polar
}
