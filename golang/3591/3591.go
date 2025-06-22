package main

func checkPrimeFrequency(nums []int) bool {
	occurrences := make(map[int]int)
	for _, num := range nums {
		occurrences[num]++
	}
	for _, count := range occurrences {
		if isPrime(count) {
			return true
		}
	}
	return false
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
