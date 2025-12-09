package main

import "testing"

func TestSpecialTriplets(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1: Basic case with one triplet",
			nums:     []int{6, 3, 6},
			expected: 1,
		},
		{
			name:     "Example 2: Zero values",
			nums:     []int{0, 1, 0, 0},
			expected: 1,
		},
		{
			name:     "Example 3: Multiple triplets",
			nums:     []int{8, 4, 2, 8, 4},
			expected: 2,
		},
		{
			name:     "No triplets possible",
			nums:     []int{1, 2, 3},
			expected: 0,
		},
		{
			name:     "All same values",
			nums:     []int{4, 4, 4, 4},
			expected: 0,
		},
		{
			name:     "Multiple zeros",
			nums:     []int{0, 0, 0, 0, 0},
			expected: 10,
		},
		{
			name:     "Large values with pattern",
			nums:     []int{100000, 50000, 100000, 50000, 100000},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := specialTriplets(tt.nums)
			if result != tt.expected {
				t.Errorf("specialTriplets(%v) = %d; want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
