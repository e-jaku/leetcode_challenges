package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHIndex(t *testing.T) {
	tests := []struct {
		name      string
		citations []int
		hIndex    int
	}{
		{name: "Test1", citations: []int{1, 2, 3, 3, 7, 7}, hIndex: 3},
		{name: "Test2", citations: []int{7, 6, 4, 3, 1}, hIndex: 3},
		{name: "Test3", citations: []int{}, hIndex: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := hIndex(tc.citations)
			require.Equal(t, tc.hIndex, res)
		})
	}
}
