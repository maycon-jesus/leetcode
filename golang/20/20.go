package main

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	opens := map[byte]byte{'(': ')', '{': '}', '[': ']'}
	closes := map[byte]byte{')': '(', '}': '{', ']': '['}
	stack := []byte{}

	for _, cha := range s {
		char := byte(cha)
		if open, ok := opens[char]; ok {
			stack = append(stack, open)
		} else if _, ok := closes[char]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != char {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
