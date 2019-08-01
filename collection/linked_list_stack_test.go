package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedListStack(t *testing.T) {
	s := NewLinkedListStack()
	for i := 0; i < 20; i++ {
		s.Push(i)
		assert.Equal(t, int64(i) + 1, s.Size())
	}

	for i := 19; i >= 0; i-- {
		item := s.Pop()
		assert.Equal(t, i, item)
		assert.Equal(t, int64(i), s.Size())
	}

	assert.True(t, s.IsEmpty())
}
