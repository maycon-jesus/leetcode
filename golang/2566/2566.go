package main

import (
	"strconv"
	"strings"
)

// Problem: Min Max Difference
// Link: https://leetcode.com/problems/min-max-difference-in-an-array/
// Difficulty: Easy
// Tags: Array, Math
// Time Complexity: O(d), where d is the number of digits in the integer num.
// Space Complexity: O(d), for storing the modified strings representing the max and min numbers.

func minMaxDifference(num int) int {
	// Convert the number to a string to manipulate digits
	numStr := strconv.Itoa(num)

	// --- Calculate the maximum possible number ---
	// Find the first digit from the left that is not '9'. This digit will be replaced by '9'.
	// If all digits are '9', no replacement is done.
	maxDigit := '0' // Initialize with '0', will be updated if a non-'9' digit is found.
	for _, digit := range numStr {
		if digit == '9' { // Skip '9's as they are already the max possible digit.
			continue
		}
		maxDigit = digit // Found the first non-'9' digit.
		break            // No need to check further for maxDigit.
	}

	// Replace all occurrences of maxDigit with '9' to get the largest possible number.
	// If maxDigit remained '0' (e.g. num was 0 or all 9s), this effectively changes nothing or changes '0' to '9'.
	maxNumStr := strings.ReplaceAll(numStr, string(maxDigit), "9")

	// --- Calculate the minimum possible number ---
	// The first digit of the number (numStr[0]) will be replaced by '0'.
	minDigit := rune(numStr[0]) // This is the digit to be replaced by '0'.

	// Replace all occurrences of minDigit with '0' to get the smallest possible number.
	minNumStr := strings.ReplaceAll(numStr, string(minDigit), "0")

	// Convert the modified strings back to integers.
	maxNum, _ := strconv.Atoi(maxNumStr) // Error handling for Atoi is omitted for brevity in LeetCode problems.
	minNum, _ := strconv.Atoi(minNumStr) // Error handling for Atoi is omitted for brevity in LeetCode problems.

	// Return the difference between the maximum and minimum possible numbers.
	return maxNum - minNum
}
