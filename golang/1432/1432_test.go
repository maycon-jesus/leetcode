package main

import "testing"

func TestMinMaxDifference(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		expected int
	}{
		{
			name:     "example 1",
			num:      555,
			expected: 888,
		},
		{
			name:     "example 2",
			num:      9,
			expected: 8,
		},
		{
			name:     "example 2",
			num:      123456,
			expected: 820000,
		},
		{
			name:     "example 2",
			num:      9288,
			expected: 8700,
		},
		{
			name:     "example 2",
			num:      1101057,
			expected: 8808050,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := maxDiff(tc.num)
			if result != tc.expected {
				t.Errorf("maxDiff(%d) = %d; want %d", tc.num, result, tc.expected)
			}
		})
	}
}
