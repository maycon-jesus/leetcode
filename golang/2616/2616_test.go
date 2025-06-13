package main

import "testing"

func TestMinimizeMax(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		p        int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{10, 1, 2, 7, 1, 3},
			p:        2,
			expected: 1,
		},
		{
			name:     "example 2",
			nums:     []int{4, 2, 1, 2},
			p:        1,
			expected: 0,
		},
		{
			name:     "example 3",
			nums:     []int{3, 4, 2, 3, 2, 1, 2},
			p:        3,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimizeMax(tt.nums, tt.p)
			if result != tt.expected {
				t.Errorf("minimizeMax(%v, %d) = %d; expected %d", tt.nums, tt.p, result, tt.expected)
			}
		})
	}
}
