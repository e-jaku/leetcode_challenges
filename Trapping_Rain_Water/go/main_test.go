package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJumpGame(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		water   int
	}{
		{name: "Test1", heights: []int{2, 3, 0, 1, 4}, water: 5},
		{name: "Test2", heights: []int{0}, water: 0},
		{name: "Test3", heights: []int{2, 2, 1}, water: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := trap(tc.heights)
			require.Equal(t, tc.water, res)
		})
	}
}
