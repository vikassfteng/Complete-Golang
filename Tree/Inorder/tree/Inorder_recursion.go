package tree

import (
	"fmt"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func NewNode(val int) *Node {
	return &Node{
		Data: val,
	}
}

func Inorder(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.Data, " ")
	Inorder(root.Left)
	Inorder(root.Right)
}
