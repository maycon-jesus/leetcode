package main

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{
			name:   "Example 1: Simple increment",
			digits: []int{1, 2, 3},
			want:   []int{1, 2, 4},
		},
		{
			name:   "Example 2: Four digit number",
			digits: []int{4, 3, 2, 1},
			want:   []int{4, 3, 2, 2},
		},
		{
			name:   "Example 3: Single digit 9",
			digits: []int{9},
			want:   []int{1, 0},
		},
		{
			name:   "All 9s - overflow case",
			digits: []int{9, 9, 9},
			want:   []int{1, 0, 0, 0},
		},
		{
			name:   "Carry propagation from last digit",
			digits: []int{1, 9, 9},
			want:   []int{2, 0, 0},
		},
		{
			name:   "No carry needed",
			digits: []int{5, 6, 7},
			want:   []int{5, 6, 8},
		},
		{
			name:   "Single digit 0",
			digits: []int{0},
			want:   []int{1},
		},
		{
			name:   "Partial carry propagation",
			digits: []int{8, 9, 9},
			want:   []int{9, 0, 0},
		},
		{
			name:   "Large number with 9s at end",
			digits: []int{1, 2, 3, 4, 9, 9},
			want:   []int{1, 2, 3, 5, 0, 0},
		},
		{
			name:   "Two digit number ending in 9",
			digits: []int{1, 9},
			want:   []int{2, 0},
		},
		{
			name:   "Maximum constraint - all 9s (100 digits)",
			digits: make99s(100),
			want:   append([]int{1}, make0s(100)...),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plusOne(tt.digits)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne(%v) = %v, want %v", tt.digits, got, tt.want)
			}
		})
	}
}

// Helper function to create a slice of n 9s
func make99s(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = 9
	}
	return result
}

// Helper function to create a slice of n 0s
func make0s(n int) []int {
	return make([]int, n)
}

// Benchmark tests
func BenchmarkPlusOne(b *testing.B) {
	digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		plusOne(digits)
	}
}

func BenchmarkPlusOneWorstCase(b *testing.B) {
	digits := make99s(100)
	for i := 0; i < b.N; i++ {
		plusOne(digits)
	}
}
