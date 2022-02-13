package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Prepend(n *Node) {
	temp := l.head
	n.next = temp
	l.head = n
	l.length++
}

func (l *LinkedList) Append(n *Node) {
	last := l.GetLastNode()
	if last == nil {
		l.head = n
	} else {
		last.next = n
	}
}

func (l *LinkedList) GetLastNode() *Node {
	current := l.head
	for {
		if current == nil || current.next == nil {
			return current
		}
		current = current.next
	}
}

func (l *LinkedList) PrintAllValues() {
	current := l.head
	for {
		if current == nil {
			return
		}
		fmt.Println(current.value)
		current = current.next
	}
}
