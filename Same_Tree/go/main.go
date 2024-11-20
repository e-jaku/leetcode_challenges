package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p == nil && q != nil) || (q == nil && p != nil) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) {
		return true
	}

	return false
}

func main() {

	treeOne := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	treeTwo := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println("Tree", treeOne, "and", treeTwo, "are the same", isSameTree(treeOne, treeTwo))
}
