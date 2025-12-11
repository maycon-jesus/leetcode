# Roman to Integer

**Difficulty:** Easy

## Problem Description

Roman numerals are represented by seven different symbols:

| Symbol | Value |
|--------|-------|
| I      | 1     |
| V      | 5     |
| X      | 10    |
| L      | 50    |
| C      | 100   |
| D      | 500   |
| M      | 1000  |

Roman numerals are usually written largest to smallest from left to right. However, there are special cases where subtraction is used:

- **I** can be placed before **V** (5) and **X** (10) to make 4 and 9
- **X** can be placed before **L** (50) and **C** (100) to make 40 and 90
- **C** can be placed before **D** (500) and **M** (1000) to make 400 and 900

Given a roman numeral string, convert it to an integer.

## Examples

### Example 1
```
Input: s = "III"
Output: 3
Explanation: III = 3
```

### Example 2
```
Input: s = "LVIII"
Output: 58
Explanation: L = 50, V = 5, III = 3
```

### Example 3
```
Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90, IV = 4
```

## Constraints

- `1 <= s.length <= 15`
- `s` contains only the characters `('I', 'V', 'X', 'L', 'C', 'D', 'M')`
- It is guaranteed that `s` is a valid roman numeral in the range `[1, 3999]`

## Solution Approach

The key insight is to traverse the string and:
1. If the current symbol's value is less than the next symbol's value, subtract it (e.g., IV = -1 + 5 = 4)
2. Otherwise, add the current symbol's value to the result

This handles both the standard additive cases and the subtractive notation cases in a single pass.

