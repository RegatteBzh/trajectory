package mapper

//Element is an element in a map
type Element interface {
	Compare(a Element, b Element) int
}
