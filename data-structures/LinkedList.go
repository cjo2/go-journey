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
	Head   *Node
	length int
}

func (l *LinkedList) Prepend(n *Node) {
	temp := l.Head
	n.Next = temp
	l.Head = n
	l.length++
}

func (l *LinkedList) Append(n *Node) {
	last := l.GetLastNode()
	if last == nil {
		l.Head = n
	} else {
		last.Next = n
	}
	l.length++
}

func (l *LinkedList) GetLastNode() *Node {
	current := l.Head
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
		temp := l.Head
		l.Head = l.Head.Next
		temp.Next = nil
		l.length = l.length - 1

		return temp, nil
	}

	var previous *Node
	counter := 1

	current := l.Head

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
	current := l.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

func (l *LinkedList) Reverse() {
	var reversed *Node
	current := l.Head
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
	l.Head = reversed
}

func (l *LinkedList) Pop() (*Node, error) {
	var popped *Node

	if l.Head == nil {
		return nil, errors.New("no nodes in list")
	}

	temp := l.Head.Next
	popped = l.Head
	popped.Next = nil
	l.Head = temp

	return popped, nil
}

func (l *LinkedList) GetNth(n int) (*Node, error) {
	current := l.Head
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
