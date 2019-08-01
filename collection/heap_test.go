package collection

import (
	"fmt"
	"github.com/imulab/go-review/object"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinPriorityQueue(t *testing.T) {
	pq := NewMinPriorityQueue()

	// insert 1 to 10
	for _, i := range []TestInt{3, 9, 6, 10, 1, 4, 2, 7, 8, 5} {
		pq.Enqueue(i)
	}

	assert.Equal(t, int64(10), pq.Size())

	for _, i := range []TestInt{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		assert.Equal(t, i, pq.Dequeue())
	}

	assert.True(t, pq.IsEmpty())
}

func TestMaxPriorityQueue(t *testing.T) {
	pq := NewMaxPriorityQueue()

	// insert 1 to 10
	for _, i := range []TestInt{3, 9, 6, 10, 1, 4, 2, 7, 8, 5} {
		pq.Enqueue(i)
	}

	assert.Equal(t, int64(10), pq.Size())

	for _, i := range []TestInt{10, 9, 8, 7, 6, 5, 4, 3, 2, 1} {
		assert.Equal(t, i, pq.Dequeue())
	}

	assert.True(t, pq.IsEmpty())
}

type TestInt int

func (i TestInt) CompareTo(another interface{}) (object.Comparison, error) {
	if j, ok := another.(TestInt); !ok {
		return 0, errors.New("not comparable")
	} else {
		if i < j {
			return object.ComparedLess, nil
		} else if i > j {
			return object.ComparedGreater, nil
		} else {
			return object.ComparedEqual, nil
		}
	}
}

func (i TestInt) String() string {
	return fmt.Sprintf("%d", i)
}
