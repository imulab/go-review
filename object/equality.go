package object

type Equality interface {
	// Returns true if this object equals another object.
	Equals(another interface{}) bool
}
