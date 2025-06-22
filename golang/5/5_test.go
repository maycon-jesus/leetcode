package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantOpts []string // Accept multiple valid answers
	}{
		{
			name:     "Example 1",
			input:    "babad",
			wantOpts: []string{"bab", "aba"},
		},
		{
			name:     "Example 2",
			input:    "cbbd",
			wantOpts: []string{"bb"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindrome(tt.input)
			found := false
			for _, want := range tt.wantOpts {
				if got == want {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("longestPalindrome(%q) = %q; want one of %v", tt.input, got, tt.wantOpts)
			}
		})
	}
}
