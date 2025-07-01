package main

func possibleStringCount(word string) int {
	len := len(word)
	lastletter := word[0]
	count := 1
	for i := 1; i < len; i++ {
		if word[i] == lastletter {
			count++
		} else {
			lastletter = word[i]
		}
	}
	return count
}
