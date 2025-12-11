# Longest Common Prefix

**Difficulty:** Easy

## Problem Description

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string `""`.

## Examples

### Example 1
```
Input: strs = ["flower","flow","flight"]
Output: "fl"
```

### Example 2
```
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
```

## Constraints

- `1 <= strs.length <= 200`
- `0 <= strs[i].length <= 200`
- `strs[i]` consists of only lowercase English letters if it is non-empty

## Solution Approach

There are several approaches to solve this problem:

### 1. Horizontal Scanning
- Start with the first string as the prefix
- Compare it with each subsequent string and trim the prefix until it matches
- Time: O(S) where S is sum of all characters, Space: O(1)

### 2. Vertical Scanning
- Compare characters from the same index across all strings
- Stop when characters don't match or reach the end of any string
- Time: O(S) where S is sum of all characters, Space: O(1)

### 3. Divide and Conquer
- Recursively find the common prefix of left and right halves
- Merge the results by comparing the two prefixes
- Time: O(S) where S is sum of all characters, Space: O(m log n) for recursion

### 4. Binary Search
- Find the minimum length string
- Binary search on the prefix length (0 to min length)
- Time: O(S log m) where m is minimum string length, Space: O(1)

