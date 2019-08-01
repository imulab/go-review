package object

type Comparison int

const (
	ComparedLess    Comparison = -1
	ComparedEqual   Comparison = 0
	ComparedGreater Comparison = 1
)

type Comparable interface {
	// Compare this object to another. If no comparable, return a non-nil error.
	CompareTo(another interface{}) (Comparison, error)
}