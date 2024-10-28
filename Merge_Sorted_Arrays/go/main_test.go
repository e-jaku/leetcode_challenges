package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {

	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		m, n     int
		expected []int
	}{
		{name: "Test1", nums1: []int{1, 6, 7, 0, 0, 0}, nums2: []int{2, 6, 8}, m: 3, n: 3, expected: []int{1, 2, 6, 6, 7, 8}},
		{name: "Test2", nums1: []int{0}, nums2: []int{2}, m: 0, n: 1, expected: []int{2}},
		{name: "Test1", nums1: []int{8, 8, 9, 0, 0, 0}, nums2: []int{2}, m: 3, n: 1, expected: []int{2, 8, 8, 9, 0, 0}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			merge(tc.nums1, tc.m, tc.nums2, tc.n)
			require.Equal(t, tc.expected, tc.nums1)
		})
	}
}
