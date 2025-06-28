package main

import (
	"reflect"
	"testing"
)

func TestMaxSubsequence(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{2, 1, 3, 3},
			k:    2,
			want: []int{3, 3},
		},
		{
			name: "Example 2",
			nums: []int{-1, -2, 3, 4},
			k:    3,
			want: []int{-1, 3, 4},
		},
		{
			name: "Example 3",
			nums: []int{3, 4, 3, 3},
			k:    2,
			want: []int{3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSubsequence(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSubsequence(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func BenchmarkMaxSubsequence(b *testing.B) {
	cases := []struct {
		name string
		nums []int
		k    int
	}{
		{"Example 1", []int{2, 1, 3, 3}, 2},
		{"Example 2", []int{-1, -2, 3, 4}, 3},
		{"Example 3", []int{3, 4, 3, 3}, 2},
	}
	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = maxSubsequence(c.nums, c.k)
			}
		})
	}
}
