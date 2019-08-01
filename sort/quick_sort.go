package sort

import (
	"errors"
	"github.com/imulab/go-review/object"
)

func QuickSort(target []interface{}, direction Direction) error {
	for _, each := range target {
		if _, ok := each.(object.Comparable); !ok {
			return errors.New("elements not comparable")
		}
	}

	if len(target) <= 1 {
		return nil
	}

	quickSort(0).Sort(target, direction)

	return nil
}

type quickSort int

func (s quickSort) Sort(target []interface{}, direction Direction) {
	s.sort(target, 0, len(target) - 1, direction)
}

func (s quickSort) sort(target []interface{}, lowInclusive int, highInclusive int, direction Direction) {
	if highInclusive <= lowInclusive {
		return
	}

	k := s.partition(target, lowInclusive, highInclusive, direction)
	s.sort(target, lowInclusive, k-1, direction)
	s.sort(target, k+1, highInclusive, direction)
}

func (s quickSort) partition(target []interface{}, lowInclusive int, highInclusive int, direction Direction) int {
	i, j := lowInclusive, highInclusive

	for {
		switch direction {
		case Ascending:
			for less(target[i], target[lowInclusive]) {
				i++
			}
			for greater(target[j], target[lowInclusive]) {
				j--
			}
		case Descending:
			for greater(target[i], target[lowInclusive]) {
				i++
			}
			for less(target[j], target[lowInclusive]) {
				j--
			}

		}

		if j <= i {
			break
		}

		swap(target, i, j)
	}

	swap(target, lowInclusive, j)
	return j
}