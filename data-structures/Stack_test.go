package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := &Stack{}

	assert.Equal(t, 0, len(stack.Items))
	stack.Push(1)
	stack.Push(2)
	assert.Equal(t, 1, stack.Items[0])
	assert.Equal(t, 2, stack.Items[1])
}

func TestStack_Pop(t *testing.T) {
	stack := &Stack{}

	assert.Equal(t, 0, len(stack.Items))
	stack.Push(1)
	stack.Push(2)
	val := stack.Pop()
	assert.Equal(t, 2, val)
}
