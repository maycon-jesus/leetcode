package main

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "example 1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "example 2",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "example 3",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "no solution",
			nums:     []int{1, 2, 3},
			target:   7,
			expected: nil, // ou []int{}, dependendo da sua implementação
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			target:   -8,
			expected: []int{2, 4},
		},
		{
			name:     "array of size 2",
			nums:     []int{1, 1},
			target:   2,
			expected: []int{0, 1},
		},
		{
			name:     "multiple solutions",
			nums:     []int{1, 2, 3, 4, 4},
			target:   8,
			expected: []int{3, 4}, // ou qualquer solução válida
		},
		{
			name:     "empty array",
			nums:     []int{},
			target:   0,
			expected: nil,
		},
	}

	equal := func(a, b []int) bool {
		if a == nil && b == nil {
			return true
		}
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := twoSum(tc.nums, tc.target)
			if !equal(result, tc.expected) {
				t.Errorf("twoSum(%v, %d) = %v; want %v", tc.nums, tc.target, result, tc.expected)
			}
		})
	}
}
