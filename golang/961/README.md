# 961. N-Repeated Element in Size 2N Array

**Difficulty:** Easy
**Topics:** Array, Hash Table

## Problem Description

You are given an integer array `nums` with the following properties:

- `nums.length == 2 * n`
- `nums` contains `n + 1` unique elements
- Exactly one element of `nums` is repeated `n` times

Return the element that is repeated `n` times.

### Examples

**Example 1:**
```
Input: nums = [1,2,3,3]
Output: 3
```

**Example 2:**
```
Input: nums = [2,1,2,5,3,2]
Output: 2
```

**Example 3:**
```
Input: nums = [5,1,5,2,5,3,5,4]
Output: 5
```

### Constraints

- `2 <= n <= 5000`
- `nums.length == 2 * n`
- `0 <= nums[i] <= 10^4`
- `nums` contains `n + 1` unique elements and one of them is repeated exactly `n` times

## Solution

### Approach: Hash Set

The solution uses a hash set to track seen elements. Since exactly one element appears `n` times and all others appear at most once, we can iterate through the array and:

1. Check if the current element has been seen before
2. If yes, return it (this is the n-repeated element)
3. If no, add it to the seen set and continue

This works because the repeated element will be the first element we encounter twice.

### Complexity Analysis

- **Time Complexity:** O(n) - We traverse the array once, and hash set operations are O(1) on average
- **Space Complexity:** O(n) - In the worst case, we store n+1 unique elements in the hash set

### Implementation

```go
func repeatedNTimes(nums []int) int {
    seen := make(map[int]bool)
    for _, num := range nums {
        if seen[num] {
            return num
        }
        seen[num] = true
    }
    return -1
}
```

## Running Tests

To run the unit tests:

```bash
go test -v
```

To run tests with coverage:

```bash
go test -v -cover
```

## Test Cases

The test suite includes:
- Basic examples from the problem description
- Edge cases (minimal array size, consecutive repeated elements)
- Various element positions (beginning, middle, end)
- Different number ranges (negative, large positive numbers)

