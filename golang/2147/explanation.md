# Algorithm Explanation - Problem 2147

## Problem Overview

Given a corridor represented by a string where:
- `'S'` = Seat
- `'P'` = Plant

We need to divide the corridor into sections where **each section has exactly 2 seats**. The goal is to count how many different ways we can make this division.

## Key Insight

The key to solving this problem is to realize that:

1. **Each section must have exactly 2 seats** - this is fixed
2. **Between each pair of sections**, we can have plants
3. **The divider can be placed at any position** among the plants

### Visual Example

For `corridor = "SSPPSPS"`:
```
S S | P P | S P S
[---]     [-----]
pair 1    pair 2
```

Between the first pair (SS) and the second pair (SPS), we have "PP". The divider can be placed:
- Before the first P: `SS|PPSPS`
- Between the two Ps: `SSP|PSPS`
- After the second P: `SSPP|SPS`

That is, **3 possible positions** = `plantCount + 1` = `2 + 1 = 3`

## Step-by-Step Algorithm

### 1. Control Variables

```go
ways := 1           // Number of ways (result)
seatCount := 0      // Seats in current section
plantCount := 0     // Plants between pairs of sections
totalSeatCount := 0 // Total seats in corridor
```

### 2. Iterating Through the Corridor

For each character in the corridor:

#### Case 1: We Find a Seat (`'S'`)

```go
if ch == 'S' {
    // If we already have 2 seats in current section
    if seatCount == 2 {
        // Multiply by the ways to divide (plants + 1)
        ways = ways * (plantCount + 1) % MOD
        // Reset counters for next section
        seatCount = 0
        plantCount = 0
    }

    seatCount++
    totalSeatCount++
}
```

**What happens:**
- When `seatCount == 2`, we complete a section
- We calculate how many positions we can place the divider: `plantCount + 1`
- We multiply the total result by these possibilities
- We reset the counters to start the next section

#### Case 2: We Find a Plant (`'P'`)

```go
else if seatCount == 2 {
    plantCount++
}
```

**What happens:**
- We only count plants **after completing a pair of seats**
- These plants represent the positions where we can place dividers

### 3. Final Validation

```go
if totalSeatCount == 0 || totalSeatCount%2 != 0 {
    return 0
}
```

**Impossible cases:**
- `totalSeatCount == 0`: No seats available
- `totalSeatCount % 2 != 0`: Odd number of seats (can't form exact pairs)

## Detailed Example

For `corridor = "SSPPSPS"`:

| Iteration | Char | seatCount | plantCount | ways | Action |
|-----------|------|-----------|------------|------|--------|
| 1 | S | 1 | 0 | 1 | First seat |
| 2 | S | 2 | 0 | 1 | Pair complete! |
| 3 | P | 2 | 1 | 1 | Plant between pairs |
| 4 | P | 2 | 2 | 1 | Another plant |
| 5 | S | 1 | 0 | 3 | `ways = 1 * (2+1) = 3`, new pair |
| 6 | P | 2 | 0 | 3 | Second seat of pair |
| 7 | S | 2 | 0 | 3 | Pair complete! |

**Result:** 3 different ways

## Why `(plantCount + 1)`?

If we have `n` consecutive plants between two pairs of seats:
```
...SS [P P P] SS...
```

We can place the divider in `n + 1` positions:
```
|P P P    (before all)
P|P P     (between 1st and 2nd)
P P|P     (between 2nd and 3rd)
P P P|    (after all)
```

## Complexity

- **Time:** O(n) - single pass through the string
- **Space:** O(1) - only auxiliary variables

## Special Cases

1. **Empty or very short corridor:** `return 0`
2. **Odd number of seats:** `return 0` (impossible to form pairs)
3. **No seats:** `return 0`
4. **Exactly 2 seats:** `return 1` (only one way)

## Modular Arithmetic

We use `MOD = 10^9 + 7` because:
- The result can be extremely large
- The `% MOD` operation keeps the number within bounds
- It's applied at each multiplication to avoid overflow
