package data_structures

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Prepend(n *Node) {
	temp := l.head
	n.Next = temp
	l.head = n
	l.length++
}

func (l *LinkedList) Append(n *Node) {
	last := l.GetLastNode()
	if last == nil {
		l.head = n
	} else {
		last.Next = n
	}
	l.length++
}

func (l *LinkedList) GetLastNode() *Node {
	current := l.head
	for current != nil && current.Next != nil {
		current = current.Next
	}
	return current
}

// RemoveNth removes the node at the "n"th position.
// This function counts from 1...n
func (l *LinkedList) RemoveNth(n int) (*Node, error) {
	if n > l.length {
		return nil, errors.New("n is larger than LinkedList length")
	}

	if n == 1 {
		temp := l.head
		l.head = l.head.Next
		temp.Next = nil
		l.length = l.length - 1

		return temp, nil
	}

	var previous *Node
	counter := 1

	current := l.head

	for counter < n {
		previous = current
		current = current.Next
		counter = counter + 1
	}

	toReturn := current
	previous.Next = current.Next
	current.Next = nil

	l.length = l.length - 1

	return toReturn, nil
}

func (l *LinkedList) PrintAllValues() {
	current := l.head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

func (l *LinkedList) Reverse() {
	var reversed *Node
	current := l.head
	for current != nil {
		if reversed == nil {
			reversed = current
			current = current.Next
			reversed.Next = nil
		} else {
			temp := current.Next
			current.Next = reversed
			reversed = current
			current = temp
		}
	}
	l.head = reversed
}

func (l *LinkedList) Pop() (*Node, error) {
	var popped *Node

	if l.head == nil {
		return nil, errors.New("no nodes in list")
	}

	temp := l.head.Next
	popped = l.head
	popped.Next = nil
	l.head = temp

	return popped, nil
}

func (l *LinkedList) GetNth(n int) (*Node, error) {
	current := l.head
	index := 0
	for current != nil && index <= n {
		if index == n {
			return current, nil
		}
		current = current.Next
		index = index + 1
	}

	return nil, errors.New("n out of bounds")
}
