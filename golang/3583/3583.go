package main

func specialTriplets(nums []int) int {
	const MOD = 1000000007
	tripletCount := 0

	leftFreq := make(map[int]int)
	rightFreq := make(map[int]int)

	// Initialize right frequency map
	for _, num := range nums {
		rightFreq[num]++
	}

	for j := 0; j < len(nums); j++ {
		rightFreq[nums[j]]--

		eq := nums[j] * 2
		leftCount := leftFreq[eq]
		rightCount := rightFreq[eq]

		tripletCount = (tripletCount + leftCount*rightCount) % MOD

		leftFreq[nums[j]]++
	}

	return tripletCount
}
