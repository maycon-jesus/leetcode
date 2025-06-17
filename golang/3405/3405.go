package main

// Problem: Count the Number of Arrays with K Matching Adjacent Elements
// Link: https://leetcode.com/problems/count-the-number-of-arrays-with-k-matching-adjacent-elements
// Difficulty: Hard
// Tags: Combinatorics, Math
// Time Complexity: O(n), where n is the length of the array
// Space Complexity: O(n), where n is the length of the array

const MOD = 1_000_000_007

var fact []int
var invFact []int
var precomputedMaxN = -1 // Tracks the max N for which factorials are computed

// power calculates (base^exp) % mod
func power(base, exp int) int {
	res := 1
	base %= MOD
	if base < 0 { // Ensure base is non-negative for modulo operations
		base += MOD
	}
	for exp > 0 {
		if exp%2 == 1 {
			res = int((int64(res) * int64(base)) % MOD) // Use int64 for intermediate product
		}
		base = int((int64(base) * int64(base)) % MOD) // Use int64 for intermediate product
		exp /= 2
	}
	return res
}

// inverse calculates modular multiplicative inverse of n under modulo MOD
// using Fermat's Little Theorem: n^(MOD-2) % MOD
func inverse(n int) int {
	return power(n, MOD-2)
}

// precomputeFactorials precomputes factorials and inverse factorials up to maxN
func precomputeFactorials(maxN int) {
	// Only recompute if maxN is greater than previously computed, or if not computed yet
	if maxN <= precomputedMaxN && precomputedMaxN != -1 {
		return
	}

	fact = make([]int, maxN+1)
	invFact = make([]int, maxN+1)

	fact[0] = 1
	invFact[0] = 1

	for i := 1; i <= maxN; i++ {
		fact[i] = int((int64(fact[i-1]) * int64(i)) % MOD) // Use int64
	}

	if maxN >= 0 && fact[maxN] == 0 && maxN >= MOD {
		// This case implies MOD is not prime or maxN is extremely large,
		// inverse(0) would be an issue. Given MOD is prime, fact[maxN] is non-zero for maxN < MOD.
		// If maxN >= MOD, fact[maxN] will be 0. C(N,K) would be 0 if N >= MOD and K < MOD.
		// This is handled by combinations if N >= MOD.
		// For this problem, N = n-1 <= 10^5 -1, which is < MOD.
	}

	if maxN >= 0 { // Ensure invFact[maxN] is safe to access
		invFact[maxN] = inverse(fact[maxN])
		for i := maxN - 1; i >= 1; i-- { // Iterate down to 1, invFact[0] is already 1
			invFact[i] = int((int64(invFact[i+1]) * int64(i+1)) % MOD) // Use int64
		}
	}
	precomputedMaxN = maxN // Update the record of what's been computed
}

// combinations calculates C(N, K) % MOD
func combinations(N, K int) int {
	if K < 0 || K > N {
		return 0
	}
	if K == 0 || K == N {
		return 1
	}
	if K > N/2 { // Optimization: C(N, K) = C(N, N-K)
		K = N - K
	}

	// Ensure factorials are computed up to N
	// This check might be redundant if precomputeFactorials is called from numberOfGoodArrays
	if N > precomputedMaxN || precomputedMaxN == -1 {
		precomputeFactorials(N)
	}

	// C(N, K) = fact[N] * invFact[K] * invFact[N-K] % MOD
	numerator := fact[N]
	denominator1 := invFact[K]
	denominator2 := invFact[N-K]

	res := (int64(numerator) * int64(denominator1)) % MOD
	res = (res * int64(denominator2)) % MOD
	return int(res)
}

// numberOfGoodArrays is the main solution function
func numberOfGoodArrays(n int, m int, k int) int {
	if k < 0 || k > n-1 {
		return 0 // k is out of bounds for matches in an array of length n
	}
	if n == 0 { // Constraint is n >= 1, but defensive
		if k == 0 {
			return 1
		} // Empty array, 0 matches
		return 0
	}
	if n == 1 { // Single element array
		if k == 0 {
			return m % MOD
		} // m choices, 0 matches
		return 0 // Cannot have matches
	}

	// Max N for C(N,K) is n-1.
	precomputeFactorials(n - 1)

	comb := combinations(n-1, k)

	power_val_m_minus_1 := 0
	exponent := n - 1 - k

	if m-1 == 0 { // Base of the power is 0 (i.e., m=1)
		if exponent == 0 {
			power_val_m_minus_1 = 1 // 0^0 = 1
		} else {
			power_val_m_minus_1 = 0 // 0^positive_exponent = 0
		}
	} else {
		power_val_m_minus_1 = power(m-1, exponent)
	}

	ans_part1 := (int64(comb) * int64(m)) % MOD
	final_ans := (ans_part1 * int64(power_val_m_minus_1)) % MOD

	if final_ans < 0 { // Ensure result is non-negative
		final_ans += MOD
	}

	return int(final_ans)
}

func countGoodArrays(n int, m int, k int) int {
	return numberOfGoodArrays(n, m, k)
}
