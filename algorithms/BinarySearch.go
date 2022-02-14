package algorithms

// BinarySearch executes binary search on a sorted integer array with unique values
// BinarySearch will return the index of the integer or -1 if the integer cannot be found
func BinarySearch(array []int, value int) int {
	left := 0
	right := len(array) - 1

	for left <= right {
		middle := (right + left) / 2
		currentVal := array[middle]

		if currentVal == value {
			return middle
		} else if currentVal < value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}
