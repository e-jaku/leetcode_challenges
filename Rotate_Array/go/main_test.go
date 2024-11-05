package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotateArray(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		value  int
		result []int
	}{
		{name: "Test1", nums: []int{1, 2, 3, 4, 5}, value: 2, result: []int{4, 5, 1, 2, 3}},
		{name: "Test2", nums: []int{1, 2, 3, 4, 5}, value: 6, result: []int{5, 1, 2, 3, 4}},
		{name: "Test3", nums: []int{}, value: 6, result: []int{}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rotate(tc.nums, tc.value)
			require.Equal(t, tc.result, tc.nums)
		})
	}
}
