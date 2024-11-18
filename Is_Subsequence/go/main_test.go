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
		{name: "Test1", s: "abx", t: "afgbxc", isSubsequence: true},
		{name: "Test2", s: "rak", t: "akr", isSubsequence: false},
		{name: "Test3", s: "", t: "asdf", isSubsequence: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := isSubsequence(tc.s, tc.t)
			require.Equal(t, tc.isSubsequence, res)
		})
	}

}
