# Longest Palindromic Substring

**LeetCode Problem 5**

## Problem Statement
Given a string `s`, return the longest palindromic substring in `s`.

### Example 1:
```
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
```

### Example 2:
```
Input: s = "cbbd"
Output: "bb"
```

## Constraints
- 1 <= s.length <= 1000
- s consists of only digits and English letters.

## Solution Approach
The solution uses the expand-around-center technique to find the longest palindromic substring in O(n^2) time and O(1) space.

## Files
- `5.go`: Go implementation optimized of the solution
- `5-v1.go`: Original Go implementation of the solution
- `5_test.go`: Unit tests for the solution
