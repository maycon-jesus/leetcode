package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		expected bool
	}{
		{
			name:     "palindrome 121",
			num:      121,
			expected: true,
		},
		{
			name:     "not palindrome -121",
			num:      -121,
			expected: false,
		},
		{
			name:     "not palindrome 10",
			num:      10,
			expected: false,
		},
		{
			name:     "palindrome 0",
			num:      0,
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := isPalindrome(tc.num)
			if result != tc.expected {
				t.Errorf("isPalindrome(%d) = %v; want %v", tc.num, result, tc.expected)
			}
		})
	}
}
