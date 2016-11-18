package wind

import "image"

//Map is a 2D maps containing winds
type Map struct {
	Width  int // Map width
	Height int // Map Height
	CellW  int // Cell Width in minutes
	CellH  int // Cell Height in minutes
	Data   []Speed
}

// Speed is the speed of the wind (m/sec)
type Speed struct {
	speedU float32
	speedV float32
}

// SetWind set a wind speed
func (buffer Map) SetWind(loc image.Point, speed Speed) {
	buffer.Data[loc.Y*buffer.Width+loc.X] = speed
}

// GetWind get a wind speed
func (buffer Map) GetWind(loc image.Point) Speed {
	return buffer.Data[loc.Y*buffer.Width+loc.X]
}
