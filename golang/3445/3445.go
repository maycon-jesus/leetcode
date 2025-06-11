package main

import "fmt"

func main() {
	fmt.Println(maxDifference("12233", 4))   // Output: -1
	fmt.Println(maxDifference("1122211", 3)) // Output: 1
}

func maxDifference(s string, k int) int {
	n := len(s)
	if n < 2 || k <= 0 {
		return -1
	}

	for 
}
