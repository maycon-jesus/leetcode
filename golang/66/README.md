# 66. Plus One

**Difficulty:** Easy

## Problem Description

You are given a large integer represented as an integer array `digits`, where each `digits[i]` is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.

## Examples

### Example 1
```
Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].
```

### Example 2
```
Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Incrementing by one gives 4321 + 1 = 4322.
Thus, the result should be [4,3,2,2].
```

### Example 3
```
Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].
```

### Example 4
```
Input: digits = [9,9,9]
Output: [1,0,0,0]
Explanation: The array represents the integer 999.
Incrementing by one gives 999 + 1 = 1000.
Thus, the result should be [1,0,0,0].
```

## Constraints

- `1 <= digits.length <= 100`
- `0 <= digits[i] <= 9`
- `digits` does not contain any leading 0's

## Solution Approach

This problem simulates **addition with carry**, similar to how we add numbers manually:

1. Start from the rightmost digit (least significant)
2. Add 1 to it and handle any carry
3. Propagate the carry to the left if needed
4. If there's still a carry after processing all digits, prepend it to the result

### Algorithm Steps

```
1. Create a result array with length = digits.length + 1 (to handle potential overflow)
2. Initialize carry = 1 (since we're adding 1)
3. Iterate from right to left through the digits:
   a. Add current digit + carry
   b. Store (sum % 10) in result array
   c. Update carry = sum / 10
4. If carry > 0 after the loop, set result[0] = carry
5. Otherwise, return result[1:] (skip the leading zero)
```

### Key Cases to Handle

- **Normal case:** `[1,2,3]` → `[1,2,4]` (no carry propagation)
- **Carry propagation:** `[1,9,9]` → `[2,0,0]` (carry propagates through multiple digits)
- **Overflow case:** `[9,9,9]` → `[1,0,0,0]` (need extra digit)

## Complexity Analysis

- **Time Complexity:** O(n)
  - We iterate through the array once, where n is the number of digits
  - Each operation (addition, modulo, division) is O(1)

- **Space Complexity:** O(n)
  - We create a new array of size n+1
  - The space needed grows linearly with input size

## Key Insights

- The problem is essentially simulating manual addition with carry
- We need to handle the edge case where all digits are 9 (requires an extra digit)
- Pre-allocating array with size n+1 avoids costly array resizing
- Using modulo (%) and integer division (/) elegantly handles carry calculation
- Early termination is possible: if carry becomes 0, remaining digits stay the same

## Related Topics

- Array
- Math
- Simulation
