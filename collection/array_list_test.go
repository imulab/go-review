package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayList(t *testing.T) {
	l := NewArrayList()
	for i := 0; i < 20; i++ {
		assert.Nil(t, l.Add(i))
	}

	assert.Equal(t, int64(20), l.Size())

	for i := 0; i < 20; i++ {
		assert.Equal(t, int64(i), l.IndexOf(i))

		item, err := l.Get(int64(i))
		assert.Nil(t, err)
		assert.Equal(t, i, item)
	}

	for i := 0; i < 20; i++ {
		assert.Nil(t, l.Remove(i))
	}
}
