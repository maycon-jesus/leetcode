# 20. Valid Parentheses

**Difficulty:** Easy

## Problem Description

Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

## Examples

### Example 1
```
Input: s = "()"
Output: true
```

### Example 2
```
Input: s = "()[]{}"
Output: true
```

### Example 3
```
Input: s = "(]"
Output: false
Explanation: Mismatched bracket types
```

### Example 4
```
Input: s = "([])"
Output: true
Explanation: Properly nested brackets
```

### Example 5
```
Input: s = "([)]"
Output: false
Explanation: Brackets are not closed in the correct order
```

## Constraints

- `1 <= s.length <= 10^4`
- `s` consists of parentheses only `'()[]{}'`

## Solution Approach

This problem is best solved using a **stack** data structure:

1. Iterate through each character in the string
2. If the character is an opening bracket `(`, `{`, or `[`, push it onto the stack
3. If the character is a closing bracket `)`, `}`, or `]`:
   - Check if the stack is empty (no matching opening bracket) → return false
   - Pop from the stack and verify it matches the closing bracket type
   - If it doesn't match → return false
4. After processing all characters, the stack should be empty (all brackets matched)

### Algorithm Steps

```
1. Create an empty stack
2. For each character c in string s:
   a. If c is opening bracket: push to stack
   b. If c is closing bracket:
      - If stack is empty: return false
      - Pop from stack and check if it matches c
      - If no match: return false
3. Return true if stack is empty, false otherwise
```

## Complexity Analysis

- **Time Complexity:** O(n)
  - We iterate through the string once, where n is the length of the string
  - Each stack operation (push/pop) is O(1)

- **Space Complexity:** O(n)
  - In the worst case, all characters are opening brackets
  - The stack would contain all n characters

## Key Insights

- Stack is ideal for matching pairs and maintaining order
- LIFO (Last In, First Out) property naturally handles nested brackets
- Early termination when mismatch is found improves average performance
- Edge cases: empty string (valid), odd length (always invalid), unmatched opening brackets

## Related Topics

- Stack
- String
- Data Structures

