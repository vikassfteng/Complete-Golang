package main

import (
	"fmt"
)


type Node struct{
	Data int
	Left *Node
	Right *Node
}


func NewNode(val int) *Node{
	return &Node{
		Data: val,
		Left: nil,
		Right: nil,
	}
}


func main(){

	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Right = NewNode(4)


	fmt.Println("root : ", root.Data)
	fmt.Println("Left child : ", root.Left.Data)
	fmt.Println("Right child : ", root.Right.Data)
	fmt.Println("Rgiht->Left : ", root.Left.Right.Data)
}
