package main

func getDescentPeriods(prices []int) int64 {
	ans := 0
	n := len(prices)

	combo := 1
	for i := 1; i < n; i++ {
		if prices[i-1]-prices[i] == 1 {
			combo++
		} else {
			ans += combo * (combo + 1) / 2
			combo = 1
		}
	}

	ans += combo * (combo + 1) / 2

	return int64(ans)
}
