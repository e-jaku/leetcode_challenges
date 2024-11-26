package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindClosestElements(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		k        int
		x        int
		expected []int
	}{
		{name: "Test1", arr: []int{1, 2, 3, 3, 7, 7}, k: 3, x: 5, expected: []int{3, 3, 7}},
		{name: "Test2", arr: []int{0, 1, 5, 9, 11, 24}, k: 3, x: 5, expected: []int{1, 5, 9}},
		{name: "Test3", arr: []int{-1, 0, 1, 2, 5}, k: 2, x: 0, expected: []int{-1, 0}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := findClosestElements(tc.arr, tc.k, tc.x)
			require.Equal(t, tc.expected, res)
		})
	}
}
