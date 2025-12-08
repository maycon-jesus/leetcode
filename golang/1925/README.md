# 1925. Count Square Sum Triples

**Difficulty:** Easy

## Problem Description

A **square triple** `(a, b, c)` is a triple where `a`, `b`, and `c` are integers and satisfy the Pythagorean theorem:

```
a² + b² = c²
```

Given an integer `n`, return the number of square triples such that `1 <= a, b, c <= n`.

## Examples

### Example 1

```
Input: n = 5
Output: 2
Explanation: The square triples are (3,4,5) and (4,3,5).
```

### Example 2

```
Input: n = 10
Output: 4
Explanation: The square triples are (3,4,5), (4,3,5), (6,8,10), and (8,6,10).
```

## Constraints

- `1 <= n <= 250`

## Solution Approach

### Algorithm

The solution uses a brute force approach with optimization:

1. **Triple Nested Loop**: Iterate through all possible combinations of `a`, `b`, and `c`
   - `a` ranges from `1` to `n`
   - `b` ranges from `1` to `n`
   - `c` ranges from `1` to `n`

2. **Check Pythagorean Condition**: For each triple `(a, b, c)`, check if `a² + b² = c²`

3. **Count Valid Triples**: Increment counter when a valid triple is found

### Time Complexity

- **O(n³)**: Three nested loops, each running up to `n` times

### Space Complexity

- **O(1)**: Only using constant extra space for the counter

### Key Insights

- The problem counts ordered pairs, so `(3,4,5)` and `(4,3,5)` are counted separately
- Common Pythagorean triples include: `(3,4,5)`, `(5,12,13)`, `(8,15,17)`, and their multiples
- Since `a² + b² = c²`, we know that `c` must be greater than both `a` and `b`

### Optimization Possibilities

- Start `c` loop from `max(a, b)` instead of `1` to reduce iterations
- Use a set of perfect squares to avoid repeated square calculations
- Break early when `a² + b²` exceeds the maximum possible `c²`
