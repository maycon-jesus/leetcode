# 2014. Longest Subsequence Repeated k Times

**Difficulty:** Hard  
**Topics:** String, Dynamic Programming, Backtracking, Enumeration  

## Problem Description

You are given a string `s` of length `n`, and an integer `k`. You are tasked to find the longest subsequence repeated `k` times in string `s`.

A **subsequence** is a string that can be derived from another string by deleting some or no characters without changing the order of the remaining characters.

A subsequence `seq` is **repeated k times** in the string `s` if `seq * k` is a subsequence of `s`, where `seq * k` represents a string constructed by concatenating `seq` k times.

- For example, `"bba"` is repeated 2 times in the string `"bababcba"`, because the string `"bbabba"`, constructed by concatenating `"bba"` 2 times, is a subsequence of the string `"bababcba"`.

Return the **longest subsequence repeated k times** in string `s`. If multiple such subsequences are found, return the **lexicographically largest** one. If there is no such subsequence, return an empty string.

## Examples

### Example 1:
```
Input: s = "letsleetcode", k = 2
Output: "let"
Explanation: There are two longest subsequences repeated 2 times: "let" and "ete".
"let" is the lexicographically largest one.
```

### Example 2:
```
Input: s = "bb", k = 2
Output: "b"
Explanation: The longest subsequence repeated 2 times is "b".
```

### Example 3:
```
Input: s = "ab", k = 2
Output: ""
Explanation: There is no subsequence repeated 2 times. Empty string is returned.
```

## Constraints

- `n == s.length`
- `2 <= n, k <= 2000`
- `2 <= n < k * 8`
- `s` consists of lowercase English letters.

## Solution Approach

The solution uses a **BFS (Breadth-First Search)** approach to systematically build subsequences:

### Algorithm Steps:

1. **Character Frequency Analysis**: Count the frequency of each character in the string `s`.

2. **Filter Valid Characters**: Only consider characters that appear at least `k` times, since they're the only ones that can potentially be part of a subsequence repeated `k` times.

3. **Sort Characters**: Sort valid characters in descending order to ensure we find the lexicographically largest result.

4. **BFS Construction**: Use BFS to build subsequences incrementally:
   - Start with an empty string
   - For each current subsequence, try appending each valid character
   - Check if the new subsequence can be repeated `k` times
   - Keep track of the longest valid subsequence found

5. **Validation Function**: For each candidate subsequence, verify:
   - Each character appears enough times (frequency check)
   - The subsequence repeated `k` times is actually a subsequence of the original string

### Key Insights:

- **Frequency Check**: A character must appear at least `k` times to be part of any valid subsequence
- **Lexicographic Ordering**: By processing characters in descending order, we naturally find the lexicographically largest result
- **BFS Guarantees**: BFS ensures we find the longest possible subsequence

## Time Complexity

- **Time**: O(n × m × k) where n is the length of string s, m is the number of valid subsequences generated, and k is the repetition factor
- **Space**: O(m) for the BFS queue and character frequency map

## Running the Solution

```bash
cd golang/2014
go test -v
```

The test file `2014_test.go` contains test cases that verify the solution against the provided examples.
