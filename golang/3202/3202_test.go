package main

import "testing"

func TestMaximumLength(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		expect int
	}{
		{nums: []int{1, 2, 3, 4, 5}, k: 2, expect: 5},
		{nums: []int{1, 4, 2, 3, 1, 4}, k: 3, expect: 4},
		{nums: []int{1, 2, 3, 10, 2}, k: 6, expect: 3},
	}

	for i, tc := range tests {
		got := maximumLength(tc.nums, tc.k)
		if got != tc.expect {
			t.Errorf("case %d: maximumLength(%v, %d) = %d; want %d", i+1, tc.nums, tc.k, got, tc.expect)
		}
	}
}
