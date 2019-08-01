package object

import (
	"errors"
	"fmt"
)

type Int int

func (i Int) CompareTo(another interface{}) (Comparison, error) {
	if j, ok := another.(Int); !ok {
		return 0, errors.New("not comparable")
	} else {
		if i < j {
			return ComparedLess, nil
		} else if i > j {
			return ComparedGreater, nil
		} else {
			return ComparedEqual, nil
		}
	}
}

func (i Int) String() string {
	return fmt.Sprintf("%d", i)
}

