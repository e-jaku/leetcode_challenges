package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCandy(t *testing.T) {
	tests := []struct {
		name    string
		ratings []int
		candy   int
	}{
		{name: "Test1", ratings: []int{2, 3, 0, 1, 4}, candy: 9},
		{name: "Test2", ratings: []int{0}, candy: 1},
		{name: "Test3", ratings: []int{2, 2, 1}, candy: 4},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := candy(tc.ratings)
			if len(tc.ratings) > 0 {
				require.Equal(t, tc.candy, res)
			}
		})
	}
}
