package main

import "testing"

func TestCheckPrimeFrequency(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"one prime count", []int{1, 1, 1, 2, 2, 3}, true},                     // 1:3, 2:2, 3:1
		{"all unique", []int{1, 2, 3, 4}, false},                               // all counts = 1
		{"prime and non-prime counts", []int{1, 1, 2, 2, 2, 3, 3, 3, 3}, true}, // 1:2, 2:3, 3:4
		{"empty", []int{}, false},
		{"large prime count", []int{5, 5, 5, 5, 5, 5, 5}, true}, // 7 times
		{"example 1", []int{1, 2, 3, 4, 5, 4}, true},            // 4:2 (prime)
		{"example 2", []int{1, 2, 3, 4, 5}, false},              // all counts = 1
		{"example 3", []int{2, 2, 2, 4, 4}, true},               // 2:3, 4:2 (both prime)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPrimeFrequency(tt.nums); got != tt.want {
				t.Errorf("checkPrimeFrequency(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	nonPrimes := []int{-1, 0, 1, 4, 6, 8, 9, 10, 12, 15}
	for _, n := range primes {
		if !isPrime(n) {
			t.Errorf("isPrime(%d) = false, want true", n)
		}
	}
	for _, n := range nonPrimes {
		if isPrime(n) {
			t.Errorf("isPrime(%d) = true, want false", n)
		}
	}
}
