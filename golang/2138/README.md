# Divide a String Into Groups of Size k

LeetCode 2138. Easy

---

## Problem Statement

A string `s` can be partitioned into groups of size `k` using the following procedure:

1. The first group consists of the first `k` characters of the string, the second group consists of the next `k` characters, and so on. Each element can be a part of exactly one group.
2. For the last group, if the string does not have `k` characters remaining, a character `fill` is used to complete the group.

Note: After removing the fill character from the last group (if it exists) and concatenating all the groups in order, the resultant string should be `s`.

Given the string `s`, the size of each group `k`, and the character `fill`, return a string array denoting the composition of every group `s` has been divided into, using the above procedure.

---

## Examples

### Example 1

**Input:**

    s = "abcdefghi", k = 3, fill = "x"

**Output:**

    ["abc", "def", "ghi"]

**Explanation:**

- The first 3 characters "abc" form the first group.
- The next 3 characters "def" form the second group.
- The last 3 characters "ghi" form the third group.
- Since all groups can be completely filled by characters from the string, we do not need to use fill.

### Example 2

**Input:**

    s = "abcdefghij", k = 3, fill = "x"

**Output:**

    ["abc", "def", "ghi", "jxx"]

**Explanation:**

- The first three groups are "abc", "def", and "ghi".
- For the last group, we can only use the character 'j' from the string. To complete this group, we add 'x' twice.
- The 4 groups formed are "abc", "def", "ghi", and "jxx".

---

## Constraints

- 1 <= s.length <= 100
- s consists of lowercase English letters only.
- 1 <= k <= 100
- fill is a lowercase English letter.

---

## Related Topics
- String
- Simulation

