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
			num:      11891,
			expected: 99009,
		},
		{
			name:     "example 2",
			num:      90,
			expected: 99,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := minMaxDifference(tc.num)
			if result != tc.expected {
				t.Errorf("minMaxDifference(%d) = %d; want %d", tc.num, result, tc.expected)
			}
		})
	}
}
