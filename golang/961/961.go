package main

func repeatedNTimes(nums []int) int {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return num
		}
		seen[num] = true
	}
	return -1
}
