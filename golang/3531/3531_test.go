package main

import "testing"

func TestCountCoveredBuildings(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		buildings [][]int
		expected  int
	}{
		{
			name: "Example 1 from problem - [2,2] is covered",
			n:    3,
			buildings: [][]int{
				{1, 2}, {2, 2}, {3, 2}, {2, 1}, {2, 3},
			},
			expected: 1,
		},
		{
			name: "Example 2 from problem - No covered buildings",
			n:    3,
			buildings: [][]int{
				{1, 1}, {1, 2}, {2, 1}, {2, 2},
			},
			expected: 0,
		},
		{
			name: "Example 3 from problem - [3,3] is covered",
			n:    5,
			buildings: [][]int{
				{1, 3}, {3, 2}, {3, 3}, {3, 5}, {5, 3},
			},
			expected: 1,
		},
		{
			name:      "No buildings",
			n:         5,
			buildings: [][]int{},
			expected:  0,
		},
		{
			name: "Less than 5 buildings",
			n:    5,
			buildings: [][]int{
				{1, 1}, {2, 2}, {3, 3}, {4, 4},
			},
			expected: 0,
		},
		{
			name: "Exactly 5 buildings but none covered",
			n:    5,
			buildings: [][]int{
				{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5},
			},
			expected: 0,
		},
		{
			name: "All buildings in a line - horizontal",
			n:    5,
			buildings: [][]int{
				{1, 1}, {2, 1}, {3, 1}, {4, 1}, {5, 1},
			},
			expected: 0,
		},
		{
			name: "All buildings in a line - vertical",
			n:    5,
			buildings: [][]int{
				{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5},
			},
			expected: 0,
		},
		{
			name: "Building at corner - not covered",
			n:    5,
			buildings: [][]int{
				{1, 1}, {1, 2}, {2, 1}, {2, 2}, {3, 3},
			},
			expected: 0,
		},
		{
			name: "Building in center - covered",
			n:    5,
			buildings: [][]int{
				{2, 2}, {1, 2}, {3, 2}, {2, 1}, {2, 3},
			},
			expected: 1,
		},
		{
			name: "Multiple buildings some covered some not",
			n:    7,
			buildings: [][]int{
				{1, 2}, {2, 1}, {2, 2}, {2, 3}, {3, 2},
			},
			expected: 1, // Only {2, 2} is covered
		},
		{
			name: "All buildings covered except corners",
			n:    5,
			buildings: [][]int{
				{1, 1}, {1, 2}, {1, 3},
				{2, 1}, {2, 2}, {2, 3},
				{3, 1}, {3, 2}, {3, 3},
			},
			expected: 1, // Only {2, 2} is covered
		},
		{
			name: "Large grid with multiple covered",
			n:    10,
			buildings: [][]int{
				{1, 5}, {2, 5}, {3, 5}, {4, 5}, {5, 5},
				{3, 3}, {3, 4}, {3, 6}, {3, 7},
			},
			expected: 1, // {3, 5} is covered
		},
		{
			name: "L-shape - no covered buildings",
			n:    5,
			buildings: [][]int{
				{1, 1}, {2, 1}, {3, 1}, {3, 2}, {3, 3},
			},
			expected: 0,
		},
		{
			name: "Cross shape - center is covered",
			n:    5,
			buildings: [][]int{
				{2, 1}, {2, 2}, {2, 3}, {1, 2}, {3, 2},
			},
			expected: 1, // {2, 2} is covered
		},
		{
			name: "Single building",
			n:    1,
			buildings: [][]int{
				{1, 1},
			},
			expected: 0,
		},
		{
			name: "Two buildings",
			n:    2,
			buildings: [][]int{
				{1, 1}, {2, 2},
			},
			expected: 0,
		},
		{
			name: "Three buildings",
			n:    3,
			buildings: [][]int{
				{1, 1}, {2, 2}, {3, 3},
			},
			expected: 0,
		},
		{
			name: "Square formation - none covered",
			n:    5,
			buildings: [][]int{
				{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2},
			},
			expected: 0, // {2, 2} needs buildings on all 4 sides in same row/col
		},
		{
			name: "Complex grid pattern",
			n:    8,
			buildings: [][]int{
				{2, 2}, {2, 4}, {2, 6},
				{4, 2}, {4, 4}, {4, 6},
				{6, 2}, {6, 4}, {6, 6},
			},
			expected: 1, // {4, 4} is covered
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countCoveredBuildings(tt.n, tt.buildings)
			if result != tt.expected {
				t.Errorf("countCoveredBuildings(%d, %v) = %d; expected %d",
					tt.n, tt.buildings, result, tt.expected)
			}
		})
	}
}

func BenchmarkCountCoveredBuildings(b *testing.B) {
	buildings := [][]int{
		{1, 1}, {2, 2}, {2, 3}, {3, 2}, {4, 2},
	}
	n := 5

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		countCoveredBuildings(n, buildings)
	}
}

func BenchmarkCountCoveredBuildingsLarge(b *testing.B) {
	// Create a large grid
	buildings := [][]int{}
	for i := 1; i <= 100; i++ {
		for j := 1; j <= 100; j++ {
			if (i+j)%3 == 0 {
				buildings = append(buildings, []int{i, j})
			}
		}
	}
	n := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		countCoveredBuildings(n, buildings)
	}
}
