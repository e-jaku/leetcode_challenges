package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		value     int
		remaining []int
	}{
		{name: "Test1", nums: []int{1, 6, 7, 0, 0, 0}, value: 6, remaining: []int{1, 7, 0, 0, 0}},
		{name: "Test2", nums: []int{1, 6, 7, 0, 0, 0}, value: 9, remaining: []int{1, 6, 7, 0, 0, 0}},
		{name: "Test3", nums: []int{}, value: 6, remaining: []int{}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := removeElement(tc.nums, tc.value)
			require.Equal(t, tc.remaining, tc.nums[:res])
		})
	}
}
