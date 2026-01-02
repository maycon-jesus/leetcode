# Detailed Explanation - Problem 961

## Understanding the Problem

Given an array `nums` of size `2n`, we know that:
- There are `n + 1` unique elements
- Exactly **one** element appears `n` times
- All other elements appear **exactly once**

Our task is to find the element that appears `n` times.

### Why is this important?

Since we have an array of size `2n` with `n + 1` unique elements, and one of them appears `n` times:
- If we sum: `n` (occurrences of the repeated element) + `n` (occurrences of the other n unique elements) = `2n` ✓
- This guarantees that **only one** element repeats, and it will be the **first** one we encounter twice

## Visualization with Examples

### Example 1: `[1, 2, 3, 3]`
```
n = 2 (array size = 4 = 2*2)
Unique elements: 1, 2, 3 (total: 3 = n+1)
Repeated element: 3 (appears 2 times = n times)

Iteration:
1. num = 1 → not seen before → add to set {1}
2. num = 2 → not seen before → add to set {1, 2}
3. num = 3 → not seen before → add to set {1, 2, 3}
4. num = 3 → ALREADY SEEN! → return 3 ✓
```

### Example 2: `[2, 1, 2, 5, 3, 2]`
```
n = 3 (array size = 6 = 2*3)
Unique elements: 1, 2, 3, 5 (total: 4 = n+1)
Repeated element: 2 (appears 3 times = n times)

Iteration:
1. num = 2 → not seen before → add to set {2}
2. num = 1 → not seen before → add to set {2, 1}
3. num = 2 → ALREADY SEEN! → return 2 ✓
```

### Example 3: `[5, 1, 5, 2, 5, 3, 5, 4]`
```
n = 4 (array size = 8 = 2*4)
Unique elements: 1, 2, 3, 4, 5 (total: 5 = n+1)
Repeated element: 5 (appears 4 times = n times)

Iteration:
1. num = 5 → not seen before → add to set {5}
2. num = 1 → not seen before → add to set {5, 1}
3. num = 5 → ALREADY SEEN! → return 5 ✓
```

## The Solution: Hash Set

### Step-by-Step Algorithm

1. **Create an empty set** to track seen elements
2. **Iterate through the array**:
   - If the current element is already in the set → **return it** (we found the repeated one!)
   - Otherwise → add it to the set
3. Return -1 (if not found, but this never happens with valid inputs)

### Why Hash Set?

A Hash Set (or map in Go) provides:
- **O(1) lookup**: checking if an element exists is instantaneous
- **O(1) insertion**: adding a new element is instantaneous
- **Memory efficient**: we only store unique elements

### Go Implementation

```go
func repeatedNTimes(nums []int) int {
    seen := make(map[int]bool)  // Create a map to track seen elements

    for _, num := range nums {   // Iterate through each number
        if seen[num] {           // If we've seen this number before
            return num           // It's the element repeated n times!
        }
        seen[num] = true         // Mark as seen
    }

    return -1  // Never reached with valid inputs
}
```

## Complexity Analysis

### Time Complexity: O(n)
- **Best case**: O(2) - we find two equal elements in the first two positions
- **Average case**: O(n) - we traverse approximately half the array
- **Worst case**: O(2n) = O(n) - we traverse almost the entire array

Since each hash set operation (lookup and insertion) is O(1), the total complexity is linear.

### Space Complexity: O(n)
- **Best case**: O(1) - we find the repetition immediately
- **Worst case**: O(n+1) = O(n) - we store all n+1 unique elements

## Alternative Optimizations

### Approach 1: Brute Force (Not Recommended)
```go
// Compare each element with all others
// Time: O(n²), Space: O(1)
```
**Disadvantage**: Too slow for large arrays.

### Approach 2: Sorting (Possible but Inefficient)
```go
// Sort the array and look for adjacent equal elements
// Time: O(n log n), Space: O(1) or O(n) depending on sorting algorithm
```
**Disadvantage**: Slower than hash set and modifies the original array.

### Approach 3: Hash Set (Current Solution) ✓
```go
// Our solution
// Time: O(n), Space: O(n)
```
**Advantage**: Linear time, single pass, simple code.

## Special Cases and Edge Cases

### Minimum Case
```go
Input: [1, 1]
n = 1, array size = 2
Output: 1
```

### Consecutive Elements
```go
Input: [9, 9, 8, 7]
Output: 9  // Found at second position
```

### Scattered Elements
```go
Input: [7, 1, 2, 3, 7, 4, 5, 6]
Output: 7  // Found at fifth position
```

### Negative Numbers
```go
Input: [-1, -1, 0, 1]
Output: -1  // Works with any integer
```

## Key Go Concepts Used

### 1. Map as Set
```go
seen := make(map[int]bool)
```
Go doesn't have a native `Set` type, so we use `map[int]bool` where:
- The key is the element
- The boolean value indicates presence (always `true`)

### 2. Range Loop
```go
for _, num := range nums
```
- First variable (ignored with `_`): index
- Second variable (`num`): element value

### 3. Zero Value
```go
if seen[num]  // If doesn't exist, returns false (zero value of bool)
```
In Go, accessing a non-existent key in a map returns the "zero value" of the type (false for bool).

## Conclusion

This solution is **optimal** for the problem because:
- ✓ Linear time O(n)
- ✓ Single pass through the array
- ✓ Clean and easy to understand code
- ✓ Works for all test cases
- ✓ Uses efficient data structures

The simplicity of the "first element seen twice" logic perfectly leverages the mathematical properties of the problem!
