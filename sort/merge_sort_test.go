package sort

import (
	"github.com/imulab/go-review/object"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	target := []interface{}{
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
	assert.Nil(t, MergeSort(target, Ascending))
	assert.True(t, IsSorted(target, Ascending))
}
