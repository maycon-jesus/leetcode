package main

// Problem: Minimize Maximum Pair Sum in Array
// Link: https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/
// Difficulty: Medium
// Tags: Array, Binary Search, Greedy

import (
	"sort"
)

func minimizeMax(nums []int, p int) int {
	if p == 0 {
		return 0
	}

	// Sort in-place to avoid creating a new array
	sort.Ints(nums)

	// Use more precise bounds for binary search
	left := 0
	right := 0

	// Find the actual maximum difference we might need
	// Only consider adjacent pairs since they're optimal after sorting
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		if diff > right {
			right = diff
		}
	}

	// Binary search on the answer
	for left < right {
		mid := left + (right-left)/2
		if canFormPairs(nums, p, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// Optimized to use minimal variables and early termination
func canFormPairs(nums []int, p int, maxDiff int) bool {
	pairs := 0

	// Single loop variable, no intermediate storage
	for i := 0; i < len(nums)-1; {
		if nums[i+1]-nums[i] <= maxDiff {
			pairs++
			if pairs == p { // Early termination - no need to continue
				return true
			}
			i += 2 // Skip both paired elements
		} else {
			i++ // Move to next element
		}
	}

	return false // We know pairs < p at this point
}
