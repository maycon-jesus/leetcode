# Detailed Explanation - Valid Parentheses

## Problem Overview

The Valid Parentheses problem asks us to determine if a string containing only bracket characters is properly balanced and ordered. This is a classic stack application problem.

## Solution Strategy

### Why Use a Stack?

A stack is the perfect data structure for this problem because:
- Brackets must be closed in **reverse order** of how they were opened (LIFO - Last In, First Out)
- When we encounter a closing bracket, it must match the **most recent** unclosed opening bracket
- The stack naturally maintains this "most recent" relationship

### Key Observations

1. **Length Check**: Any valid parentheses string must have an even length. If the length is odd, we can immediately return `false`.

2. **Matching Pairs**: Each opening bracket has exactly one corresponding closing bracket:
   - `(` matches `)`
   - `{` matches `}`
   - `[` matches `]`

3. **Order Matters**: The string `"([)]"` is invalid even though it has matching pairs, because the brackets are not closed in the correct order.

## Algorithm Walkthrough

### Data Structures Used

```go
opens := map[byte]byte{'(': ')', '{': '}', '[': ']'}
closes := map[byte]byte{')': '(', '}': '{', ']': '['}
stack := []byte{}
```

- **opens**: Maps each opening bracket to its corresponding closing bracket
- **closes**: Maps each closing bracket to its corresponding opening bracket
- **stack**: Stores the expected closing brackets

### Step-by-Step Process

Let's trace through the example `"([])"``:

1. **Character '('**:
   - It's an opening bracket
   - Push its corresponding closing bracket `')'` onto stack
   - Stack: `[')']`

2. **Character '['**:
   - It's an opening bracket
   - Push its corresponding closing bracket `']'` onto stack
   - Stack: `[')', ']']`

3. **Character ']'**:
   - It's a closing bracket
   - Check stack: not empty ✓
   - Pop from stack and compare: `']'` == `']'` ✓
   - Stack: `[')']`

4. **Character ')'**:
   - It's a closing bracket
   - Check stack: not empty ✓
   - Pop from stack and compare: `')'` == `')'` ✓
   - Stack: `[]`

5. **End of string**:
   - Stack is empty ✓
   - Return `true`

### Why This Approach Works

**Clever Design Choice**: Instead of storing the opening brackets themselves on the stack, we store their **expected closing brackets**. This simplifies the comparison logic:

```go
// When we encounter an opening bracket
stack = append(stack, open)  // 'open' is the expected closing bracket

// When we encounter a closing bracket
if stack[len(stack)-1] != char {  // Direct comparison!
    return false
}
```

## Edge Cases Handled

1. **Empty stack on closing bracket**: `")("`
   - First character is `)`, but stack is empty
   - Return `false` immediately

2. **Mismatched bracket types**: `"(]"`
   - When we encounter `]`, top of stack is `)` (not `]`)
   - Return `false`

3. **Unclosed brackets**: `"(("`
   - After processing all characters, stack still has elements
   - Return `false`

4. **Odd length strings**: `"(()"`
   - Cannot be valid (early termination optimization)
   - Return `false`

## Code Analysis

### Line-by-Line Breakdown

```go
if len(s)%2 != 0 {
    return false  // Optimization: odd length can't be valid
}
```

```go
for _, cha := range s {
    char := byte(cha)  // Convert rune to byte for map lookup
```

```go
    if open, ok := opens[char]; ok {
        // It's an opening bracket
        stack = append(stack, open)  // Push expected closing bracket
```

```go
    } else if _, ok := closes[char]; ok {
        // It's a closing bracket
        if len(stack) == 0 || stack[len(stack)-1] != char {
            return false  // No matching opening or wrong type
        }
        stack = stack[:len(stack)-1]  // Pop from stack
    }
}
```

```go
return len(stack) == 0  // Valid only if all brackets were closed
```

## Complexity Analysis

### Time Complexity: O(n)
- Single pass through the string: O(n)
- Each character is processed exactly once
- Stack operations (push/pop) are O(1)
- Map lookups are O(1) average case

### Space Complexity: O(n)
- **Stack space**: In worst case (all opening brackets), we store n/2 elements
- **Map space**: O(1) - fixed size maps with 3 entries each
- Overall: O(n)

## Alternative Approaches

### 1. Without Length Check
You could skip the odd-length check, but it's a simple optimization that can save processing time.

### 2. Storing Opening Brackets
Instead of storing expected closing brackets, you could store the opening brackets and do a reverse lookup when encountering closing brackets. This is less elegant but equally correct.

### 3. Using Switch/Case
Instead of maps, you could use switch statements to determine bracket types. This might be slightly faster but less maintainable.

## Common Mistakes to Avoid

1. **Forgetting to check stack emptiness** before popping
2. **Not checking if stack is empty** at the end (unclosed brackets)
3. **Comparing wrong values** (e.g., comparing opening with closing directly)
4. **Not handling odd-length strings** efficiently

## Why This Solution is Elegant

1. **Clear separation**: Maps clearly define the bracket relationships
2. **Smart storage**: Storing expected closing brackets simplifies comparison
3. **Early termination**: Multiple exit points prevent unnecessary processing
4. **Readable**: The code's intent is clear and self-documenting
