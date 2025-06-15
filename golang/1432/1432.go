package main

import (
	"strconv"
)

// Problem: Max Difference You Can Get From Changing an Integer
// Link: https://leetcode.com/problems/max-difference-you-can-get-from-changing-an-integer/
// Difficulty: Medium
// Tags: Math, Greedy
// Time Complexity: O(n), where n is the number of digits in the integer num.
// Space Complexity: O(n), where n is the number of digits in the integer num, due to string manipulation.

func maxDiff(num int) int {
	numStr := strconv.Itoa(num) // Convert integer to string for easy manipulation
	n := len(numStr)            // Get the number of digits

	// --- Calculate the maximum possible number ---
	// Find the first digit that is not '9' and replace all its occurrences with '9'
	maxNumBytes := []byte(numStr) // Convert string to byte slice for in-place modification
	maxDigit := byte(0)           // Variable to store the digit to be replaced
	for i := 0; i < n; i++ {
		if maxNumBytes[i] != '9' { // Find the first non-'9' digit
			maxDigit = maxNumBytes[i]
			break
		}
	}
	for i := 0; i < n; i++ {
		if maxNumBytes[i] == maxDigit { // Replace all occurrences of maxDigit with '9'
			maxNumBytes[i] = '9'
		}
	}

	maxNum, _ := strconv.Atoi(string(maxNumBytes)) // Convert modified byte slice back to integer

	// --- Calculate the minimum possible number ---
	// If the first digit is not '1', replace all its occurrences with '1'
	// Else, find the first digit (not at index 0) that is not '0' or '1', and replace all its occurrences with '0'
	minNumBytes := []byte(numStr) // Convert string to byte slice for in-place modification
	if minNumBytes[0] != '1' {
		minDigit := minNumBytes[0] // The first digit to be replaced with '1'
		for i := 0; i < n; i++ {
			if minNumBytes[i] == minDigit { // Replace all occurrences of minDigit with '1'
				minNumBytes[i] = '1'
			}
		}
	} else {
		minDigit := byte(0)
		for i := 1; i < n; i++ {
			if minNumBytes[i] != '0' && minNumBytes[i] != '1' { // Find the first digit (not at index 0) that is not '0' or '1'
				minDigit = minNumBytes[i]
				break
			}
		}
		if minDigit != 0 {
			for i := 1; i < n; i++ {
				if minNumBytes[i] == minDigit { // Replace all occurrences of minDigit with '0'
					minNumBytes[i] = '0'
				}
			}
		}
	}
	minNum, _ := strconv.Atoi(string(minNumBytes)) // Convert modified byte slice back to integer

	return maxNum - minNum // Return the difference between the maximum and minimum
}
