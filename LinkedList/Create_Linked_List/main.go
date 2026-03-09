package main

import "fmt"

// Node represents a node in a linked list
type Node struct {
	Val  int
	Next *Node
}

// PrintLinkedList prints a linked list in the format val1->val2->val3
// and returns the head pointer (following Go conventions for chainable methods)
func PrintLinkedList(head *Node) *Node {
	if head == nil {
		return head
	}

	current := head
	for current != nil {
		fmt.Print(current.Val)
		if current.Next != nil {
			fmt.Print("->")
		}
		current = current.Next
	}
	fmt.Println() // Add newline for better output formatting
	return head
}

func deleteHead(head *Node) *Node {
	if head == nil {
		return nil
	}
	return head.Next
}

func deleteTail(head *Node) *Node {
	if head == nil {
		return head
	}
	// If there's only one node, delete it
	if head.Next == nil {
		return nil
	}
	current := head
	// Find the second-to-last node
	for current.Next.Next != nil {
		current = current.Next
	}
	current.Next = nil
	return head
}
func deleteAtPosition(head *Node, position int) *Node {
	if head == nil || position < 1 {
		return head
	}
	// If deleting the first position (head)
	if position == 1 {
		return head.Next
	}
	current := head
	// Navigate to the node before the one to delete
	for i := 1; i < position-1 && current.Next != nil; i++ {
		current = current.Next
	}
	// Check if the position exists
	if current.Next != nil {
		current.Next = current.Next.Next
	}
	return head
}
func deleteValue(head *Node, value int) *Node {
	if head == nil {
		return head
	}
	if head.Val == value {
		return head.Next
	}
	current := head
	for current.Next != nil {
		if current.Next.Val == value {
			current.Next = current.Next.Next
			break
		}
		current = current.Next
	}
	return head
}

func insertAtHead(head *Node, val int) *Node {
	current := &Node{Val: val}
	current.Next = head
	head = current
	return head
}
func insertAtPos(head *Node, k, val int) *Node {
	if head == nil {
		if k == 1 {
			newNode := &Node{Val: val}
			return newNode
		} else {
			return head
		}
	}
	if k == 1 {
		newNode := &Node{Val: val}
		newNode.Next = head.Next
		head.Next = newNode
		return head
	}
	current := head
	cnt := 0
	for current != nil {
		cnt++
		if cnt == k-1 {
			newNode := &Node{Val: val}
			newNode.Next = current.Next
			current.Next = newNode
			break
		}
		current = current.Next
	}
	return head
}
func main() {
	head := &Node{Val: 0}
	current := head

	for i := 1; i < 5; i++ {
		current.Next = &Node{Val: i}
		current = current.Next
	}

	PrintLinkedList(head)
	// head = deleteTail(head)
	// PrintLinkedList(head)
	// head = deleteAtPosition(head, 1)
	// PrintLinkedList(head)
	// head = deleteAtPosition(head, 2)
	// PrintLinkedList(head)
	// head = deleteValue(head, 0)
	// PrintLinkedList(head)
	// head = insertAtHead(head, 9)
	// PrintLinkedList(head)
	head = insertAtPos(head, 1, 9)
	PrintLinkedList(head)

}
