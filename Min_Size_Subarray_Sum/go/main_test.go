package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		min    int
	}{
		{name: "Test1", nums: []int{2, 7, 11, 15}, target: 14, min: 1},
		{name: "Test2", nums: []int{2, 3, 4}, target: 9, min: 3},
		{name: "Test3", nums: []int{2, 0}, target: 3, min: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := minSubArrayLen(tc.target, tc.nums)
			require.Equal(t, tc.min, res)
		})
	}

}
