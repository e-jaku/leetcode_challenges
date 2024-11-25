package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name          string
		s             string
		t             string
		isSubsequence bool
	}{
		{name: "Test1", s: "abx", t: "akx", isSubsequence: true},
		{name: "Test2", s: "rak", t: "rake", isSubsequence: true},
		{name: "Test2", s: "ake", t: "rake", isSubsequence: true},
		{name: "Test3", s: "", t: "asdf", isSubsequence: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := isOneEditDistance(tc.s, tc.t)
			require.Equal(t, tc.isSubsequence, res)
		})
	}

}
