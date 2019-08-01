package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedListQueue(t *testing.T) {
	q := NewLinkedListQueue()
	for i := 0; i < 20; i++ {
		q.Enqueue(i)
		assert.Equal(t, int64(i) + 1, q.Size())
	}

	for i := 0; i < 20; i++ {
		item := q.Dequeue()
		assert.Equal(t, i, item)
		assert.Equal(t, 20-1-int64(i), q.Size())
	}

	assert.True(t, q.IsEmpty())
}
