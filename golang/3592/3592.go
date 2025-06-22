package main

func findCoins(numWays []int) []int {
	coins := make([]int, 0)
	error := false
	for i, ways := range numWays {
		amount := i + 1
		count := possibilities(coins, amount)
		if count > ways {
			error = true
			break
		}
		if ways > 0 {
			if ways-count > 1 {
				error = true
				break
			}
			if count < ways {
				coins = append(coins, amount)
			}
		}
	}

	if error {
		return []int{}
	}
	return coins
}

func possibilities(coins []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= target; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[target]
}
