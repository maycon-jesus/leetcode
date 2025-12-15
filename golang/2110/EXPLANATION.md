# Algorithm Explanation - LeetCode 2110

## Problem
Find the number of "smooth descent periods" in an array of stock prices. A smooth descent period is a contiguous sequence where each price is exactly 1 unit lower than the previous one.

## Approach

The algorithm uses a **sliding window** technique with combinatorial counting.

### Main Idea

If we have a descent sequence of size `n`, the total number of possible contiguous subsets is given by the formula:

```
n * (n + 1) / 2
```

For example, for the sequence `[3, 2, 1]` (size 3):
- Possible subsets: `[3]`, `[2]`, `[1]`, `[3,2]`, `[2,1]`, `[3,2,1]`
- Total: 3 * 4 / 2 = 6 subsets

### Step-by-Step Operation

1. **Initialization**:
   - `combo = 1`: tracks the size of the current descent sequence
   - `ans = 0`: accumulates the total number of periods

2. **Array iteration**:
   - For each position `i` (starting at 1):
     - If `prices[i-1] - prices[i] == 1`: the descent continues, increment `combo`
     - Otherwise: the sequence has ended
       - Add `combo * (combo + 1) / 2` to the result
       - Reset `combo = 1` to start a new sequence

3. **Finalization**:
   - Add the last sequence to the result

## Execution Example

For `prices = [3, 2, 1, 4]`:

```
i=1: 3-2=1 ✓ → combo=2
i=2: 2-1=1 ✓ → combo=3
i=3: 1-4≠1 ✗ → ans += 3*4/2 = 6, combo=1
Final: ans += 1*2/2 = 1
Result: 6 + 1 = 7
```

Periods found: `[3]`, `[2]`, `[1]`, `[4]`, `[3,2]`, `[2,1]`, `[3,2,1]`

## Complexity

- **Time**: O(n) - single pass through the array
- **Space**: O(1) - uses only auxiliary variables

## Why It Works

The formula `n * (n + 1) / 2` is the sum of the first `n` natural numbers, which represents all the ways to choose a contiguous subarray from a sequence of size `n`. By identifying each descent sequence and applying this formula, we count all valid periods without duplication.
