package main

import "fmt"

// Problem: Palindrome Number
// Link: https://leetcode.com/problems/palindrome-number/
// Difficulty: Easy
// Tags: Math, String
// Time Complexity: O(log10(n)), where n is the number of digits in the integer x.
// Space Complexity: O(1), since we are using a constant amount of space for variables.

// main demonstrates the usage of the isPalindrome function by printing the results
// of checking whether several integers are palindromes. It outputs the results for
// the numbers 121 (expected: true), -121 (expected: false), and 10 (expected: false).
func main() {
	fmt.Println(isPalindrome(121))  // true
	fmt.Println(isPalindrome(-121)) // false
	fmt.Println(isPalindrome(10))   // false
}

// isPalindrome checks whether the given integer x is a palindrome.
// A palindrome is a number that reads the same backward as forward.
// Negative numbers are not considered palindromes.
// Returns true if x is a palindrome, false otherwise.
func isPalindrome(x int) bool {
	// Negative numbers are not palindromes
	if x < 0 {
		return false
	}

	// Single digit numbers are palindromes
	if x < 10 {
		return true
	}

	// Reverse the number and compare with original
	original := x
	reversed := 0

	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return original == reversed
}
