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
		{name: "Test1", prices: []int{1, 2, 3, 3, 7, 7}, profit: 6},
		{name: "Test2", prices: []int{7, 6, 4, 3, 1}, profit: 0},
		{name: "Test3", prices: []int{}, profit: 0},
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
