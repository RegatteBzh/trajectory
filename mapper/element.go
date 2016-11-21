package mapper

//Element is an element in a map
type Element interface {
	Compare(b Element) int
}
