package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_Prepend(t *testing.T) {
	// Starting with a new empty list
	l := LinkedList{}
	l.Prepend(&Node{Value: 1})
	assert.Equal(t, 1, l.head.Value)
	assert.Equal(t, 1, l.length)

	// Adding to a list with existing node
	l.Prepend(&Node{Value: 2})
	assert.Equal(t, 2, l.head.Value)
	assert.Equal(t, 2, l.length)

	// Adding to a list with existing nodes
	l.Prepend(&Node{Value: 3})
	assert.Equal(t, 3, l.head.Value)
	assert.Equal(t, 3, l.length)
}

func TestLinkedList_Append(t *testing.T) {
	l := LinkedList{}

	// Add a node to a list without a node
	l.Append(&Node{Value: 1})
	assert.Equal(t, 1, l.head.Value)
	assert.Equal(t, 1, l.length)

	// Adding a node to a list with a node
	l.Append(&Node{Value: 2})
	assert.Equal(t, 2, l.GetLastNode().Value)
	assert.Equal(t, 2, l.length)
}

func TestLinkedList_GetLastNode(t *testing.T) {
	// Starting with a new empty list
	l := LinkedList{}

	// Getting the last node in an empty list
	got := l.GetLastNode()
	assert.Nil(t, got)

	// Adding to a list with no node at the head
	l.Prepend(&Node{Value: 1})
	got = l.GetLastNode()
	assert.Equal(t, 1, got.Value)

	// Adding to a list with existing node
	l.Prepend(&Node{Value: 2})
	got = l.GetLastNode()
	assert.Equal(t, 1, got.Value)
}
