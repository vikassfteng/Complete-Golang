package main

import (
	"fmt"
	"levelorder/tree"
)

func main() {
	root := tree.NewNode(1)
	root.Left = tree.NewNode(2)
	root.Right = tree.NewNode(3)
	root.Left.Left = tree.NewNode(4)
	root.Left.Right = tree.NewNode(5)
	root.Right.Left = tree.NewNode(6)
	root.Right.Right = tree.NewNode(7)

	fmt.Println("Level Order Traversal:")
	result := tree.Levelorder(root)
	for _, level := range result {
		fmt.Println(level)
	}
}
