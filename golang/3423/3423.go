package main

import (
	"math"
)

func main() {
	nums := []int{1, 2, 4}
	result := maxAdjacentDistance(nums)
	println(result) // Output: 3

	nums = []int{-5, -10, -5}
	result = maxAdjacentDistance(nums)
	println(result) // Output: 5

	nums = []int{2, 1, 0}
	result = maxAdjacentDistance(nums)
	println(result) // Output: 2
}

func maxAdjacentDistance(nums []int) int {
	l := len(nums)
	if l < 2 {
		return 0
	}
	maxDiff := int(math.Abs(float64(nums[0] - nums[l-1])))
	for i := 1; i < l-1; i++ {
		diffPrev := absInt(nums[i] - nums[i-1])
		diffNext := absInt(nums[i] - nums[i+1])
		if diffPrev > maxDiff {
			maxDiff = diffPrev
		}
		if diffNext > maxDiff {
			maxDiff = diffNext
		}
	}
	return maxDiff
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
