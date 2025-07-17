package main

func maximumLength(nums []int, k int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k)
		for rem := 0; rem < k; rem++ {
			dp[i][rem] = 1 // every element is a subsequence of length 1
		}
	}
	ans := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			rem := (nums[j] + nums[i]) % k
			if dp[j][rem]+1 > dp[i][rem] {
				dp[i][rem] = dp[j][rem] + 1
				if dp[i][rem] > ans {
					ans = dp[i][rem]
				}
			}
		}
	}
	return ans
}
