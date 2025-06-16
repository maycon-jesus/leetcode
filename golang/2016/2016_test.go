package main

import (
	"testing"
)

func TestMaximumDifference(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "simple increasing sequence",
			nums:     []int{1, 2},
			expected: 1,
		},
		{
			name:     "longer increasing sequence",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 4, // 5-1=4
		},
		{
			name:     "decreasing sequence",
			nums:     []int{5, 4, 3, 2, 1},
			expected: -1, // No valid pairs where second > first
		},
		{
			name:     "mixed sequence with positive difference",
			nums:     []int{7, 1, 5, 3, 6, 4},
			expected: 5, // 6-1=5
		},
		{
			name:     "V-shaped sequence",
			nums:     []int{9, 7, 5, 3, 1, 2, 4, 6, 8},
			expected: 7, // 8-1=7
		},
		{
			name:     "sawtooth pattern",
			nums:     []int{1, 5, 2, 6, 3, 7},
			expected: 6, // 7-1=6
		},
		{
			name:     "equal elements",
			nums:     []int{5, 5, 5, 5},
			expected: -1, // No positive differences
		},
		{
			name:     "single element",
			nums:     []int{5},
			expected: -1, // Need at least two elements
		},
		{
			name:     "negative numbers with positive difference",
			nums:     []int{-5, -4, -3, -2, -1},
			expected: 4, // -1-(-5)=4
		},
		{
			name:     "negative numbers with no valid difference",
			nums:     []int{-1, -2, -3, -4, -5},
			expected: -1, // No valid pairs
		},
		{
			name:     "mixed positive and negative",
			nums:     []int{-10, 5, -5, 10},
			expected: 20, // 10-(-10)=20
		},
		{
			name:     "minimum at the end, no valid difference",
			nums:     []int{9, 8, 7, 6, 5},
			expected: -1, // No valid pairs
		},
		{
			name:     "large test case - decreasing",
			nums:     []int{999, 997, 980, 976, 948, 940, 938, 928, 924, 917, 907, 907, 881, 878, 864, 862, 859, 857, 848, 840, 824, 824, 824, 805, 802, 798, 788, 777, 775, 766, 755, 748, 735, 732, 727, 705, 700, 697, 693, 679, 676, 644, 634, 624, 599, 596, 588, 583, 562, 558, 553, 539, 537, 536, 509, 491, 485, 483, 454, 449, 438, 425, 403, 368, 345, 327, 287, 285, 270, 263, 255, 248, 235, 234, 224, 221, 201, 189, 187, 183, 179, 168, 155, 153, 150, 144, 107, 102, 102, 87, 80, 57, 55, 49, 48, 45, 26, 26, 23, 15},
			expected: -1, // Strictly decreasing sequence
		},
		{
			name:     "large difference in middle",
			nums:     []int{10, 5, 100, 0, 20},
			expected: 95, // 100-5=95
		},
		{
			name:     "repeating pattern",
			nums:     []int{1, 3, 1, 3, 1, 3},
			expected: 2, // 3-1=2
		},
		{
			name:     "edge case - exactly two elements with positive diff",
			nums:     []int{0, 10},
			expected: 10, // 10-0=10
		},
		{
			name:     "edge case - exactly two elements with negative diff",
			nums:     []int{10, 0},
			expected: -1, // No valid pairs
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := maximumDifference(tc.nums)
			if result != tc.expected {
				t.Errorf("maximumDifference(%v) = %d; want %d", tc.nums, result, tc.expected)
			}
		})
	}
}
