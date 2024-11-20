package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name            string
		sentence        string
		substringLength int
	}{
		{name: "Test1", sentence: "abcabcbb", substringLength: 3},
		{name: "Test2", sentence: "race a bike", substringLength: 6},
		{name: "Test3", sentence: "", substringLength: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := lengthOfLongestSubstring(tc.sentence)
			require.Equal(t, tc.substringLength, res)
		})
	}

}
