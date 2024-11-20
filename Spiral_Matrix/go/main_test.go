package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		result []int
	}{
		{name: "Test1", matrix: [][]int{
			[]int{1, 2},
			[]int{3, 4},
		}, result: []int{1, 2, 4, 3}},
		{name: "Test2", matrix: [][]int{
			[]int{1, 2},
		}, result: []int{1, 2}},
		{name: "Test3", matrix: [][]int{}, result: []int{}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := spiralOrder(tc.matrix)
			require.Equal(t, tc.result, res)
		})
	}
}
