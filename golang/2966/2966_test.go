package main

import (
	"reflect"
	"testing"
)

func TestDivideArray(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		expect [][]int
	}{
		{
			nums:   []int{1, 3, 4, 8, 7, 9, 3, 5, 1},
			k:      2,
			expect: [][]int{{1, 1, 3}, {3, 4, 5}, {7, 8, 9}},
		},
		{
			nums:   []int{2, 4, 2, 2, 5, 2},
			k:      2,
			expect: [][]int{},
		},
		{
			nums:   []int{4, 2, 9, 8, 2, 12, 7, 12, 10, 5, 8, 5, 5, 7, 9, 2, 5, 11},
			k:      14,
			expect: [][]int{{2, 2, 2}, {4, 5, 5}, {5, 5, 7}, {7, 8, 8}, {9, 9, 10}, {11, 12, 12}},
		},
	}

	for _, tc := range tests {
		result := divideArray(tc.nums, tc.k)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Errorf("divideArray(%v, %d) = %v; want %v", tc.nums, tc.k, result, tc.expect)
		}
	}
}
