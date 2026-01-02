package main

import "testing"

func TestRepeatedNTimes(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1 - repeated element at beginning",
			nums:     []int{1, 2, 3, 3},
			expected: 3,
		},
		{
			name:     "example 2 - repeated element scattered",
			nums:     []int{2, 1, 2, 5, 3, 2},
			expected: 2,
		},
		{
			name:     "example 3 - repeated element at end",
			nums:     []int{5, 1, 5, 2, 5, 3, 5, 4},
			expected: 5,
		},
		{
			name:     "minimal case - 2 elements",
			nums:     []int{1, 1},
			expected: 1,
		},
		{
			name:     "repeated element consecutive",
			nums:     []int{9, 9, 8, 7},
			expected: 9,
		},
		{
			name:     "repeated element far apart",
			nums:     []int{7, 1, 2, 3, 7, 4, 5, 6},
			expected: 7,
		},
		{
			name:     "large numbers",
			nums:     []int{10000, 1, 2, 10000},
			expected: 10000,
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -1, 0, 1},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := repeatedNTimes(tt.nums)
			if result != tt.expected {
				t.Errorf("repeatedNTimes(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}
