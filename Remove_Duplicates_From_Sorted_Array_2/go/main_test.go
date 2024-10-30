package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates2(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		remaining []int
	}{
		{name: "Test1", nums: []int{1, 2, 3, 3, 3, 7, 7}, remaining: []int{1, 2, 3, 3, 7, 7}},
		{name: "Test2", nums: []int{0, 0, 1, 2, 3, 3}, remaining: []int{0, 0, 1, 2, 3, 3}},
		{name: "Test3", nums: []int{}, remaining: []int{}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := removeDuplicates(tc.nums)
			if len(tc.nums) > 0 {
				require.Equal(t, tc.remaining, tc.nums[:res])
			}
		})
	}
}
