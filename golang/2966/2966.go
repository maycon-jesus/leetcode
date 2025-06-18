package main

// Problem: Divide Array Into Arrays With Maximum Difference
// Link: https://leetcode.com/problems/divide-array-into-arrays-with-max-difference
// Difficulty: Medium
// Tags: Array, Sorting, Greedy
// Time Complexity: O(n log n), where n is the number of elements in the array
// Space Complexity: O(n), for the result array

import "sort"

func divideArray(nums []int, k int) [][]int {
	n := len(nums)
	sort.Ints(nums)

	result := make([][]int, n/3)
	for i := 0; i < n; i += 3 {
		if nums[i+2]-nums[i] > k {
			return [][]int{}
		}
		result[i/3] = nums[i : i+3]
	}
	return result
}
