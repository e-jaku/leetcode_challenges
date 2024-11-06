package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		profit int
	}{
		{name: "Test1", prices: []int{7, 1, 5, 3, 6, 4}, profit: 7},
		{name: "Test2", prices: []int{1, 2, 3, 4, 5}, profit: 4},
		{name: "Test3", prices: []int{7, 6, 4, 3, 1}, profit: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := maxProfit(tc.prices)
			if len(tc.prices) > 0 {
				require.Equal(t, tc.profit, res)
			}
		})
	}
}
