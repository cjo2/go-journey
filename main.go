package main

func main() {
	mylist := LinkedList{}
	node1 := &Node{value: 48}
	node2 := &Node{value: 18}
	node3 := &Node{value: 16}
	mylist.Append(node1)
	mylist.Append(node2)
	mylist.Append(node3)
	mylist.PrintAllValues()
}
