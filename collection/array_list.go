package collection

import (
	"errors"
	"github.com/imulab/go-review/object"
	"math"
)

func NewArrayList() List {
	return newArrayList()
}

func newArrayList() *arrayList {
	return &arrayList{
		array: make([]interface{}, 0, arrayListDefaultCapacity),
		size:  0,
	}
}

const (
	arrayListDefaultCapacity = 10
)

type arrayList struct {
	array []interface{}
	size  int
}

func (l *arrayList) Add(item interface{}) error {
	if err := l.ensureCapacity(l.size + 1); err != nil {
		return errors.New("maximum size reached")
	}
	l.array = append(l.array, item)
	l.size++
	return nil
}

func (l *arrayList) Remove(item interface{}) error {
	index := l.IndexOf(item)
	if index < 0 {
		return errors.New("item not found")
	}
	return l.RemoveIndex(index)
}

func (l *arrayList) RemoveIndex(index int64) error {
	if err := l.checkIndex(index); err != nil {
		return err
	}
	l.array = append(l.array[:index], l.array[index+1:]...)
	l.ensureCapacity(l.size - 1)
	l.size--
	return nil
}

func (l *arrayList) Size() int64 {
	return int64(l.size)
}

func (l *arrayList) IsEmpty() bool {
	return l.size == 0
}

func (l *arrayList) Contains(item interface{}) bool {
	return l.IndexOf(item) >= 0
}

func (l *arrayList) Clear() {
	l.array = make([]interface{}, 0, arrayListDefaultCapacity)
}

func (l *arrayList) ToArray() []interface{} {
	array := make([]interface{}, len(l.array), cap(l.array))
	copy(array, l.array)
	return array
}

func (l *arrayList) Get(index int64) (interface{}, error) {
	if err := l.checkIndex(index); err != nil {
		return nil, err
	}
	return l.array[index], nil
}

func (l *arrayList) Set(index int64, item interface{}) error {
	if err := l.checkIndex(index); err != nil {
		return err
	}
	l.array[index] = item
	return nil
}

func (l *arrayList) IndexOf(item interface{}) int64 {
	for i, each := range l.array {
		var equals bool
		{
			if obj, ok := each.(object.Equality); ok {
				equals = obj.Equals(item)
			} else {
				equals = each == item
			}
		}

		if equals {
			return int64(i)
		}
	}
	return -1
}

func (l *arrayList) checkIndex(index int64) error {
	if index < 0 || index >= int64(l.size) {
		return errors.New("index out of bounds")
	}
	return nil
}

func (l *arrayList) ensureCapacity(newSize int) error {
	if newSize < 0 {
		return errors.New("size overflow")
	}

	var newCap = cap(l.array)
	{
		// shrink
		if newSize <= (cap(l.array) >> 1) {
			// shrink to 0.75
			newCap = cap(l.array) - cap(l.array) >> 2
			// prevent over-shrink
			if newCap < arrayListDefaultCapacity {
				newCap = arrayListDefaultCapacity
			}
		}

		// grow
		if newSize > (cap(l.array) >> 1 + cap(l.array) >> 2) {
			// grow to 1.5
			newCap = cap(l.array) + cap(l.array) >> 1
			// prevent overflow
			if newCap < 0 {
				newCap = math.MaxInt32
			}
		}
	}

	if newCap == cap(l.array) {
		return nil
	}

	newArray := make([]interface{}, len(l.array), newCap)
	copy(newArray, l.array)
	l.array = newArray

	return nil
}
