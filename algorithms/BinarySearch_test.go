package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	// BinarySearch when array is of size 1
	array := []int{0}
	index := BinarySearch(array, 0)
	assert.Equal(t, 0, index)

	// BinarySearch when array is of size 2 and value is on left
	//array = []int{0, 1}
	//index = BinarySearch(array, 1)
	//assert.Equal(t, 1, index)

	// BinarySearch when array is of size 2 and value is on right
	//array = []int{0, 1}
	//index = BinarySearch(array, 0)
	//assert.Equal(t, 0, index)

	// BinarySearch when array is of size 3 and value is in the middle
	//array = []int{0, 1, 2}
	//index = BinarySearch(array, 1)
	//assert.Equal(t, 1, index)

	// Return -1 if the value cannot be found
	array = []int{1}
	index = BinarySearch(array, 2)
	assert.Equal(t, -1, index)
}
