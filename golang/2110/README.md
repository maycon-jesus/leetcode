# 2110. Number of Smooth Descent Periods of a Stock

**Difficulty:** Medium

## Problem Description

You are given an integer array `prices` representing the daily price history of a stock, where `prices[i]` is the stock price on the ith day.

A **smooth descent period** of a stock consists of one or more contiguous days such that the price on each day is lower than the price on the preceding day by **exactly 1**. The first day of the period is exempted from this rule.

Return the number of smooth descent periods.

## Examples

### Example 1:
```
Input: prices = [3,2,1,4]
Output: 7
```
**Explanation:** There are 7 smooth descent periods:
- `[3]`, `[2]`, `[1]`, `[4]`, `[3,2]`, `[2,1]`, and `[3,2,1]`
- Note that a period with one day is a smooth descent period by the definition.

### Example 2:
```
Input: prices = [8,6,7,7]
Output: 4
```
**Explanation:** There are 4 smooth descent periods: `[8]`, `[6]`, `[7]`, and `[7]`
- Note that `[8,6]` is **not** a smooth descent period as 8 - 6 ≠ 1.

### Example 3:
```
Input: prices = [1]
Output: 1
```
**Explanation:** There is 1 smooth descent period: `[1]`

## Constraints

- `1 <= prices.length <= 10^5`
- `1 <= prices[i] <= 10^5`

## Solution

### Approach: Sliding Window with Combinatorial Counting

The key insight is that for a descent sequence of length `n`, the total number of contiguous subarrays is `n * (n + 1) / 2`.

**Algorithm:**
1. Track the length of the current descent sequence (`combo`)
2. When the descent breaks or we reach the end, calculate all subarrays using the formula
3. Reset the counter and continue

### Complexity Analysis

- **Time Complexity:** O(n) - Single pass through the array
- **Space Complexity:** O(1) - Only using constant extra space

### Implementation

See [2110.go](2110.go) for the complete implementation.

For a detailed explanation of the algorithm, see [EXPLANATION.md](EXPLANATION.md).

## Topics

- Array
- Math
- Dynamic Programming
- Sliding Window