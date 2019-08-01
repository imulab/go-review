package sort

import "github.com/imulab/go-review/object"

type Sort interface {
	// Sort the given target in place with direction.
	Sort(target []interface{}, direction Direction)
}

type Direction bool

const (
	Ascending  Direction = true
	Descending Direction = false
)

func IsSorted(target []interface{}, direction Direction) bool {
	if len(target) <= 1 {
		return true
	}

	for i := 1; i < len(target); i++ {
		switch direction {
		case Ascending:
			if greater(target[i - 1], target[i]) {
				return false
			}
		case Descending:
			if less(target[i - 1], target[i]) {
				return false
			}
		default:
			panic("impossible direction")
		}
	}

	return true
}

func less(a, b interface{}) bool {
	if i, ok1 := a.(object.Comparable); ok1 {
		if j, ok2 := b.(object.Comparable); ok2 {
			r, _ := i.CompareTo(j)
			return r == object.ComparedLess
		}
	}
	panic("not comparable")
}

func greater(a, b interface{}) bool {
	if i, ok1 := a.(object.Comparable); ok1 {
		if j, ok2 := b.(object.Comparable); ok2 {
			r, _ := i.CompareTo(j)
			return r == object.ComparedGreater
		}
	}
	panic("not comparable")
}

func swap(target []interface{}, i, j int) {
	temp := target[i]
	target[i] = target[j]
	target[j] = temp
}