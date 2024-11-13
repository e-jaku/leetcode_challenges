package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPalindrome(t *testing.T) {
	tests := []struct {
		name         string
		sentence     string
		isPalindrome bool
	}{
		{name: "Test1", sentence: "abba , is i abba", isPalindrome: true},
		{name: "Test2", sentence: "race a bike", isPalindrome: false},
		{name: "Test3", sentence: " ", isPalindrome: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := isPalindrome(tc.sentence)
			require.Equal(t, tc.isPalindrome, res)
		})
	}

}
