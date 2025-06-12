package main

import (
	"testing"
)

func TestMaxAdjacentDistance(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "increasing sequence",
			nums:     []int{1, 2, 4},
			expected: 3,
		},
		{
			name:     "negative numbers",
			nums:     []int{-5, -10, -5},
			expected: 5,
		},
		{
			name:     "decreasing sequence",
			nums:     []int{2, 1, 0},
			expected: 2,
		},
		{
			name:     "single element",
			nums:     []int{7},
			expected: 0,
		},
		{
			name:     "two elements",
			nums:     []int{10, 15},
			expected: 5,
		},
		{
			name:     "all same elements",
			nums:     []int{3, 3, 3},
			expected: 0,
		},
		{
			name:     "large difference at end",
			nums:     []int{1, 2, 100},
			expected: 99,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := maxAdjacentDistance(tc.nums)
			if result != tc.expected {
				t.Errorf("maxAdjacentDistance(%v) = %d; want %d", tc.nums, result, tc.expected)
			}
		})
	}
}