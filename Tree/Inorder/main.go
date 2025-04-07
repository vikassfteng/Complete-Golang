package main

import (
	"fmt"
	"inorder/tree"
)

func main() {
	root := tree.NewNode(1)
	root.Left = tree.NewNode(2)
	root.Right = tree.NewNode(3)
	root.Left.Right = tree.NewNode(4)
	root.Left.Left = tree.NewNode(5)
	root.Right.Left = tree.NewNode(6)
	root.Right.Right = tree.NewNode(7)

	fmt.Print("Inorder traversal: ")
	tree.Inorder(root)
}