package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	left := 1 + maxDepth(root.Left)
	right := 1 + maxDepth(root.Right)

	if left > right {
		return left
	} else {
		return right
	}
}

func main() {
	node := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println("The maximum depth of the tree is:", maxDepth(node))
}
