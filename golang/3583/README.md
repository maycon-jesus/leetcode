# 3583. Count Special Triplets

**Difficulty:** Medium

## Problem Description

You are given an integer array `nums`.

A **special triplet** is defined as a triplet of indices `(i, j, k)` such that:
- `0 <= i < j < k < n`, where `n = nums.length`
- `nums[i] == nums[j] * 2`
- `nums[k] == nums[j] * 2`

Return the total number of special triplets in the array.

Since the answer may be large, return it modulo `10^9 + 7`.

## Examples

### Example 1:
```
Input: nums = [6,3,6]
Output: 1
Explanation: The only special triplet is (i, j, k) = (0, 1, 2), where:
- nums[0] = 6, nums[1] = 3, nums[2] = 6
- nums[0] = nums[1] * 2 = 3 * 2 = 6
- nums[2] = nums[1] * 2 = 3 * 2 = 6
```

### Example 2:
```
Input: nums = [0,1,0,0]
Output: 1
Explanation: The only special triplet is (i, j, k) = (0, 2, 3), where:
- nums[0] = 0, nums[2] = 0, nums[3] = 0
- nums[0] = nums[2] * 2 = 0 * 2 = 0
- nums[3] = nums[2] * 2 = 0 * 2 = 0
```

### Example 3:
```
Input: nums = [8,4,2,8,4]
Output: 2
Explanation: There are exactly two special triplets:
1. (i, j, k) = (0, 1, 3): nums[0] = 8, nums[1] = 4, nums[3] = 8
2. (i, j, k) = (1, 2, 4): nums[1] = 4, nums[2] = 2, nums[4] = 4
```

## Constraints

- `3 <= n == nums.length <= 10^5`
- `0 <= nums[i] <= 10^5`

## Solution Approach

### Algorithm: Two-Pointer with Frequency Maps

The key insight is to iterate through the array treating each element as the middle element `j` of a potential triplet, and count how many valid `i` and `k` pairs exist.

**Steps:**
1. For each position `j`, we need to find:
   - How many elements to the left equal `nums[j] * 2`
   - How many elements to the right equal `nums[j] * 2`
   - The product gives us the number of valid triplets with `j` as the middle

2. Use two frequency maps:
   - `leftFreq`: Tracks frequency of elements to the left of `j`
   - `rightFreq`: Tracks frequency of elements to the right of `j`

3. Initialize `rightFreq` with all elements, then:
   - For each position `j`, remove `nums[j]` from `rightFreq`
   - Count triplets: `leftFreq[nums[j]*2] * rightFreq[nums[j]*2]`
   - Add `nums[j]` to `leftFreq`

**Time Complexity:** O(n) - Single pass through the array
**Space Complexity:** O(n) - For the frequency maps

### Edge Cases
- Zero values: `0 * 2 = 0`, so zeros can form triplets with each other
- Large numbers: Apply modulo `10^9 + 7` to prevent overflow
- No valid triplets: Return 0
