package main

// Problem: Two Sum
// Link: https://leetcode.com/problems/two-sum/
// Difficulty: Easy
// Tags: Array, Hash Table

import "fmt"

// main demonstrates the usage of the twoSum function by printing the indices of the two numbers
// in various integer slices that add up to a specified target value.
func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // Output: [0, 1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // Output: [1, 2]
	fmt.Println(twoSum([]int{3, 3}, 6))         // Output: [0, 1]
}

// twoSum finds two indices of the numbers in the given slice 'nums' that add up to the specified 'target'.
// It returns a slice containing the indices of the two numbers. If no such pair exists, it returns an empty slice.
//
// Parameters:
//
//	nums   - A slice of integers to search through.
//	target - The target sum to find.
//
// Returns:
//
//	A slice of two integers representing the indices of the numbers that add up to 'target'.
//	If no such pair exists, returns an empty slice.
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i1, v1 := range nums {
		diff := target - v1
		v2Index, ok := m[diff]
		if ok {
			return []int{v2Index, i1}
		}
		m[v1] = i1
	}
	return []int{}
}
