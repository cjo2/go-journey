package main

import "github.com/cjo2/go-journey/data-structures"

func main() {
	mylist := data_structures.LinkedList{}
	node1 := &data_structures.LinkedListNode{Value: 48}
	node2 := &data_structures.LinkedListNode{Value: 18}
	node3 := &data_structures.LinkedListNode{Value: 16}
	mylist.Append(node1)
	mylist.Append(node2)
	mylist.Append(node3)
	mylist.PrintAllValues()
}
