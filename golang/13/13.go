package main

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	n := len(s)

	for i := 0; i < n; i++ {
		currentChar := s[i]
		currentValue := romanMap[currentChar]
		total += currentValue

		if i+1 < n {
			nextChar := s[i+1]
			nextCharValue := romanMap[nextChar]

			if currentValue < nextCharValue {
				total = total - (currentValue * 2) + nextCharValue
				i++
				continue
			}
		}
	}

	return total
}
