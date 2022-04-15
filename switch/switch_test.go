package _switch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEvenNumbers(t *testing.T) {
	allEvens := []int{2, 4, 6, 8, 10, 12}
	evens, oddCount := GetEvenNumbers(allEvens)
	assert.Equal(t, 6, len(evens))
	assert.Equal(t, 0, oddCount)

	allOdds := []int{1, 3, 5, 15, 25}
	evens, oddCount = GetEvenNumbers(allOdds)
	assert.Equal(t, 0, len(evens))
	assert.Equal(t, 5, oddCount)

	mix := []int{1, 2, 3, 4, 5}
	evens, oddCount = GetEvenNumbers(mix)
	assert.Equal(t, 2, len(evens))
	assert.Equal(t, 3, oddCount)
}

func TestPlane_Navigate(t *testing.T) {
	plane := &Plane{}
	plane.Navigate("up")
	assert.Equal(t, 1, plane.XPosition)
	plane.Navigate("down")
	assert.Equal(t, 0, plane.XPosition)
	plane.Navigate("left")
	assert.Equal(t, 0, plane.XPosition)
	assert.Equal(t, -1, plane.YPosition)
	plane.Navigate("right")
	assert.Equal(t, 100, plane.XPosition)
	assert.Equal(t, 0, plane.YPosition)
}
