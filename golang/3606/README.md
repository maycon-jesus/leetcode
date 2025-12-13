# 3606. Coupon Code Validator

![Difficulty: Easy](https://img.shields.io/badge/Difficulty-Easy-green)
![Status: Solved](https://img.shields.io/badge/Status-Solved-success)
![Premium](https://img.shields.io/badge/Premium-locked-orange)

## Problem Description

You are given three arrays of length `n` that describe the properties of `n` coupons: `code`, `businessLine`, and `isActive`. The i-th coupon has:

- **code[i]**: a string representing the coupon identifier
- **businessLine[i]**: a string denoting the business category of the coupon
- **isActive[i]**: a boolean indicating whether the coupon is currently active

### Validation Rules

A coupon is considered **valid** if all of the following conditions hold:

1. `code[i]` is non-empty and consists only of alphanumeric characters (a-z, A-Z, 0-9) and underscores (_)
2. `businessLine[i]` is one of: `"electronics"`, `"grocery"`, `"pharmacy"`, `"restaurant"`
3. `isActive[i]` is `true`

### Task

Return an array of the codes of all valid coupons, sorted:
1. First by `businessLine` in the order: **electronics → grocery → pharmacy → restaurant**
2. Then by `code` in **lexicographical (ascending) order** within each category

## Examples

### Example 1

```
Input:
  code = ["SAVE20","","PHARMA5","SAVE@20"]
  businessLine = ["restaurant","grocery","pharmacy","restaurant"]
  isActive = [true,true,true,true]

Output: ["PHARMA5","SAVE20"]
```

**Explanation:**
- First coupon: valid ✓
- Second coupon: empty code (invalid ✗)
- Third coupon: valid ✓
- Fourth coupon: special character @ (invalid ✗)

### Example 2

```
Input:
  code = ["GROCERY15","ELECTRONICS_50","DISCOUNT10"]
  businessLine = ["grocery","electronics","invalid"]
  isActive = [false,true,true]

Output: ["ELECTRONICS_50"]
```

**Explanation:**
- First coupon: inactive (invalid ✗)
- Second coupon: valid ✓
- Third coupon: invalid business line (invalid ✗)

## Constraints

- `n == code.length == businessLine.length == isActive.length`
- `1 <= n <= 100`
- `0 <= code[i].length, businessLine[i].length <= 100`
- `code[i]` and `businessLine[i]` consist of printable ASCII characters
- `isActive[i]` is either `true` or `false`

## Solution Approach

### Algorithm

1. **Validation Phase**: Filter valid coupons by checking:
   - Code is non-empty
   - Code contains only alphanumeric characters and underscores
   - Business line is one of the four valid categories
   - Coupon is active

2. **Sorting Phase**: Sort valid coupons using a custom comparator:
   - Primary sort: by business line order (electronics < grocery < pharmacy < restaurant)
   - Secondary sort: lexicographically by code within the same business line

3. **Optimization**: Store indices instead of duplicating strings during sorting

### Implementation Details

- Uses `map[string]bool` for O(1) business line validation
- Uses `map[string]int` for efficient business line ordering
- Custom sort function with `slices.SortFunc` for combined sorting logic
- Manual character validation for coupon code (alphanumeric + underscore)

## Complexity Analysis

- **Time Complexity**: O(n × m + n log n)
  - n × m: validating n coupons with average length m
  - n log n: sorting valid coupons
- **Space Complexity**: O(n)
  - Storage for valid indices and result array

## Usage

```go
code := []string{"SAVE20", "", "PHARMA5", "SAVE@20"}
businessLine := []string{"restaurant", "grocery", "pharmacy", "restaurant"}
isActive := []bool{true, true, true, true}

result := validateCoupons(code, businessLine, isActive)
// Output: ["PHARMA5", "SAVE20"]
```

## Key Takeaways

- Use maps for efficient category validation and ordering
- Sort by indices to avoid unnecessary string duplication
- Custom comparators enable multi-level sorting in a single pass
- Character validation can be done efficiently with range checks

