package main

import (
	"slices"
)

func maxSubsequence(nums []int, k int) []int {
	// Find the k largest numbers and their indices
	type pair struct {
		val int
		idx int
	}
	pairs := make([]pair, len(nums))
	for i, v := range nums {
		pairs[i] = pair{v, i}
	}
	// Sort by value descending
	slices.SortFunc(pairs, func(a, b pair) int {
		return b.val - a.val
	})
	// Take the first k pairs
	selected := pairs[:k]
	// Sort selected by original index to keep order
	slices.SortFunc(selected, func(a, b pair) int {
		return a.idx - b.idx
	})
	// Collect the values
	resp := make([]int, k)
	for i, p := range selected {
		resp[i] = p.val
	}
	return resp
}
