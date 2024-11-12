package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		jumps int
	}{
		{name: "Test1", nums: []int{2, 3, 0, 1, 4}, jumps: 2},
		{name: "Test2", nums: []int{0}, jumps: 0},
		{name: "Test3", nums: []int{2, 2, 1}, jumps: 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := jump(tc.nums)
			if len(tc.nums) > 0 {
				require.Equal(t, tc.jumps, res)
			}
		})
	}
}
