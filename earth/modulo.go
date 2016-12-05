package earth

import "math"

// Modulo compute the modulo keeping result positive
func Modulo(x, y int) int {
	x = int(math.Mod(float64(x), float64(y))) + y
	return int(math.Mod(float64(x), float64(y)))
}
