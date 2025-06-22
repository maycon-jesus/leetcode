package main

// func longestPalindrome(s string) string {
// 	biggestPalindrome := ""
// 	for start, _ := range s {
// 		for end := len(s); end >= start; end-- {
// 			possiblePalindrome := s[start:end]
// 			if isPalindrome(possiblePalindrome) && len(possiblePalindrome) > len(biggestPalindrome) {
// 				biggestPalindrome = possiblePalindrome
// 			}
// 		}
// 	}
// 	return biggestPalindrome
// }

// func isPalindrome(s string) bool {
// 	left, right := 0, len(s)-1
// 	for left < right {
// 		if s[left] != s[right] {
// 			return false
// 		}
// 		left++
// 		right--
// 	}
// 	return true
// }
