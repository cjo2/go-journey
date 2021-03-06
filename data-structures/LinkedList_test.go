package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_Prepend(t *testing.T) {
	// Starting with a new empty list
	l := LinkedList{}
	l.Prepend(&LinkedListNode{Value: 1})
	assert.Equal(t, 1, l.Head.Value)
	assert.Equal(t, 1, l.length)

	// Adding to a list with existing node
	l.Prepend(&LinkedListNode{Value: 2})
	assert.Equal(t, 2, l.Head.Value)
	assert.Equal(t, 2, l.length)

	// Adding to a list with existing nodes
	l.Prepend(&LinkedListNode{Value: 3})
	assert.Equal(t, 3, l.Head.Value)
	assert.Equal(t, 3, l.length)
}

func TestLinkedList_Append(t *testing.T) {
	l := LinkedList{}

	// Add a node to a list without a node
	l.Append(&LinkedListNode{Value: 1})
	assert.Equal(t, 1, l.Head.Value)
	assert.Equal(t, 1, l.length)

	// Adding a node to a list with a node
	l.Append(&LinkedListNode{Value: 2})
	assert.Equal(t, 2, l.GetLastNode().Value)
	assert.Equal(t, 2, l.length)
}

func TestLinkedList_GetLastNode(t *testing.T) {
	// Starting with a new empty list
	l := LinkedList{}

	// Getting the last node in an empty list
	got := l.GetLastNode()
	assert.Nil(t, got)

	// Adding to a list with no node at the Head
	l.Prepend(&LinkedListNode{Value: 1})
	got = l.GetLastNode()
	assert.Equal(t, 1, got.Value)

	// Adding to a list with existing node
	l.Prepend(&LinkedListNode{Value: 2})
	got = l.GetLastNode()
	assert.Equal(t, 1, got.Value)
}

func TestLinkedList_RemoveNth(t *testing.T) {
	l := LinkedList{}

	// Attempt to remove from an empty LinkedList and return an error
	removed, err := l.RemoveNth(1)
	assert.Error(t, err)
	assert.Equal(t, 0, l.length)

	// Add one node, then remove it
	l.Append(&LinkedListNode{Value: 1})
	assert.Equal(t, 1, l.Head.Value)
	assert.Equal(t, 1, l.length)
	removed, err = l.RemoveNth(1)
	assert.Nil(t, l.Head)
	assert.Nil(t, err)
	assert.Equal(t, 0, l.length)
	assert.Equal(t, 1, removed.Value)

	// Add two nodes, then remove the first one
	l = LinkedList{}
	l.Append(&LinkedListNode{Value: 1})
	l.Append(&LinkedListNode{Value: 2})
	removed, err = l.RemoveNth(1)
	assert.Equal(t, 1, removed.Value)
	assert.Equal(t, 2, l.Head.Value)
	assert.Equal(t, 1, l.length)

	// Add two nodes, then remove the last one
	l = LinkedList{}
	l.Append(&LinkedListNode{Value: 1})
	l.Append(&LinkedListNode{Value: 2})
	removed, err = l.RemoveNth(2)
	assert.Equal(t, 2, removed.Value)
	assert.Equal(t, 1, l.Head.Value)
	assert.Equal(t, 1, l.length)

	// Add three nodes, then remove the second one
	l = LinkedList{}
	l.Append(&LinkedListNode{Value: 1})
	l.Append(&LinkedListNode{Value: 2})
	l.Append(&LinkedListNode{Value: 3})
	removed, err = l.RemoveNth(2)
	assert.Equal(t, 2, removed.Value)
	assert.Equal(t, 1, l.Head.Value)
	assert.Equal(t, 2, l.length)
	assert.Equal(t, 3, l.GetLastNode().Value)
}

func TestLinkedList_Reverse(t *testing.T) {
	l := &LinkedList{}
	l.Append(&LinkedListNode{Value: 1})
	l.Append(&LinkedListNode{Value: 2})
	l.Append(&LinkedListNode{Value: 3})
	l.Append(&LinkedListNode{Value: 4})

	headNode := l.Head

	assert.Equal(t, 1, headNode.Value)
	assert.Equal(t, 2, headNode.Next.Value)
	assert.Equal(t, 3, headNode.Next.Next.Value)
	assert.Equal(t, 4, headNode.Next.Next.Next.Value)

	l.Reverse()

	// headNode points to the last LinkedListNode after calling reverse
	// so set headNode to the new Head
	headNode = l.Head

	assert.Equal(t, 4, headNode.Value)
	assert.Equal(t, 3, headNode.Next.Value)
	assert.Equal(t, 2, headNode.Next.Next.Value)
	assert.Equal(t, 1, headNode.Next.Next.Next.Value)
}

func TestLinkedList_Pop(t *testing.T) {
	l := &LinkedList{}
	l.Append(&LinkedListNode{Value: 1})
	l.Append(&LinkedListNode{Value: 2})

	popped, _ := l.Pop()
	assert.Equal(t, 1, popped.Value)
	assert.Equal(t, 2, l.Head.Value)

	popped, _ = l.Pop()
	assert.Equal(t, 2, popped.Value)

	popped, err := l.Pop()
	assert.Error(t, err)
	assert.Nil(t, popped)

}

func TestLinkedList_GetNth(t *testing.T) {
	l := &LinkedList{}
	l.Append(&LinkedListNode{Value: 1})
	l.Append(&LinkedListNode{Value: 2})
	l.Append(&LinkedListNode{Value: 3})
	l.Append(&LinkedListNode{Value: 4})

	node, err := l.GetNth(0)
	assert.Equal(t, 1, node.Value)
	node, err = l.GetNth(1)
	assert.Equal(t, 2, node.Value)
	node, err = l.GetNth(2)
	assert.Equal(t, 3, node.Value)
	node, err = l.GetNth(3)
	assert.Equal(t, 4, node.Value)
	node, err = l.GetNth(4)
	assert.Error(t, err)
}
