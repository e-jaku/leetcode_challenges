package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name   string
		p      *TreeNode
		q      *TreeNode
		isSame bool
	}{
		{name: "Test1", p: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 4}},
			q: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 4}}, isSame: true},
		{name: "Test2", p: &TreeNode{Val: 1}, isSame: false},
		{name: "Test3", p: nil, q: nil, isSame: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := isSameTree(tc.p, tc.q)
			require.Equal(t, tc.isSame, res)
		})
	}
}
