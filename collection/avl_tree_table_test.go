package collection

import (
	"github.com/imulab/go-review/object"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAVLTreeTable(t *testing.T) {
	table := NewAVLTreeTable()
	data := []object.Comparable{
		object.Int(4),
		object.Int(6),
		object.Int(1),
		object.Int(7),
		object.Int(2),
		object.Int(3),
		object.Int(9),
		object.Int(8),
		object.Int(10),
		object.Int(5),
	}

	for i, each := range data {
		table.Put(each, each)
		assert.Equal(t, each, table.Get(each))
		assert.Equal(t, int64(i+1), table.Size())
	}

	for i, each := range data {
		table.Del(each)
		assert.Nil(t, table.Get(each))
		assert.Equal(t, int64(len(data)-1-i), table.Size())
	}

	assert.True(t, table.IsEmpty())
}
