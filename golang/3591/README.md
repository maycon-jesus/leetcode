# 3591. Check if Any Element Has Prime Frequency

## Problem Description

Given an integer array `nums`, return `true` if the frequency of any element of the array is prime, otherwise, return `false`.

The frequency of an element `x` is the number of times it occurs in the array.
A prime number is a natural number greater than 1 with only two factors, 1 and itself.

---

## Examples

### Example 1
- **Input:** `nums = [1,2,3,4,5,4]`
- **Output:** `true`
- **Explanation:**
  - The number 4 appears twice. 2 is a prime number.

### Example 2
- **Input:** `nums = [1,2,3,4,5]`
- **Output:** `false`
- **Explanation:**
  - All elements appear once. 1 is not a prime number.

### Example 3
- **Input:** `nums = [2,2,2,4,4]`
- **Output:** `true`
- **Explanation:**
  - The number 2 appears three times (3 is prime), and 4 appears twice (2 is prime).

---

## Constraints
- `1 <= nums.length <= 100`
- `0 <= nums[i] <= 100`

---

## Hints
- Count the frequency of each element in the array.
- Check if any frequency is a prime number.

---

## Related Topics
- Array
- Hash Table
- Math
