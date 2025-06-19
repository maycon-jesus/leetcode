package main

// Problem: Median of Two Sorted Arrays
// LeetCode: https://leetcode.com/problems/median-of-two-sorted-arrays/
// Difficulty: Hard
// Tags: Array, Binary Search, Divide and Conquer
// Time Complexity: O(m + n), where m and n are the lengths of the two input arrays.
// Space Complexity: O(m + n), where m and n are the lengths of the two input arrays.

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sizeMerged := len(nums1) + len(nums2)
	numsMerged := make([]int, sizeMerged)
	i, j, k := 0, 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			numsMerged[k] = nums1[i]
			i++
		} else {
			numsMerged[k] = nums2[j]
			j++
		}
		k++
	}
	if i < len(nums1) {
		for _, n := range nums1[i:] {
			numsMerged[k] = n
			k++
		}
	}
	if j < len(nums2) {
		for _, n := range nums2[j:] {
			numsMerged[k] = n
			k++
		}
	}

	if sizeMerged%2 == 0 {
		return float64(numsMerged[sizeMerged/2-1]+numsMerged[sizeMerged/2]) / 2
	}
	return float64(numsMerged[sizeMerged/2])
}
