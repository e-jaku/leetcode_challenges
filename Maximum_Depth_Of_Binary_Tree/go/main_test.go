package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name  string
		nodes *TreeNode
		depth int
	}{
		{name: "Test1", nodes: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 4}}, depth: 3},
		{name: "Test2", nodes: &TreeNode{Val: 1}, depth: 1},
		{name: "Test3", nodes: nil, depth: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := maxDepth(tc.nodes)
			require.Equal(t, tc.depth, res)
		})
	}
}
