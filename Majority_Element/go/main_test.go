package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		majority int
	}{
		{name: "Test1", nums: []int{1, 3, 3, 3, 7, 7}, majority: 3},
		{name: "Test2", nums: []int{0, 0, 1, 2, 2, 2}, majority: 2},
		{name: "Test3", nums: []int{}, majority: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := majorityElement(tc.nums)
			if len(tc.nums) > 0 {
				require.Equal(t, tc.majority, res)
			}
		})
	}
}
