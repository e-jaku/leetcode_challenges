package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		area    int
	}{
		{name: "Test1", heights: []int{1, 2, 3, 3, 7, 7}, area: 9},
		{name: "Test2", heights: []int{7, 6, 4, 3, 1}, area: 9},
		{name: "Test3", heights: []int{}, area: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := maxArea(tc.heights)
			require.Equal(t, tc.area, res)
		})
	}
}
