package mapper

import (
	"image"

	"github.com/regattebzh/trajectory/mapper"
)

// Map is a map
type Map struct {
	Width  int // Map width
	Height int // Map Height
	CellW  int // Cell Width in minutes
	CellH  int // Cell Height in minutes
	Data   []Element
	Max    Element
	Min    Element
}

// Set sets a map value
func (buffer Map) Set(loc image.Point, element Element) {
	buffer.Data[loc.Y*buffer.Width+loc.X] = element
}

// Get gets a map value
func (buffer Map) Get(loc image.Point) Element {
	return buffer.Data[loc.Y*buffer.Width+loc.X]
}

// GetMax gets the maximum value
func (buffer Map) GetMax() Element {
	var max Element
	for _, w := range buffer.Data {
		comp, ok := w.Compare(max)
		if ok && comp > 0 {
			max = w
		}
	}
	return max
}

//GetMin gets the minimum value
func (buffer Map) GetMin() Element {
	var min Element
	for _, w := range buffer.Data {
		comp, ok := w.Compare(min)
		if ok && comp < 0 {
			min = w
		}
	}
	return min
}

// ComputeParameters compute diff, min and max
func (buffer Map) ComputeParameters() {
	buffer.Max = buffer.GetMax()
	buffer.Min = buffer.GetMin()
}

// New create a new Mapper
func New(r image.Rectangle, cellW int, cellH int) {
	return mapper.Map{
		Width:  r.Dx(),
		Height: r.Dy(),
		Data:   make([]mapper.Element, r.Dx()*r.Dy()),
		CellH:  cellH,
		CellW:  cellW,
	}
}
