package sort

import (
	"github.com/imulab/go-review/object"
	"github.com/pkg/errors"
)

func MergeSort(target []interface{}, direction Direction) error {
	for _, each := range target {
		if _, ok := each.(object.Comparable); !ok {
			return errors.New("elements not comparable")
		}
	}

	if len(target) <= 1 {
		return nil
	}

	mergeSort(0).Sort(target, direction)

	return nil
}

type mergeSort int

func (s mergeSort) Sort(target []interface{}, direction Direction) {
	s.sort(target, 0, len(target) - 1, direction)
}

func (s mergeSort) sort(target []interface{}, lowInclusive int, highInclusive int, direction Direction) {
	if highInclusive <= lowInclusive {
		return
	}

	mid := lowInclusive + (highInclusive - lowInclusive) / 2

	s.sort(target, lowInclusive, mid, direction)
	s.sort(target, mid+1, highInclusive, direction)
	s.merge(target, lowInclusive, mid, highInclusive, direction)
}

func (s mergeSort) merge(target []interface{}, lowInclusive int, mid int, highInclusive int, direction Direction) {
	first := s.copy(target, lowInclusive, mid)
	second := s.copy(target, mid+1, highInclusive)
	i, j := 0, 0

	for k := lowInclusive; k <= highInclusive; k++ {
		if i >= len(first) {
			target[k] = second[j]
			j++
			continue
		}

		if j >= len(second) {
			target[k] = first[i]
			i++
			continue
		}

		if i < len(first) && j < len(second) {
			switch direction {
			case Ascending:
				if less(first[i], second[j]) {
					target[k] = first[i]
					i++
				} else {
					target[k] = second[j]
					j++
				}
			case Descending:
				if greater(first[i], second[j]) {
					target[k] = first[i]
					i++
				} else {
					target[k] = second[j]
					j++
				}
			default:
				panic("impossible direction")
			}
		}
	}
}

func (s mergeSort) copy(target []interface{}, lowInclusive int, highInclusive int) []interface{} {
	if highInclusive < lowInclusive {
		return []interface{}{}
	}

	c := make([]interface{}, 0, highInclusive - lowInclusive + 1)
	for i := lowInclusive; i <= highInclusive; i++ {
		c = append(c, target[i])
	}
	return c
}

