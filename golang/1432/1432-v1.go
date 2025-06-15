package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// // Problem: Max Difference You Can Get From Changing an Integer
// // Link: https://leetcode.com/problems/max-difference-you-can-get-from-changing-an-integer/
// // Difficulty: Medium
// // Tags: Math, String
// // Time Complexity: O(n), where n is the number of digits in the integer num.
// // Space Complexity: O(n), where n is the number of digits in the integer num, due to string manipulation.

// func maxDiff(num int) int {
// 	fmt.Println("================:")
// 	// Convert the number to a string to manipulate digits
// 	numStr := strconv.Itoa(num)

// 	// --- Calculate the maximum possible number ---
// 	// Find the first digit from the left that is not '9'. This digit will be replaced by '9'.
// 	// If all digits are '9', no replacement is done.
// 	maxDigit := '0' // Initialize with '0', will be updated if a non-'9' digit is found.
// 	for _, digit := range numStr {
// 		if digit == '9' { // Skip '9's as they are already the max possible digit.
// 			continue
// 		}
// 		maxDigit = digit // Found the first non-'9' digit.
// 		break            // No need to check further for maxDigit.
// 	}

// 	// Replace all occurrences of maxDigit with '9' to get the largest possible number.
// 	// If maxDigit remained '0' (e.g. num was 0 or all 9s), this effectively changes nothing or changes '0' to '9'.
// 	maxNumStr := strings.ReplaceAll(numStr, string(maxDigit), "9")
// 	fmt.Println("MaxNumStr:", maxNumStr)

// 	// --- Calculate the minimum possible number ---
// 	// The first digit of the number (numStr[0]) will be replaced by '0'.
// 	minDigit := rune(numStr[0]) // This is the digit to be replaced by '0'.
// 	itsZero := false            // Flag to check if we can replace the first digit with '0'.

// 	for i, digit := range numStr {
// 		if i == 0 && digit != '1' {
// 			fmt.Println("aaaa")
// 			minDigit = digit
// 			itsZero = false
// 			break
// 		} else if digit != '0' && digit != rune(numStr[0]) {
// 			fmt.Println("aaa", string(digit), i)
// 			minDigit = digit
// 			itsZero = true
// 			break
// 		} else if digit == rune(numStr[0]) && digit != '1' && digit != '0' {
// 			fmt.Println("aa")
// 			minDigit = digit
// 			if digit == '1' {
// 				fmt.Println("b")
// 				itsZero = false
// 			} else {
// 				fmt.Println("bb")
// 				itsZero = true
// 			}
// 			break
// 		}
// 	}

// 	var replacement string
// 	if itsZero {
// 		replacement = "0"
// 	} else {
// 		replacement = "1"
// 	}
// 	minNumStr := strings.ReplaceAll(numStr, string(minDigit), replacement)
// 	fmt.Println("Replacement Digit:", string(minDigit), "Replacement String:", replacement)
// 	fmt.Println("minNumStr:", minNumStr)

// 	// Convert the modified strings back to integers.
// 	maxNum, _ := strconv.Atoi(maxNumStr) // Error handling for Atoi is omitted for brevity in LeetCode problems.
// 	minNum, _ := strconv.Atoi(minNumStr) // Error handling for Atoi is omitted for brevity in LeetCode problems.
// 	fmt.Println("Max Number:", maxNum, "Min Number:", minNum)

// 	// Return the difference between the maximum and minimum possible numbers.
// 	return maxNum - minNum
// }
