package main

import "testing"

func TestLongestSubsequenceRepeatedK(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected string
	}{
		{
			name:     "Example 1 - letsleetcode with k=2",
			s:        "letsleetcode",
			k:        2,
			expected: "let",
		},
		{
			name:     "Example 2 - bb with k=2",
			s:        "bb",
			k:        2,
			expected: "b",
		},
		{
			name:     "Example 3 - ab with k=2",
			s:        "ab",
			k:        2,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestSubsequenceRepeatedK(tt.s, tt.k)
			if result != tt.expected {
				t.Errorf("longestSubsequenceRepeatedK(%q, %d) = %q, want %q",
					tt.s, tt.k, result, tt.expected)
			}
		})
	}
}
