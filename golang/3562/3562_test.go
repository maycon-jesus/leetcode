package main

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		present   []int
		future    []int
		hierarchy [][]int
		budget    int
		expected  int
	}{
		{
			name:      "Example 1: Simple parent-child with discount",
			n:         2,
			present:   []int{1, 2},
			future:    []int{4, 3},
			hierarchy: [][]int{{1, 2}},
			budget:    3,
			expected:  5,
		},
		{
			name:      "Example 2: Budget constraint - only child",
			n:         2,
			present:   []int{3, 4},
			future:    []int{5, 8},
			hierarchy: [][]int{{1, 2}},
			budget:    4,
			expected:  4,
		},
		{
			name:      "Example 3: Two children, choose one",
			n:         3,
			present:   []int{4, 6, 8},
			future:    []int{7, 9, 11},
			hierarchy: [][]int{{1, 2}, {1, 3}},
			budget:    10,
			expected:  10,
		},
		{
			name:      "Example 4: Chain of employees",
			n:         3,
			present:   []int{5, 2, 3},
			future:    []int{8, 5, 6},
			hierarchy: [][]int{{1, 2}, {2, 3}},
			budget:    7,
			expected:  12,
		},
		{
			name:      "Single employee (CEO only)",
			n:         1,
			present:   []int{10},
			future:    []int{20},
			hierarchy: [][]int{},
			budget:    10,
			expected:  10,
		},
		{
			name:      "Single employee - insufficient budget",
			n:         1,
			present:   []int{10},
			future:    []int{20},
			hierarchy: [][]int{},
			budget:    5,
			expected:  0,
		},
		{
			name:      "No profit possible - budget too low",
			n:         2,
			present:   []int{10, 10},
			future:    []int{15, 15},
			hierarchy: [][]int{{1, 2}},
			budget:    0,
			expected:  0,
		},
		{
			name:      "Negative profit stocks",
			n:         2,
			present:   []int{10, 10},
			future:    []int{5, 5},
			hierarchy: [][]int{{1, 2}},
			budget:    20,
			expected:  0,
		},
		{
			name:      "Large budget - buy all",
			n:         3,
			present:   []int{5, 2, 3},
			future:    []int{8, 5, 6},
			hierarchy: [][]int{{1, 2}, {2, 3}},
			budget:    100,
			expected:  12,
		},
		{
			name:      "Discount makes difference",
			n:         2,
			present:   []int{2, 10},
			future:    []int{5, 20},
			hierarchy: [][]int{{1, 2}},
			budget:    7,
			expected:  18,
		},
		{
			name:      "Deep tree - 4 levels",
			n:         4,
			present:   []int{1, 2, 3, 4},
			future:    []int{5, 6, 7, 8},
			hierarchy: [][]int{{1, 2}, {2, 3}, {3, 4}},
			budget:    5,
			expected:  21, // All 4 with discounts: 1 + 1 + 1 + 2 = 5, profit = (5-1)+(6-1)+(7-1)+(8-2) = 21
		},
		{
			name:      "Wide tree - CEO with 3 children",
			n:         4,
			present:   []int{2, 4, 6, 8},
			future:    []int{5, 8, 10, 12},
			hierarchy: [][]int{{1, 2}, {1, 3}, {1, 4}},
			budget:    10,
			expected:  18, // Buy CEO (2) + child1 (2) + child2 (3) + child3 (4) = 11, but limited by budget
		},
		{
			name:      "Minimal profit scenario",
			n:         2,
			present:   []int{1, 1},
			future:    []int{2, 2},
			hierarchy: [][]int{{1, 2}},
			budget:    1,
			expected:  3, // Can buy both with discount: parent 1, child 0 (floor(1/2)=0), total cost 1, profit = (2-1)+(2-0) = 3
		},
		{
			name:      "Max constraints - high values",
			n:         3,
			present:   []int{50, 50, 50},
			future:    []int{50, 50, 50},
			hierarchy: [][]int{{1, 2}, {1, 3}},
			budget:    160,
			expected:  50, // Buy all 3: 50 + 25 + 25 = 100, profit = 0+0+0 = 0, but actually can buy more optimally
		},
		{
			name:      "Asymmetric tree",
			n:         5,
			present:   []int{3, 2, 4, 1, 5},
			future:    []int{8, 7, 9, 6, 10},
			hierarchy: [][]int{{1, 2}, {1, 3}, {2, 4}, {2, 5}},
			budget:    15,
			expected:  32, // Optimal combination with discounts
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProfit(tt.n, tt.present, tt.future, tt.hierarchy, tt.budget)
			if result != tt.expected {
				t.Errorf("maxProfit() = %d, expected %d", result, tt.expected)
			}
		})
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	// Benchmark with a moderately complex tree
	n := 10
	present := []int{5, 10, 15, 20, 25, 8, 12, 18, 22, 30}
	future := []int{15, 25, 30, 35, 40, 18, 22, 28, 32, 45}
	hierarchy := [][]int{
		{1, 2}, {1, 3}, {2, 4}, {2, 5}, {3, 6}, {3, 7}, {4, 8}, {5, 9}, {6, 10},
	}
	budget := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxProfit(n, present, future, hierarchy, budget)
	}
}

func BenchmarkMaxProfitLarge(b *testing.B) {
	// Benchmark with larger input
	n := 20
	present := make([]int, n)
	future := make([]int, n)
	hierarchy := make([][]int, 0, n-1)

	for i := 0; i < n; i++ {
		present[i] = (i + 1) * 2
		future[i] = (i + 1) * 4
		if i > 0 {
			hierarchy = append(hierarchy, []int{(i / 2) + 1, i + 1})
		}
	}
	budget := 160

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxProfit(n, present, future, hierarchy, budget)
	}
}
