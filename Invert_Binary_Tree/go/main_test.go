package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name     string
		tree     *TreeNode
		inverted *TreeNode
	}{
		{name: "Test1", tree: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 4}},
			inverted: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Left: &TreeNode{Val: 4}}},
		{name: "Test2", tree: &TreeNode{Val: 1}, inverted: &TreeNode{Val: 1}},
		{name: "Test3", tree: nil, inverted: nil},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := invertTree(tc.tree)
			require.Equal(t, tc.inverted, res)
		})
	}
}
