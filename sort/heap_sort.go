package sort

import (
	"errors"
	"github.com/imulab/go-review/object"
)

func HeapSort(target []interface{}, direction Direction) error {
	for _, each := range target {
		if _, ok := each.(object.Comparable); !ok {
			return errors.New("elements not comparable")
		}
	}

	if len(target) <= 1 {
		return nil
	}

	(&heapSort{N: len(target), dir: !direction}).Sort(target, direction)

	return nil
}

type heapSort struct {
	// the effective size of the heap
	N int

	// heap direction, reverse of sort direction
	dir Direction
}

func (s *heapSort) Sort(target []interface{}, _ Direction) {
	s.N = len(target)

	// construct heap
	for i := s.N / 2; i >= 0; i-- {
		s.demote(target, i, s.dir)
	}

	// put each in place
	for s.N > 0 {
		swap(target, 0, s.N - 1)
		s.N--
		s.demote(target, 0, s.dir)
	}
}

func (s *heapSort) promote(target []interface{}, i int, direction Direction) {
	for i != s.parent(i) {
		switch direction {
		case Ascending:
			if less(target[i], target[s.parent(i)]) {
				swap(target, i, s.parent(i))
			}
		case Descending:
			if greater(target[i], target[s.parent(i)]) {
				swap(target, i, s.parent(i))
			}
		default:
			panic("impossible direction")
		}
		i = s.parent(i)
	}
}

func (s *heapSort) demote(target []interface{}, i int, direction Direction) {
	for {
		candidate := s.leftChild(i)
		{
			if candidate >= s.N {
				return
			}

			if s.rightSibling(candidate) < s.N {
				switch direction {
				case Ascending:
					if less(target[s.rightSibling(candidate)], target[candidate]) {
						candidate = s.rightSibling(candidate)
					}
				case Descending:
					if greater(target[s.rightSibling(candidate)], target[candidate]) {
						candidate = s.rightSibling(candidate)
					}
				default:
					panic("impossible direction")
				}
			}
		}

		switch direction {
		case Ascending:
			if less(target[candidate], target[i]) {
				swap(target, candidate, i)
				i = candidate
			} else {
				return
			}
		case Descending:
			if greater(target[candidate], target[i]) {
				swap(target, candidate, i)
				i = candidate
			} else {
				return
			}
		default:
			panic("impossible direction")
		}
	}
}

func (s *heapSort) leftChild(i int) int {
	return i * 2 + 1
}

func (s *heapSort) rightChild(i int) int {
	return i * 2 + 2
}

func (s *heapSort) rightSibling(i int) int {
	return i + 1
}

func (s *heapSort) parent(i int) int {
	return (i - 1) / 2
}