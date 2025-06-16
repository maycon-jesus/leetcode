package main

// Problem: Maximum Difference Between a Pair of Values
// Time Complexity: O(n), where n is the length of the input array
// Space Complexity: O(1), using only constant extra space
func maximumDifference(nums []int) int {
	maxDiff := -1
	minVal := nums[0]

	// Single pass through the array
	for i := 1; i < len(nums); i++ {
		// If current element is greater than min so far
		// calculate difference and update maxDiff if needed
		if nums[i] > minVal {
			diff := nums[i] - minVal
			if diff > maxDiff {
				maxDiff = diff
			}
		} else {
			// Update min value if current element is smaller
			minVal = nums[i]
		}
	}

	return maxDiff
}
