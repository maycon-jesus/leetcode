package main

import (
	"testing"
)

// Original quadratic implementation for benchmark comparison
func maximumDifferenceOriginal(nums []int) int {
	diff := -1

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if n := nums[j] - nums[i]; n > 0 && n > diff {
				diff = n
			}
		}
	}

	return diff
}

// Benchmark the optimized implementation
func BenchmarkMaximumDifference(b *testing.B) {
	smallInput := []int{1, 2, 3, 4, 5}
	mediumInput := make([]int, 100)
	largeInput := make([]int, 1000)

	// Initialize test data
	for i := 0; i < len(mediumInput); i++ {
		mediumInput[i] = i % 20
	}

	for i := 0; i < len(largeInput); i++ {
		largeInput[i] = i % 50
	}

	b.Run("Small Input (5 elements)", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			maximumDifference(smallInput)
		}
	})

	b.Run("Medium Input (100 elements)", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			maximumDifference(mediumInput)
		}
	})

	b.Run("Large Input (1000 elements)", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			maximumDifference(largeInput)
		}
	})
}

// Benchmark the original implementation for comparison
func BenchmarkMaximumDifferenceOriginal(b *testing.B) {
	smallInput := []int{1, 2, 3, 4, 5}
	mediumInput := make([]int, 100)

	// Initialize test data
	for i := 0; i < len(mediumInput); i++ {
		mediumInput[i] = i % 20
	}

	b.Run("Small Input (5 elements)", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			maximumDifferenceOriginal(smallInput)
		}
	})

	b.Run("Medium Input (100 elements)", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			maximumDifferenceOriginal(mediumInput)
		}
	})

	// Not benchmarking large input for original algorithm as it would be too slow
}
