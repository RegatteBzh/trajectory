package mapper

import (
	"errors"
	"image"

	"github.com/regattebzh/trajectory/earth"
)

// Map is a map
type Map struct {
	Rect  image.Rectangle
	CellW int // Cell Width in minutes
	CellH int // Cell Height in minutes
	Data  [][]Element
	Max   Element
	Min   Element
}

// NewRectangle create a rectangle on the trigonometric earth
func NewRectangle(x0, y0, x1, y1 int) image.Rectangle {
	return image.Rect(
		earth.Modulo(x0, 360),
		earth.Modulo(y0, 180),
		earth.Modulo(x1, 360),
		earth.Modulo(y1, 180),
	)
}

// Left put a map at the left of the other
func (buffer Map) Left(right Map) (result Map, err error) {
	if buffer.Rect.Dy() != right.Rect.Dy() {
		err = errors.New("Both rectangle do not have the same height")
		return
	}

	data := [][]Element{}
	for _, column := range buffer.Data {
		data = append(data, column)
	}
	for _, column := range right.Data {
		data = append(data, column)
	}

	result = Map{
		Rect: image.Rect(
			buffer.Rect.Min.X,
			buffer.Rect.Min.Y,
			buffer.Rect.Max.X+right.Rect.Dx(),
			buffer.Rect.Max.Y,
		),
		Data:  data,
		CellH: buffer.CellH,
		CellW: buffer.CellW,
	}
	return
}

// Set sets a map value
func (buffer Map) Set(loc image.Point, element Element) {
	buffer.Data[loc.X][loc.Y] = element
}

// Get gets a map value
func (buffer Map) Get(loc image.Point) Element {
	return buffer.Data[loc.X][loc.Y]
}

// GetMax gets the maximum value
func (buffer Map) GetMax() Element {
	var max Element
	for _, column := range buffer.Data {
		for _, w := range column {
			comp, ok := w.Compare(max)
			if ok && comp > 0 {
				max = w
			}
		}
	}
	return max
}

//GetMin gets the minimum value
func (buffer Map) GetMin() Element {
	var min Element
	for _, column := range buffer.Data {
		for _, w := range column {
			comp, ok := w.Compare(min)
			if ok && comp < 0 {
				min = w
			}
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
func New(r image.Rectangle, cellW int, cellH int) Map {
	data := [][]Element{}
	for i := 0; i < r.Dx(); i++ {
		data = append(data, make([]Element, r.Dy()))
	}

	return Map{
		Rect:  r,
		Data:  data,
		CellH: cellH,
		CellW: cellW,
	}
}
