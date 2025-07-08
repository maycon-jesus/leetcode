package main

import "testing"

func TestMaxValue(t *testing.T) {
	tests := []struct {
		events [][]int
		k      int
		expect int
	}{
		{
			events: [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 1}},
			k:      2,
			expect: 7,
		},
		{
			events: [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 10}},
			k:      2,
			expect: 10,
		},
		{
			events: [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}, {4, 4, 4}},
			k:      3,
			expect: 9,
		},
	}

	for i, tc := range tests {
		got := maxValue(tc.events, tc.k)
		if got != tc.expect {
			t.Errorf("Test %d failed: got %d, want %d", i+1, got, tc.expect)
		}
	}
}
