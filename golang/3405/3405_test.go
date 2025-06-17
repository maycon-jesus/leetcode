package main

import "testing"

// MOD_TEST is the same as MOD in the main file, defined here for clarity in tests if needed,
// or tests can rely on the main file's MOD.
// const MOD_TEST = 1_000_000_007

func TestNumberOfGoodArrays(t *testing.T) {
	tests := []struct {
		name string
		n    int
		m    int
		k    int
		want int
	}{
		{"n=3, m=2, k=1", 3, 2, 1, 4},
		{"n=4, m=3, k=2", 4, 3, 2, 18},

		// Edge cases based on problem constraints (1<=n<=1000, 1<=m<=1000, 0<=k<=n-1)
		// and function's internal handling
		{"k < 0 (func should handle)", 3, 2, -1, 0},
		{"k > n-1 (func should handle)", 2, 2, 2, 0}, // n=2, n-1=1, so k=2 is > n-1

		{"n=1, k=0", 1, 5, 0, 5},
		{"n=1, k > n-1 (e.g., k=1 for n=1)", 1, 5, 1, 0}, // n-1 = 0, k=1 is > n-1

		{"m=1, k=n-1 (all elements must be 1, all match)", 3, 1, 2, 1}, // n=3, k=2. (1,1,1)
		{"m=1, k < n-1 (impossible if m=1)", 3, 1, 1, 0},               // n=3, k=1. If m=1, all are same, so k must be n-1.
		{"m=1, k=0, n=2 (impossible if m=1 and n>1)", 2, 1, 0, 0},      // If m=1, (1,1), k=1. So k=0 is impossible.

		{"n=10, m=10, k=5", 10, 10, 5, 8266860},

		{"n=2, m=1, k=1 (all match, m=1, involves 0^0)", 2, 1, 1, 1}, // (1,1)

		{"k=0 (no adjacent matches), n=3, m=2", 3, 2, 0, 2}, // C(2,0)*2*(2-1)^2 = 1*2*1 = 2. Arrays: (1,2,1), (2,1,2)
		{"k=0 (no adjacent matches), n=2, m=3", 2, 3, 0, 6}, // C(1,0)*3*(3-1)^1 = 1*3*2 = 6

		{"k=n-1 (all adjacent matches), n=3, m=2", 3, 2, 2, 2}, // C(2,2)*2*(2-1)^0 = 1*2*1 = 2. Arrays: (1,1,1), (2,2,2)
		{"k=n-1 (all adjacent matches), n=4, m=5", 4, 5, 3, 5}, // C(3,3)*5*(5-1)^0 = 1*5*1 = 5

		// Constraints: 1 <= n <= 1000, 1 <= m <= 1000, 0 <= k <= n - 1
		{"max_n_k=n-1, m=2", 1000, 2, 999, 2}, // C(999,999) * 2 * (1)^0 = 1 * 2 * 1 = 2
		{"max_n_k=0, m=2", 1000, 2, 0, 2},     // C(999,0) * 2 * (1)^999 = 1 * 2 * 1 = 2
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset global precomputation state for each test run
			precomputedMaxN = -1
			fact = nil
			invFact = nil

			got := numberOfGoodArrays(tt.n, tt.m, tt.k)
			if got != tt.want {
				t.Errorf("numberOfGoodArrays(n=%d, m=%d, k=%d) = %v, want %v", tt.n, tt.m, tt.k, got, tt.want)
			}
		})
	}
}

func TestCombinations(t *testing.T) {
	tests := []struct {
		name string
		N    int
		K    int
		want int
	}{
		{"C(5,2)", 5, 2, 10},
		{"C(5,0)", 5, 0, 1},
		{"C(5,5)", 5, 5, 1},
		{"C(5,3)", 5, 3, 10}, // C(N,K) = C(N, N-K)
		{"C(10,1)", 10, 1, 10},
		{"C(10,9)", 10, 9, 10},
		{"C(0,0)", 0, 0, 1},
		{"C(5,6) K>N", 5, 6, 0},
		{"C(5,-1) K<0", 5, -1, 0},
		{"C(20,10)", 20, 10, 184756}, // A larger combination
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			precomputedMaxN = -1 // Reset state
			fact = nil
			invFact = nil

			got := combinations(tt.N, tt.K)
			if got != tt.want {
				// Re-precompute for error message context if N was valid
				// This is mainly for debugging the test itself if it fails due to factorial issues.
				if tt.N >= 0 && (precomputedMaxN == -1 || tt.N > precomputedMaxN) {
					// precomputeFactorials(tt.N) // Let combinations handle its own precomputation
				}
				t.Errorf("combinations(N=%d, K=%d) = %v, want %v. (Factorials were computed up to %d by the call to combinations)", tt.N, tt.K, got, tt.want, precomputedMaxN)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name string
		base int
		exp  int
		want int
	}{
		{"2^3", 2, 3, 8},
		{"5^0", 5, 0, 1},
		{"0^5 (0 to positive power)", 0, 5, 0},
		{"0^0 (as per problem context)", 0, 0, 1},
		{"7^1", 7, 1, 7},
		{"3^5", 3, 5, 243}, // 243 % MOD = 243
		{"(3^3)%MOD", 3, 3, 27},
		{"large_base_large_exp", 12345, 678, 0}, // Placeholder, will be calculated
		{"base < 0, exp even", -2, 2, 4},
		{"base < 0, exp odd (-2)^3", -2, 3, (MOD - 8 + MOD) % MOD}, // Ensure positive result for -8 % MOD
	}

	// Calculate expected for "large_base_large_exp"
	// This uses the same logic as power, so it's more of an integration check for this specific case
	expectedPowerLarge := 1
	baseVal := 12345
	expVal := 678

	tempBase := baseVal % MOD
	if tempBase < 0 { // Ensure base is non-negative for modulo operations
		tempBase += MOD
	}

	for curExp := expVal; curExp > 0; curExp /= 2 {
		if curExp%2 == 1 {
			expectedPowerLarge = int((int64(expectedPowerLarge) * int64(tempBase)) % MOD)
		}
		tempBase = int((int64(tempBase) * int64(tempBase)) % MOD)
	}
	tests[len(tests)-3].want = expectedPowerLarge // Update the "large_base_large_exp" test case

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := power(tt.base, tt.exp)
			if got != tt.want {
				t.Errorf("power(base=%d, exp=%d) = %v, want %v", tt.base, tt.exp, got, tt.want)
			}
		})
	}
}
