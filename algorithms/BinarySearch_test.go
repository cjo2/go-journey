package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	// BinarySearch when slice is of size 1
	slice := []int{0}
	index := BinarySearch(slice, 0)
	assert.Equal(t, 0, index)

	// BinarySearch when slice is of size 2 and value is on left
	slice = []int{0, 1}
	index = BinarySearch(slice, 1)
	assert.Equal(t, 1, index)

	// BinarySearch when slice is of size 2 and value is on right
	slice = []int{0, 1}
	index = BinarySearch(slice, 0)
	assert.Equal(t, 0, index)

	// BinarySearch when slice is of size 3 and value is in the middle
	slice = []int{0, 1, 2}
	index = BinarySearch(slice, 1)
	assert.Equal(t, 1, index)

	// Return -1 if the value cannot be found
	slice = []int{1}
	index = BinarySearch(slice, 2)
	assert.Equal(t, -1, index)
}
