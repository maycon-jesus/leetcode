package main

// Title: Candy
// Problem: https://leetcode.com/problems/candy/
// Difficulty: Hard
// Tags: Array, Greedy, Dynamic Programming

func main() {
	// Example usage
	ratings := []int{1, 0, 2}
	result := candy(ratings)
	println(result) // Output: 5

	// Example usage
	ratings = []int{1, 2, 2}
	result = candy(ratings)
	println(result) // Output: 4

	// Example usage
	ratings = []int{1, 3, 2, 2, 1}
	result = candy(ratings)
	println(result) // Output: 7
}

func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}
	// Passagem da esquerda para a direita
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	// Passagem da direita para a esquerda
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}
	total := 0
	for _, c := range candies {
		total += c
	}
	return total
}
