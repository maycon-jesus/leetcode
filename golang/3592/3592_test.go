package main

import (
	"testing"
)

func TestFindCoins(t *testing.T) {
	testCases := []struct {
		name     string
		numWays  []int
		expected []int
	}{
		{
			name:     "Example 1",
			numWays:  []int{0, 1, 0, 2, 0, 3, 0, 4, 0, 5},
			expected: []int{2, 4, 6},
		},
		{
			name:     "Example 2",
			numWays:  []int{1, 2, 2, 3, 4},
			expected: []int{1, 2, 5},
		},
		{
			name:     "Example 3",
			numWays:  []int{1, 2, 3, 4, 15},
			expected: []int{},
		},
		{
			name:     "numWays 1,0 impossible",
			numWays:  []int{1, 0},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			coins := findCoins(tc.numWays)
			if len(coins) != len(tc.expected) {
				t.Errorf("findCoins(%v) = %v; want %v", tc.numWays, coins, tc.expected)
				return
			}
			for i := range coins {
				if coins[i] != tc.expected[i] {
					t.Errorf("findCoins(%v) = %v; want %v", tc.numWays, coins, tc.expected)
					return
				}
			}
		})
	}
}
