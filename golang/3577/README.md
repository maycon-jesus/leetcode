# 3577. Count the Number of Computer Unlocking Permutations

**Difficulty:** Medium

## Problem Description

You are given an array `complexity` of length `n`.

There are `n` locked computers in a room with labels from `0` to `n - 1`, each with its own unique password. The password of computer `i` has a complexity `complexity[i]`.

### Unlocking Rules

1. The password for computer `0` is already decrypted and serves as the **root**
2. All other computers must be unlocked using computer `0` or another previously unlocked computer
3. You can decrypt the password for computer `i` using the password from computer `j` where:
   - `j < i` (computer `j` has a lower index)
   - `complexity[j] < complexity[i]` (computer `j` has lower complexity)

### Task

Find the **number of valid permutations** of `[0, 1, 2, ..., (n - 1)]` that represent a valid order in which the computers can be unlocked, starting from computer `0` as the only initially unlocked one.

Since the answer may be large, return it **modulo 10⁹ + 7**.

**Note:** Computer with label `0` is always unlocked first (not necessarily the computer at the first position in the permutation).

## Examples

### Example 1

**Input:** `complexity = [1,2,3]`

**Output:** `2`

**Explanation:**

The valid permutations are:

1. `[0, 1, 2]`
   - Unlock computer 0 (root, complexity=1)
   - Unlock computer 1 using computer 0 (1 < 2)
   - Unlock computer 2 using computer 1 (2 < 3)

2. `[0, 2, 1]`
   - Unlock computer 0 (root, complexity=1)
   - Unlock computer 2 using computer 0 (1 < 3)
   - Unlock computer 1 using computer 0 (1 < 2)

### Example 2

**Input:** `complexity = [3,3,3,4,4,4]`

**Output:** `0`

**Explanation:**

There are no valid permutations. Computer 0 has complexity 3, but computers 1 and 2 also have complexity 3. Since we need `complexity[j] < complexity[i]` to unlock computer `i`, and no computer has complexity strictly less than 3 except computer 0, we cannot unlock all computers.

## Constraints

- `2 <= complexity.length <= 10⁵`
- `1 <= complexity[i] <= 10⁹`

## Solution Approach

The solution uses dynamic programming with a map to track the number of ways to unlock computers at each complexity level.

### Algorithm

1. Start with computer 0 already unlocked
2. For each subsequent computer, count how many previously unlocked computers have strictly lower complexity
3. Use a map to store the count of unlocked computers at each complexity level
4. Accumulate the total number of valid permutations modulo 10⁹ + 7

### Time Complexity

- **O(n × m)** where `n` is the number of computers and `m` is the number of unique complexity values
- In the worst case, this could be O(n²) if all complexities are unique

### Space Complexity

- **O(m)** where `m` is the number of unique complexity values
