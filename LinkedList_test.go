package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_Prepend(t *testing.T) {
	// Starting with a new empty list
	l := LinkedList{}
	l.Prepend(&Node{value: 1})
	assert.Equal(t, 1, l.head.value)
	assert.Equal(t, 1, l.length)

	// Adding to a list with existing node
	l.Prepend(&Node{value: 2})
	assert.Equal(t, 2, l.head.value)
	assert.Equal(t, 2, l.length)

	// Adding to a list with existing nodes
	l.Prepend(&Node{value: 3})
	assert.Equal(t, 3, l.head.value)
	assert.Equal(t, 3, l.length)
}

func TestLinkedList_Append(t *testing.T) {
	l := LinkedList{}

	// Add a node to a list without a node
	l.Append(&Node{value: 1})
	assert.Equal(t, 1, l.head.value)
	assert.Equal(t, 1, l.length)

	// Adding a node to a list with a node
	l.Append(&Node{value: 2})
	assert.Equal(t, 2, l.GetLastNode().value)
	assert.Equal(t, 2, l.length)
}

func TestLinkedList_GetLastNode(t *testing.T) {
	// Starting with a new empty list
	l := LinkedList{}

	// Getting the last node in an empty list
	got := l.GetLastNode()
	assert.Nil(t, got)

	// Adding to a list with no node at the head
	l.Prepend(&Node{value: 1})
	got = l.GetLastNode()
	assert.Equal(t, 1, got.value)

	// Adding to a list with existing node
	l.Prepend(&Node{value: 2})
	got = l.GetLastNode()
	assert.Equal(t, 1, got.value)
}
