package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {

	tests := []struct {
		name      string
		nums      []int
		maxLength int
	}{
		{name: "Test1", nums: []int{1, 0, 1, 1, 0}, maxLength: 3},
		{name: "Test2", nums: []int{0}, maxLength: 0},
		{name: "Test3", nums: []int{1, 0, 0, 1, 1}, maxLength: 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := longestSubarray(tc.nums)
			require.Equal(t, tc.maxLength, res)
		})
	}
}
