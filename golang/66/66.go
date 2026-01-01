package main

import "fmt"

func plusOne(digits []int) []int {
	arr := make([]int, len(digits)+1)
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		arr[i+1] = sum % 10
		carry = sum / 10
	}
	if carry > 0 {
		arr[0] = carry
	} else if arr[0] == 0 {
		return arr[1:]
	}
	return arr
}

func main() {
	fmt.Println(plusOne([]int{9}))
}
