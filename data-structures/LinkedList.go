package data_structures

import "fmt"

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

func (l *LinkedList) PrintAllValues() {
	current := l.head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}
