package main

func countOdds(low int, high int) int {
	lowEven := low%2 == 0
	highEven := high%2 == 0

	if lowEven && highEven {
		return (high - low) / 2
	}

	return (high-low)/2 + 1
}
