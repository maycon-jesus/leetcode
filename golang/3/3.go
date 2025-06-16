package main

// Problem: Longest Substring Without Repeating Characters
// Link: https://leetcode.com/problems/longest-substring-without-repeating-characters/
// Difficulty: Medium
// Tags: Hash Table, String, Sliding Window
// Time Complexity: O(n), where n is the length of the string s.
// Space Complexity: O(min(n, m)), where m is the size of the character set

// Function that returns the length of the longest substring without repeating characters
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxLength := 0
	start := 0                         // Start of the current sliding window
	charIndexMap := make(map[rune]int) // Map to store the last seen index of each character

	for end, char := range s {
		// If the character has already been seen and its last occurrence is within the current window (start <= lastSeenIndex)
		if lastSeenIndex, ok := charIndexMap[char]; ok && lastSeenIndex >= start {
			// Move the start of the window to the right of the previous occurrence of the repeated character
			start = lastSeenIndex + 1
		}

		// Update the last seen index for the current character
		charIndexMap[char] = end

		// Calculate the current window size (end - start + 1) and update the maximum size
		currentLength := end - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}
