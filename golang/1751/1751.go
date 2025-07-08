package main

import (
	"sort"
)

func maxValue(events [][]int, k int) int {
	n := len(events)
	// Sort events by end time
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	// Precompute the previous non-overlapping event for each event
	prev := make([]int, n)
	for i := 0; i < n; i++ {
		l, r := 0, i-1
		for l <= r {
			m := (l + r) / 2
			if events[m][1] < events[i][0] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		prev[i] = r
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			// Not take event i-1
			dp[i][j] = dp[i-1][j]
			// Take event i-1
			if v := dp[prev[i-1]+1][j-1] + events[i-1][2]; v > dp[i][j] {
				dp[i][j] = v
			}
		}
	}
	return dp[n][k]
}
