package main

import "testing"

func TestCandy(t *testing.T) {
	tests := []struct {
		name     string
		ratings  []int
		expected int
	}{
		{
			name:     "example 1",
			ratings:  []int{1, 0, 2},
			expected: 5,
		},
		{
			name:     "example 2",
			ratings:  []int{1, 2, 2},
			expected: 4,
		},
		{
			name:     "example 3",
			ratings:  []int{1, 3, 2, 2, 1},
			expected: 7,
		},
		{
			name:     "empty",
			ratings:  []int{},
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := candy(tc.ratings)
			if result != tc.expected {
				t.Errorf("candy(%v) = %d; want %d", tc.ratings, result, tc.expected)
			}
		})
	}
}
