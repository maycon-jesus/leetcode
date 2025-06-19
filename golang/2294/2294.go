package main

// Problem: Partition Array Such That Maximum Difference Is K
// Link: https://leetcode.com/problems/partition-array-such-that-maximum-difference-is-k/
// Difficulty: Medium
// Tags: Array, Greedy, Sorting
// Time Complexity: O(n + m), where n is the length of nums and m is the range of values in nums
// Space Complexity: O(m), where m is the range of values in nums

func partitionArray(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	minV, maxV := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < minV {
			minV = nums[i]
		}
		if nums[i] > maxV {
			maxV = nums[i]
		}
	}

	present := make([]bool, maxV+1)
	for _, num := range nums {
		present[num] = true
	}

	count := 0
	i := minV
	for i <= maxV {
		if present[i] {
			count++
			i += k + 1
		} else {
			i++
		}
	}
	return count
}
