package main

var MOD = 1_000_000_007

func countPermutations(complexity []int) int {
	n := len(complexity)
	root := complexity[0]

	for i := 1; i < n; i++ {
		if complexity[i] <= root {
			return 0
		}
	}

	ans := 1
	for i := 2; i < n; i++ {
		ans = (ans * i) % MOD
	}

	return ans
}
