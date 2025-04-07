package tree

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

func Levelorder(root *Node) [][]int {
	result := [][]int{}

	queue := []*Node{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Data)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}
