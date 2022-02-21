package algorithms

import (
	datastructures "github.com/cjo2/go-journey/data-structures"
)

// Rearrange updates the LinkedList in-place
func Rearrange(l datastructures.LinkedList) {
	var rearranged *datastructures.Node
	fast := l.Head
	slow := l.Head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	secondList := slow.Next
	slow.Next = nil
	secondList = reverse(secondList)
	firstList := l.Head

	var current *datastructures.Node

	for firstList != nil && secondList != nil {
		if rearranged == nil {
			temp := firstList.Next
			insert := firstList
			insert.Next = nil
			rearranged = insert
			firstList = temp
			current = rearranged
		} else {
			temp := firstList.Next
			insert := firstList
			insert.Next = nil
			firstList = temp
			current.Next = insert
			current = current.Next
		}

		temp := secondList.Next
		insert := secondList
		insert.Next = nil
		current.Next = insert
		secondList = temp
		current = current.Next
	}

	if firstList != nil {
		current.Next = firstList
	}

	l.Head = rearranged
}

func reverse(node *datastructures.Node) *datastructures.Node {
	current := node
	var reversed *datastructures.Node
	for current != nil {
		temp := current.Next
		current.Next = reversed
		reversed = current
		current = temp
	}

	return reversed
}
