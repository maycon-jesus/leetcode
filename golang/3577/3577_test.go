package main

import "testing"

func TestCountPermutations(t *testing.T) {
	tests := []struct {
		name       string
		complexity []int
		want       int
	}{
		{
			name:       "example 1 - valid increasing sequence",
			complexity: []int{1, 2, 3},
			want:       2,
		},
		{
			name:       "example 2 - first element not smallest",
			complexity: []int{5, 1, 2, 3},
			want:       0,
		},
		{
			name:       "example 3 - all elements greater than first",
			complexity: []int{1, 5, 3, 4, 2},
			want:       24,
		},
		{
			name:       "single element",
			complexity: []int{1},
			want:       1,
		},
		{
			name:       "two elements - valid",
			complexity: []int{1, 2},
			want:       1,
		},
		{
			name:       "two elements - invalid",
			complexity: []int{2, 1},
			want:       0,
		},
		{
			name:       "equal elements after first",
			complexity: []int{5, 5, 6},
			want:       0,
		},
		{
			name:       "first element equal to second",
			complexity: []int{3, 3, 4},
			want:       0,
		},
		{
			name:       "large valid sequence",
			complexity: []int{1, 10, 20, 30, 40, 50},
			want:       120,
		},
		{
			name:       "minimum valid with negative numbers",
			complexity: []int{-10, -5, 0, 5},
			want:       6,
		},
		{
			name:       "one element not greater than first",
			complexity: []int{10, 20, 30, 5, 40},
			want:       0,
		},
		{
			name:       "all elements same value",
			complexity: []int{5, 5, 5, 5},
			want:       0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPermutations(tt.complexity)
			if got != tt.want {
				t.Errorf("countPermutations(%v) = %v, want %v", tt.complexity, got, tt.want)
			}
		})
	}
}

func TestCountPermutationsModulo(t *testing.T) {
	// Test that modulo operation is applied correctly
	// For n=3, factorial(2) = 2 (no modulo needed)
	// For n=4, factorial(3) = 6 (no modulo needed)
	// For n=10, factorial(9) should be modulo MOD

	tests := []struct {
		name       string
		complexity []int
		want       int
	}{
		{
			name:       "factorial that needs modulo - large sequence",
			complexity: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			want:       3628800, // 10! = 3628800
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPermutations(tt.complexity)
			if got != tt.want {
				t.Errorf("countPermutations(%v) = %v, want %v", tt.complexity, got, tt.want)
			}
		})
	}
}

func TestCountPermutationsEdgeCases(t *testing.T) {
	tests := []struct {
		name       string
		complexity []int
		want       int
	}{
		{
			name:       "boundary value - first element at minimum",
			complexity: []int{-1000000, 0, 1, 2},
			want:       6,
		},
		{
			name:       "boundary value - large first element fails",
			complexity: []int{1000000, 999999, 1000001},
			want:       0,
		},
		{
			name:       "mixed values all greater than first",
			complexity: []int{1, 10, 9, 8, 7},
			want:       24, // (n-1)! = 4! = 24, all elements > first
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPermutations(tt.complexity)
			if got != tt.want {
				t.Errorf("countPermutations(%v) = %v, want %v", tt.complexity, got, tt.want)
			}
		})
	}
}

func BenchmarkCountPermutations(b *testing.B) {
	benchmarks := []struct {
		name       string
		complexity []int
	}{
		{
			name:       "small valid sequence",
			complexity: []int{1, 2, 3, 4, 5},
		},
		{
			name:       "small invalid sequence",
			complexity: []int{5, 1, 2, 3, 4},
		},
		{
			name:       "medium valid sequence",
			complexity: []int{1, 10, 20, 30, 40, 50, 60, 70, 80, 90},
		},
		{
			name:       "large valid sequence",
			complexity: make([]int, 100),
		},
	}

	// Initialize large sequence
	for i := 0; i < 100; i++ {
		benchmarks[3].complexity[i] = i + 1
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				countPermutations(bm.complexity)
			}
		})
	}
}
