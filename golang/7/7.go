package main

const (
	minInt = -1 << 31
	maxInt = (1 << 31) - 1
)

func reverse(x int) int {
	var result int
	for x != 0 {
		digit := x % 10
		x /= 10

		if result > maxInt/10 || (result == maxInt/10 && digit > 7) {
			return 0 // Overflow
		}
		if result < minInt/10 || (result == minInt/10 && digit < -8) {
			return 0 // Underflow
		}

		result = result*10 + digit
	}

	return result
}
