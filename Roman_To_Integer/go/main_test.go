package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name    string
		roman   string
		integer int
	}{
		{name: "Test1", roman: "III", integer: 3},
		{name: "Test2", roman: "LVIII", integer: 58},
		{name: "Test3", roman: "MCMXCIV", integer: 1994},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := romanToInt(tc.roman)
			require.Equal(t, tc.integer, res)
		})
	}

}
