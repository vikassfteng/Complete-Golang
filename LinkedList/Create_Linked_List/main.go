package main

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func NewNode(val int) *Node {
	return &Node{
		value: val,
		next:  nil,
		prev:  nil,
	}
}

func PrintLL(head *Node) {
	current := head
	for current != nil {
		fmt.Print(current.value)
		if current.next != nil {
			fmt.Print("->")
		}
		current = current.next
	}
}

func PrintReverse(tail *Node) {
	current := tail
	for current != nil {
		fmt.Print(current.value)
		if current.prev != nil {
			fmt.Print("->")
		}
		current = current.prev
	}
}

func main() {

	node := NewNode(0)
	current := node
	temp := current

	for i := 1; i < 5; i++ {
		current.next = NewNode(i)
		current = current.next
		current.prev = temp
		temp = current
	}
	PrintLL(node)
	fmt.Println("\n")
	PrintReverse(current)
}
