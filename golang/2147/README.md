# 2147. Number of Ways to Divide a Long Corridor

**Difficulty:** Hard

## Problem Description

Along a long library corridor, there is a line of seats and decorative plants. You are given a 0-indexed string `corridor` of length `n` consisting of letters `'S'` and `'P'` where:
- `'S'` represents a seat
- `'P'` represents a plant

One room divider has already been installed to the left of index 0, and another to the right of index `n - 1`. Additional room dividers can be installed. For each position between indices `i - 1` and `i` (1 <= i <= n - 1), at most one divider can be installed.

**Task:** Divide the corridor into non-overlapping sections, where each section has exactly **two seats** with any number of plants. There may be multiple ways to perform the division. Two ways are different if there is a position with a room divider installed in the first way but not in the second way.

Return the number of ways to divide the corridor. Since the answer may be very large, return it modulo `10^9 + 7`. If there is no way, return `0`.

## Examples

### Example 1
```
Input: corridor = "SSPPSPS"
Output: 3
```
**Explanation:** There are 3 different ways to divide the corridor. Each section has exactly two seats.

### Example 2
```
Input: corridor = "PPSPSP"
Output: 1
```
**Explanation:** There is only 1 way to divide the corridor, by not installing any additional dividers. Installing any would create some section that does not have exactly two seats.

### Example 3
```
Input: corridor = "S"
Output: 0
```
**Explanation:** There is no way to divide the corridor because there will always be a section that does not have exactly two seats.

## Constraints

- `n == corridor.length`
- `1 <= n <= 10^5`
- `corridor[i]` is either `'S'` or `'P'`

## Solution Approach

The key insight is that each section must have exactly 2 seats. The problem becomes counting the number of ways to place dividers between consecutive pairs of seats.

### Algorithm

1. **Count total seats**: If the total number of seats is 0 or odd, return 0 (impossible to divide into pairs)

2. **Count plants between seat pairs**: After finding each pair of seats (seats 2, 3), (4, 5), etc., count the number of plants before the next pair begins

3. **Multiply choices**: For each group of plants between pairs, we have `(plantCount + 1)` ways to place a divider (before the first plant, between any plants, or after the last plant)

4. **Edge cases**:
   - Total seats = 0 → return 0
   - Total seats is odd → return 0
   - Total seats = 2 → return 1

### Complexity

- **Time Complexity:** O(n) - single pass through the corridor
- **Space Complexity:** O(1) - only using constant extra space

## Implementation Notes

The solution uses modular arithmetic (`MOD = 10^9 + 7`) to prevent integer overflow, as the result can be very large.

