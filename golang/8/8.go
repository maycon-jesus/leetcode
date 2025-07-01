package main

func myAtoi(s string) int {
	const (
		minInt = -1 << 31
		maxInt = (1 << 31) - 1
	)

	sign := 1
	result := 0
	i := 0

	// Skip leading whitespace
	for i < len(s) && s[i] == ' ' {
		i++
	}

	// Check for sign
	if i < len(s) && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// Convert digits to integer
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')
		i++

		if result > (maxInt-digit)/10 {
			return maxInt // Overflow
		}
		result = result*10 + digit
	}

	return result * sign
}
