# 1523. Count Odd Numbers in an Interval Range

**Difficulty:** Easy
**Topics:** Math

## Problem Description

Given two non-negative integers `low` and `high`, return the count of odd numbers between `low` and `high` (inclusive).

## Examples

### Example 1:
```
Input: low = 3, high = 7
Output: 3
Explanation: The odd numbers between 3 and 7 are [3,5,7].
```

### Example 2:
```
Input: low = 8, high = 10
Output: 1
Explanation: The odd numbers between 8 and 10 are [9].
```

## Constraints

- `0 <= low <= high <= 10^9`

## Solution Approach

This problem can be solved efficiently using mathematical reasoning without iterating through all numbers in the range.

### Key Insights:
1. In any consecutive sequence of numbers, odd and even numbers alternate
2. The count of odd numbers depends on whether the boundaries are odd or even
3. We can calculate the result in O(1) time using modular arithmetic

### Algorithm:
1. Check if `low` is even or odd using modulo operator
2. Check if `high` is even or odd
3. Calculate the count based on boundary conditions:
   - If both `low` and `high` are even: `(high - low) / 2`
   - Otherwise (at least one is odd): `(high - low) / 2 + 1`

### Edge Cases:
- Both numbers are even (e.g., 8 and 10)
- Both numbers are odd (e.g., 3 and 7)
- One even, one odd (e.g., 2 and 7)
- Same number (e.g., 5 and 5)

## Complexity Analysis

- **Time Complexity:** O(1) - constant time operations
- **Space Complexity:** O(1) - only a few variables used
