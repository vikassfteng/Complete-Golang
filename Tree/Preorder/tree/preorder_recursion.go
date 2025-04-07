package tree

import "fmt"

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

func Preorder(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.Data, " ")
	Preorder(root.Left)
	Preorder(root.Right)
}
