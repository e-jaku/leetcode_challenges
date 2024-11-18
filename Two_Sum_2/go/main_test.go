package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		target  int
		indexes []int
	}{
		{name: "Test1", nums: []int{2, 7, 11, 15}, target: 9, indexes: []int{1, 2}},
		{name: "Test2", nums: []int{2, 3, 4}, target: 6, indexes: []int{1, 3}},
		{name: "Test3", nums: []int{-1, 0}, target: -1, indexes: []int{1, 2}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := twoSum(tc.nums, tc.target)
			require.Equal(t, tc.indexes, res)
		})
	}

}
