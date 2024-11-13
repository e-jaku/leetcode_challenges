package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanJump(t *testing.T) {
	tests := []struct {
		name       string
		nums       []int
		reachesEnd bool
	}{
		{name: "Test1", nums: []int{1, 3, 3, 3, 7, 7}, reachesEnd: true},
		{name: "Test2", nums: []int{0, 0, 1, 2, 2, 2}, reachesEnd: false},
		{name: "Test3", nums: []int{}, reachesEnd: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := canJump(tc.nums)
			if len(tc.nums) > 0 {
				require.Equal(t, tc.reachesEnd, res)
			}
		})
	}

}
