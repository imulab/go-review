package collection

type Collection interface {
	// Add an item to a collection. Returns non-nil error if item cannot be added.
	Add(item interface{}) error

	// Remove an item from a collection. Returns non-nil error if removal failed.
	Remove(item interface{}) error

	// Return the size of a collection. When return value is 0, IsEmpty() method should
	// return true.
	Size() int64

	// Express method to check if the collection is empty. When the collection is empty,
	// the Size() method should return 0.
	IsEmpty() bool

	// Check if the collection contains an item. If item implements object.Equality interface,
	// comparison will be done using that interface, otherwise the built-in equality mechanism
	// is used.
	Contains(item interface{}) bool

	// Clears the collection of all items.
	Clear()

	// Copies all items from a collection (in implementation's order) to an array.
	ToArray() []interface{}
}

type List interface {
	// inherits all method from collection
	Collection

	// Get an item from a list. If index is out of bounds, a non-nil error is returned.
	Get(index int64) (interface{}, error)

	// Set index to new item. If index is out of bounds, a non-nil error is returned.
	Set(index int64, item interface{}) error

	// Get index of an item from the list. If item does not exist, -1 will be returned.
	IndexOf(item interface{}) int64
}

type Set interface {
	// inherits all method from collection
	Collection
}