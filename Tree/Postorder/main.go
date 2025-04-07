package main

import (
	"fmt"
	"postorder/tree"
)

func main() {

	root := tree.NewNode(1)
	root.Left = tree.NewNode(2)
	root.Right = tree.NewNode(3)
	root.Left.Right = tree.NewNode(4)

	fmt.Println("Postorder traversal starting")
	tree.Postorder(root)

}
